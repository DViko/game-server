package main

import (
	"net"

	"authentication/handler"
	"authentication/helpers"
	pd "authentication/pkg"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"authentication/data_base"
)

const (
	sCrt = "certificate/server.crt"
	sKey = "certificate/server.key"
)

func main() {

	db := data_base.NewDB(context.Background(), "postgres://postgres:root@localhost:5432/authentication")
	authService := handler.NewAuthenticationService(db)

	creds, _ := credentials.NewServerTLSFromFile(sCrt, sKey)
	server := grpc.NewServer(grpc.Creds(creds))

	pd.RegisterAuthenticationServiceServer(server, authService)
	lis, err := net.Listen("tcp", "localhost:50051")

	helpers.ErrorHelper(err, "failed to listen:")

	if err := server.Serve(lis); err != nil {
		helpers.ErrorHelper(err, "failed to serve:")
	}
}
