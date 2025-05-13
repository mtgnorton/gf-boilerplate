// package main 入口
package main

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"gf-boilerplate/internal/cmd"
	_ "gf-boilerplate/internal/packed"
	"gf-boilerplate/internal/service/st"
	"gf-boilerplate/internal/service/valid"
)

func main() {
	ctx := gctx.GetInitCtx()
	cleanup := prepare(ctx)
	defer cleanup(ctx)
	cmd.Main.Run(ctx)
}

func prepare(ctx context.Context) func(ctx context.Context) {
	g.I18n().SetLanguage("zh-CN")
	st.MustInitConfigByEnv(ctx)
	st.MustInitCacheFromConfig(ctx)
	cleanup := st.MustInitTracerByConfig(ctx)
	valid.RegisterAll()
	return cleanup
}
