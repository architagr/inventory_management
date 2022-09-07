package cache

const (
	LOCAL_CACHE = iota
	REDIS_CACHE
)

type ICache interface {
	getValue(keyName string) interface{}
}

var _ ICache = new(LocalCache)

func ProvideLocalCache() *LocalCache {
	return &LocalCache{}
}
