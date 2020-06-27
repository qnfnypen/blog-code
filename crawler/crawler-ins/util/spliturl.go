package util

import (
	"regexp"
	"strings"
)

// SplitURL 通过/切割url
func SplitURL(url string,th int) string {
	// https://www.instagram.com/angelaqiqi_99/
	regex := `://(.*)`
	rp := regexp.MustCompile(regex)
	ss := rp.FindStringSubmatch(url)
	s := strings.Split(ss[1],"/")

	if th > len(s) || th <= 0 {
		return ""
	}

	return s[th-1]
}