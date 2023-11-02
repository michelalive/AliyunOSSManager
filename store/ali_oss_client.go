package store

import (
	"aliyun-oss/parameter"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Client struct {
	AliClient *oss.Client
}

// 定义函数生成一个oss client，用于操作oss
func NewOssClient(o *parameter.Options) *Client {
	client, err := oss.New(o.Endpoint, o.AccessKeyId, o.AccessKeySecret)
	if err != nil {
		panic(nil)
	}
	return &Client{AliClient: client}
}
