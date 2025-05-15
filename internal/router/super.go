package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Super(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/super", func(group *ghttp.RouterGroup) {
	})
}
