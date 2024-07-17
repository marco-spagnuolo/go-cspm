package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetEC2Client restituisce un client EC2
func GetEC2Client(sess *session.Session) *ec2.EC2 {
	return ec2.New(sess)
}

// GetS3Client restituisce un client S3
func GetS3Client(sess *session.Session) *s3.S3 {
	return s3.New(sess)
}

// GetRDSClient restituisce un client RDS
func GetRDSClient(sess *session.Session) *rds.RDS {
	return rds.New(sess)
}

// GetIAMClient restituisce un client IAM
func GetIAMClient(sess *session.Session) *iam.IAM {
	return iam.New(sess)
}
