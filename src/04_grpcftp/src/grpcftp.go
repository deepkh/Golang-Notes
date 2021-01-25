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
  "log"
  "errors"
  "strconv"
  "path/filepath"
  "os"

  "example.com/04_grpcftp"
)

/**********************************************
 *ReadListCb
 **********************************************/
type ReadListCb func(path string, write_cb func(file *protos.File) error) error

func ReadLocalFileList(path string, write_cb func(file *protos.File) error) error {
  file := &protos.File{Path: path}
  return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
      file.Size = uint64(info.Size())

      if info.IsDir() {
        file.Type = protos.File_DIRECTORY
      } else {
        file.Type = protos.File_FILE
      }

      if err := write_cb(file); err != nil {
        return err
      }
      return nil
  })
}


/**********************************************
 *WriteListCb
 **********************************************/
type WriteListCb func(segment *protos.File) error

func WriteListToConsole(segment *protos.File) error {
  log.Printf("List: %v", segment)
  return nil
}

/**********************************************
 *ReadFileSegmentCb
 **********************************************/
type ReadFileSegmentCb func(path string, segment *protos.FileSegment,
  write_cb func(segment *protos.FileSegment) error) error

func ReadFileToFileSegment(path string, segment *protos.FileSegment,
  write_cb func(segment *protos.FileSegment) error) error {

  file, err := os.Open(path)
  if err != nil {
    return err
  }

  info, err := file.Stat()
  if err != nil {
    return err
  }

  segment.FileName = info.Name()
  //segment.Buf = make([]byte, 1)
  segment.AvailableSize = uint64(info.Size())

  var buf []byte;
  for segment.AvailableSize > 0 {
    if segment.AvailableSize > 4096 {
      buf = make([]byte, 4096)
    } else {
      buf = make([]byte, segment.AvailableSize)
    }

    n, err := file.Read(buf)

    if err != nil {
      return err
    }

    if uint64(n) != segment.AvailableSize {
      buf = buf[:n]
    }

    segment.Buf = buf

    if err := write_cb(segment); err != nil {
      return err
    }

    segment.AvailableSize -= uint64(n)
  }

  return nil
}


/**********************************************
 *WriteFileSegmentCb
 **********************************************/
type WriteFileSegmentCb func (path string, read_cb func() (*protos.FileSegment,error)) error

func WriteFileSegmentToFile(path string, read_cb func() (*protos.FileSegment,error)) error {
  var file *os.File = nil
  var segment *protos.FileSegment;
  var err error
  var n int

  for {

    if segment,err = read_cb(); err != nil {
      return err
    }

    if segment == nil {
      return nil
    }

    if file == nil {
      os.MkdirAll(path, 0755)
      file, err = os.Create(path+"/"+segment.FileName)
      if err != nil {
        log.Printf("err2: %v", err)
        return err
      }
    }

    if segment.AvailableSize == 0 {
      return errors.New("segment.AvailableSize is zero")
    }

    n, err = file.Write(segment.Buf)

    if err != nil {
      log.Printf("err3: %v", err)
      return err
    }

    if n != len(segment.Buf) {
      err = errors.New("file.Write len "+strconv.Itoa(n)+" != "+strconv.Itoa(len(segment.Buf)))
      log.Printf("err4: %v", err)
      return err
    }
  }

  return nil
}

