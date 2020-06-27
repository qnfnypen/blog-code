package pkg

import (
	"fmt"
	"regexp"
)

// GetAter 获取after的值
func GetAter(content string) string {
	// "end_cursor":"(.*)"},"edges":
	regex := `"end_cursor":"(.*?)"},"edges":`
	rp := regexp.MustCompile(regex)
	ss := rp.FindStringSubmatch(content)

	return ss[1]
}

// GetID 获取ID值
func GetID(content string) string {
	// {"id":"436245394","username":
	regex := `{"id":"(.*?)","username":`
	rp := regexp.MustCompile(regex)
	ss := rp.FindStringSubmatch(content)

	return ss[1]
}

// JoinURL 拼接url
func JoinURL(id, after string) string {
	return fmt.Sprintf(`https://www.instagram.com/graphql/query/?query_hash=%s&variables={"id":"%s","first":12,"after":"%s"}`,
		"eddbde960fed6bde675388aac39a3657", id, after)
}
