package cli

import (
	"aliyun-oss/store"
	"aliyun-oss/util"
	"errors"
	"github.com/spf13/cobra"
)

var PullCommond = &cobra.Command{
	Use:   "pull",
	Short: "Download files from Alibaba Cloud OSS",
	Long:  "Download files from Alibaba Cloud OSS",
	RunE: func(cmd *cobra.Command, args []string) error {
		option := util.ReadConfig(config)

		//赋值bucket
		if bucket != "" {
			option.BucketName = bucket
		}

		// 赋值oss文件
		if cloudfile == "" {
			return errors.New("no files need to be downloaded specified")
		}
		option.CloudFileName = cloudfile

		// 赋值本地文件
		var cloudfilename = util.GetLastPartAfterSplit(cloudfile)

		if localfile == "" {
			option.LocalFileName = cloudfilename
		} else {
			file := util.DownloadFileSpell(localfile, cloudfilename)
			option.LocalFileName = file
		}

		// 调用下载方法
		client := store.NewOssClient(option)
		err := client.PullFileFromBucket(option)
		if err != nil {
		}
		return err
	},
}

func init() {
	PullCommond.PersistentFlags().StringVarP(&config, "config", "f", "config/config.yaml", "指定一个配置文件")
	PullCommond.PersistentFlags().StringVarP(&bucket, "name", "", "", "可手动指定需要查看的bucket")
	PullCommond.PersistentFlags().StringVarP(&localfile, "local", "", "", "指定需要上传的本地文件")
	PullCommond.PersistentFlags().StringVarP(&cloudfile, "cloud", "", "", "指定上传至阿里云oss后的文件路径")
}
