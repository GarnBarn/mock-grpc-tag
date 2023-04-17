package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/GarnBarn/common-go/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedTagServer
}

func (p *server) GetTag(context.Context, *proto.TagRequest) (*proto.TagPublic, error) {
	return &proto.TagPublic{
		Id: "1",
	}, nil
}

func (p *server) IsTagExists(context.Context, *proto.TagRequest) (*proto.TagExistsResponse, error) {
	return &proto.TagExistsResponse{
		IsExists: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5002))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterTagServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
