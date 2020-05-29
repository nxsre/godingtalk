package main

import (
	dingtalk "github.com/nxsre/godingtalk"
)

func main() {
	dd := dingtalk.NewDingTalkClient("", "")
	dd.RobotSecret = "SECcdb07xxxxxxxx30a6095f2"
	dd.RobotToken = "6ca3baxxxxxxxxxxx82bf79a"

	msg:=`
标题
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题

引用
> A man who stands for nothing will fall for anything.

文字加粗、斜体
**bold**
*italic*

链接
[this is a link](http://www.baidu.com)

图片
![](http://a2.peoplecdn.cn/xxxx.png)

无序列表
- item1
- item2

有序列表
1. item1
2. item2
`
	dd.SendRobotMarkdownMessage("test",msg,  &dingtalk.RobotAtList{
		AtMobiles: []string{"153xxxxxxxx"},
		IsAtAll: true,
	})

	//blackfriday.New(blackfriday.)

	//dd.SendRobotTextMessage("test", &dingtalk.RobotAtList{
	//	AtMobiles: []string{"15311519572"},
	//	//IsAtAll: true,
	//})
}
