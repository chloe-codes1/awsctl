package main

import (
	"lambda-manager/internal/config"
	"lambda-manager/pkg/asg"
	"os"
	"strings"
)

func init() {
	if len(os.Args) < 2 {
		os.Exit(-1)
	}

	configFile := strings.Join([]string{"internal/config/config.", os.Args[1], ".yml"}, "")

	conf, err := config.ReadConfig(configFile)
	if err != nil {
		panic(err)
	}
	config.VPCConf = &conf.VPC
	config.SubnetConf = &conf.Subnet
}

func main() {
	//lambda.GetLambdaWithVPC()
	//lambda.ModifyLambdaVpcConfig()
	asg.ListAsgInSpecificSubnet()
}