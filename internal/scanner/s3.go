package scanner

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

func ScanS3Buckets(svc s3iface.S3API) ([]*s3.Bucket, error) {
	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return result.Buckets, nil
}
