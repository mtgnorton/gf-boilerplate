package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"gf-boilerplate/internal/router"
	"gf-boilerplate/internal/service/st"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start super, admin, broker, client server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				router.Super(ctx, group)
				router.Admin(ctx, group)
				router.Broker(ctx, group)
				router.Client(ctx, group)
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
