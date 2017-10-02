package main

import (
	"os"
	"s3tree/s3utl"

	"github.com/urfave/cli"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var class         string
var bucket        string
var folder        string
var showFileAttrs bool
var s3Credentials *credentials.Credentials
var s3Client      *s3.S3

func main() {
	app := buildApp()
	app.Run(os.Args)
}

func buildApp() *cli.App {
	app := cli.NewApp()
	app.Name    = "s3tree"
	app.Version = "0.5.2"
	app.Usage   = "List contents of s3 buckets in a tree-like format."
	app.Action  = func(c *cli.Context) error {
		if class == "" {
			class = "default"
		}
		s3LocalCredentials := s3utl.S3LocalCreds{Class: class}
		s3Credentials = s3LocalCredentials.Set()

		s3c := s3utl.S3Client{Credentials: s3Credentials}
		s3Client = s3c.Fetch()

		nodes := s3utl.FetchNodes(s3Client, bucket, folder)

		nodes.IterTree(showFileAttrs)
		return nil
	}
	app.Flags = []cli.Flag{
	        cli.StringFlag{
	                Name:  "c",
			Value: "default",
	                Usage: "-c </.aws/credentials [class]>",
	                Destination: &class,
	        },
	        cli.StringFlag{
	                Name:  "b",
	                Usage: "-b <bucket>",
	                Destination: &bucket,
	        },
	        cli.StringFlag{
	                Name:  "f",
	                Usage: "-f <folder>",
	                Destination: &folder,
	        },
	        cli.BoolFlag{
	                Name:  "s",
	                Usage: "-s | Include file size/date in output",
	                Destination: &showFileAttrs,
	        },
	}
	return app
}
