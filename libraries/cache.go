package libraries

/*
基于go-cache 写以下方法. 使用中文注释. 内容判断详细.
1. 添加缓存.  字段有 key, value, 有效期(默认0 为永久有效).
2. 根据key 删除缓存.
3. 修改缓存内容. 同时可以修改有效期
4. 根据key 获得缓存内容
*/

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var c = cache.New(5*time.Minute, 10*time.Minute) // 创建一个新的缓存实例，过期时间为5分钟，清理时间为10分钟

// AddCache 添加缓存，key为缓存的键，value为缓存的值，expiration为有效期，默认0为永久有效
func AddCache(key string, value interface{}, expiration time.Duration) {
	if expiration == 0 {
		c.Set(key, value, cache.DefaultExpiration) // 如果有效期为0，使用默认有效期
	} else {
		c.Set(key, value, expiration) // 设置指定有效期的缓存
	}
}

// DeleteCache 根据key删除缓存
func DeleteCache(key string) {
	c.Delete(key) // 删除指定key的缓存
}

// UpdateCache 修改缓存内容，同时可以修改有效期
func UpdateCache(key string, value interface{}, expiration time.Duration) {
	if expiration == 0 {
		c.Set(key, value, cache.DefaultExpiration) // 如果有效期为0，使用默认有效期
	} else {
		c.Set(key, value, expiration) // 设置指定有效期的缓存
	}
}

// GetCache 根据key获得缓存内容
func GetCache(key string) (interface{}, bool) {
	return c.Get(key) // 返回缓存内容和是否存在的布尔值
}
