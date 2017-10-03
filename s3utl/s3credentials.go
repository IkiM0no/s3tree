package s3utl

import (
	"os"
	"fmt"
	"reflect"

	"github.com/mitchellh/go-homedir"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type S3LocalCreds struct {
	HomeDir       string
	AwsDir        string `default:"/.aws/credentials"`
	Class         string `default:"default"`
}

// Return an AWS credentials object from local ~/.aws/credentials
func (s S3LocalCreds) Set() (*credentials.Credentials) {
	typ := reflect.TypeOf(s)
	if s.Class == "" {
		z, _ := typ.FieldByName("Class")
		s.Class = z.Tag.Get("default")
	}
	if s.AwsDir == "" {
		z, _ := typ.FieldByName("AwsDir")
		s.AwsDir = z.Tag.Get("default")
	}
	if s.HomeDir == "" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println("Err retrieving HomeDir and none provided:", err)
			os.Exit(1)
		}
		s.HomeDir = home
	}
	credsPath := fmt.Sprintf("%s%s", s.HomeDir, s.AwsDir)
	creds := credentials.NewSharedCredentials(credsPath, s.Class)
	return creds
}
