package kernel

import (
	"github.com/ArtisanCloud/PowerLibs/cache"
	"time"
)

type InteractsWithCache struct {
	Cache cache.CacheInterface
}

func NewRedisClient(options *cache.RedisOptions) cache.CacheInterface {
	if options == nil {
		return nil
	}

	if options.Host == "" {
		return nil
	}

	return cache.NewGRedis(options)
}

func NewInteractsWithCache(client cache.CacheInterface) *InteractsWithCache {

	interact := &InteractsWithCache{
		Cache: client,
	}
	if client == nil {
		interact.Cache = interact.createDefaultCache()
	}

	return interact
}

func (interactCache *InteractsWithCache) GetCache() cache.CacheInterface {
	if interactCache.Cache != nil {
		return interactCache.Cache
	}

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
