package main

import (
	"context"
	"fmt"
	"github.com/thiago5171/grpc_learn/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.DialOption())
	if err != nil {
		log.Fatalf("Não foi possivel se conectar ao servidor gRPC: %v\n", err)
	}
	closeConnection(connection)
	client := pb.NewUserServiceClient(connection)
	AddUser(client)
}
func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "1",
		Name:  "thiago",
		Email: "t@email.com",
	}
	user, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("erro ao realizar requisição:%v", err)
	}

	fmt.Println(user)
}
func closeConnection(connection *grpc.ClientConn) {
	err := connection.Close()
	if err != nil {
		log.Fatalf("não foi possivel fechar a conecxão:%v\n", err)
	}
}
