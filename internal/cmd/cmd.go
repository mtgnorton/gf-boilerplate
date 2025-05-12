package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"gf-boilerplate/internal/controller/backend/auth"
	"gf-boilerplate/internal/service/middleware"
	"gf-boilerplate/internal/service/st"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start backend server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server("api_backend")
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.HandlerResponse)
				group.Middleware(middleware.HandleError)
				group.Bind(
					auth.NewRole(),
					auth.NewMember(),
					auth.NewMenu(),
				)
			})
			provider := st.MustInitPrometheusByConfig(ctx, s)
			defer func() {
				err = provider.Shutdown(ctx)
				if err != nil {
					g.Log().Warning(ctx, err)
				}
			}()
			s.Run()
			return nil
		},
	}
)
