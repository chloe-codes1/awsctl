/*
Copyright © 2021 chloe-codes1 <juhyun.kim@lindsey.edu>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"awsctl/internal/config"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	env     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awsctl",
	Short: "Manage your AWS resources easily and efficiently",
	Long: `
Features:

    Lambda
    - List Functions
      - Returns only the first 50 lambdas
    - List Functions Pages
      - Return all lambdas
    - Get Function
      - Check if the function is in 
    - Update Function Configuration


    Auto Scaling Groups
    - List Auto Scaling Groups
      - Return all ASGs
    - List ASG in specific subnet`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aws-manager.yaml)")
	rootCmd.PersistentFlags().StringVar(&env, "env", "", "set environment (default is devel)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".aws-manager" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".aws-manager")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// Set default environment
	if env == "" {
		env = "devel"
	}

	// Read config file
	configFile := strings.Join([]string{"internal/config/config.", env, ".yml"}, "")

	conf, err := config.ReadConfig(configFile)
	if err != nil {
		panic(err)
	}
	config.VPCConf = &conf.VPC
	config.SubnetConf = &conf.Subnet
}
