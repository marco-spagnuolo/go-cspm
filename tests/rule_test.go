package tests

import (
	"cspm-go/internal/mocks"
	"cspm-go/internal/rules"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/mock/gomock"
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEC2 := mocks.NewMockEC2API(ctrl)
	instance := &ec2.Instance{
		BlockDeviceMappings: []*ec2.InstanceBlockDeviceMapping{
			{
				Ebs: &ec2.EbsInstanceBlockDevice{
					// This line is removed since Encrypted is not a field in EbsInstanceBlockDevice
				},
			},
		},
	}
	if rules.CheckEBSVolumeEncryption(instance, mockEC2) {
		t.Fatalf("Expected instance to fail EBS volume encryption check")
	}
}

func TestCheckS3BucketEncryption(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockS3 := mocks.NewMockS3API(ctrl)
	bucketName := "test-bucket"

	mockS3.EXPECT().GetBucketEncryption(&s3.GetBucketEncryptionInput{
		Bucket: aws.String(bucketName),
	}).Return(&s3.GetBucketEncryptionOutput{}, nil).AnyTimes()

	if rules.CheckS3BucketEncryption(bucketName, mockS3) {
		t.Fatalf("Expected bucket to fail encryption check")
	}
}

func TestCheckS3BucketPublicAccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockS3 := mocks.NewMockS3API(ctrl)
	bucketName := "test-bucket"

	mockS3.EXPECT().GetBucketPolicyStatus(&s3.GetBucketPolicyStatusInput{
		Bucket: aws.String(bucketName),
	}).Return(&s3.GetBucketPolicyStatusOutput{}, nil).AnyTimes()

	if rules.CheckS3BucketPublicAccess(bucketName, mockS3) {
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIAM := mocks.NewMockIAMAPI(ctrl)
	user := &iam.User{
		UserName: aws.String("test-user"),
	}

	mockIAM.EXPECT().ListMFADevices(&iam.ListMFADevicesInput{
		UserName: user.UserName,
	}).Return(&iam.ListMFADevicesOutput{}, nil).AnyTimes()

	if rules.CheckIAMUserMFA(user, mockIAM) {
		t.Fatalf("Expected user to fail MFA check")
	}
}
