package godingtalk

import (
	"net/url"
)

type RobotAtList struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

//SendRobotMarkdownMessage can send a text message to a group chat
// 如果需要 at , msg 中必须包含at手机号的文本如： @153xxxxxx , atall 的消息不需要在 msg 包含 @
func (c *DingTalkClient) SendRobotMarkdownMessage(title string, msg string, at *RobotAtList) (data MessageResponse, err error) {
	params := url.Values{}
	request := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  msg,
		},
	}
	if at != nil {
		request["at"] = at
	}
	err = c.httpRPC("robot/send", params, request, &data)
	return data, err
}

// SendRobotTextMessage can send a text message and at user to a group chat
func (c *DingTalkClient) SendRobotTextMessage(msg string, at *RobotAtList) (data OAPIResponse, err error) {
	params := url.Values{}
	request := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	if at != nil {
		request["at"] = at
	}
	err = c.httpRPC("robot/send", params, request, &data)
	return data, err
}


// SendRobotOAMessage can send a text message and at user to a group chat
func (c *DingTalkClient) SendRobotOAMessage(msg OAMessage, at *RobotAtList) (data OAPIResponse, err error) {
	params := url.Values{}
	request := map[string]interface{}{
		"msgtype": "oa",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	if at != nil {
		request["at"] = at
	}
	err = c.httpRPC("robot/send", params, request, &data)
	return data, err
}
