package main

import (
	"context"
	"fmt"
	pb "grpcPractice/proto"
	"log"

	"google.golang.org/grpc"
)

var grpcClient pb.MyprotoServiceClient

func main() {
	grpcConn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connec: %v", err)
	}
	defer grpcConn.Close()
	grpcClient = pb.NewMyprotoServiceClient(grpcConn)

	AddUser("test1", "123456")
	AddUser("test2", "123456")
	LoginUser("test1", "123456")
	UserList()
}

func AddUser(username, userpwd string) {
	res, err := grpcClient.AddUser(context.Background(),
		&pb.UserRequest{
			UserName: username,
			UserPwd:  userpwd,
		})
	if err != nil {
		log.Fatalf("failed to add user: %v", err)
	}
	fmt.Println(res.Result)
}

func LoginUser(username, userpwd string) {
	res, err := grpcClient.LoginUser(context.Background(),
		&pb.UserRequest{
			UserName: username,
			UserPwd:  userpwd,
		})
	if err != nil {
		log.Fatalf("failed to login user: %v", err)
	}
	fmt.Println(username, res.Result)
}

func UserList() {
	res, err := grpcClient.UserList(context.Background(), &pb.UserListRequest{})
	if err != nil {
		log.Fatalf("User List Error: ", err)
	}
	fmt.Println(res.Result, " / ", res.UserName)
}
