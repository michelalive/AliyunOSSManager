package util

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
	"unicode"
)

// FormatObject 格式化bucketfilielist的输出
func FormatObject(object oss.ObjectProperties) string {
	key := adjustWidth(object.Key, 100)
	lastModified := adjustWidth(object.LastModified.Format("2006-01-02 15:04:05"), 20)
	size := adjustWidth(fmt.Sprintf("%d", object.Size), 10)

	return key + " " + lastModified + " " + size
}

// adjustWidth 调整字段宽度以考虑中英文字符宽度差异
func adjustWidth(s string, width int) string {
	totalWidth := 0
	var wg sync.WaitGroup

	for _, r := range s {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if unicode.Is(unicode.Han, r) {
				totalWidth += 2 // 中文字符宽度为2
			} else {
				totalWidth++
			}
		}()

		wg.Wait()
	}

	for totalWidth < width {
		s += " "
		totalWidth++
	}

	return s
}
