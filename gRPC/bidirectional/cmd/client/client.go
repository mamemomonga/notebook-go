package main

import (
	"flag"
	"log"
	"time"

	"example.com/grpc-souhoukou/pb"
	"google.golang.org/grpc"
)

var (
	serial int32
	id     int32
)

func main() {
	log.Println("info: SouHouKou Client")

	h := flag.String("host", "localhost:50001", "Server Address")
	i := flag.Int("id", 1, "Your ID")
	flag.Parse()

	log.Printf("info: Server: %s / ID: %d", *h, *i)
	serial = 1
	id = int32(*i)

	// サーバがダウンしてもリトライし続ける
	for {
		conn, err := grpc.Dial(*h, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("alert: did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewSouHouKouClient(conn)
		chatRunner(c)
		conn.Close()
		time.Sleep(time.Second * 1)
	}
}
