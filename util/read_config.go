package util

import (
	"aliyun-oss/parameter"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var config *parameter.Options

func ReadConfig(configfile string) *parameter.Options {
	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// 解析 YAML 数据到 Config 结构体
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		return nil
	}

	return config
}
