package container

import "time"

// IContainer 容器相关方法
type IContainer interface {
	Get(key string) any
	Set(key string, value any) bool
	Has(key string) bool
}

// ICache config cache
type ICache interface {
	IContainer
	FuzzyDelete(key string)
}

// IConfig 缓存接口
// 应用到redis、cache等cache db
type IConfig interface {
	Get(key string) any
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetDuration(key string) time.Duration
	GetStringSlice(key string) []string
}

type IDriver interface {
	IContainer
	Listen()
}
