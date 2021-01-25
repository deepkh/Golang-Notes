/* fork form https://github.com/grpc/grpc/blob/master/examples/cpp/helloworld/greeter_client.cc
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

static protos::ListRequest MakeListRequest(std::string path) {
  protos::ListRequest request;
  request.mutable_file()->set_type(protos::File_Type::File_Type_DIRECTORY);
  request.mutable_file()->set_path(path);
  return request;
};

static protos::PullRequest MakePullRequest(std::string path) {
  protos::PullRequest request;
  request.mutable_file()->set_type(protos::File_Type::File_Type_FILE);
  request.mutable_file()->set_path(path);
  return request;
};

class Client {
public:
  static std::string Hello(
    std::shared_ptr<protos::Ftp::Stub> stub
      , const std::string& message) {
    protos::HelloRequest request;
    request.set_message(message);

    protos::HelloResponse response;
    grpc::ClientContext context;
    
    grpc::Status status = stub->Hello(&context, request, &response);

    // Act upon its status.
    if (status.ok()) {
      return response.message();
    } else {
      std::cout << status.error_code() << ": " << status.error_message()
                << std::endl;
      return "RPC failed";
    }
  }
 
  static grpc::Status List(std::shared_ptr<protos::Ftp::Stub> stub
      , std::string request_path
      , WriteListCb write_cb) {
    grpc::ClientContext context;
    protos::ListRequest request = MakeListRequest(request_path);
    protos::ListResponse response;
    std::unique_ptr<grpc::ClientReader<protos::ListResponse> > reponse_reader(
        stub->List(&context, request));
    while(reponse_reader->Read(&response)) {
      auto status = write_cb(response.file());
      if (!status.ok()) {
        reponse_reader->Finish();
        return status;
      }
    }
    printf("List: Done: \n");
    return reponse_reader->Finish();
  }

  static grpc::Status Pull(std::shared_ptr<protos::Ftp::Stub> stub
      , std::string &request_path
      , void *arg
      , GrpcFtp::WriteFileSegmentCb write_cb) {
    grpc::ClientContext context;
    protos::PullRequest request = MakePullRequest(request_path);
    protos::PullResponse response;

    std::unique_ptr<grpc::ClientReader<protos::PullResponse> > reponse_reader(
        stub->Pull(&context, request));

    auto read_cb = [&]() -> protos::FileSegment * {
      if (reponse_reader->Read(&response)) {
        return response.mutable_segment();
      }
      return nullptr;
    };

    auto status = write_cb(arg, read_cb);
    if (!status.ok()) {
      reponse_reader->Finish();
      return status;
    }

    printf("Pull: Done: \n");
    return reponse_reader->Finish();
  }

  static grpc::Status Push(std::shared_ptr<protos::Ftp::Stub> stub
      , std::string &request_path
      , GrpcFtp::ReadFileSegmentCb read_cb) {
    grpc::Status status;
    grpc::ClientContext context;
    protos::PushRequest request;
    protos::PushResponse response;
    protos::FileSegment *segment = request.mutable_segment();

    std::unique_ptr<grpc::ClientWriter<protos::PushRequest> > request_writer(
        stub->Push(&context, &response));

    auto write_cb = [&](protos::FileSegment *segment) -> grpc::Status {
      printf("Push:%s %lu %lu\n", segment->file_name().c_str()
          , segment->mutable_buf()->size(), segment->available_size());
      if (!request_writer->Write(request)) {
        return grpc::Status(grpc::StatusCode::INVALID_ARGUMENT, "failed to write file: writer->Write err");
      }
      return grpc::Status::OK;
    };

    if (!(status = read_cb(
      segment
      , (void *) request_path.c_str()
      , write_cb
    )).ok()) {
      return status;
    }

    request_writer->WritesDone();
    status = request_writer->Finish();
    if (status.ok()) {
    }
    
    printf("Push: Done: %d\n", response.status());
    return status;
  }

};
};

int main(int argc, char** argv) {
  std::string host = "localhost:50051";
  std::string cmd;
  std::string request_path;

  if (argc != 4) {
    printf("%s localhost:50051 hello message\n", argv[0]);
    printf("%s localhost:50051 ls from/path\n", argv[0]);
    printf("%s localhost:50051 pull from/path\n", argv[0]);
    printf("%s localhost:50051 push from/path\n", argv[0]);
    return -1;
  }
  
  host = argv[1];
  cmd = argv[2];
  request_path = argv[3];
  
  std::shared_ptr<grpc::Channel> channel = grpc::CreateChannel(
      host, grpc::InsecureChannelCredentials());
  std::shared_ptr<protos::Ftp::Stub> stub(protos::Ftp::NewStub(channel));
  
  //Hello
  if (cmd.compare("hello") == 0) {
    std::string message = GrpcFtp::Client::Hello(stub, request_path);
    printf("hello: %s\n", message.c_str());
  //List directory
  } else if (cmd.compare("ls") == 0) {
    grpc::Status status = GrpcFtp::Client::List(stub, request_path, GrpcFtp::WriteListToConsole);
    if (!status.ok()) {
      std::string msg = status.error_message();
      printf("failed to FtpClient::List err:%d msg:%s\n", status.error_code(), msg.c_str());
      return -1;
    }
  //Pull a file from remote server
  } else if (cmd.compare("pull") == 0) {

    std::string save_path = "pull";    
    grpc::Status status = GrpcFtp::Client::Pull(stub
        , request_path
        , (void *)save_path.c_str()
        , GrpcFtp::WriteFileSegmentToFile);
    if (!status.ok()) {
      std::string msg = status.error_message();
      printf("failed to FtpClient::Pull err:%d msg:%s\n", status.error_code(), msg.c_str());
      return -1;
    }
  //Push a local file to server
  } else if (cmd.compare("push") == 0) {
    grpc::Status status = GrpcFtp::Client::Push(stub
        , request_path
        , GrpcFtp::ReadFileToFileSegment);
    if (!status.ok()) {
      std::string msg = status.error_message();
      printf("failed to FtpClient::Push err:%d msg:%s\n", status.error_code(), msg.c_str());
      return -1;
    }
  }

  return 0;
}

