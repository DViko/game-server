package main

import (
	"log"
	"net"

	"authentication/handler"
	pd "authentication/pkg"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	credentials, err := credentials.NewServerTLSFromFile("certificate/server.crt", "certificate/server.key")
	if err != nil {
		log.Fatal("failed to load TLS keys", err)
	}

	server := grpc.NewServer(grpc.Creds(credentials))
	pd.RegisterAutentificationServiceServer(server, &handler.Authentication{})
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
