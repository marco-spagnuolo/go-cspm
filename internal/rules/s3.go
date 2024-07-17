package rules

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Verifica che i bucket S3 abbiano la crittografia abilitata
func CheckS3BucketEncryption(bucketName string, svc *s3.S3) bool {
	input := &s3.GetBucketEncryptionInput{
		Bucket: aws.String(bucketName),
	}
	_, err := svc.GetBucketEncryption(input)
	return err == nil
}

// Verifica che i bucket S3 non siano pubblicamente accessibili
func CheckS3BucketPublicAccess(bucketName string, svc *s3.S3) bool {
	input := &s3.GetBucketAclInput{
		Bucket: aws.String(bucketName),
	}
	result, err := svc.GetBucketAcl(input)
	if err != nil {
		return false
	}
	for _, grant := range result.Grants {
		if *grant.Grantee.URI == "http://acs.amazonaws.com/groups/global/AllUsers" && *grant.Permission == "READ" {
			return false
		}
	}
	return true
}
