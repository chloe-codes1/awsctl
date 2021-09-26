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
	"awsctl/pkg/asg"

	"github.com/spf13/cobra"
)

// listAsgInSpecificSubnetCmd represents the listAsgInSpecificSubnet command.
var listAsgInSpecificSubnetCmd = &cobra.Command{
	Use:   "list-asg-in-specific-subnet",
	Short: "List Auto Scaling Group in specific subnet",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		asg.ListAsgInSpecificSubnet()
	},
}

func init() {
	rootCmd.AddCommand(listAsgInSpecificSubnetCmd)
}
