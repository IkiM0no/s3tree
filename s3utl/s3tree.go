package s3utl

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	. "github.com/logrusorgru/aurora"
)

var reFileNoPath = regexp.MustCompile(`.*/`)
var rePathNoFile = regexp.MustCompile(`/[^/]*$`)

type TreeNodes []TreeNode

type TreeNode struct {
	NodeName     string
	IsFolder     bool
	Size         int64
	LastModified time.Time
}

// fetch slice of Tree Nodes - convenience wrapper for s3 response.Contents
func FetchNodes(svc *s3.S3, bucket, folder string) (TreeNodes, error) {
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		//Prefix: aws.String(folder),
		StartAfter: aws.String(folder),
	}
	resp, err := svc.ListObjectsV2(params)
	if err != nil {
		return nil, err
	}
	var nodes TreeNodes
	for _, key := range resp.Contents {
		var Node = TreeNode{
			NodeName:     *key.Key,
			IsFolder:     strings.HasSuffix(*key.Key, "/"),
			Size:         *key.Size,
			LastModified: *key.LastModified,
		}
		nodes = append(nodes, Node)
	}
	return nodes, nil
}

// build and print the tree
func (nodes TreeNodes) IterTree(showFileAttrs, dirOnly bool) {
	const (
		indentChar  = "|    "
		prefixChar  = "%s|__%s"
		verboseForm = "%s [%d] [%s]"
		timeFmt     = "2006-01-02 15:04:05"
	)
	var (
		nFolders = 0
		nFiles   = 0
	)
	for _, node := range nodes {
		var nodeStr string
		nDeep := strings.Count(node.NodeName, "/")
		if node.IsFolder {
			nFolders += 1
			nodeStr = fmt.Sprintf(prefixChar,
				strings.Repeat(indentChar, nDeep-1),
				Bold(Green(lastNode(node.NodeName))),
			)
			fmt.Println(nodeStr)
		} else if !dirOnly {
			nFiles += 1
			nodeStr = fmt.Sprintf(prefixChar,
				strings.Repeat(indentChar, nDeep),
				stripPath(node.NodeName),
			)
			if showFileAttrs {
				nodeStr = fmt.Sprintf(verboseForm,
					nodeStr,
					node.Size,
					node.LastModified.Format(timeFmt),
				)
			}
			fmt.Println(nodeStr)
		}
	}
	fmt.Printf("\n%d directories, %d files\n", nFolders, nFiles)
}

func lastNode(inString string) string {
	return strings.Split(inString, "/")[len(strings.Split(inString, "/"))-2]
}

func stripPath(inString string) string {
	return reFileNoPath.ReplaceAllString(inString, "")
}

func stripFile(inString string) string {
	return rePathNoFile.ReplaceAllString(inString, "")
}
