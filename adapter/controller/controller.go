package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"ginweb/domain"

	"github.com/gin-gonic/gin"
	"github.com/glassonion1/logz"
)

func parsePathInt(ctx *gin.Context, key string) (int, bool) {
	value, err := strconv.Atoi(ctx.Param(key))
	fmt.Println("value: ", value)
	if err == nil {
		return value, true
	}
	res := ctx.Error(err).SetType(gin.ErrorTypePublic)
	logz.Errorf(ctx, "Error:%v", res)
	return 0, false
}

func response(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, domain.Response{
		Message: http.StatusText(code),
		Data:    data,
	})
}

func responseCreated(ctx *gin.Context, data interface{}) {
	response(ctx, http.StatusCreated, data)
}

func responseNoContent(ctx *gin.Context) {
	response(ctx, http.StatusNoContent, nil)
}

func responseOK(ctx *gin.Context, data interface{}) {
	response(ctx, http.StatusOK, data)
}

func checkError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	res := ctx.Error(err).SetType(gin.ErrorTypePublic)
	logz.Errorf(ctx, "Error:%v", res)
	return true
}

func bindJSON(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindJSON(obj)
	if err == nil {
		return true
	}

	res := ctx.Error(err).SetType(gin.ErrorTypeBind)
	logz.Errorf(ctx, "Error:%v", res)
	return false
}
