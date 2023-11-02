package util

import (
	"fmt"
	"os"
	"strings"
)

// CheckFileOrDirectoryExistence 检查文件路径是否存在，并返回相应的错误。
func CheckFileOrDirectoryExistence(filePath string) error {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s does not exist", filePath)
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("%s is a directory", filePath)
	}
	return nil
}

// DownloadFileSpell 拼接下载后的路径
func DownloadFileSpell(localfile string, cloudfile string) string {
	fileInfo, err := os.Stat(localfile)
	if os.IsNotExist(err) {
		return localfile
	}
	if fileInfo.IsDir() {
		if strings.HasSuffix(localfile, "/") {
			return localfile + cloudfile
		}
		return localfile + "/" + cloudfile
	}
	return ""
}

// 获取带目录的文件路径下的文件名
func GetLastPartAfterSplit(s string) string {
	parts := strings.Split(s, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}
