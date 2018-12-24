package aliyunoss

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Client struct {
	key        string
	sec        string
	region     string
	bucketName string
	internal   bool
}

func NewAliyunOSSClient(key, sec, region, bucketName string, internal bool) (*Client, error) {
	return &Client{key, sec, region, bucketName, internal}, nil
}

func (c *Client) load() (bucket *oss.Bucket, err error) {

	var endpoint string

	if c.internal {
		// oss-cn-hangzhou-internal.aliyuncs.com
		endpoint = fmt.Sprintf("https://oss-%s-internal.aliyuncs.com", c.region)
	} else {
		endpoint = fmt.Sprintf("https://oss-%s.aliyuncs.com", c.region)
	}

	client, err := oss.New(endpoint, c.key, c.sec)
	if err != nil {
		return nil, err
	}

	bucket, err = client.Bucket(c.bucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

// Put file to remote
func (c *Client) Put(object, filepath string) error {
	// return bucket.PutObjectFromFile(object, filepath)
	bucket, err := c.load()

	if err != nil {
		return err
	}

	return bucket.PutObjectFromFile(object, filepath)

}
