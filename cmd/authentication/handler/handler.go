package handler

import (
	"context"

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
		Token:    jwt.GenerateToken(uData.UserId, uData.Username),
		Error:    0,
	}, nil
}

func (s *Authentication) SignIn(ctx context.Context, req *pd.AuthenticationRequest) (*pd.AuthenticationResponse, error) {

	uData := s.DB.SigningIn(ctx, req)

	return &pd.AuthenticationResponse{
		UserId:   uData.UserId,
		Username: uData.Username,
		Token:    jwt.GenerateToken(uData.UserId, uData.Username),
		Error:    0,
	}, nil
}

func (s *Authentication) Update(ctx context.Context, req *pd.AuthenticationRequest) (*pd.AuthenticationResponse, error) {
	uData := s.DB.Update(ctx, req)

	return &pd.AuthenticationResponse{
		UserId:   uData.UserId,
		Username: req.Username,
		Token:    jwt.GenerateToken(uData.UserId, uData.Username),
		Error:    0,
	}, nil
}
