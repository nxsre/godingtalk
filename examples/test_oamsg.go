package main

import (
	"github.com/nxsre/godingtalk"
	"os"
)

func main() {
	dd := godingtalk.NewDingTalkClient(os.Getenv("appkey"), os.Getenv("appsecret"))
	dd.RefreshAccessToken()
	uid, _ := dd.UseridByMobile("15800000000")
	dd.SendToConversationOAMessage(uid, "chatxxxxxxxxxxxxxxxxxx", godingtalk.OAMessage{
		Head: godingtalk.OAMessageHead{
			BgColor: "FFBBBBBB",
			Text:    "测试消息",
		},
		URL: "http://dingtalk.com",
		Body: godingtalk.OAMessageBody{
			Title: "Test Title",
			Form: []godingtalk.OAMessageForm{
				godingtalk.OAMessageForm{
					Key:   "状态",
					Value: "启动",
				},
			}}})
}
