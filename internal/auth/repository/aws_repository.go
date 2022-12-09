package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"go-rest-api/internal/auth"
	"go-rest-api/internal/models"
)

type authAWSRepository struct {
	client *minio.Client
}

// PutObject Upload file to AWS
func (a authAWSRepository) PutObject(ctx context.Context, input models.UploadInput) (*minio.UploadInfo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authAWSRepository.PutObject")
	defer span.Finish()

	options := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	uploadInfo, err := a.client.PutObject(ctx, input.BucketName, a.generateFileName(input.Name), input.File, input.Size, options)
	if err != nil {
		return nil, errors.Wrap(err, "authAWSRepository.FileUpload.PutObject")
	}

	return &uploadInfo, err
}

// GetObject Download file from AWS
func (a authAWSRepository) GetObject(ctx context.Context, bucket string, fileName string) (*minio.Object, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authAWSRepository.GetObject")
	defer span.Finish()

	object, err := a.client.GetObject(ctx, bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "authAWSRepository.FileDownload.GetObject")
	}
	return object, nil
}

// RemoveObject Delete file from AWS
func (a authAWSRepository) RemoveObject(ctx context.Context, bucket string, fileName string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authAWSRepository.RemoveObject")
	defer span.Finish()

	if err := a.client.RemoveObject(ctx, bucket, fileName, minio.RemoveObjectOptions{}); err != nil {
		return errors.Wrap(err, "authAWSRepository.RemoveObject")
	}
	return nil
}

// NewAuthAWSRepository Auth AWS S3 repository constructor
func NewAuthAWSRepository(awsClient *minio.Client) auth.AWSRepository {
	return &authAWSRepository{client: awsClient}
}

func (a *authAWSRepository) generateFileName(fileName string) string {
	uid := uuid.New().String()
	return fmt.Sprintf("%s-%s", uid, fileName)
}
