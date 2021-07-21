package asg

import (
	"awsctl/internal/config"
	"awsctl/internal/connector"
	"fmt"
	"strings"
)

func ListAsgInSpecificSubnet(){
	var asgNames []string
	for _, asg := range connector.DescribeAutoScalingGroups(){
		for _, subnet := range strings.Split(*asg.VPCZoneIdentifier, ","){

			switch subnet {
			case
				config.SubnetConf.PrivateSubnet.AZone,
				config.SubnetConf.PrivateSubnet.BZone,
				config.SubnetConf.PrivateSubnet.CZone,
				config.SubnetConf.PrivateSubnet.DZone:

				asgName := *asg.AutoScalingGroupName
				var lbName string
				if len(asg.TargetGroupARNs) > 0 {
					lbName = *asg.TargetGroupARNs[0]
				}
				asgNames = append(asgNames, asgName)
				fmt.Println("ASG: ", asgName, " LB: ", lbName)
				break
			}
		}
	}
	fmt.Println("Total:", len(asgNames))
}
