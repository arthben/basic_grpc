package main

import (
	"context"
	"fmt"

	"github.com/arthben/basic_grpc/grpc_client/pkg/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var grpcServer = "127.0.0.1:8111"

func main() {
	conn, err := getConnection(grpcServer)
	if err != nil {
		panic(err)
	}

	client := contract.NewThingsTodoClient(conn)

	// Add new todo
	_, err = client.AddTask(context.Background(), &contract.AddTodoReq{
		Todo:    "Try this",
		Details: "Is it working?",
	})
	if err != nil {
		fmt.Printf(">> err while add new task: %v\n", err)
		panic(err)
	}
	fmt.Println(">> Success Added 1 todo")

	// List todo
	v, err := client.TaskList(context.Background(), &emptypb.Empty{})
	if err != nil {
		fmt.Printf(">> err while list todo: %v\n", err)
		panic(err)
	}
	fmt.Printf("v: %v\n", v)
}

func getConnection(srvAddr string) (*grpc.ClientConn, error) {
	// connection only implement tcp protocol
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.NewClient(srvAddr, opts...)
}
