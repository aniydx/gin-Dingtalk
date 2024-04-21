package config

const webHookAlert = "https://oapi.dingtalk.com/robot/send?access_token=614f6b6d7f91fb77de01e012df3b17751c7fd125d06fe4af482b562e3bda3b97"

type DingTalkMsg struct {
	MsgType  string `json:"msgtype"`
	Markdown `json:"markdown"`
	At       `json:"at"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}
