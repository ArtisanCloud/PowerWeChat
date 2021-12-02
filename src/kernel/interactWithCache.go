package kernel

import (
	"github.com/ArtisanCloud/PowerLibs/cache"
	"time"
)

type InteractsWithCache struct {
	Cache CacheInterface
}

type CacheInterface cache.CacheInterface
type RedisOptions cache.RedisOptions

func NewRedisClient(options *RedisOptions) CacheInterface {
	if options == nil {
		return nil
	}

	if options.Host == "" {
		return nil
	}

	return cache.NewGRedis(options)
}

func NewInteractsWithCache(client CacheInterface) *InteractsWithCache {

	interact := &InteractsWithCache{
		Cache: client,
	}
	if client == nil {
		interact.Cache = interact.createDefaultCache()
	}

	return interact
}

func (interactCache *InteractsWithCache) GetCache() CacheInterface {
	if interactCache.Cache != nil {
		return interactCache.Cache
	}

	// create default cache
	interactCache.Cache = interactCache.createDefaultCache()

	return interactCache.Cache
}

func (interactCache *InteractsWithCache) setCache(cache CacheInterface) *InteractsWithCache {

	interactCache.Cache = cache

	return interactCache
}

func (interactCache *InteractsWithCache) createDefaultCache() CacheInterface {

	interactCache.Cache = cache.NewMemCache("ac.go.power", time.Duration(1500), "")

	return interactCache.Cache
}
