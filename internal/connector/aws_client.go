package connector

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var c = AWSClient{}

type AWSClient struct {
	LambdaService *lambda.Lambda
	ASGService    *autoscaling.AutoScaling
}

func getSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func GetLambdaService() *lambda.Lambda {
	if c.LambdaService == nil {
		c.LambdaService = lambda.New(getSession())
	}
	return c.LambdaService
}

func GetASGService() *autoscaling.AutoScaling {
	if c.ASGService == nil {
		c.ASGService = autoscaling.New(getSession())
	}
	return c.ASGService
}
