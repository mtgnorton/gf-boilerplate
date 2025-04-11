// package main 入口
package main

import (
	_ "gf-boilerplate/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-boilerplate/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
