package handler

import (
	"context"

	db "authentication/data_base"
	jwt "authentication/jwt"
	pd "authentication/pkg"
)

type Authentication struct {
	pd.UnimplementedRegistrationServiceServer
	DB *db.DataBase
}

func NewAuthenticationService(db *db.DataBase) *Authentication {

	return &Authentication{DB: db}
}

func (s *Authentication) Registration(ctx context.Context, req *pd.RegistrationRequest) (*pd.RegistrationResponse, error) {
	test := s.DB.Registration(ctx, req)
	token := jwt.GenerateToken(test)

	return &pd.RegistrationResponse{
		Username: test,
		Token:    token,
		Error:    0,
	}, nil
}
