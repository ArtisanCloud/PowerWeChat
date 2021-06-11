package kernel

import (
	"github.com/ArtisanCloud/go-libs/cache"
	"time"
)

type CacheToken struct {
	Cache cache.CacheInterface
}

func (cacheToken *CacheToken) getCache() cache.CacheInterface {
	if cacheToken.Cache != nil {
		return cacheToken.Cache
	}

	// tbd global redis exist

	// create default cache
	cacheToken.Cache = cacheToken.createDefaultCache()

	return cacheToken.Cache
}

func (cacheToken *CacheToken) setCache(cache cache.CacheInterface) *CacheToken {

	cacheToken.Cache = cache

	return cacheToken
}

func (cacheToken *CacheToken) createDefaultCache() cache.CacheInterface {

	cacheToken.Cache = cache.NewMemCache("ac.go.wx", time.Duration(1500), "")

	return cacheToken.Cache
}
