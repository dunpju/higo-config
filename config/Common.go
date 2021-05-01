package config

import "strings"

func Set(path string, value interface{}) {
	container.Set(path, value)
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

// 获取配置
func Get(key string) interface{} {
	value := container.Get(key)
	if nil == value {
		return nil
	}
	return value
}

// Env配置
func Env(key string) interface{} {
	return Get(EnvPath + "." + key)
}

// Serve配置
func Serve(key string) interface{} {
	return Get(ServePath + "." + key)
}

// App配置
func App(key string) interface{} {
	return Get(AppPath + "." + key)
}

// Db配置
func Db(key string) interface{} {
	return Get(DatabasePath + "." + key)
}

// Auth配置
func Auth(key string) interface{} {
	return Get(AuthPath + "." + key)
}

// 所有配置
func All() *Configure {
	return container.All()
}

func parse(path string) interface{} {
	keys := strings.Split(path, ".")
	previous := make([]string, 0)
	var current interface{}
	current = container
	for _, key := range keys {
		previous = append(previous, key)
		if cur, ok := current.(*Configure); ok {
			current = get(cur, key, previous)
		}
	}
	return current
}

func get(this *Configure, key string, previous []string) interface{} {
	v, ok := this.Value[key]
	if !ok {
		panic(strings.Join(previous, ".") + " non-existent")
	}
	return v
}

func remove(this *Configure, key string) bool {
	if _, ok := this.Value[key]; ok {
		var tmp []string
		for _, value := range this.Sort {
			if value != key {
				tmp = append(tmp, value)
			}
		}
		delete(this.Value, key)
		this.Sort = tmp
	} else {
		return false
	}
	return true
}

func exist(this *Configure, key string) bool {
	if _, ok := this.Value[key]; ok {
		return true
	}
	return false
}
