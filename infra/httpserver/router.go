package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/aws-cognito-example/infra/cognitolib"
)

func newRouter(cognitoApp *cognitolib.CognitoApp) *gin.Engine {
	router := gin.Default()

	ctrl := Controller{
		cognitoApp: cognitoApp,
	}

	router.POST("/signup", ctrl.Signup)

	return router
}
