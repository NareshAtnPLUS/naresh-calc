package ports

import (
	"github.com/MiniKartV1/calc/internal/models"
	"github.com/MiniKartV1/calc/internal/types"
	user_models "github.com/MiniKartV1/minikart-auth/pkg/models"
)

type DBPort interface {
	CloseDBConnection()
	FindUserByEmail(email *string) (*user_models.User, error)
	SaveCompute(parms *types.CalcParameters, result int, operation string) (*models.Calculation, error)
}
