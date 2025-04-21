package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "gateway/pkg"
)

const (
	tlsCertFile = "certificate/server.crt"
	tlsKeyFile  = "certificate/server.key"
)

func main() {

	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Регистрируем gRPC Gateway хендлер
	err := pb.RegisterAutentificationServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", dialOpts)
	if err != nil {
		log.Fatalf("failed to register handler: %v", err)
	}

	log.Println("gRPC Gateway listening on https://locallhost:8080")
	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
