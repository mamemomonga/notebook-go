package main

import (
	"flag"
	"log"
	"net"

	"example.com/grpc-souhoukou/pb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("info: SouHouKou Server")

	l := flag.String("listen", ":50001", "Server Address")
	flag.Parse()

	lis, err := net.Listen("tcp", *l)
	if err != nil {
		log.Fatalf("alert: failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSouHouKouServer(s, &server{})
	log.Printf("info: server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("alert: failed to serve: %v", err)
	}

}
