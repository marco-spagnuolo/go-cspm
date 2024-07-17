package scanner

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

func ScanEC2Instances(svc ec2iface.EC2API) ([]*ec2.Instance, error) {
	result, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{})
	if err != nil {
		return nil, err
	}
	var instances []*ec2.Instance
	for _, reservation := range result.Reservations {
		instances = append(instances, reservation.Instances...)
	}
	return instances, nil
}
