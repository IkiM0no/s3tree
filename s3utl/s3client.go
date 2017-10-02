package s3utl

import (
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type S3Client struct {
	Region        string `default:"us-west-2"`
	Credentials   *credentials.Credentials
}

// Fetch S3 Client object
func (s3c S3Client) Fetch() (*s3.S3) {
	typ := reflect.TypeOf(s3c)
	if s3c.Region == "" {
		z, _ := typ.FieldByName("Region")
		s3c.Region = z.Tag.Get("default")
	}
	s3Client := s3.New(session.New(), &aws.Config{
	        Region:      aws.String(s3c.Region),
	        Credentials: s3c.Credentials,
	})
	return s3Client
}
