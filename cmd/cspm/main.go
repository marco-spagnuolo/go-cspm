package main

import (
	"cspm-go/internal/rules"
	"cspm-go/internal/scanner"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess := session.Must(session.NewSession())

	// Create AWS service clients
	ec2Svc := ec2.New(sess)
	s3Svc := s3.New(sess)
	rdsSvc := rds.New(sess)
	iamSvc := iam.New(sess)

	// Scansione delle risorse EC2
	ec2Instances, err := scanner.ScanEC2Instances(ec2Svc)
	if err != nil {
		log.Fatalf("Failed to scan EC2 instances: %v", err)
	}
	for _, instance := range ec2Instances {
		if !rules.CheckInstanceSecurityGroups(instance) {
			log.Printf("EC2 Instance ID: %s has default security group", *instance.InstanceId)
		}
		if !rules.CheckEBSVolumeEncryption(instance, ec2Svc) {
			log.Printf("EC2 Instance ID: %s has unencrypted EBS volumes", *instance.InstanceId)
		}
	}

	// Scansione dei bucket S3
	s3Buckets, err := scanner.ScanS3Buckets(s3Svc)
	if err != nil {
		log.Fatalf("Failed to scan S3 buckets: %v", err)
	}
	for _, bucket := range s3Buckets {
		if !rules.CheckS3BucketEncryption(*bucket.Name, s3Svc) {
			log.Printf("S3 Bucket Name: %s does not have encryption enabled", *bucket.Name)
		}
		if !rules.CheckS3BucketPublicAccess(*bucket.Name, s3Svc) {
			log.Printf("S3 Bucket Name: %s is publicly accessible", *bucket.Name)
		}
	}

	// Scansione delle istanze RDS
	rdsInstances, err := scanner.ScanRDSInstances(rdsSvc)
	if err != nil {
		log.Fatalf("Failed to scan RDS instances: %v", err)
	}
	for _, dbInstance := range rdsInstances {
		if !rules.CheckRDSInstanceEncryption(dbInstance) {
			log.Printf("RDS Instance ID: %s does not have encryption enabled", *dbInstance.DBInstanceIdentifier)
		}
		if !rules.CheckRDSInstanceBackup(dbInstance) {
			log.Printf("RDS Instance ID: %s does not have automatic backups configured", *dbInstance.DBInstanceIdentifier)
		}
	}

	// Scansione degli utenti IAM
	iamUsers, err := scanner.ScanIAMUsers(iamSvc)
	if err != nil {
		log.Fatalf("Failed to scan IAM users: %v", err)
	}
	for _, user := range iamUsers {
		if !rules.CheckIAMUserMFA(user, iamSvc) {
			log.Printf("IAM User Name: %s does not have MFA enabled", *user.UserName)
		}
	}
}
