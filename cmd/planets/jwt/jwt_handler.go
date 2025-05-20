package jwt

import (
	"context"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var key = []byte("the_game_aythentication")

type contextKey string

const userIdKey contextKey = "userId"

type Claims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(uId string, uName string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   uId,
		"username": uName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, err
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, err
	}

	return claims, nil
}

func Interceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		AHeader := md.Get("authorization")

		if len(AHeader) == 0 || !strings.HasPrefix(AHeader[0], "Bearer ") {
			return nil, status.Error(codes.Unauthenticated, "missing or invalid token")
		}

		token := strings.TrimPrefix(AHeader[0], "Bearer ")
		claims, err := ValidateToken(token)

		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		ctx = context.WithValue(ctx, userIdKey, claims.UserId)

		return handler(ctx, req)
	}
}
