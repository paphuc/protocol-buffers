In Employee to run the new code generation will be necessary to install the following gRPC gen plugins:

go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
Then use the following command to generate the code.

  # generate the messages
 protoc --go_out="$GO_GEN_PATH" -I "$dependecies" "$proto"

 # generate the services
 protoc --go-grpc_out="$GO_GEN_PATH" -I "$dependecies" "$proto"


 Example: protoc --go_out=paths=source_relative:.  --go-grpc_out=paths=source_relative:. ./transaction/transaction.proto