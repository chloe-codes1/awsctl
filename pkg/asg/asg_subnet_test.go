package asg

import (
	"github.com/stretchr/testify/suite"
	"lambda-manager/internal/config"
	"strings"
	"testing"
)

type ASGSubnetSuite struct {
	suite.Suite
}

func TestASGSubnet(t *testing.T){
	asgSubnetSuite := new(ASGSubnetSuite)
	suite.Run(t, asgSubnetSuite)
}

func (suite *ASGSubnetSuite) SetupSuite(){
	configFile := strings.Join([]string{"internal/config/config.devel.yml"}, "")

	conf, err := config.ReadConfig(configFile)
	if err != nil {
		panic(err)
	}
	config.VPCConf = &conf.VPC
	config.SubnetConf = &conf.Subnet
}

func (suite *ASGSubnetSuite) TestListAsgInProductPrivateSubnet() {
	ListAsgInSpecificSubnet()
}
