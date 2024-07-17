package tests

import (
	"cspm-go/internal/mocks"
	"cspm-go/internal/scanner"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/mock/gomock"
)

func TestScanEC2InstancesWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEC2 := mocks.NewMockEC2API(ctrl)
	mockEC2.EXPECT().DescribeInstances(gomock.Any()).Return(&ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{
			{
				Instances: []*ec2.Instance{
					{InstanceId: aws.String("i-1234567890abcdef0")},
				},
			},
		},
	}, nil)

	instances, err := scanner.ScanEC2Instances(mockEC2)
	if err != nil {
		t.Fatalf("Failed to scan EC2 instances: %v", err)
	}
	if len(instances) == 0 {
		t.Fatalf("No EC2 instances found")
	}
}

func TestScanS3BucketsWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockS3 := mocks.NewMockS3API(ctrl)
	mockS3.EXPECT().ListBuckets(gomock.Any()).Return(&s3.ListBucketsOutput{
		Buckets: []*s3.Bucket{
			{Name: aws.String("test-bucket")},
		},
	}, nil)

	buckets, err := scanner.ScanS3Buckets(mockS3)
	if err != nil {
		t.Fatalf("Failed to scan S3 buckets: %v", err)
	}
	if len(buckets) == 0 {
		t.Fatalf("No S3 buckets found")
	}
}

func TestScanRDSInstancesWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRDS := mocks.NewMockRDSAPI(ctrl)
	mockRDS.EXPECT().DescribeDBInstances(gomock.Any()).Return(&rds.DescribeDBInstancesOutput{
		DBInstances: []*rds.DBInstance{
			{DBInstanceIdentifier: aws.String("db-1234567890abcdef0")},
		},
	}, nil)

	instances, err := scanner.ScanRDSInstances(mockRDS)
	if err != nil {
		t.Fatalf("Failed to scan RDS instances: %v", err)
	}
	if len(instances) == 0 {
		t.Fatalf("No RDS instances found")
	}
}

func TestScanIAMUsersWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIAM := mocks.NewMockIAMAPI(ctrl)
	mockIAM.EXPECT().ListUsers(gomock.Any()).Return(&iam.ListUsersOutput{
		Users: []*iam.User{
			{UserName: aws.String("test-user")},
		},
	}, nil)

	users, err := scanner.ScanIAMUsers(mockIAM)
	if err != nil {
		t.Fatalf("Failed to scan IAM users: %v", err)
	}
	if len(users) == 0 {
		t.Fatalf("No IAM users found")
	}
}
