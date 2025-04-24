// package main 入口
package main

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"gf-boilerplate/internal/cmd"
	_ "gf-boilerplate/internal/packed"

	"gf-boilerplate/internal/service/global"
	"gf-boilerplate/internal/service/valid"
)

func main() {
	ctx := gctx.GetInitCtx()
	prepare(ctx)
	cmd.Main.Run(ctx)

}

func prepare(ctx context.Context) {
	g.I18n().SetLanguage("zh-CN")
	global.GetConfig().InitConfigFromEnv(ctx)
	valid.RegisterAll()
}
