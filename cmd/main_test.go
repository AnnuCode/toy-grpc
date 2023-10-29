package main

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	userpb "github.com/AnnuCode/toy-grpc/gen/proto/user/v1"
	"github.com/AnnuCode/toy-grpc/internal/controller/user"
	grpchandler "github.com/AnnuCode/toy-grpc/internal/handler/grpc"
	"github.com/AnnuCode/toy-grpc/internal/repository/memory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func newServer(t *testing.T, register func(srv *grpc.Server)) *grpc.ClientConn {
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop()
	})

	register(srv)

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("srv.Serve %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	t.Cleanup(func() {
		conn.Close()
	})
	if err != nil {
		t.Fatalf("grpc.DialContext %v", err)
	}

	return conn
}
func TestUserService_GetUser(t *testing.T) {
	repo := memory.New()
	ctrl := user.New(repo)
	h := grpchandler.New(ctrl)
	conn := newServer(t, func(srv *grpc.Server) {
		userpb.RegisterGetUserServiceServer(srv, h)
	})

	client := userpb.NewGetUserServiceClient(conn)
	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{Id: 2})
	if err != nil {
		t.Fatalf("client.GetUser %v", err)
	}

	if res.User.Id != 2 && res.User.Fname != "Cristiano Ronaldo" {
		t.Fatalf("Unexpected values %v", res.User)
	}
}
func TestUserService_GetUsers(t *testing.T) {
	repo := memory.New()
	ctrl := user.New(repo)
	h := grpchandler.New(ctrl)
	conn := newServer(t, func(srv *grpc.Server) {
		userpb.RegisterGetUsersServiceServer(srv, h)
	})

	client := userpb.NewGetUsersServiceClient(conn)
	res, err := client.GetUsers(context.Background(), &userpb.GetUsersRequest{Id: []int64{2, 3}})
	if err != nil {
		t.Fatalf("client.GetUsers %v", err)
	}
	if len(res.Users) != 2 {
		t.Errorf("Expected 2 Users, got %d", len(res.Users))
	}
	expectedUsersDetails := []struct {
		Id      int64
		Fname   string
		City    string
		Phone   int64
		Height  float32
		Married bool
	}{
		// {1, "Lionel Messi", "Barcelona", 1234567891, 5.7, true},
		{2, "Cristiano Ronaldo", "Madrid", 1234737891, 6.1, false},
		{3, "Neymar", "Paris", 1234888891, 5.9, false},
		// {4, "Robert Lewandowski", "Munich", 1234544491, 6.1, true},
	}
	for i, user := range res.Users {
		expected := expectedUsersDetails[i]
		if user.Id != expected.Id || user.Fname != expected.Fname || user.City != expected.City || user.Phone != expected.Phone || user.Height != expected.Height || user.Married != expected.Married {
			t.Errorf("Wrong User details. Expected: %+v\n, Got: %+v\n", expected, user)
		}
	}
}
