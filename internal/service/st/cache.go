package st

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

var cache *gcache.Cache

func GetCache() *gcache.Cache {
	if cache == nil {
		panic("cache not init")
	}
	return cache
}

func MustInitCacheFromConfig(ctx context.Context) {
	var adapter gcache.Adapter

	cacheAdapter := g.Cfg().MustGet(ctx, "cache.adapter").String()
	switch cacheAdapter {
	case "redis":
		adapter = gcache.NewAdapterRedis(g.Redis())
	default:
		adapter = gcache.NewAdapterMemory()
	}

	// 数据库缓存，默认和通用缓冲驱动一致，如果你不想使用默认的，可以自行调整
	g.DB().GetCache().SetAdapter(adapter)
	cache = gcache.New()
	cache.SetAdapter(adapter)
}
