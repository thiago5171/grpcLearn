package main

import (
	"github.com/thiago5171/grpc_learn/pb"
	"github.com/thiago5171/grpc_learn/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Erro ao se conectar: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, service.NewUserService())
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Não foi possíel se conectar: %v\n", err)
	}
}
