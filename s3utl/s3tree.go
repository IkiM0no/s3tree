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
func FetchNodes(svc *s3.S3, bucket, folder string) TreeNodes {
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(folder),
	}
	resp, _ := svc.ListObjectsV2(params)
	var Nodes TreeNodes
	for _, key := range resp.Contents {
		var Node = TreeNode{
			NodeName:     *key.Key,
			IsFolder:     strings.HasSuffix(*key.Key, "/"),
			Size:         *key.Size,
			LastModified: *key.LastModified,
		}
		Nodes = append(Nodes, Node)
	}
	return Nodes
}

// build and print the tree
func (nodes TreeNodes) IterTree(showFileAttrs bool) {
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
		} else {
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
