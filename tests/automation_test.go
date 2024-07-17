package tests

import (
	"cspm-go/internal/automation"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
)

func TestApplyEC2SecurityGroupFix(t *testing.T) {
	sess := session.Must(session.NewSession())
	err := automation.ApplyEC2SecurityGroupFix(sess, "i-1234567890abcdef0")
	if err != nil {
		t.Fatalf("Failed to apply EC2 security group fix: %v", err)
	}
}

func TestApplyS3BucketPolicyFix(t *testing.T) {
	sess := session.Must(session.NewSession())
	err := automation.ApplyS3BucketPolicyFix(sess, "test-bucket")
	if err != nil {
		t.Fatalf("Failed to apply S3 bucket policy fix: %v", err)
	}
}

func TestApplyRDSInstanceFix(t *testing.T) {
	sess := session.Must(session.NewSession())
	err := automation.ApplyRDSInstanceFix(sess, "db-1234567890abcdef0")
	if err != nil {
		t.Fatalf("Failed to apply RDS instance fix: %v", err)
	}
}

func TestApplyIAMUserPolicyFix(t *testing.T) {
	sess := session.Must(session.NewSession())
	err := automation.ApplyIAMUserPolicyFix(sess, "test-user")
	if err != nil {
		t.Fatalf("Failed to apply IAM user policy fix: %v", err)
	}
}
