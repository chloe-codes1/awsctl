package lambda

import (
	"awsctl/internal/config"
	"awsctl/internal/connector"
	"fmt"
)

func GetLambdaWithVpc() []string{
	 var functions []string

	 for _, function := range connector.ListFunctionsPages(){
	 	functionName := *function.FunctionName
	 	funcInfo := connector.GetFunction(functionName)
	 	vpcConfig := funcInfo.Configuration.VpcConfig
	 	if vpcConfig != nil && *vpcConfig.VpcId == config.VPCConf.VpcID  {
	 		for _, subnetID := range vpcConfig.SubnetIds{
	 			switch *subnetID {
				case
					config.SubnetConf.PrivateSubnet.AZone,
					config.SubnetConf.PrivateSubnet.BZone,
					config.SubnetConf.PrivateSubnet.CZone,
					config.SubnetConf.PrivateSubnet.DZone:

					functions = append (functions, functionName)
					fmt.Println(functionName)
	 			}
			}
		}
	 }
	 fmt.Println("Number of Lambdas in Private Subnet?", len(functions))
	 return functions
}

func ModifyLambdaVpcConfig(){
	functions := GetLambdaWithVpc()
	for _, function := range functions{
		connector.UpdateFunctionConfiguration(function)
	}
}
