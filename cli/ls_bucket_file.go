package cli

import (
	"aliyun-oss/store"
	"aliyun-oss/util"
	"github.com/spf13/cobra"
)

var LsCommond = &cobra.Command{
	Use:   "ls",
	Short: "View the list of files in the bucket",
	Long:  "View the list of files in the bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		option := util.ReadConfig(config)
		if bucket != "" {
			option.BucketName = bucket
		}
		client := store.NewOssClient(option)
		err := client.GetBucketFileList(option)
		if err != nil {
		}
		return err
	},
}

func init() {
	LsCommond.PersistentFlags().StringVarP(&config, "config", "f", "config/config.yaml", "指定一个配置文件")
	LsCommond.PersistentFlags().StringVarP(&bucket, "name", "", "", "可手动指定需要查看的bucket")
}
