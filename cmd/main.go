package main

import (
	"log"
	"net"

	userpb "github.com/AnnuCode/toy-grpc/gen/proto/user/v1"
	"github.com/AnnuCode/toy-grpc/internal/controller/user"
	grpchandler "github.com/AnnuCode/toy-grpc/internal/handler/grpc"
	"github.com/AnnuCode/toy-grpc/internal/repository/memory"
	"google.golang.org/grpc"
)

func main() {
	repo := memory.New()
	ctrl := user.New(repo)
	h := grpchandler.New(ctrl)

	lis, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userpb.RegisterGetUserServiceServer(grpcServer, h)
	userpb.RegisterGetUsersServiceServer(grpcServer, h)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
