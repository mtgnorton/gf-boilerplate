// Package cmd 路由配置
package cmd

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"gf-boilerplate/internal/controller"
	"gf-boilerplate/internal/service/middleware"
)

// InitRouter 初始化路由
func InitRouter() {
	s := g.Server()
	s.Use(middleware.HandlerResponse)

	// 公共路由
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Auth{},
		)
	})

	// 需要认证的路由
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.Bind(
			controller.Auth{},
		)
	})

	// 需要RBAC权限验证的路由
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth, middleware.RBAC)
		group.Bind(
			controller.Member{},
			controller.Role{},
			controller.Menu{},
		)
	})
}
