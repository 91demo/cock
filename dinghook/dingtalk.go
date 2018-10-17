package dinghook

import (
	"errors"
	"log"
)

// DingTalk 钉钉结构体
type DingTalk struct {
}

// Send 发送钉钉消息
func (d *DingTalk) Send(token string, content string) error {
	if token == "" {
		return errors.New("need dingding token")
	}

	// 发送钉钉
	ding := NewDing(token)
	result := ding.SendMessage(Message{Content: content, AtPersion: atMobiles, AtAll: isAtAll})

	if !result.Success {
		log.Println("token:", token)
		log.Fatalln("Error", result.ErrMsg)
	}

	return nil
}

// NewDingTalk 初始化DingTalk结构体
func NewDingTalk() *DingTalk {
	return &DingTalk{}
}
