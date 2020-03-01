package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {

	access_key_id := "ZVTEVUOMQOURSFMQXHOK"
	access_secret_key := "BUVOAXq9Z75QHN8cizKH9jSePJtb0lQJ21yUQzPt"
	end_point := "https://s3.pek3b.qingstor.com"
	zone := "pek3b"
	session, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(access_key_id, access_secret_key, ""),
		Endpoint:         aws.String(end_point),
		Region:           aws.String(zone),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	})
	if err != nil {
		panic(err)
	}
	svc := s3.New(session)

	// ListBuckets(svc)
	// CreateBucket(svc, "test-bucket")
	// ListBucketKeys(svc, "majun-test")
	// UploadObject(svc, "majun-test", "test-xg.jpeg")
	DownloadObject(svc, "majun-test", "test-xg.jpeg")

}

func ListBuckets(svc *s3.S3) {
	result, err := svc.ListBuckets(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("%s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

func CreateBucket(svc *s3.S3, bucket string) {
	params := &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	}
	_, err := svc.CreateBucket(params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Waiting for bucket %q to be created...\n", bucket)
	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bucket %q successfully created\n", bucket)
}

func ListBucketKeys(svc *s3.S3, bucket string) {
	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}
	resp, err := svc.ListObjects(params)

	if err != nil {
		panic(err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}
}

func UploadObject(svc *s3.S3, bucket string, filePath string) {
	fp, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filePath),
		Body:   fp,
	}

	_, err = svc.PutObjectWithContext(context.Background(), putObjectInput)
	if err != nil {
		panic(err)
	}
}

func DownloadObject(svc *s3.S3, bucket string, key string) {

	getObjectOutput, err := svc.GetObjectWithContext(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		panic(err)
	}
	defer getObjectOutput.Body.Close()

	fp, err := os.Create(key)
	if err != nil {
		panic(err)
	}

	io.Copy(fp, getObjectOutput.Body)
}
