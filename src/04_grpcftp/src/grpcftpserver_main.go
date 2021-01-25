/* fork from https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
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
package main

import (
  "context"
  "log"
  "net"
  "io"

  "google.golang.org/grpc"
  "example.com/04_grpcftp"
)

const (
  port = ":50051"
)

type GrpcFtpService struct {
  read_list_cb ReadListCb
  read_file_segment_cb ReadFileSegmentCb
  write_file_segment_cb WriteFileSegmentCb
  protos.UnimplementedFtpServer
}

func NewGrpcFtpService(read_list_cb ReadListCb,
  read_file_segment_cb ReadFileSegmentCb,
  write_file_segment_cb WriteFileSegmentCb) *GrpcFtpService {
  return &GrpcFtpService {
    read_list_cb, read_file_segment_cb, write_file_segment_cb, protos.UnimplementedFtpServer{},
  }
}

func (s *GrpcFtpService) Hello(ctx context.Context, request *protos.HelloRequest) (*protos.HelloResponse, error) {
  log.Printf("Hello: %v", request.GetMessage())
  return &protos.HelloResponse{Message: "Hello: " + request.GetMessage()}, nil
}

func (s *GrpcFtpService) List(request *protos.ListRequest, response_writer protos.Ftp_ListServer) error {
  log.Printf("List: %v", request.File)

  response := &protos.ListResponse{}
  write_cb := func(file *protos.File) error {
    response.File = file
    if err := response_writer.Send(response); err != nil {
      return err
    }
    return nil
  }

  return s.read_list_cb(request.GetFile().GetPath(), write_cb)
}


func (s *GrpcFtpService) Pull(request *protos.PullRequest, response_writer protos.Ftp_PullServer) error {
  log.Printf("Pull: %v", request.File)
  response := &protos.PullResponse{
    Segment: &protos.FileSegment{},
  }

  write_cb := func(segment *protos.FileSegment) error {
    log.Printf("Pull: %v %v", segment.FileName, segment.AvailableSize)
    if err := response_writer.Send(response); err != nil {
      return err
    }
    return nil
  }

  if err := s.read_file_segment_cb(request.File.Path, response.Segment, write_cb); err != nil {
    return err
  }

  log.Printf("Pull: done")
  return nil
}


func (s *GrpcFtpService) Push(request_reader protos.Ftp_PushServer) error {
  var err error
  var request *protos.PushRequest = nil

  read_cb := func() (*protos.FileSegment, error) {
    request, err = request_reader.Recv()
    if err == io.EOF {
      return nil, request_reader.SendAndClose(&protos.PushResponse{Status: 1})
    } else if err != nil {
      log.Printf("err1: %v", err)
      return nil, err
    }
    log.Printf("Push: %v %v", request.Segment.FileName, request.Segment.AvailableSize)
    return request.Segment, nil
  }

  if err = s.write_file_segment_cb("push", read_cb); err != nil {
    return err
  }

  log.Printf("Push: done")
  return nil
}


func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  protos.RegisterFtpServer(s,
    NewGrpcFtpService(ReadLocalFileList, ReadFileToFileSegment, WriteFileSegmentToFile))
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}

