package handler

import (
	"context"
	"log"

	db "authentication/data_base"
	jwt "authentication/jwt"
	pd "authentication/pkg"
)

type Authentication struct {
	pd.UnimplementedAuthenticationServiceServer
	DB *db.DataBase
}

func NewAuthenticationService(db *db.DataBase) *Authentication {
	return &Authentication{DB: db}
}

func (s *Authentication) SignUp(ctx context.Context, req *pd.AuthenticationRequest) (*pd.AuthenticationResponse, error) {

	uData := s.DB.Registration(ctx, req)

	return &pd.AuthenticationResponse{
		UserId:   uData.UserId,
		Username: req.Username,
		Token:    jwt.GenerateToken(uData.UserId),
		Error:    0,
	}, nil
}

func (s *Authentication) SignIn(ctx context.Context, req *pd.AuthenticationRequest) (*pd.AuthenticationResponse, error) {
	log.Println("result")
	result := s.DB.SigningIn(ctx, req)
	log.Println("result", result)

	return &pd.AuthenticationResponse{
		UserId:   result.UserId,
		Username: result.Username,
		Token:    jwt.GenerateToken(result.UserId),
		Error:    0,
	}, nil
}
