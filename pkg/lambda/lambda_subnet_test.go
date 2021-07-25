package lambda

import (
	"awsctl/internal/config"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type LambdaSubnetSuite struct {
	suite.Suite
}

func TestLambdaSubnet(t *testing.T){
	lambdaSubnetSuite:= new(LambdaSubnetSuite)
	suite.Run(t, lambdaSubnetSuite)
}

func (suite *LambdaSubnetSuite) SetupSuite() {
	configFile := strings.Join([]string{"internal/config/config.devel.yml"}, "")

	conf, err := config.ReadConfig(configFile)
	if err != nil {
		panic(err)
	}
	config.VPCConf = &conf.VPC
	config.SubnetConf = &conf.Subnet
}

func (suite *LambdaSubnetSuite) TestGetLambdaWithVPC() {
	GetLambdaWithVpc()
}
