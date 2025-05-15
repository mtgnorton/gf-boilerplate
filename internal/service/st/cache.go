// package st 缓存单例模式
package st

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

var (
	jwtCache  *gcache.Cache
	cacheLock sync.RWMutex
)

// GetJWTCache 获取缓存实例
func GetJWTCache() *gcache.Cache {
	cacheLock.RLock()
	defer cacheLock.RUnlock()

	if jwtCache == nil {
		panic("cache not initialized")
	}
	return jwtCache
}

func MustInitCacheFromConfig(ctx context.Context) {
	var adapter gcache.Adapter

	cacheAdapter := g.Cfg().MustGet(ctx, "jwt.cache.adapter").String()
	switch cacheAdapter {
	case "redis":
		adapter = gcache.NewAdapterRedis(g.Redis())
	default:
		adapter = gcache.NewAdapterMemory()
	}

	// 数据库缓存，默认和通用缓冲驱动一致，如果你不想使用默认的，可以自行调整
	g.DB().GetCache().SetAdapter(adapter)
	cacheLock.Lock()
	defer cacheLock.Unlock()
	jwtCache = gcache.New()
	jwtCache.SetAdapter(adapter)
}
