package store

import "aliyun-oss/parameter"

type Aliyunosser interface {
	// 获取bucket列表
	GetBucketList() ([]string, error)

	// 获取bucket文件列表
	GetBucketFileList(o *parameter.Options) error

	// 向bucket中推送文件
	PushFileToBucket(o *parameter.Options) error

	// 从bucket中下载文件
	PullFileFromBucket(o *parameter.Options) error

	// 获取bucket的存储空间大小
	GetBucketStorage(o *parameter.Options) error
}
