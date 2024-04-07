package rest

import (
	"log"
	"net/http"

	"github.com/MiniKartV1/calc/internal/ports"
	"github.com/MiniKartV1/minikart-auth/pkg/middlewares"
	user_types "github.com/MiniKartV1/minikart-auth/pkg/types"
	"github.com/MiniKartV1/minikart-auth/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

var SERVER *gin.Engine

func (rest Adapter) Run() {
	var err error

	SERVER = gin.Default()
	SERVER.Use(middlewares.CORSMiddleware())
	apiRoutes := SERVER.Group("/api/calc")
	SERVER.GET("/api-calc/health", rest.Health)
	apiRoutes.Use(middlewares.JwtMiddleware())
	registerCalcRoutes(apiRoutes, &rest)
	err = SERVER.Run(":3100")
	if err != nil {
		log.Fatalf("Cannot Start the rest server")
	}
}

func registerCalcRoutes(router *gin.RouterGroup, rest *Adapter) {

	router.POST("/addition", rest.Addition)
	router.POST("/subtraction", rest.Subtraction)
	router.POST("/multiplication", rest.Multiplication)
	router.POST("/division", rest.Division)
}
func (rest Adapter) Health(ctx *gin.Context) {
	if claims, exists := ctx.Get("user"); exists {
		userClaims := claims.(*user_types.UserClaims) // Type assertion
		isActive, userErr := rest.api.Health(&userClaims.Email)
		if !isActive || userErr != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "user is not active. please contact dev",
			})
		}
	}

	clientType := ctx.GetHeader("X-Client-Type")
	if clientType == "web-app" {
		authRoutes := utils.GetRoutes(SERVER, "calc")
		ctx.JSON(http.StatusOK, gin.H{"status": "UP", "operations": authRoutes})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "UP"})
	return
}
