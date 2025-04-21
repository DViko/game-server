package main

import (
	"fmt"
	"log"
	"net"

	"authentication/handler"
	pd "authentication/pkg"

	"google.golang.org/grpc"

	"authentication/data_base"
)

func main() {
	server := grpc.NewServer()

	pd.RegisterAutentificationServiceServer(server, &handler.Authentication{})
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("server listening at", lis.Addr())

	d := data_base.DataBase{}
	d.ConnectDB()

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
