package scanner

import (
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

func ScanIAMUsers(svc iamiface.IAMAPI) ([]*iam.User, error) {
	result, err := svc.ListUsers(&iam.ListUsersInput{})
	if err != nil {
		return nil, err
	}
	return result.Users, nil
}
