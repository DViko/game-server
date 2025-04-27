package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "gateway/pkg"
)

const (
	sCrt = "certificate/server.crt"
	sKey = "certificate/server.key"
)

func main() {

	creds, _ := credentials.NewServerTLSFromFile(sCrt, sKey)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	pb.RegisterAuthenticationServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", dialOpts)

	http.ListenAndServeTLS("localhost:8080", sCrt, sKey, mux)
}
