package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Broker(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/broker", func(group *ghttp.RouterGroup) {
	})
}
