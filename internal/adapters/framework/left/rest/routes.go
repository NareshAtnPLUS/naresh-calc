package rest

import (
	"net/http"

	"github.com/MiniKartV1/calc/internal/types"
	"github.com/gin-gonic/gin"
)

func (rest Adapter) Addition(ctx *gin.Context) {
	var calculationParameters types.CalcParameters
	if err := ctx.ShouldBind(&calculationParameters); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	result, err := rest.api.Addition(&calculationParameters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "computation successful",
		"result":  result,
	})
	return

}
func (rest Adapter) Subtraction(ctx *gin.Context) {
	var calculationParameters types.CalcParameters
	if err := ctx.ShouldBind(&calculationParameters); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	result, err := rest.api.Subtraction(&calculationParameters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "computation successful",
		"result":  result,
	})
	return
}
func (rest Adapter) Multiplication(ctx *gin.Context) {
	var calculationParameters types.CalcParameters
	if err := ctx.ShouldBind(&calculationParameters); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	result, err := rest.api.Multiplication(&calculationParameters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "computation successful",
		"result":  result,
	})
	return
}
func (rest Adapter) Division(ctx *gin.Context) {
	var calculationParameters types.CalcParameters
	if err := ctx.ShouldBind(&calculationParameters); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	result, err := rest.api.Division(&calculationParameters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "computation successful",
		"result":  result,
	})
	return
}
