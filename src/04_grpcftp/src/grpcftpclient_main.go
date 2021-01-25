/* fork from https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
  "context"
  "log"
  "os"
  "time"
  "io"
  "errors"
  _"strconv"

  "google.golang.org/grpc"
  "example.com/04_grpcftp"
)

type GrpcFtpClient struct {
  grpc protos.FtpClient
}

func NewGrpcFtpClient(cc grpc.ClientConnInterface) *GrpcFtpClient {
  return &GrpcFtpClient {
    grpc: protos.NewFtpClient(cc),
  }
}

func (c *GrpcFtpClient) Hello(ctx context.Context, request string) (error, string) {
  r, err := c.grpc.Hello(ctx, &protos.HelloRequest{Message: request})
  if err != nil {
    log.Fatalf("Hello: failed0: %v", err)
    return err, ""
  }
  return nil, r.GetMessage()
}

func (c *GrpcFtpClient) List(ctx context.Context, path string, writ_cb WriteListCb) error {
  request_reader, err0 := c.grpc.List(ctx, &protos.ListRequest{File: &protos.File{Path: path, Type: protos.File_FILE, Size:0}})
  if err0 != nil {
    log.Fatalf("List: failed0: %v", err0)
    return err0
  }

  for {
    response, err := request_reader.Recv()
    if err == io.EOF {
      log.Printf("List: Done")
      return nil
    }

    if response == nil {
      err1 := errors.New("List: response is nil")
      log.Fatalf("%v", err1)
      return err1
    }

    writ_cb(response.File)
  }

  return nil
}

func (c *GrpcFtpClient) Pull(ctx context.Context, path string, write_cb WriteFileSegmentCb) error {
  request_reader, err0 := c.grpc.Pull(ctx, &protos.PullRequest{File: &protos.File{Path: path, Type: protos.File_FILE, Size:0}})
  if err0 != nil {
    log.Fatalf("Pull: failed0: %v", err0)
    return err0
  }

  read_cb := func() (*protos.FileSegment, error) {
    response, err := request_reader.Recv()
    if err == io.EOF {
      return nil, nil
    } else if err != nil {
      log.Printf("Pull: failed1: %v", err)
      return nil, err
    }
    log.Printf("Pull: %v %v", response.Segment.FileName, response.Segment.AvailableSize)
    return response.Segment, nil
  }

  if err := write_cb("pull", read_cb); err != nil {
    log.Printf("Pull: failed2: %v", err)
    return err
  }
  log.Printf("Pull: done")

  return nil
}

func (c *GrpcFtpClient) Push(ctx context.Context, path string, read_cb ReadFileSegmentCb) error {
  request_writer, err0 := c.grpc.Push(ctx)
  if err0 != nil {
    log.Fatalf("Push: failed0: %v", err0)
    return err0
  }

  request := &protos.PushRequest{
    Segment: &protos.FileSegment{},
  }

  write_cb := func(segment *protos.FileSegment) error {
    log.Printf("Push: %v %v", segment.FileName, segment.AvailableSize)
    if err := request_writer.Send(request); err != nil {
      return err
    }
    return nil
  }

  if err1 := read_cb(path, request.Segment, write_cb); err1 != nil {
    log.Printf("Push: failed1: %v", err1)
    return err1
  }

  response, err2 := request_writer.CloseAndRecv()
  if err2 != nil {
    log.Fatalf("Push: failed2: request_writer.CloseAndRecv: err:%v", err2)
    return err2
  }

  log.Printf("Push: done: response:%v", response)
  return nil
}

func main() {

  if len(os.Args) != 4 {
    log.Printf("%v localhost:50051 hello message\n", os.Args[0]);
    log.Printf("%v localhost:50051 ls from/path\n", os.Args[0]);
    log.Printf("%v localhost:50051 pull from/path\n", os.Args[0]);
    log.Printf("%v localhost:50051 push from/path\n", os.Args[0]);
    return
  }

  host := os.Args[1];
  cmd := os.Args[2];
  arg := os.Args[3];

  conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()

  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  grpcFtpClient := NewGrpcFtpClient(conn)

  if cmd == "hello" {
    err, response := grpcFtpClient.Hello(ctx, arg)
    if err != nil {
      log.Printf("Hello: failed: %v", err)
      return
    }
    log.Printf("Hello: %v", response)
  } else if cmd == "ls" {
    err := grpcFtpClient.List(ctx, arg, WriteListToConsole)
    if err != nil {
      log.Printf("List: failed: %v", err)
      return
    }
  } else if cmd == "pull" {
    err := grpcFtpClient.Pull(ctx, arg, WriteFileSegmentToFile)
    if err != nil {
      log.Printf("Pull: failed: %v", err)
      return
    }
  } else if cmd == "push" {
    err := grpcFtpClient.Push(ctx, arg, ReadFileToFileSegment)
    if err != nil {
      log.Printf("Push: failed: %v", err)
      return
    }
  }
}

