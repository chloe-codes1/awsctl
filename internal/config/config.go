package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	VPCConf *VPCConfig
	SubnetConf *SubnetConfig
)

type AppConfig struct {
	VPC       VPCConfig `yaml:"vpc"`
	Subnet 	SubnetConfig 	`yaml:"subnet"`
}

type VPCConfig struct{
	VpcID string `yaml:"VpcID"`
}

type SubnetConfig struct {
	PrivateSubnet PrivateSubnetConfig `yaml:"privateSubnet"`
	ServerlessSubnet ServerlessSubnetConfig `yaml:"serverlessSubnet"`
}

type PrivateSubnetConfig struct {
	AZone string `yaml:"AZone"`
	BZone string `yaml:"BZone"`
	CZone string `yaml:"CZone"`
	DZone string `yaml:"DZone"`
}

type ServerlessSubnetConfig struct {
	AZone string `yaml:"AZone"`
	BZone string `yaml:"BZone"`
	CZone string `yaml:"CZone"`
	DZone string `yaml:"DZone"`
}

func ReadConfig(fileName string) (*AppConfig, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "file read error")
	}
	return readConfigBytes(file)
}

func readConfigBytes(data []byte) (*AppConfig, error) {
	var appConfig AppConfig
	err := yaml.Unmarshal(data, &appConfig)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal error")
	}
	return &appConfig, nil
}