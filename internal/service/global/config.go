// Package global 全局组件-灵活配置
package global

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

const defaultConfigName = "config"

// Config 一些可以动态配置的配置
type Config struct {
}

// GetConfig 获取配置
func GetConfig() *Config {
	//nolint:errcheck
	return globalMapping.GetOrSetFuncLock(defaultConfigName, func() interface{} {
		return &Config{}
	}).(*Config)
}

// GetDebug 获取debug模式
func (c *Config) GetDebug(ctx context.Context) bool {
	debug, err := g.Cfg().Get(ctx, "system.debug")
	if err != nil {
		g.Log().Infof(ctx, "获取配置debug失败: %v", err)
		return true
	}
	return debug.Bool()
}

// SetDebug 动态设置debug模式
func (c *Config) SetDebug(ctx context.Context, debug bool) error {
	//nolint:errcheck
	err := g.Cfg().GetAdapter().(*gcfg.AdapterFile).Set("system.debug", debug)
	if err != nil {
		return gerror.Wrapf(err, "设置配置debug失败")
	}
	return nil
}
