package kernel

import (
	"github.com/ArtisanCloud/PowerLibs/cache"
	"time"
)

type InteractsWithCache struct {
	Cache cache.CacheInterface
}

func (interactCache *InteractsWithCache) GetCache() cache.CacheInterface {
	if interactCache.Cache != nil {
		return interactCache.Cache
	}

	// tbd global redis exist

	// create default cache
	interactCache.Cache = interactCache.createDefaultCache()

	return interactCache.Cache
}

func (interactCache *InteractsWithCache) setCache(cache cache.CacheInterface) *InteractsWithCache {

	interactCache.Cache = cache

	return interactCache
}

func (interactCache *InteractsWithCache) createDefaultCache() cache.CacheInterface {

	interactCache.Cache = cache.NewMemCache("ac.go.power", time.Duration(1500), "")

	return interactCache.Cache
}
