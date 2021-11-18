package cognitolib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoApp struct {
	client    *cognito.CognitoIdentityProvider
	region    string
	poolID    string
	appID     string
	appSecret string
}

func SetupCognito(region, poolID, appID string) (*CognitoApp, error) {
	// new session
	newSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	return &CognitoApp{
		client: cognito.New(newSession),
		region: region,
		poolID: poolID,
		appID:  appID,
	}, nil
}
