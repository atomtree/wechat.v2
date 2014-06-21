package wechat

import (
	"errors"
	"github.com/chanxuehong/wechat/message"
)

// 发送客服消息功能都一样, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) msgCustomSend(msg interface{}) (err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := clientMessageCustomSendURL(token)

	var result Error
	if err = c.postJSON(_url, msg, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		return &result
	}
	return
}

// 发送客服消息, 文本.
func (c *Client) MsgCustomSendText(msg *message.TextResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 图片.
func (c *Client) MsgCustomSendImage(msg *message.ImageResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 语音.
func (c *Client) MsgCustomSendVoice(msg *message.VoiceResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 视频.
func (c *Client) MsgCustomSendVideo(msg *message.VideoResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 音乐.
func (c *Client) MsgCustomSendMusic(msg *message.MusicResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 图文.
func (c *Client) MsgCustomSendNews(msg *message.NewsResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}
