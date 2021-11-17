package httputils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/aws-cognito-example/infra/httpserver/dto"
)

func BindJSONAbort(c *gin.Context, req interface{}) {
	resp := new(dto.DefaultRespWithData)
	resp.SetError(dto.Code(http.StatusNotAcceptable), dto.Message("invalid request"))
	resp.SetData(req)

	c.JSON(http.StatusNotAcceptable, resp)
}

func InternalErrorAbort(c *gin.Context, message string) {
	resp := new(dto.DefaultResp)
	resp.SetError(dto.Code(http.StatusInternalServerError), dto.Message(message))

	c.JSON(http.StatusInternalServerError, resp)
}

func RespOK(c *gin.Context) {
	resp := new(dto.DefaultResp)
	resp.SetError(dto.Code(http.StatusOK), dto.Message("ok"))

	c.JSON(http.StatusOK, resp)
}
