package handler

import (
	"context"

	db "planets/data_base"
	pkg "planets/pkg"
)

type PlanetsService struct {
	pkg.UnimplementedPlanetsServiceServer
	DB *db.DataBase
}

func NewPlanetsService(db *db.DataBase) *PlanetsService {
	return &PlanetsService{DB: db}
}

func (s *PlanetsService) AllPlanets(ctx context.Context, req *pkg.PlanetsRequest) (*pkg.PlanetsResponse, error) {

	result, err := s.DB.AllPlanets(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pkg.PlanetsResponse{
		Name:    result.Name,
		Type:    result.Type,
		TempMin: result.TempMin,
		TempMax: result.TempMax,
	}, nil
}

func (s *PlanetsService) HomePlanet(ctx context.Context, req *pkg.PlanetsRequest) (*pkg.PlanetsResponse, error) {

	uData, err := s.DB.HomePlanet(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pkg.PlanetsResponse{
		Name: uData.Name,
	}, nil
}

func (s *PlanetsService) Update(ctx context.Context, req *pkg.PlanetsRequest) (*pkg.PlanetsResponse, error) {

	uData, err := s.DB.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pkg.PlanetsResponse{
		Name: uData.Name,
	}, nil
}
