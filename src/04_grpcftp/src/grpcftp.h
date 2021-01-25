/* fork from https://github.com/grpc/grpc/blob/master/examples/cpp/helloworld/greeter_server.cc
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
#ifndef _GRPCFTP_H_
#define _GRPCFTP_H_
#include <iostream>
#include <fstream>
#include <string>
#include <functional>
#include <grpcpp/grpcpp.h>
#include <grpcpp/health_check_service_interface.h>
#include <grpcpp/ext/proto_server_reflection_plugin.h>
#include <grpcftp.grpc.pb.h>
#include <ghc/filesystem.hpp>
#include "file_stream.h"
namespace fs = ghc::filesystem;

namespace GrpcFtp {

/**********************************************
 *ReadListCb
 **********************************************/
typedef std::function<grpc::Status(const std::string &path, std::function<bool(protos::File &file)> write_cb)> ReadListCb;

static grpc::Status ReadLocalFileList(const std::string &path, std::function<bool(protos::File &file)> write_cb) {
  fs::path dir = fs::u8path(path);
  grpc::Status status = grpc::Status::OK;
  protos::File file;

  try {
    auto rdi = fs::recursive_directory_iterator(dir);
      for(fs::directory_entry de : rdi) {
        file.set_type(de.is_directory() ? 
            protos::File_Type::File_Type_DIRECTORY
            : protos::File_Type::File_Type_FILE);
        file.set_path(de.path().string());
        file.set_size(de.file_size());
        write_cb(file);
      }
  }
  catch(fs::filesystem_error fe) {
    status = grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, fe.what());
  }

  return status;
}

/**********************************************
 *WriteListCb
 **********************************************/
typedef std::function<grpc::Status(const protos::File &file)> WriteListCb;

static grpc::Status WriteListToConsole(const protos::File &file) {
  printf("%d %s %lu\n", file.type(), file.path().c_str(), file.size());
  return grpc::Status::OK;
}

/**********************************************
 *ReadFileSegmentCb
 **********************************************/
typedef std::function<grpc::Status(protos::FileSegment *segment, void *arg
  , std::function<grpc::Status(protos::FileSegment *)> write_cb)> ReadFileSegmentCb;

static grpc::Status ReadFileToFileSegment(protos::FileSegment *segment, void *arg
    , std::function<grpc::Status(protos::FileSegment *)> write_cb) {
  grpc::Status status;

  fs2::FileReadStream frs;
  fs::path file = fs::u8path((const char *)arg);

  if (frs.Open(file.string())) {
    printf("failed to open %s for read\n", file.string().c_str());
    return grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, "failed to open file for read");
  }

  segment->set_file_name(file.filename());
  segment->set_available_size((uint64_t) frs.Length());

  while(segment->available_size() > 0) {
    if (segment->available_size() < 4096) {
      segment->mutable_buf()->resize(segment->available_size());
    } else {
      segment->mutable_buf()->resize(4096);
    }

    if (frs.Read(&(*segment->mutable_buf())[0], segment->mutable_buf()->size()) <= 0) {
      return grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, "failed to read file");
    }
    
    if (!(status = write_cb(segment)).ok()) {
      return grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, "failed to read file: read_cb err");
    }

    segment->set_available_size(
        segment->available_size() - segment->mutable_buf()->size());
  }

  return grpc::Status::OK;
};


/**********************************************
 *WriteFileSegmentCb
 **********************************************/
typedef std::function<grpc::Status(void *arg, std::function<protos::FileSegment *()> read_cb)> WriteFileSegmentCb;

static grpc::Status WriteFileSegmentToFile(void *arg, std::function<protos::FileSegment *()> read_cb)
{
  fs2::FileWriteStream fws;
  protos::FileSegment *segment;

  while ((segment = read_cb())) {
    if (!fws.IsOpen()) {
      fs::path path = fs::u8path((const char *)arg);
      fs::create_directory(path.string());
      path = path / fs::u8path(segment->file_name());
      if (fws.Open(path.string())) {
        printf("%s\n", path.string().c_str());
        return grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, "failed to open file for write");
      }
    }

    if (segment->available_size() == 0) {
      return grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, "failed to read file: available is zero");
    }

    printf("pull:%s %lu %lu\n", segment->file_name().c_str(), segment->buf().size(), segment->available_size());

    if (fws.Write(segment->buf().data(), segment->buf().size()) <= 0) {
      return grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, "failed to write");
    }
  }
  return grpc::Status::OK;
};


};

#endif
