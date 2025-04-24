package main

import (
	"fmt"
	"log"
	"net"

	"authentication/handler"
	"authentication/helpers"
	pd "authentication/pkg"
	"context"

	"google.golang.org/grpc"

	"authentication/data_base"
)

func main() {
	server := grpc.NewServer()

	db := data_base.NewDB(context.Background(), "postgres://postgres:root@localhost:5432/authentication")
	authService := handler.NewAuthenticationService(db)

	pd.RegisterRegistrationServiceServer(server, authService)
	lis, err := net.Listen("tcp", "localhost:50051")

	helpers.ErrorHelper(err, "failed to listen:")

	fmt.Println("server listening at", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
