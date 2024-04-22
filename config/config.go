package config

const webHookAlert = "https://oapi.dingtalk.com/robot/send?access_token=123456"

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
