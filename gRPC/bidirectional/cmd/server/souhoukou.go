package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"example.com/grpc-souhoukou/pb"
	"google.golang.org/grpc/peer"
)

var (
	serial int32
)

type server struct {
	pb.UnimplementedSouHouKouServer
}

func getClientIP(ctx context.Context) string {
	if pr, ok := peer.FromContext(ctx); ok {
		addr := pr.Addr.String()
		return addr[:strings.LastIndex(addr, ":")]
	}
	return ""
}

func (s *server) Chat(stream pb.SouHouKou_ChatServer) error {
	for {
		// 受信
		in, err := stream.Recv()
		if err == io.EOF {
			continue
		}
		if err != nil {
			return err
		}
		ser := in.GetSerial()
		id := in.GetId()
		mes := in.GetMessage()
		log.Printf("info: SERVER [RECV] IP: %s / ID: %d / Serial: %d / Message: %s",
			getClientIP(stream.Context()),
			id, ser, mes)

		// 返答
		err = stream.Send(&pb.ChatReply{
			Serial:  serial,
			Message: fmt.Sprintf("ID %d さん、合計 %d 番目のこんにちは", id, serial),
		})
		serial++

		if err != nil {
			return err
		}
	}
}
