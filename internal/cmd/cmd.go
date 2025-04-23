package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"gf-boilerplate/internal/controller/backend/auth"
	"gf-boilerplate/internal/service/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.HandlerResponse)
				group.Middleware(middleware.HandleError)
				group.Bind(
					auth.NewRole(),
					auth.NewMember(),
				)
			})

			s.Run()
			return nil
		},
	}
)
