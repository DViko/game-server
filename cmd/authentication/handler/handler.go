package handler

import (
	pd "authentication/pkg"
	"context"
)

type Authentication struct {
	pd.UnimplementedAutentificationServiceServer
}

func (s *Authentication) Login(ctx context.Context, req *pd.LoginRequest) (*pd.LoginResponse, error) {
	return &pd.LoginResponse{
		Token: "token5647",
	}, nil
}
