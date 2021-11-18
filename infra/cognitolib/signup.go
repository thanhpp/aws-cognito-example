package cognitolib

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func (app CognitoApp) SignUp(email, password, phonenumber, gender, name string, birthdate time.Time) error {
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
			{
				Name:  aws.String("birthdate"),
				Value: aws.String(birthdate.Format("2006-01-02")),
			},
		},
	}

	_, err := app.client.SignUp(newUser)
	if err != nil {
		return err
	}

	return nil
}
