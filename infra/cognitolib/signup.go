package cognitolib

import (
	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func (app CognitoApp) SignUp(email, password, phonenumber, gender, name string) error {
	newUser := &cognito.SignUpInput{
		Username: aws.String(email),
		Password: aws.String(password),
		ClientId: aws.String(app.appID),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("gender"),
				Value: aws.String(gender),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(phonenumber),
			},
			{
				Name:  aws.String("name"),
				Value: aws.String(name),
			},
		},
	}

	_, err := app.client.SignUp(newUser)
	if err != nil {
		return err
	}

	return nil
}
