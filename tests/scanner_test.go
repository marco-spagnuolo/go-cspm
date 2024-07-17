package tests

import (
	"cspm-go/internal/scanner"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestScanEC2Instances(t *testing.T) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
	ec2Svc := ec2.New(sess)
	instances, err := scanner.ScanEC2Instances(ec2Svc)
	if err != nil {
		t.Fatalf("Failed to scan EC2 instances: %v", err)
	}
	if len(instances) == 0 {
		t.Fatalf("No EC2 instances found")
	}
}

func TestScanS3Buckets(t *testing.T) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
	s3Svc := s3.New(sess)
	buckets, err := scanner.ScanS3Buckets(s3Svc)
	if err != nil {
		t.Fatalf("Failed to scan S3 buckets: %v", err)
	}
	if len(buckets) == 0 {
		t.Fatalf("No S3 buckets found")
	}
}

func TestScanRDSInstances(t *testing.T) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
	rdsSvc := rds.New(sess)
	instances, err := scanner.ScanRDSInstances(rdsSvc)
	if err != nil {
		t.Fatalf("Failed to scan RDS instances: %v", err)
	}
	if len(instances) == 0 {
		t.Fatalf("No RDS instances found")
	}
}

func TestScanIAMUsers(t *testing.T) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
	iamSvc := iam.New(sess)
	users, err := scanner.ScanIAMUsers(iamSvc)
	if err != nil {
		t.Fatalf("Failed to scan IAM users: %v", err)
	}
	if len(users) == 0 {
		t.Fatalf("No IAM users found")
	}
}
