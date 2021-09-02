all: server client

server: pb/souhoukou.pb.go
	go build -o bin/server ./cmd/server

client: pb/souhoukou.pb.go
	go build -o bin/client ./cmd/client

pb:
	mkdir -p pb

pb/souhoukou.pb.go: pb protos/souhoukou.proto
	protoc \
	  --proto_path=./protos \
	  --go_out=./pb \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=./pb \
	  --go-grpc_opt=paths=source_relative \
	  souhoukou.proto

clean:
	rm -rf bin vendor

PHONY: all server client
