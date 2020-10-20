package s3

import (
	"fmt"

	"github.com/yuhang2/initsequence/server/config"
)

func newS3Client() Client {
	return &awsClient{
		configName: config.Config.Name,
	}
}

type Client interface {
	Upload(objectID string)
}

type awsClient struct {
	configName string
}

func (client *awsClient) Upload(objectID string) {
	fmt.Println("Output: upload - ", client.configName)
	fmt.Println("The expect output should be \"upload - load config\"")
}

func InitializeSomething() {}

var S3Client = newS3Client()
