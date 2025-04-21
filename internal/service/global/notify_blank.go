package global

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

const defaultNotifyName = "notify"

// Notifier 定义通知发送器接口
type Notifier interface {
	Send(ctx context.Context, content string) error
}

func GetNotifier(ctx context.Context) Notifier {
	//nolint:errcheck
	return globalMapping.GetOrSetFuncLock(defaultNotifyName, func() interface{} {
		defaultNotify, err := g.Cfg().Get(ctx, "notify.default")
		if err != nil {
			g.Log().Warningf(ctx, "获取默认通知器失败: %v", err)
			return &BlankBot{}
		}
		switch defaultNotify.String() {
		case "fs":
			webhook := g.Cfg().MustGet(ctx, "notify.fsWebhook").String()
			return &FeiShuBot{
				webhook: webhook,
			}
		default:
			return &BlankBot{}
		}
	}).(Notifier)
}

// BlankBot 空实现，不发送任何消息
type BlankBot struct{}

// Send 空实现，不发送任何消息
func (b *BlankBot) Send(ctx context.Context, content string) error {
	// 不执行任何操作，直接返回nil
	return nil
}
