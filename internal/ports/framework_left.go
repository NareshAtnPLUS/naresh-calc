package ports

import "github.com/gin-gonic/gin"

type RESTPort interface {
	Run()
	Health(ctx *gin.Context)
	Addition(ctx *gin.Context)
	Subtraction(ctx *gin.Context)
	Multiplication(ctx *gin.Context)
	Division(ctx *gin.Context)
}
