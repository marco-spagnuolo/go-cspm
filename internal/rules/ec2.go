package rules

import (
	"log"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// Verifica che le istanze EC2 non utilizzino il gruppo di sicurezza predefinito
func CheckInstanceSecurityGroups(instance *ec2.Instance) bool {
	for _, sg := range instance.SecurityGroups {
		if *sg.GroupName == "default" {
			return false
		}
	}
	return true
}

// Verifica che le istanze EC2 abbiano crittografia abilitata sui volumi EBS
func CheckEBSVolumeEncryption(instance *ec2.Instance, svc *ec2.EC2) bool {
	for _, blockDevice := range instance.BlockDeviceMappings {
		if blockDevice.Ebs != nil {
			volumeID := blockDevice.Ebs.VolumeId
			input := &ec2.DescribeVolumesInput{
				VolumeIds: []*string{volumeID},
			}

			result, err := svc.DescribeVolumes(input)
			if err != nil {
				log.Printf("Unable to describe volume %s, %v", *volumeID, err)
				return false
			}

			for _, volume := range result.Volumes {
				if volume.Encrypted == nil || !*volume.Encrypted {
					return false
				}
			}
		}
	}
	return true
}
