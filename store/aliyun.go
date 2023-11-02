package store

import (
	"aliyun-oss/parameter"
	"aliyun-oss/util"
	"fmt"
	"log"
)

// 实现 Aliyunosser interface{} 的 GetBucketList() 方法
func (c *Client) GetBucketList() ([]string, error) {
	var bucket_list []string

	lsRes, err := c.AliClient.ListBuckets()
	if err != nil {
		return nil, err
	}

	for _, bucket := range lsRes.Buckets {
		bucket_list = append(bucket_list, bucket.Name)
	}
	return bucket_list, nil
}

// 实现 Aliyunosser interface{} 的 GetBucketFileList(o *parameter.Options) 方法
func (c *Client) GetBucketFileList(o *parameter.Options) error {
	bucket, err := c.AliClient.Bucket(o.BucketName)
	if err != nil {
		log.Fatal(err)
	}

	lsRes, err := bucket.ListObjects()
	if err != nil {
		log.Fatal(err)
	}

	for _, object := range lsRes.Objects {
		re := util.FormatObject(object)
		fmt.Println(re)
	}
	fmt.Printf("\nBucket: %s 中共有%d个文件\n", o.BucketName, len(lsRes.Objects))

	return nil
}

// 实现 Aliyunosser interface{} 的 GetBucketStorage(o *parameter.Options) 方法
func (c *Client) GetBucketStorage(o *parameter.Options) error {
	var (
		size       int64
		total_size string
	)
	bucket, err := c.AliClient.Bucket(o.BucketName)
	if err != nil {
		return err
	}

	lsRes, err := bucket.ListObjects()
	if err != nil {
		return err
	}

	for _, object := range lsRes.Objects {
		size += object.Size
		fmt.Printf("\r%s 已使用: %d", o.BucketName, size)
	}

	total_size = util.FormatStorageSize(size)

	fmt.Printf("\r%s 已使用: %s\n", o.BucketName, total_size)

	return nil
}

// 实现 Aliyunosser interface{} 的 PushFileToBucket(o *parameter.Options) 方法
func (c *Client) PushFileToBucket(o *parameter.Options) error {
	bucket, err := c.AliClient.Bucket(o.BucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(o.CloudFileName, o.LocalFileName)
	if err != nil {
		return err
	}
	log.Printf("文件%s已成功上传至oss://%s/%s.\n", o.LocalFileName, o.BucketName, o.CloudFileName)

	return nil
}

func (c *Client) PullFileFromBucket(o *parameter.Options) error {
	bucket, err := c.AliClient.Bucket(o.BucketName)
	if err != nil {
		return err
	}

	err = bucket.GetObjectToFile(o.CloudFileName, o.LocalFileName)
	if err != nil {
		return err
	}
	log.Printf("文件oss://%s/%s已成功下载至%s", o.BucketName, o.CloudFileName, o.LocalFileName)
	return nil
}
