package app

import (
	"github.com/gin-gonic/gin"
	"github.com/liurongdev/firstGoWeb/model"
	"net/http"
)

func OK(ctx *gin.Context, data interface{}, msg string) {
	var response model.Response
	if msg != "" {
		response.Msg = msg
	} else {
		response.Msg = "success"
	}
	response.Data = data
	ctx.JSON(http.StatusOK, response.OK())
}

func ERROR(ctx *gin.Context, data interface{}, msg string, code int) {
	var response model.Response
	if msg != "" {
		response.Msg = msg
	}
	response.Data = data
	ctx.JSON(http.StatusOK, response.ERROR(code))
}
