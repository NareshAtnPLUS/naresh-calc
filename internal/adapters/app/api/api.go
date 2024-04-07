package api

import (
	"fmt"

	"github.com/MiniKartV1/calc/internal/models"
	"github.com/MiniKartV1/calc/internal/ports"
	"github.com/MiniKartV1/calc/internal/types"
)

type Adapter struct {
	calc ports.CalculationPort
	db   ports.DBPort
}

func NewAdapter(calc ports.CalculationPort, db ports.DBPort) *Adapter {
	return &Adapter{
		calc: calc,
		db:   db,
	}
}

func (api Adapter) Addition(params *types.CalcParameters) (*models.Calculation, error) {
	result, err := api.calc.Addition(params.X, params.Y)
	if err != nil {
		return nil, err
	}
	dbResult, err := api.db.SaveCompute(params, result, "+")
	if err != nil {
		return &models.Calculation{}, err
	}
	return dbResult, nil
}
func (api Adapter) Subtraction(params *types.CalcParameters) (*models.Calculation, error) {
	result, err := api.calc.Subtraction(params.X, params.Y)
	if err != nil {
		return nil, err
	}
	dbResult, err := api.db.SaveCompute(params, result, "-")
	if err != nil {
		return &models.Calculation{}, err
	}
	return dbResult, nil
}
func (api Adapter) Multiplication(params *types.CalcParameters) (*models.Calculation, error) {
	result, err := api.calc.Multiplication(params.X, params.Y)
	if err != nil {
		return nil, err
	}
	dbResult, err := api.db.SaveCompute(params, result, "*")
	if err != nil {
		return &models.Calculation{}, err
	}
	return dbResult, nil
}
func (api Adapter) Division(params *types.CalcParameters) (*models.Calculation, error) {
	result, remainder, err := api.calc.Division(params.X, params.Y)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Remainder: %d", remainder)
	dbResult, err := api.db.SaveCompute(params, result, "/")
	if err != nil {
		return &models.Calculation{}, err
	}
	return dbResult, nil
}

func (api Adapter) Health(email *string) (bool, error) {

	dbUser, err := api.db.FindUserByEmail(email)
	if !dbUser.IsActive || err != nil {
		return false, nil
	}
	return true, nil
}
