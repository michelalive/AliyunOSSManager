package cli

import (
	"aliyun-oss/store"
	"aliyun-oss/util"
	"errors"
	"github.com/spf13/cobra"
	"strings"
)

var PushCommond = &cobra.Command{
	Use:   "push",
	Short: "Used to push local files to Alibaba Cloud OSS",
	Long:  "Used to push local files to Alibaba Cloud OSS",
	RunE: func(cmd *cobra.Command, args []string) error {
		option := util.ReadConfig(config)

		//赋值bucket
		if bucket != "" {
			option.BucketName = bucket
		}

		// 赋值本地文件
		if localfile == "" {
			return errors.New("no files need to be uploaded specified")
		}
		err := util.CheckFileOrDirectoryExistence(localfile)
		if err != nil {
			return err
		}
		option.LocalFileName = localfile

		// 赋值oss文件
		var (
			localfilename = util.GetLastPartAfterSplit(option.LocalFileName)
			cloudfilename = util.GetLastPartAfterSplit(cloudfile)
		)
		if cloudfile == "" {
			option.CloudFileName = localfilename
		} else if cloudfilename != localfilename && !strings.HasSuffix(cloudfile, "/") {
			option.CloudFileName = cloudfile
		} else {
			option.CloudFileName = cloudfile
		}

		// 调用上传方法
		client := store.NewOssClient(option)
		err = client.PushFileToBucket(option)
		if err != nil {
		}
		return err
	},
}

func init() {
	PushCommond.PersistentFlags().StringVarP(&config, "config", "f", "config/config.yaml", "指定一个配置文件")
	PushCommond.PersistentFlags().StringVarP(&bucket, "name", "", "", "可手动指定需要查看的bucket")
	PushCommond.PersistentFlags().StringVarP(&localfile, "local", "", "", "指定需要上传的本地文件")
	PushCommond.PersistentFlags().StringVarP(&cloudfile, "cloud", "", "", "指定上传至阿里云oss后的文件路径")
}
