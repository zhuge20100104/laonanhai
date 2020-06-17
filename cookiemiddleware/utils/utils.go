package utils

import (
	"sync"
)

var (
	// Data 全局的Data对象，处理Gin框架跳转时Context Keys丢失的问题
	Data *data
)

type data struct {
	Keys map[string]interface{}
	mu   sync.RWMutex
}

func init() {
	Data = new(data)
	Data.Keys = make(map[string]interface{}, 20)
}

// Set 设置Data的key和value值
func (d *data) Set(key string, value interface{}) {
	d.mu.Lock()
	Data.Keys[key] = value
	d.mu.Unlock()
}

// Get 获取Data的key值对应的value值
func (d *data) Get(key string) (ele interface{}, exists bool) {
	d.mu.RLock()
	ele, exists = Data.Keys[key]
	d.mu.RUnlock()
	return
}

// GetString 获取string类型的value值
func (d *data) GetString(key string) string {
	ele, exists := d.Get(key)
	if !exists {
		return ""
	}
	return ele.(string)
}

// GetInt 获取int类型的value值
func (d *data) GetInt(key string) int {
	ele, exists := d.Get(key)
	if !exists {
		return 0
	}
	return ele.(int)
}
