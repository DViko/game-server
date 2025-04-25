package handler

import (
	"context"

	db "authentication/data_base"
	jwt "authentication/jwt"
	pd "authentication/pkg"
)

type Authentication struct {
	pd.UnimplementedRegistrationServiceServer
	pd.UnimplementedSigningInServiceServer
	DB *db.DataBase
}

func NewAuthenticationService(db *db.DataBase) *Authentication {

	return &Authentication{DB: db}
}

func (s *Authentication) Registration(ctx context.Context, req *pd.AuthenticationRequest) (*pd.AuthenticationResponse, error) {
	uData := s.DB.Registration(ctx, req)
	token := jwt.GenerateToken(uData)

	return &pd.AuthenticationResponse{
		UserId:   uData[0],
		Username: uData[1],
		Token:    token,
		Error:    0,
	}, nil
}

func (s *Authentication) SigningIn(ctx context.Context, req *pd.AuthenticationRequest) (*pd.AuthenticationResponse, error) {
	//test := s.DB.SigningIn(ctx, req)

	return &pd.AuthenticationResponse{
		//UserId:   test,
		//Username: test,
		//Token:    jwt.GenerateToken(test),
		Error: 0,
	}, nil
}
