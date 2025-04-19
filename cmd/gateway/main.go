package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "gateway/pkg"
)

const (
	tlsCertFile = "certificate/server.crt"
	tlsKeyFile  = "certificate/server.key"
)

func main() {

	cert, err := tls.LoadX509KeyPair(tlsCertFile, tlsKeyFile)
	if err != nil {
		log.Fatalf("failed to load TLS certificates: %v", err)
	}

	tlsCreds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
	})

	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCreds)}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Регистрируем gRPC Gateway хендлер
	err = pb.RegisterAutentificationServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", dialOpts)
	if err != nil {
		log.Fatalf("failed to register handler: %v", err)
	}

	log.Println("gRPC Gateway listening on https://localhost:8080")
	err = http.ListenAndServeTLS(
		":8080",
		"certificate/server.crt",
		"certificate/server.key",
		mux,
	)
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
