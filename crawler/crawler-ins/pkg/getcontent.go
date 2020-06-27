package pkg

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	// log初始化
	_ "github.com/qnfnypen/crawler-ins/global"
	"github.com/rs/zerolog/log"
)

// GetContent 获取页面内容
func GetContent(html string) string {
	// 解析代理地址
	proxy, err := url.Parse("http://127.0.0.1:10809")
	if err != nil {
		log.Fatal().Str("Error", err.Error()).Msg("代理地址解析错误")
	}
	transport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		ResponseHeaderTimeout: 5 * time.Second,
		MaxIdleConnsPerHost:   10,
	}
	// 定义client
	client := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
	}
	// 创建请求
	req, err := http.NewRequest(http.MethodGet, html, nil)
	req.Header.Add("user-agent",`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.36`)
	req.Header.Set("cookie",`ig_did=7157A668-1F19-4CF3-B16C-2995DF9B7FE3; mid=XvWMfQALAAF4uXAVNPgscDGhxfQ4; csrftoken=DLWL073LwyEGEFppJTaA0Cngio0SwZEO; ds_user_id=38041720444; sessionid=38041720444%3AEDFjitOs7zE1xu%3A1; rur=ASH; urlgen="{\"8.210.7.43\": 45102\054 \"47.240.49.254\": 45102\054 \"35.221.222.156\": 15169\054 \"8.210.0.106\": 45102}:1jokTE:JN9jM5VeFNjwYxw3ocFqpv07bNA"`)
	if err != nil {
		log.Fatal().Str("Error", err.Error()).Msg("创建请求失败")
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal().Str("Error", err.Error()).Msg("请求链接失败")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal().Msg("状态码不为200")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Str("Error", err.Error()).Msg("获取response失败")
	}

	return string(body)
}
