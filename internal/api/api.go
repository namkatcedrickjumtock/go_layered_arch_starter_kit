package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/namkatcedrickjumtock/go_layered_arch_starter_kit/internal/models"
	"github.com/namkatcedrickjumtock/go_layered_arch_starter_kit/internal/services/echo"
)

func NewAPIListener(echoer echo.Echoer) (*gin.Engine, error) {
	route := gin.Default()
	route.Use(cors.Default())

	route.GET("/echo", func(ctx *gin.Context) {
		msg := ctx.Query("msg")

		resp, err := echoer.Echo(ctx, msg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, resp)
	})

	return route, nil
}
