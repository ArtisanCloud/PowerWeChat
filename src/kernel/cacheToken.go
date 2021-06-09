package kernel

import "github.com/ArtisanCloud/go-libs/cache"

type CacheToken struct {
	Cache *cache.CacheInterface
}

func (cacheToken *CacheToken) getCache() *cache.CacheInterface {
	if cacheToken.Cache !=nil{
		return cacheToken.Cache
	}

	cacheToken.Cache = cacheToken.createDefaultCache()

	return cacheToken.Cache
}

func (cacheToken *CacheToken) setCache(cache *cache.CacheInterface){
	cacheToken.Cache = cache
}

func (cacheToken *CacheToken) createDefaultCache() *cache.CacheInterface {

	return cacheToken.Cache
}
