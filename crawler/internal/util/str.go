package util

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strings"
)

func StripHTML(s string) string {
	// 把块级标签替换为换行，保留语义分割
	s = strings.NewReplacer(
		"</p>", "\n",
		"</div>", "\n",
		"</li>", "\n",
		"<br>", "\n",
		"<br/>", "\n",
		"<br />", "\n",
	).Replace(s)
	// 去掉所有HTML标签
	s = regexp.MustCompile(`<[^>]*>`).ReplaceAllString(s, "")
	// 将换行、制表符、连续空白替换为单个空格
	s = regexp.MustCompile(`[\n\r\t]+`).ReplaceAllString(s, " ")
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	// 去掉首尾空白
	return strings.TrimSpace(s)
}

func MD5Hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func Truncate(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) > maxLen {
		return string(runes[:maxLen])
	}
	return s
}
