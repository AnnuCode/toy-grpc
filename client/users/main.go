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
	client := userpb.NewGetUsersServiceClient(conn)

	users, err := client.GetUsers(context.Background(), &userpb.GetUsersRequest{
		Id: []int64{2, 3, 4},
	})
	if err != nil {
		log.Fatalf("failed to GetUser: %v", err)
	}
	fmt.Println("Users Details:")
	for _, user := range users.Users {
		printUser(user)
	}
	// fmt.Printf("%v+\n", users)
}

func printUser(user *userpb.User) {
	fmt.Printf("ID: %d\n", user.Id)
	fmt.Printf("Name: %s\n", user.Fname)
	fmt.Printf("City: %s\n", user.City)
	fmt.Printf("Phone: %d\n", user.Phone)
	fmt.Printf("Height: %.2f\n", user.Height)
	fmt.Printf("Married: %t\n", user.Married)
	fmt.Printf("----------\n")
}
