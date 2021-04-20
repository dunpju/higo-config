package config

import "strings"

type Configure map[interface{}]interface{}

type CacheConfigure map[string]interface{}

func New() Configure {
	c := make(Configure)
	return c
}

func Set(path string, value interface{}) {
	paths := strings.Split(path, ".")
	current := container
	for level, key := range paths {
		current = current.set(key, level, len(paths), value).(Configure)
	}
}

func (this Configure) set(key string, level, length int, value interface{}) interface{} {
	if level < length-1 {
		if exist, ok := this[key]; ok {
			this = exist.(Configure)
		} else {
			this[key] = New()
			this = this[key].(Configure)
		}
	} else {
		this[key] = value
	}
	return this
}

func String(key string) string {
	configure := container.Get(key)
	if nil == configure {
		return ""
	}
	return configure.(string)
}

func Int(key string) int {
	configure := container.Get(key)
	if nil == configure {
		return 0
	}
	return configure.(int)
}

// 外部获取配置
func Get(key string) interface{} {
	config := container.Get(key)
	if nil == config {
		return nil
	}
	return config
}

// 外部获取所有配置
func All() Configure {
	return container.All()
}

func parse(path string) interface{} {
	paths := strings.Split(path, ".")
	var current interface{}
	previous := make([]string, 0)
	current = container
	for _, key := range paths {
		previous = append(previous, key)
		if cur, ok := current.(Configure); ok {
			current = get(cur, key, previous)
		}
	}
	return current
}

func get(current Configure, key string, previous []string) interface{} {
	v, ok := current[key]
	if !ok {
		panic(strings.Join(previous, ".") + " non-existent")
	}
	return v
}

// 设置
func (this Configure) Set(key string, value interface{}) Configure {
	this[key] = value
	return this
}

// 获取所有配置
func (this Configure) All() Configure {
	return this
}

// 获取配置
func (this Configure) Get(key string) interface{} {
	paths := strings.Split(key, ".")
	if len(paths) > 1 {
		return parse(key)
	}
	if value, ok := this[key]; ok {
		return value
	} else {
		panic(key + " non-existent")
	}
}

// 第一个字符串元素
func FirstString() string {
	var first string
	for _, v := range container {
		first = v.(string)
		break
	}
	return first
}

// 第一个数字元素
func FirstInt() int {
	var first int
	for _, v := range container {
		first = v.(int)
		break
	}
	return first
}
