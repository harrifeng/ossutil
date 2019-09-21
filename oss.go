package ossutil

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// GetObjectTimedURL is used to get oss url which will expire (set by the expireSecond)
func GetObjectTimedURL(
	endpoint string,
	accessID string,
	accessKey string,
	bucketName string,
	objectName string,
	expireSecond int64,
) string {

	client, err := oss.New(endpoint, accessID, accessKey)
	if err != nil {
		fmt.Println("client err:", err)
		return ""
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("bucket err", err)
		return ""
	}
	url, err := bucket.SignURL(objectName, oss.HTTPGet, expireSecond)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return url
}
