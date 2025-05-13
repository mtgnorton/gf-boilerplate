// Package st 全局单例-灵活配置
package st

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/text/gstr"
)

var defaultConfig = &config{}

// config 一些可以动态配置的配置
type config struct{}

// GetConfig 获取配置
func GetConfig() *config {
	return defaultConfig
}

// GetDebug 获取debug模式
func (c *config) GetDebug(ctx context.Context) bool {
	debug, err := g.Cfg().Get(ctx, "system.debug")
	if err != nil {
		g.Log().Infof(ctx, "获取配置debug失败: %v", err)
		return true
	}
	return debug.Bool()
}

// SetDebug 动态设置debug模式
func (c *config) SetDebug(ctx context.Context, debug bool) error {
	//nolint:errcheck
	err := g.Cfg().GetAdapter().(*gcfg.AdapterFile).Set("system.debug", debug)
	if err != nil {
		return gerror.Wrapf(err, "设置配置debug失败")
	}
	return nil
}

// MustInitConfigByEnv 收集以WLINK_开头的环境变量并设置到配置中
func MustInitConfigByEnv(ctx context.Context) {
	keys := []string{}
	for key, val := range genv.Map() {
		if gstr.HasPrefix(key, "WLINK_") {
			key = gstr.ToLower(gstr.Replace(gstr.Replace(key, "WLINK_", ""), "_", "."))
			keys = append(keys, key)
			//nolint:errcheck
			err := g.Cfg().GetAdapter().(*gcfg.AdapterFile).Set(key, val)
			if err != nil {
				g.Log().Panicf(ctx, "从环境变量初始化配置失败 err:%+v", err)
			}
		}
	}
	g.Log().Infof(ctx, "从环境变量初始化配置: %v", keys)
}
