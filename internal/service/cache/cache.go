// Package cache 缓存服务
package cache

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
)

// Service 缓存服务
type Service struct {
	cache *gcache.Cache
}

var (
	instance = newService()
)

// newService 创建服务实例
func newService() *Service {
	return &Service{
		cache: gcache.New(),
	}
}

// Instance 获取服务实例
func Instance() *Service {
	return instance
}

// Set 设置缓存
func (s *Service) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	return s.cache.Set(ctx, key, value, duration)
}

// Get 获取缓存
func (s *Service) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	return s.cache.Get(ctx, key)
}

// Remove 删除缓存
func (s *Service) Remove(ctx context.Context, key interface{}) error {
	return s.cache.Remove(ctx, key)
}

// Clear 清空缓存
func (s *Service) Clear(ctx context.Context) error {
	return s.cache.Clear(ctx)
}
