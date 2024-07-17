package scanner

import (
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
)

func ScanRDSInstances(svc rdsiface.RDSAPI) ([]*rds.DBInstance, error) {
	result, err := svc.DescribeDBInstances(&rds.DescribeDBInstancesInput{})
	if err != nil {
		return nil, err
	}
	return result.DBInstances, nil
}
