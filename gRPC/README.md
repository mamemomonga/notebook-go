# gRPC

* [双方向通信](./bidirectional)

# クイックスタート

* https://grpc.io/docs/languages/go/quickstart/
* https://github.com/grpc/grpc-go/tree/master/examples/helloworld

## *.pb.go, *_grpc.pb.go を生成する

protocのインストール

	$ mkdir protoc
	$ cd protoc
	$ curl -o protoc.zip -L https://github.com/protocolbuffers/protobuf/releases/download/v3.17.3/protoc-3.17.3-osx-x86_64.zip
	$ unzip protoc.zip
	$ cp bin/protoc ~/bin/
	$ cd ..

プラグインのインストール

	$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	$ go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

grpc-hello の作成

	$ mkdir grpc-hello
	$ cd grpc-hello
	$ go mod init example.com/grpc-hello

protoファイルの作成

	$ mkdir protos

protos/hello.proto

	$ cat > protos/hello.proto << 'EOS'
	syntax = "proto3";
	option go_package = "example.com/grpc-hello/pb";

	package helloworld;

	service Greeter {
	  rpc SayHello (HelloRequest) returns (HelloReply) {}
	}
	message HelloRequest {
	  string name = 1;
	}
	message HelloReply {
	  string message = 1;
	}
	EOS

ProtocolBufferとgRPCコードの生成

	$ mkdir pb
	
protocの実行
	
	$ protoc \
	  --proto_path=./protos \
	  --go_out=./pb \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=./pb \
	  --go-grpc_opt=paths=source_relative \
	  hello.proto

## サーバの作成
  
	$ mkdir server
	
server/server.go

	$ cat > server/server.go << 'EOS'
	package main
	
	import (
		"context"
		"log"
		"net"
	
		"google.golang.org/grpc"
		pb "example.com/grpc-hello/pb"
	)
	
	const (
		port = ":50051"
	)
	type server struct {
		pb.UnimplementedGreeterServer
	}
	func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
		log.Printf("Received: %v", in.GetName())
		return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
	}
	
	func main() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterGreeterServer(s, &server{})
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
	EOS

## クライアントの作成

	$ mkdir client

client/client.go

	$ cat > client/client.go << 'EOS'
	package main
	
	import (
		"context"
		"log"
		"os"
		"time"
	
		"google.golang.org/grpc"
		pb "example.com/grpc-hello/pb"
	)
	
	const (
		address     = "localhost:50051"
		defaultName = "world"
	)
	
	func main() {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)
	
		name := defaultName
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
	EOS

###  ライブラリの取得

	$ go mod vendor

### サーバの起動

	$ go run ./server/

### 別ターミナルでクライアントの実行

	$ go run ./client/