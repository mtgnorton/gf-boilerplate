// Package middleware 中间件-权限验证
package middleware

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"gf-boilerplate/internal/consts"
	"gf-boilerplate/internal/model/entity"
	"gf-boilerplate/internal/service/cache"
)

// Auth 认证中间件
func Auth(r *ghttp.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		r.Response.WriteJsonExit(DefaultHandlerResponse{
			Code:    401,
			Message: g.I18n().T(r.Context(), "auth.unauthorized"),
		})
		return
	}

	// 去掉Bearer前缀
	token = strings.TrimPrefix(token, consts.TokenType+" ")

	// 获取用户信息
	var member *entity.Member
	cacheValue, err := cache.Instance().Get(r.Context(), fmt.Sprintf(consts.CacheKeyUser, token))
	if err != nil {
		r.Response.WriteJsonExit(DefaultHandlerResponse{
			Code:    401,
			Message: g.I18n().T(r.Context(), "auth.unauthorized"),
		})
		return
	}

	if cacheValue == nil {
		r.Response.WriteJsonExit(DefaultHandlerResponse{
			Code:    401,
			Message: g.I18n().T(r.Context(), "auth.token_expired"),
		})
		return
	}

	if err = cacheValue.Struct(&member); err != nil {
		r.Response.WriteJsonExit(DefaultHandlerResponse{
			Code:    401,
			Message: g.I18n().T(r.Context(), "auth.unauthorized"),
		})
		return
	}

	// 设置用户信息到上下文
	r.SetCtxVar(consts.ContextKeyUser, member)
	r.SetCtxVar(consts.ContextKeyRoleID, member.RoleId)

	r.Middleware.Next()
}

// RBAC RBAC权限验证中间件
func RBAC(r *ghttp.Request) {
	// 获取当前用户角色
	roleID := r.GetCtxVar(consts.ContextKeyRoleID).Uint64()
	if roleID == 0 {
		r.Response.WriteJsonExit(DefaultHandlerResponse{
			Code:    403,
			Message: g.I18n().T(r.Context(), "auth.forbidden"),
		})
		return
	}

	// 获取角色信息
	var role entity.Role
	err := g.DB().Model("role").Where("id", roleID).Scan(&role)
	if err != nil {
		r.Response.WriteJsonExit(DefaultHandlerResponse{
			Code:    500,
			Message: g.I18n().T(r.Context(), "auth.role_not_exist"),
		})
		return
	}

	// 检查角色状态
	if !role.Status {
		r.Response.WriteJsonExit(DefaultHandlerResponse{
			Code:    403,
			Message: g.I18n().T(r.Context(), "auth.role_disabled"),
		})
		return
	}

	// 获取当前请求的权限标识
	path := r.URL.Path
	method := r.Method

	// 查询权限
	count, err := g.DB().Model("casbin").
		Where("p_type", "p").
		Where("v0", role.Key).
		Where("v1", path).
		Where("v2", method).
		Count()
	if err != nil {
		r.Response.WriteJsonExit(DefaultHandlerResponse{
			Code:    500,
			Message: gerror.New(g.I18n().T(r.Context(), "auth.check_permission_failed")).Error(),
		})
		return
	}

	if count == 0 {
		r.Response.WriteJsonExit(DefaultHandlerResponse{
			Code:    403,
			Message: g.I18n().T(r.Context(), "auth.forbidden"),
		})
		return
	}

	r.Middleware.Next()
}
