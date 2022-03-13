package main

import (
	"ManagementTwitter/config"
	"ManagementTwitter/template"
	"ManagementTwitter/tweetapi"
	"strings"
	"time"
)

func main() {
	//fmt.Println(tweetapi.PostTweet("test"))

	//予約投稿処理
	//(ツイッターの予約投稿は使わず、ゴルーチンで試してみたかった)
	layout := "2006/01/02 15:04:05"
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		//本来は排他制御が必要
		for range ticker.C {
			data, err := config.Readfile()
			if err == nil {
				arData := strings.Split(string(data), ",")
				t, _ := time.Parse(layout, arData[0])
				if time.Now().Before(t) && arData[2] == "0" {
					//Now < t & ツイートフラグ=0

					//ツイート
					tweetapi.PostTweet(arData[1])

					//ツイート完了
					data = []byte(arData[0] + "," + arData[1] + ",1")
					err = config.Savefile(data)
					if err != nil {
						return
					}
				}

			}
		}
	}()
	// err = config.Savefile([]byte(data))
	// if err != nil {
	// 	return
	// }
	template.HandlerCall()

}
