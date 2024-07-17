package automation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Applicazione delle correzioni per le istanze EC2
func ApplyEC2SecurityGroupFix(sess *session.Session, instanceID string) error {
	svc := ec2.New(sess)
	_, err := svc.ModifyInstanceAttribute(&ec2.ModifyInstanceAttributeInput{
		InstanceId: aws.String(instanceID),
		Groups:     []*string{aws.String("sg-12345678")}, // Assicurarsi di usare un Security Group configurato correttamente
	})
	return err
}
