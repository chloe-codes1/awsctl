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
	"awsctl/pkg/lambda"

	"github.com/spf13/cobra"
)

// getLambdaWithVpcCmd represents the getLambdaWithVpc command.
var getLambdaWithVpcCmd = &cobra.Command{
	Use:   "get-lambda-with-vpc",
	Short: "Get Lambda functions with VPC configured",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		lambda.GetLambdaWithVpc()
	},
}

func init() {
	rootCmd.AddCommand(getLambdaWithVpcCmd)
}
