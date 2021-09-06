package connector

import (
	"awsctl/internal/config"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// Returns only the first 50 lambdas
func ListFunctions() lambda.ListFunctionsOutput {
	svc := GetLambdaService()
	result, err := svc.ListFunctions(&lambda.ListFunctionsInput{
		MaxItems: aws.Int64(123),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lambda.ErrCodeServiceException:
				fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
			case lambda.ErrCodeTooManyRequestsException:
				fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
			case lambda.ErrCodeInvalidParameterValueException:
				fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	return *result
}

// Return all lambdas
func ListFunctionsPages() []*lambda.FunctionConfiguration {
	svc := GetLambdaService()
	var result []*lambda.FunctionConfiguration
	err := svc.ListFunctionsPages(&lambda.ListFunctionsInput{},
		func(page *lambda.ListFunctionsOutput, lastPage bool) bool {
			result = append(result, page.Functions...)
			return !lastPage
		})
	if err != nil {
		return nil
	}
	return result
}

func GetFunction(funcName string) lambda.GetFunctionOutput {
	svc := GetLambdaService()
	result, err := svc.GetFunction(&lambda.GetFunctionInput{
		FunctionName: aws.String(funcName),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lambda.ErrCodeServiceException:
				fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
			case lambda.ErrCodeResourceNotFoundException:
				fmt.Println(lambda.ErrCodeResourceNotFoundException, aerr.Error())
			case lambda.ErrCodeTooManyRequestsException:
				fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
			case lambda.ErrCodeInvalidParameterValueException:
				fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {

			fmt.Println(err.Error())
		}
	}
	return *result
}

func UpdateFunctionConfiguration(funcName string) *lambda.FunctionConfiguration {
	svc := GetLambdaService()
	input := &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String(funcName),
		VpcConfig: &lambda.VpcConfig{
			SubnetIds: []*string{
				aws.String(config.SubnetConf.ServerlessSubnet.AZone),
				aws.String(config.SubnetConf.ServerlessSubnet.CZone),
				aws.String(config.SubnetConf.ServerlessSubnet.BZone),
				aws.String(config.SubnetConf.ServerlessSubnet.DZone),
			},
		},
	}

	result, err := svc.UpdateFunctionConfiguration(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lambda.ErrCodeServiceException:
				fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
			case lambda.ErrCodeResourceNotFoundException:
				fmt.Println(lambda.ErrCodeResourceNotFoundException, aerr.Error())
			case lambda.ErrCodeInvalidParameterValueException:
				fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
			case lambda.ErrCodeTooManyRequestsException:
				fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
			case lambda.ErrCodeResourceConflictException:
				fmt.Println(lambda.ErrCodeResourceConflictException, aerr.Error())
			case lambda.ErrCodePreconditionFailedException:
				fmt.Println(lambda.ErrCodePreconditionFailedException, aerr.Error())
			case lambda.ErrCodeCodeVerificationFailedException:
				fmt.Println(lambda.ErrCodeCodeVerificationFailedException, aerr.Error())
			case lambda.ErrCodeInvalidCodeSignatureException:
				fmt.Println(lambda.ErrCodeInvalidCodeSignatureException, aerr.Error())
			case lambda.ErrCodeCodeSigningConfigNotFoundException:
				fmt.Println(lambda.ErrCodeCodeSigningConfigNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	return result
}
