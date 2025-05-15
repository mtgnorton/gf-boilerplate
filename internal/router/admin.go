package router

import (
	"context"
	"gf-boilerplate/internal/controller/backend/auth"
	"gf-boilerplate/internal/service/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Admin(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Bind(auth.NewAccess())
		group.Middleware(middleware.HandlerResponse)
		group.Middleware(middleware.HandleError)
		group.Bind(
			auth.NewRole(),
			auth.NewMember(),
			auth.NewMenu(),
		)
	})
}
