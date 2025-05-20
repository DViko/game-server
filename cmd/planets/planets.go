package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"planets/data_base"
	"planets/handler"

	//"planets/jwt"
	pkg "planets/pkg"
)

const (
	sCrt = "certificate/server.crt"
	sKey = "certificate/server.key"
)

func main() {
	db, err := data_base.NewDB(context.Background(), "postgres://postgres:root@localhost:5432/planets_data")
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.ConnectClose()

	authService := handler.NewPlanetsService(db)

	creds, _ := credentials.NewServerTLSFromFile(sCrt, sKey)
	//server := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(jwt.Interceptor()))
	server := grpc.NewServer(grpc.Creds(creds))

	pkg.RegisterPlanetsServiceServer(server, authService)
	lis, err := net.Listen("tcp", "localhost:50052")

	if err != nil {
		log.Fatal("failed to create listen:", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
