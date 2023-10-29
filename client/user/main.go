package main

import (
	"context"
	"fmt"
	"log"

	userpb "github.com/AnnuCode/toy-grpc/gen/proto/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial("localhost:9879", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := userpb.NewGetUserServiceClient(conn)

	user, err := client.GetUser(context.Background(), &userpb.GetUserRequest{
		Id: 2,
	})
	if err != nil {
		log.Fatalf("failed to GetUser: %v", err)
	}
	fmt.Println("User Details:")
	printUser(user)
}
func printUser(user *userpb.GetUserResponse) {
	fmt.Printf("ID: %d\n", user.User.Id)
	fmt.Printf("Name: %s\n", user.User.Fname)
	fmt.Printf("City: %s\n", user.User.City)
	fmt.Printf("Phone: %d\n", user.User.Phone)
	fmt.Printf("Height: %.2f\n", user.User.Height)
	fmt.Printf("Married: %t\n", user.User.Married)
	fmt.Printf("----------\n")
}
