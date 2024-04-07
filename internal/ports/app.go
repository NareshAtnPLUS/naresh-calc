package ports

import (
	"github.com/MiniKartV1/calc/internal/models"
	"github.com/MiniKartV1/calc/internal/types"
)

type APIPort interface {
	Addition(params *types.CalcParameters) (*models.Calculation, error)
	Subtraction(params *types.CalcParameters) (*models.Calculation, error)
	Multiplication(params *types.CalcParameters) (*models.Calculation, error)
	Division(params *types.CalcParameters) (*models.Calculation, error)
	Health(email *string) (bool, error)
}
