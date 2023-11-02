package main

import (
	"aliyun-oss/cli"
	"fmt"
	"os"
)

func main() {
	cli.RootCmd.AddCommand(cli.LsCommond, cli.DuCommond, cli.PushCommond, cli.PullCommond, cli.BucketCommond)

	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
