package data_base

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	pkg "planets/pkg"
)

type DataBase struct {
	DBPool *pgxpool.Pool
}

func NewDB(ctx context.Context, cntStr string) (*DataBase, error) {

	conn, err := pgxpool.New(context.Background(), cntStr)

	if err != nil {
		return nil, err
	}

	return &DataBase{DBPool: conn}, nil
}

func (db *DataBase) ConnectClose() {

	db.DBPool.Close()
}

func (db *DataBase) AllPlanets(ctx context.Context, req *pkg.PlanetsRequest) (*pkg.PlanetsResponse, error) {

	var result pkg.PlanetsResponse

	err := db.DBPool.QueryRow(
		ctx,
		"SELECT planet_name, planet_type, temp_min, temp_max FROM planets WHERE owner_id = $1",
		"69d1e939-a46c-4a50-abce-a5f3722951a8").Scan(&result.Name, &result.Type, &result.TempMin, &result.TempMax)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (db *DataBase) HomePlanet(ctx context.Context, req *pkg.PlanetsRequest) (*pkg.PlanetsResponse, error) {

	var result pkg.PlanetsResponse

	err := db.DBPool.QueryRow(
		ctx,
		"SELECT planet_name, planet_type, temp_min, temp_max FROM planets WHERE owner_id = $1 AND home = $2",
		"69d1e939-a46c-4a50-abce-a5f3722951a8", true).Scan(
		&result.Name,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DataBase) Update(ctx context.Context, req *pkg.PlanetsRequest) (*pkg.PlanetsResponse, error) {

	var uData pkg.PlanetsResponse

	err := db.DBPool.QueryRow(ctx, "").Scan()
	if err != nil {
		return nil, err
	}

	return &uData, nil
}
