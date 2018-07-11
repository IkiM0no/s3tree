package main

import (
	"os"
	"s3tree/s3utl"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/urfave/cli"
)

var (
	gClass         string
	gBucket        string
	gFolder        string
	gShowFileAttrs bool
	gS3Credentials *credentials.Credentials
	gS3Client      *s3.S3
)

func main() {
	app := buildApp()
	app.Run(os.Args)
}

func buildApp() *cli.App {
	app := cli.NewApp()
	app.Name = "s3tree"
	app.Version = "1.1.0"
	app.Usage = "list contents of s3 buckets in a tree-like format."
	app.Action = func(c *cli.Context) error {
		if gClass == "" {
			gClass = "default"
		}
		s3LocalCredentials := s3utl.S3LocalCreds{Class: gClass}
		gS3Credentials = s3LocalCredentials.Set()

		s3c := s3utl.S3Client{Credentials: gS3Credentials}
		gS3Client = s3c.Fetch()

		nodes := s3utl.FetchNodes(gS3Client, gBucket, gFolder)
		nodes.IterTree(gShowFileAttrs)
		return nil
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "c",
			Value:       "default",
			Usage:       "-c </.aws/credentials [class]>",
			Destination: &gClass,
		},
		cli.StringFlag{
			Name:        "b",
			Usage:       "-b <bucket>",
			Destination: &gBucket,
		},
		cli.StringFlag{
			Name:        "f",
			Usage:       "-f <folder>",
			Destination: &gFolder,
		},
		cli.BoolFlag{
			Name:        "s",
			Usage:       "-s | include file size/date in output",
			Destination: &gShowFileAttrs,
		},
	}
	return app
}
