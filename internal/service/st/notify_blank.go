// Package st 全局单例-通知器
package st

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

var defaultNotifier Notifier

// GetNotifier 定义通知发送器接口
type Notifier interface {
	Send(ctx context.Context, content string) error
}

func GetNotifier() Notifier {
	if defaultNotifier == nil {
		panic("notify not init")
	}
	return defaultNotifier
}

// InitNotifierByConfig 根据配置获取默认通知器
func InitNotifierByConfig(ctx context.Context) {
	defaultNotify, err := g.Cfg().Get(ctx, "notify.default")
	if err != nil {
		g.Log().Warningf(ctx, "获取默认通知器失败: %v", err)
		defaultNotifier = &BlankBot{}
	}
	switch defaultNotify.String() {
	case "fs":
		webhook := g.Cfg().MustGet(ctx, "notify.fsWebhook").String()
		defaultNotifier = &FeiShuBot{
			webhook: webhook,
		}
	default:
		defaultNotifier = &BlankBot{}
	}
}

// BlankBot 空实现，不发送任何消息
type BlankBot struct{}

// Send 空实现，不发送任何消息
func (b *BlankBot) Send(ctx context.Context, content string) error {
	// 不执行任何操作，直接返回nil
	return nil
}
