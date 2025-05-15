package router

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Client(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/client", func(group *ghttp.RouterGroup) {
	})
}
