package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Storage defines the interface for file storage operations
type Storage interface {
	UploadFile(ctx context.Context, bucket, key string, file io.Reader, contentType string) (string, error)
}

type s3Storage struct {
	client   *s3.Client
	region   string
	endpoint string
}

// NewS3Storage creates a new S3-compatible storage client
func NewS3Storage(region, endpoint, accessKey, secretKey string, usePathStyle bool) (Storage, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if endpoint != "" {
			return aws.Endpoint{
				URL:           endpoint,
				SigningRegion: region,
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = usePathStyle
	})

	return &s3Storage{
		client:   client,
		region:   region,
		endpoint: endpoint,
	}, nil
}

func (s *s3Storage) UploadFile(ctx context.Context, bucket, key string, file io.Reader, contentType string) (string, error) {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", err
	}

	// For S3, the URL is usually https://bucket.s3.region.amazonaws.com/key
	// For MinIO, it depends on the endpoint.
	// We'll return a simple path or a full URL if endpoint is provided.
	if s.endpoint != "" {
		return s.endpoint + "/" + bucket + "/" + key, nil
	}

	return "https://" + bucket + ".s3." + s.region + ".amazonaws.com/" + key, nil
}
