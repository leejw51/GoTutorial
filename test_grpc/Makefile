all: convert
	go build grpc_server.go
	go build grpc_client.go
convert:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	protoc --go_out=. --go-grpc_out=. ./proto/*.proto
