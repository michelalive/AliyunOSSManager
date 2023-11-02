package cli

import (
	"aliyun-oss/store"
	"aliyun-oss/util"
	"github.com/spf13/cobra"
)

var DuCommond = &cobra.Command{
	Use:   "du",
	Short: "View the storage space used by the bucket",
	Long:  "View the storage space used by the bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		option := util.ReadConfig(config)
		if bucket != "" {
			option.BucketName = bucket
		}
		client := store.NewOssClient(option)
		err := client.GetBucketStorage(option)
		if err != nil {
		}
		return err
	},
}

func init() {
	DuCommond.PersistentFlags().StringVarP(&config, "config", "f", "config/config.yaml", "指定一个配置文件")
	DuCommond.PersistentFlags().StringVarP(&bucket, "name", "", "", "可手动指定需要查看的bucket")
}
