package util

import (
	"bufio"
	"os"

	// 初始化log
	_ "github.com/qnfnypen/crawler-ins/global"
	"github.com/rs/zerolog/log"
)

// WriteToFile 将url写入文件
func WriteToFile(path, url string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal().Str("Error", err.Error()).Msg("打开存储文件失败")
	}
	defer file.Close()

	bw := bufio.NewWriter(file)
	bw.WriteString(url + "\n")
	bw.Flush()
}
