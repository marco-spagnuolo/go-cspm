package tests

import (
	"cspm-go/internal/rules"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestCheckInstanceSecurityGroups(t *testing.T) {
	instance := &ec2.Instance{
		SecurityGroups: []*ec2.GroupIdentifier{
			{GroupName: aws.String("default")},
		},
	}
	if rules.CheckInstanceSecurityGroups(instance) {
		t.Fatalf("Expected instance to fail security group check")
	}
}

func TestCheckEBSVolumeEncryption(t *testing.T) {
	instance := &ec2.Instance{
		BlockDeviceMappings: []*ec2.InstanceBlockDeviceMapping{
			{Ebs: &ec2.EbsInstanceBlockDevice{VolumeId: aws.String("vol-12345678")}}, // Utilizza un campo valido
		},
	}
	if rules.CheckEBSVolumeEncryption(instance, &ec2.EC2{}) {
		t.Fatalf("Expected instance to fail EBS volume encryption check")
	}
}

func TestCheckS3BucketEncryption(t *testing.T) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
	svc := s3.New(sess)
	bucketName := "test-bucket"
	if rules.CheckS3BucketEncryption(bucketName, svc) {
		t.Fatalf("Expected bucket to fail encryption check")
	}
}

func TestCheckS3BucketPublicAccess(t *testing.T) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
	svc := s3.New(sess)
	bucketName := "test-bucket"
	if rules.CheckS3BucketPublicAccess(bucketName, svc) {
		t.Fatalf("Expected bucket to fail public access check")
	}
}

func TestCheckRDSInstanceEncryption(t *testing.T) {
	instance := &rds.DBInstance{
		StorageEncrypted: aws.Bool(false),
	}
	if rules.CheckRDSInstanceEncryption(instance) {
		t.Fatalf("Expected instance to fail encryption check")
	}
}

func TestCheckRDSInstanceBackup(t *testing.T) {
	instance := &rds.DBInstance{
		BackupRetentionPeriod: aws.Int64(0),
	}
	if rules.CheckRDSInstanceBackup(instance) {
		t.Fatalf("Expected instance to fail backup check")
	}
}

func TestCheckIAMUserMFA(t *testing.T) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
	svc := iam.New(sess)
	user := &iam.User{
		UserName: aws.String("test-user"),
	}
	if rules.CheckIAMUserMFA(user, svc) {
		t.Fatalf("Expected user to fail MFA check")
	}
}
