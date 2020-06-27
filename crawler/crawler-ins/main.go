package main

import (
	"encoding/json"
	"fmt"

	"github.com/qnfnypen/crawler-ins/model"
	"github.com/qnfnypen/crawler-ins/pkg"
)

func main() {
	content := pkg.GetContent("https://www.instagram.com/angelaqiqi_99/")

	// after := pkg.GetAter(content)
	id := pkg.GetID(content)

	// QVFEM2ZrRk40aXBFMmpZclhvWlB0VDdySkFJVHJBNGhBRldRcm51VVF2MHZhTlY2bzR2amFSZG4zelJseWwxeHgtWVBfY0tScG5uYkFHTHc3cHh1dnl5dw==
	url := pkg.JoinURL(id, "QVFENW44LTFtYjBYcHFOS250dThwOGQ3UHoxNTZ6OElLZkpQXzBCR1FEZlhRb3hRN3FtYkhMN3ltODdhdEdNNEFieWs2ZVJKT0NTZkVHdmZSazlrTHd2WA==")
	fmt.Println(url)

	content = pkg.GetContent(url)
	var ins model.InsInfo
	json.Unmarshal([]byte(content), &ins)
	for _, v := range ins.Data.User.EdgeOwnerToTimelineMedia.Edges {
		if v.Node.IsVideo {
			fmt.Println("video；", v.Node.VideoURL)
		} else {
			fmt.Println("img：", v.Node.DisplayURL)
		}
	}
	// for {
	// 	var ins model.InsInfo
	// 	json.Unmarshal([]byte(content), &ins)
	// 	for _, v := range ins.Data.User.EdgeOwnerToTimelineMedia.Edges {
	// 		if v.Node.IsVideo {
	// 			util.WriteToFile("doc/video.txt", v.Node.VideoURL)
	// 		} else {
	// 			util.WriteToFile("doc/img.txt", v.Node.DisplayURL)
	// 		}
	// 	}
	// 	after := pkg.GetAter(content)
	// 	url = pkg.JoinURL(id, after)
	// 	content = pkg.GetContent(url)

	// 	if !ins.Data.User.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage {
	// 		return
	// 	}
	// }
}
