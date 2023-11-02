package util

import (
	"regexp"
	"strings"
)

// 判断输入的OSS路径格式，并返回bucketname
func IsCustomFormatValid(cloudfile string) (bool, string) {
	var bucketName string
	pattern := `^oss://([^/]+/)+([^/]+)?$`

	// 编译正则表达式
	regex := regexp.MustCompile(pattern)

	// 使用正则表达式进行匹配
	if regex.MatchString(cloudfile) {
		bucketName = strings.Split(cloudfile, "/")[2]
		return true, bucketName
	}

	return false, ""
}
