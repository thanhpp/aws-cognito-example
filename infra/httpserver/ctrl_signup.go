package httpserver

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/aws-cognito-example/infra/cognitolib"
	"github.com/thanhpp/aws-cognito-example/infra/httpserver/dto"
	"github.com/thanhpp/aws-cognito-example/infra/httpserver/httputils"
)

type Controller struct {
	cognitoApp *cognitolib.CognitoApp
}

func (ctrl Controller) Signup(c *gin.Context) {
	req := new(dto.SignupReq)

	if err := c.ShouldBindJSON(req); err != nil {
		log.Printf("bind json %v \n", err)
		httputils.BindJSONAbort(c, req)
		return
	}

	err := ctrl.cognitoApp.SignUp(
		req.Email,
		req.Password,
		req.PhoneNumber,
		req.Gender,
		req.FullName,
		req.DOB,
	)
	if err != nil {
		log.Printf("cognito signup %v \n", err)
		httputils.InternalErrorAbort(c, err.Error())
		return
	}

	httputils.RespOK(c)
}
