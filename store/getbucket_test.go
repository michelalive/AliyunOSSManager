package store

import (
	"aliyun-oss/parameter"
	"github.com/stretchr/testify/assert"
	"testing"
)

var option = parameter.Options{
	Endpoint:        "your-Endpoint",
	AccessKeyId:     "your-AccessKeyId",
	AccessKeySecret: "your-AccessKeySecret",
	BucketName:      "your-BucketName",
	LocalFileName:   "your-LocalFileName",
	CloudFileName:   "your-CloudFileName",
}

var should *assert.Assertions

// 测试获取bucket列表的方法
func TestAliyunOssClient_GetBucketList(t *testing.T) {
	should = assert.New(t)
	client := NewOssClient(&option)
	bucket_list, err := client.GetBucketList()
	should.NoError(err)
	should.NotEmpty(bucket_list)
}

// 测试获取bucket中文件的方法
func TestAliyunOssClient_GetBucketFileList(t *testing.T) {
	should = assert.New(t)
	client := NewOssClient(&option)
	err := client.GetBucketFileList(&option)
	should.NoError(err)
}

// 测试获取Bucket存储空间的方法
func TestClient_GetBucketStorage(t *testing.T) {
	should = assert.New(t)
	client := NewOssClient(&option)
	err := client.GetBucketStorage(&option)
	should.NoError(err)
}

// 测试上传文件功能的方法
func TestClient_PushFileToBucket(t *testing.T) {
	should = assert.New(t)
	client := NewOssClient(&option)
	err := client.PushFileToBucket(&option)
	should.NoError(err)
}

// 测试下载文件的功能的方法
func TestClient_PullFileFromBucket(t *testing.T) {
	should = assert.New(t)
	client := NewOssClient(&option)
	err := client.PullFileFromBucket(&option)
	should.NoError(err)
}
