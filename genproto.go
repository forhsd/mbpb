package protobuf

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//go:generate protoc --proto_path=./ --go_out=:./ --go-grpc_out=:./ ./entrypoint.proto

// protoc --proto_path=./protobuf --go_out=:./pb --go-grpc_out=:./pb ./protobuf/test.proto
