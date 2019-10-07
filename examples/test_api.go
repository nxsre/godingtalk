package main

import (
	"fmt"
	"github.com/soopsio/godingtalk"
)

func main() {

	c := godingtalk.NewDingTalkClient("ding****w9m6ap", "N8Vqh3KIPD******VHBBCD4or4m")
	err := c.RefreshAccessToken()
	if err != nil {
		panic(err)
	}

	record, err := c.DepartmentList()
	if err != nil {
		panic(err)
	}

	fmt.Println(record.Departments)

	us, err := c.UserList(83000064, 0, 100)
	fmt.Printf("%+v", us)

	// 提示为《你的好友在群组里发了一条消息》
	//c.SendAppLinkMessage("297***723","1322****8304","111111222222","你是谁222","https://static-legacy.dingtalk.com/media/lADPBbCc1g87THbM18zX_215_215.jpg","http://www.baidu.com")
	// 提示为消息内容
	//c.SendAppMessage("297***723","1322****8304","99990330")
	// 提示为消息内容 Body.Title
	//c.SendAppOAMessage("297***723", "1322****8304", godingtalk.OAMessage{
	//	URL:   "https://www.baidu.com",
	//	PcURL: "https://www.csdn.net",
	//	Head: struct {
	//		BgColor string `json:"bgcolor,omitempty"`
	//		Text    string `json:"text,omitempty"`
	//	}{
	//		"red",
	//		"222222",
	//	},
	//	Body: struct {
	//		Title     string                     `json:"title,omitempty"`
	//		Form      []godingtalk.OAMessageForm `json:"form,omitempty"`
	//		Rich      godingtalk.OAMessageRich   `json:"rich,omitempty"`
	//		Content   string                     `json:"content,omitempty"`
	//		Image     string                     `json:"image,omitempty"`
	//		FileCount int                        `json:"file_count,omitempty"`
	//		Author    string                     `json:"author,omitempty"`
	//	}{
	//		Title: "4444",
	//		Form: []godingtalk.OAMessageForm{
	//			godingtalk.OAMessageForm{
	//				Key:   "aaa",
	//				Value: "aaaaval",
	//			},
	//		},
	//		Rich: godingtalk.OAMessageRich{
	//			"1111",
	//			"22223333",
	//		},
	//		Content: "testContent",
	//		Author:  "王勇敬",
	//	},
	//})

	// 提示为《你的好友在群组里发了一条消息》
	//c.SendAppActionCardMessage("297466956", "1322****8304", godingtalk.ActionCardMessage{
	//	Title: "测试 ActionCard",
	//	//SingleTitle: "测试2222",
	//	//SingleUrl:   "https://www.baidu.com",
	//	Markdown:       "### 1111122221",
	//	BtnOrientation: "1",
	//	BtnJsonList: []godingtalk.ActionCardMessageBtn{
	//		godingtalk.ActionCardMessageBtn{
	//			Title:     "按钮1",
	//			ActionUrl: "http://www.baidu.com",
	//		},
	//		godingtalk.ActionCardMessageBtn{
	//			Title:     "按钮2",
	//			ActionUrl: "http://www.csdn.net",
	//		},
	//	},
	//})

	u, _ := c.UserInfoByID("132****04")
	fmt.Printf("%+v", u.IsLeaderInDepts)

	//
	//result, err := c.ListAttendanceResult([]string{"085354234826136236"}, dataFrom, dataTo, 0, 2)
	//if err != nil {
	//	panic(err)
	//} else if len(result.Records) > 0 {
	//	fmt.Printf("%#v\n", result.Records[0])
	//}
}
