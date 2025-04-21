// Package global 全局组件-飞书通知
package global

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

const feiShuSendTimeout = time.Second * 2

// FeiShuBot 飞书机器人实现
type FeiShuBot struct {
	webhook string
}

// Send 发送消息到飞书
func (f *FeiShuBot) Send(ctx context.Context, content string) error {
	if f.webhook == "" {
		return gerror.New("webhook 未设置")
	}
	data := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]string{
			"text": content,
		},
	}

	response, err := g.Client().SetTimeout(feiShuSendTimeout).Post(ctx, f.webhook, data)
	if err != nil {
		return gerror.Wrapf(err, "飞书消息发送失败,webhook: %s", f.webhook)
	}
	defer func() {
		err = response.Close()
		if err != nil {
			g.Log().Error(ctx, "关闭飞书消息发送响应失败,webhook: %s, err: %v", f.webhook, err)
		}
	}()

	return nil
}
