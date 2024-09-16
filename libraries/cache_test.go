package libraries

import (
	"testing"
	"time"
)

func TestAddCache(t *testing.T) {
	key := "测试键"
	value := "测试值"
	expiration := 0 * time.Second

	AddCache(key, value, expiration)

	if _, found := GetCache(key); !found {
		t.Errorf("期望找到键为 %s 的缓存，但未找到。", key)
	}
}

func TestDeleteCache(t *testing.T) {
	key := "测试键"
	value := "测试值"

	AddCache(key, value, 0)

	DeleteCache(key)

	if _, found := GetCache(key); found {
		t.Errorf("期望键为 %s 的缓存被删除，但仍然存在。", key)
	}
}

func TestUpdateCache(t *testing.T) {
	key := "测试键"
	value := "初始值"
	newValue := "更新值"
	expiration := 5 * time.Second

	AddCache(key, value, 0)
	UpdateCache(key, newValue, expiration)

	cachedValue, found := GetCache(key)
	if !found {
		t.Errorf("期望在更新后找到键为 %s 的缓存，但未找到。", key)
	} else if cachedValue != newValue {
		t.Errorf("期望缓存值为 %s，但实际为 %s。", newValue, cachedValue)
	}
}

func TestGetCache(t *testing.T) {
	key := "测试键"
	value := "测试值"

	AddCache(key, value, 0)

	cachedValue, found := GetCache(key)
	if !found {
		t.Errorf("期望找到键为 %s 的缓存，但未找到。", key)
	} else if cachedValue != value {
		t.Errorf("期望缓存值为 %s，但实际为 %s。", value, cachedValue)
	}
}
