package data_base

import (
	"context"

	"authentication/helpers"
	pd "authentication/pkg"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DataBase struct {
	DBPool *pgxpool.Pool
}

func NewDB(ctx context.Context, cntStr string) *DataBase {

	conn, err := pgxpool.New(context.Background(), cntStr)

	helpers.ErrorHelper(err, "Failed to connect to database:")

	return &DataBase{DBPool: conn}
}

func (db *DataBase) ConnectClose() {

	db.DBPool.Close()
}

func (db *DataBase) Registration(ctx context.Context, req *pd.AuthenticationRequest) []string {

	var insertId string

	err := db.DBPool.QueryRow(
		ctx,
		"INSERT INTO users(email, username, password_hash) VALUES ($1, $2, $3) RETURNING id",
		req.Email, req.Username, req.Password).Scan(&insertId)

	helpers.ErrorHelper(err, "Failed to register user:")

	return []string{insertId, req.Username}
}

func (db *DataBase) SigningIn(ctx context.Context, req *pd.AuthenticationRequest) string {

	var username string

	db.DBPool.QueryRow(ctx, "SELECT username FROM users WHERE email = $1", req.Email).Scan(&username)

	return username
}
