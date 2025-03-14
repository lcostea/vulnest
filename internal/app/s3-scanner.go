package app

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	awsv "github.com/lcostea/vulnest/internal/aws"
)

func ScanBucket(bucket string, extendedOutput bool) {
	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(aws.AnonymousCredentials{}))
	if err != nil {
		log.Fatalf("unable to load AWS SDK config: %v", err)
	}
	s3Client := s3.NewFromConfig(sdkConfig)

	log.Infof("BUCKET: %s\n", bucket)

	// aws s3 ls s3://{BUCKET_NAME} --no-sign-request
	log.Info("testing list permissions on bucket")
	objects, err := awsv.ListObjects(s3Client, ctx, bucket)
	if err != nil {
		log.Errorf("unable to list objects in bucket %s: %s", bucket, err)
	} else {
		log.Infof("!!! S3 bucket is publicly accessible, found %d objects\n", len(objects))
	}

	// aws s3api get-bucket-acl --bucket {BUCKET_NAME} --no-sign-request
	log.Info("testing bucket ACL")
	res, err := awsv.BucketAcl(s3Client, ctx, bucket)
	if err != nil {
		log.Errorf("unable to get ACL for bucket %s: %s", bucket, err)
	} else {
		log.Info("!!! S3 bucket has public ACL\n")
		// Print each grant only on extended output
		if extendedOutput {
			printACLGrants(res)
		}
	}

}

func printACLGrants(res *s3.GetBucketAclOutput) {
	if res.Owner.DisplayName != nil {
		log.Infof("Owner: %s\n", *res.Owner.DisplayName)
	} else {
		log.Infof("Owner: %s\n", *res.Owner.ID)
	}
	log.Info("\nBUCKET Grants:\n")
	for i, grant := range res.Grants {
		log.Infof("Grant %d:\n", i+1)

		// Print grantee information
		if grant.Grantee.DisplayName != nil {
			log.Infof("  Grantee: %s\n", *grant.Grantee.DisplayName)
		}
		if grant.Grantee.ID != nil {
			log.Infof("  Grantee ID: %s\n", *grant.Grantee.ID)
		}
		if grant.Grantee.URI != nil {
			log.Infof("  Grantee URI: %s\n", *grant.Grantee.URI)
		}
		log.Infof("  Permission: %s\n", grant.Permission)
	}
}
