package cli

import (
	"aliyun-oss/store"
	"aliyun-oss/util"
	"fmt"
	"github.com/spf13/cobra"
)

var BucketCommond = &cobra.Command{
	Use:   "bucket",
	Short: "View the list of bucket",
	Long:  "View the list of bucket",
	RunE: func(cmd *cobra.Command, args []string) error {
		option := util.ReadConfig(config)
		client := store.NewOssClient(option)
		bucket_list, err := client.GetBucketList()
		if err != nil {
			return err
		}
		fmt.Println("你的oss中有以下bucket:")
		for _, b := range bucket_list {
			fmt.Println(b)
		}
		return nil
	},
}

func init() {
	BucketCommond.PersistentFlags().StringVarP(&config, "config", "f", "config/config.yaml", "指定一个配置文件")
}
