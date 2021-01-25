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

#include "grpcftp.h"

namespace GrpcFtp {

class Service final : public protos::Ftp::Service {
public:
  Service(GrpcFtp::ReadListCb read_list_cb
    , GrpcFtp::ReadFileSegmentCb read_file_segment_cb
    , std::string &save_path
    , GrpcFtp::WriteFileSegmentCb write_file_segment_cb) {
    read_list_cb_ = read_list_cb;
    read_file_segment_cb_ = read_file_segment_cb;
    save_path_ = save_path;
    write_file_segment_cb_ = write_file_segment_cb;
  };

  grpc::Status Hello(grpc::ServerContext* context
      , const protos::HelloRequest* request
      , protos::HelloResponse* response) override {
    static int count = 0;
    std::string message("Hello_Reponse ");
    message += std::to_string(count++);
    message += " ";
    message += request->message();
    response->set_message(message);
    printf("Hello: %s\n",message.c_str());
    return grpc::Status::OK;
  };

  grpc::Status List(grpc::ServerContext* context,
                      const protos::ListRequest* request,
                      grpc::ServerWriter<protos::ListResponse>* response_writer) override {
    protos::ListResponse response;

    auto write_cb = [&](protos::File &file) {
      *response.mutable_file() = file;
      response_writer->Write(response);
      return true;
    };

    grpc::Status status;
    if (!(status = read_list_cb_(request->file().path(), write_cb)).ok()) {
      return status;
    }

    return grpc::Status::OK;
  }

  grpc::Status Pull(grpc::ServerContext* context,
                      const protos::PullRequest* request,
                      grpc::ServerWriter<protos::PullResponse>* response_writer) override {
    grpc::Status status;
    protos::PullResponse response;
    protos::FileSegment *segment = response.mutable_segment();

    auto write_cb = [&](protos::FileSegment *segment) -> grpc::Status {
      printf("pull:%s %lu %lu\n", segment->file_name().c_str()
          , segment->mutable_buf()->size(), segment->available_size());
      response_writer->Write(response);
      return grpc::Status::OK;
    };

    if (!(status = read_file_segment_cb_(
      segment
      , (void *) request->file().path().c_str()
      , write_cb
    )).ok()) {
      return status;
    }

    return grpc::Status::OK;
  }

  grpc::Status Push(grpc::ServerContext* context
             , grpc::ServerReader<protos::PushRequest>* request_reader
             , protos::PushResponse* summary) override {
    protos::PushRequest request;

    auto read_cb = [&]() -> protos::FileSegment * {
      if (request_reader->Read(&request)) {
        return request.mutable_segment();
      }
      return nullptr;
    };

    auto status = write_file_segment_cb_((void *) save_path_.c_str(), read_cb); 

    if (!status.ok()) {
      return status;
    }

    summary->set_status(1);
    return grpc::Status::OK;
  }

private:
  GrpcFtp::ReadListCb read_list_cb_ = nullptr;
  GrpcFtp::ReadFileSegmentCb read_file_segment_cb_ = nullptr;
  std::string save_path_;
  GrpcFtp::WriteFileSegmentCb write_file_segment_cb_ = nullptr;
};
};

void RunServer() {
  std::string server_address("0.0.0.0:50051");

  std::string save_path = "push";
  GrpcFtp::Service service(GrpcFtp::ReadLocalFileList
      , GrpcFtp::ReadFileToFileSegment
      , save_path
      , GrpcFtp::WriteFileSegmentToFile);
  grpc::EnableDefaultHealthCheckService(true);
  grpc::reflection::InitProtoReflectionServerBuilderPlugin();
  grpc::ServerBuilder builder;
  // Listen on the given address without any authentication mechanism.
  builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
  // Register "service" as the instance through which we'll communicate with
  // clients. In this case it corresponds to an *synchronous* service.
  builder.RegisterService(&service);
  // Finally assemble the server.
  std::unique_ptr<grpc::Server> server(builder.BuildAndStart());
  std::cout << "Server listening on " << server_address << std::endl;

  // Wait for the server to shutdown. Note that some other thread must be
  // responsible for shutting down the server for this call to ever return.
  server->Wait();
}

int main(int argc, char** argv) {
  RunServer();
}

