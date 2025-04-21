// Package auth 认证服务
package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"gf-boilerplate/internal/consts"
	"gf-boilerplate/internal/model/entity"
	"gf-boilerplate/internal/service/cache"
)

// LoginReq 登录请求参数
type LoginReq struct {
	Username string `v:"required#请输入用户名" dc:"用户名"`
	Password string `v:"required#请输入密码" dc:"密码"`
}

// LoginRes 登录响应参数
type LoginRes struct {
	Token string `json:"token" dc:"访问令牌"`
}

// Service auth服务
type Service struct{}

var (
	instance = &Service{}
)

// Instance 获取服务实例
func Instance() *Service {
	return instance
}

// Login 用户登录
func (s *Service) Login(ctx context.Context, req *LoginReq) (*LoginRes, error) {
	var member entity.Member
	err := g.DB().Model("member").Where("username", req.Username).Scan(&member)
	if err != nil {
		return nil, err
	}
	if member.Id == 0 {
		return nil, gerror.New(g.I18n().T(ctx, "auth.username_not_exist"))
	}

	if member.Status != consts.StatusEnabled {
		return nil, gerror.New(g.I18n().T(ctx, "auth.account_disabled"))
	}

	// 验证密码
	if member.Password != s.EncryptPassword(req.Password, member.Salt) {
		return nil, gerror.New(g.I18n().T(ctx, "auth.password_error"))
	}

	// 生成token
	token, err := s.createToken(ctx, &member)
	if err != nil {
		return nil, err
	}

	return &LoginRes{Token: token}, nil
}

// GetLoginUser 获取登录用户信息
func (s *Service) GetLoginUser(ctx context.Context) (*entity.Member, error) {
	value := ctx.Value(consts.ContextKeyUser)
	if value == nil {
		return nil, gerror.New(g.I18n().T(ctx, "auth.unauthorized"))
	}

	var member entity.Member
	if err := gconv.Struct(value, &member); err != nil {
		return nil, err
	}
	return &member, nil
}

// EncryptPassword 密码加密
func (s *Service) EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(password + salt)
}

// createToken 创建token
func (s *Service) createToken(ctx context.Context, member *entity.Member) (string, error) {
	token := gmd5.MustEncryptString(fmt.Sprintf("%d%s%d", member.Id, member.Username, gtime.Now().UnixNano()))

	// 缓存用户信息
	cacheKey := fmt.Sprintf(consts.CacheKeyUser, member.Id)
	err := cache.Instance().Set(ctx, cacheKey, member, time.Duration(consts.TokenExpire)*time.Second)
	if err != nil {
		return "", err
	}

	return token, nil
}
