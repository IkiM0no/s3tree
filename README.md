# s3tree

List contents of s3 buckets in a tree-like format.

s3tree is a Go implementation of Dem Pilafian's wonderful tree.sh for use with Amazon Web Services Simple Storage Service (S3).

Usage example:
```
$ s3tree -c default -b my-bucket -f my-folder -s
```
```
|    |__tree
|    |    |__file1.txt
|    |    |__file2.txt
|    |    |__folder1
|    |    |    |__file1.txt
|    |    |    |__nestedfolder1
|    |    |    |__nestedfolder2
|    |    |    |    |__file2.txt
|    |    |    |    |__file3.txt
|    |    |__folder2
|    |    |    |__file1.txt
|    |    |    |__nestedfolder1
|    |    |    |    |__deepnestedfolder1
|    |    |    |    |    |__file1.txt
|    |    |    |    |__file3.txt
|    |    |__folder3
|    |    |    |__nestedfolder1
|    |    |    |    |__file2.txt
|    |    |    |__nestedfolder2
|    |    |    |__nestedfolder3
|    |    |    |    |__deepnestedfolder1
|    |    |    |    |    |__file1.txt
|    |    |    |    |__deepnestedfolder2
|    |    |    |    |    |__file2.txt
|    |    |    |    |    |__file3.txt
|    |    |    |    |__file1.txt
|    |    |    |    |__file2.txt
```

Get help:
```
s3tree -h
```

```
NAME:
   s3tree - List contents of s3 buckets in a tree-like format.

USAGE:
   s3tree [global options] command [command options] [arguments...]

VERSION:
   0.5.2

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -c value       -c </.aws/credentials [class]> (default: "default")
   -b value       -b <bucket>
   -f value       -f <folder>
   -s             -s | Include file size/date in output
   --help, -h     show help
   --version, -v  print the version
```
