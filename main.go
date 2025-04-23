// package main 入口
package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"gf-boilerplate/internal/cmd"
	_ "gf-boilerplate/internal/packed"
	"gf-boilerplate/internal/service/valid"
)

func main() {
	g.I18n().SetLanguage("zh-CN")
	valid.RegisterAll()
	cmd.Main.Run(gctx.GetInitCtx())
}
