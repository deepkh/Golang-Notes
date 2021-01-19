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

#include <iostream>
#include <fstream>
#include <string>
#include <grpcpp/grpcpp.h>
#include <grpchello.grpc.pb.h>

class GreeterClient {
 public:
  // Assembles the client's payload, sends it and presents the response back
  // from the server.
  static std::string SayHello(std::shared_ptr<protos::Greeter::Stub> stub, const std::string& user) {
    // Data we are sending to the server.
    protos::HelloRequest request;
    request.set_name(user);

    // Container for the data we expect from the server.
    protos::HelloReply reply;

    // Context for the client. It could be used to convey extra information to
    // the server and/or tweak certain RPC behaviors.
    grpc::ClientContext context;

    // The actual RPC.
    grpc::Status status = stub->SayHello(&context, request, &reply);

    // Act upon its status.
    if (status.ok()) {
      return reply.message();
    } else {
      std::cout << status.error_code() << ": " << status.error_message()
                << std::endl;
      return "RPC failed";
    }
  }
};

int main(int argc, char** argv) {
  std::string target_str = "localhost:50051";
  if (argc == 2) {
    target_str = std::string(argv[1]);
  }
  std::shared_ptr<grpc::Channel> channel = grpc::CreateChannel(
      target_str, grpc::InsecureChannelCredentials());
  std::shared_ptr<protos::Greeter::Stub> stub(protos::Greeter::NewStub(channel));
  std::string reply = GreeterClient::SayHello(stub, std::string("deepkh"));
  std::cout << "Greeter received: " << reply << std::endl;
  return 0;
}

