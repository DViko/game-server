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

func (db *DataBase) Registration(ctx context.Context, req *pd.RegistrationRequest) string {

	var username string

	db.DBPool.QueryRow(ctx, "SELECT username FROM users WHERE email = $1", req.Username).Scan(&username)

	return username
}
