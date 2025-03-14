package aws

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
)

func ListObjects(client *s3.Client, ctx context.Context, bucketName string) ([]types.Object, error) {
	var err error
	var output *s3.ListObjectsV2Output
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}
	var objects []types.Object
	objectPaginator := s3.NewListObjectsV2Paginator(client, input)
	for objectPaginator.HasMorePages() {
		output, err = objectPaginator.NextPage(ctx)
		if err != nil {
			var noBucket *types.NoSuchBucket
			if errors.As(err, &noBucket) {
				log.Errorf("bucket %s does not exist.\n", bucketName)
				err = noBucket
			}
			break
		} else {
			objects = append(objects, output.Contents...)
		}
	}
	return objects, err
}

func BucketAcl(client *s3.Client, ctx context.Context, bucketName string) (*s3.GetBucketAclOutput, error) {
	input := &s3.GetBucketAclInput{
		Bucket: aws.String(bucketName),
	}
	result, err := client.GetBucketAcl(ctx, input)
	if err != nil {
		log.Debugf("unable to get ACL for bucket %s: %s", bucketName, err)
	}
	return result, err

}
