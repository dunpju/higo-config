package config

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MS map[string]interface{}

type Configure struct {
	Sort  []string
	Value MS
}

func New() *Configure {
	return &Configure{Sort: make([]string, 0), Value: make(MS)}
}

func (this *Configure) set(key string, level, length int, value interface{}) *Configure {
	if level < length-1 {
		exist, ok := this.Value[key]
		if configure, ook := exist.(*Configure); ook && ok {
			this = configure
		} else {
			if !ok {
				this.Sort = append(this.Sort, key)
			}
			this.Value[key] = New()
			this = this.Value[key].(*Configure)
		}
	} else {
		if _, ok := this.Value[key]; !ok {
			this.Sort = append(this.Sort, key)
		}
		this.Value[key] = value
	}
	return this
}

func (this *Configure) Sprint() {
	fmt.Println(this.String())
}

func (this *Configure) String() string {
	jsonByte, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}
	return string(jsonByte)
}

func convert(mi map[interface{}]interface{}) *Configure {
	conf := New()
	for k, v := range mi {
		if m, ok := v.(map[interface{}]interface{}); ok {
			conf.Set(k.(string), convert(m))
		} else if m, ok := v.(map[string]interface{}); ok {
			conf.Set(k.(string), convertString(m))
		} else {
			conf.Set(k.(string), v)
		}
	}
	return conf
}

func convertString(mi map[string]interface{}) *Configure {
	conf := New()
	for k, v := range mi {
		if m, ok := v.(map[string]interface{}); ok {
			conf.Set(k, convertString(m))
		} else if m, ok := v.(map[interface{}]interface{}); ok {
			conf.Set(k, convert(m))
		} else {
			conf.Set(k, v)
		}
	}
	return conf
}

// Set 设置
func (this *Configure) Set(key string, value interface{}) *Configure {
	if m, ok := value.(map[interface{}]interface{}); ok {
		value = convert(m)
	} else if m, ok := value.(map[string]interface{}); ok {
		value = convertString(m)
	}

	paths := strings.Split(key, ".")
	current := this
	for level, key := range paths {
		current = current.set(key, level, len(paths), value)
	}
	return this
}

// All 获取所有配置
func (this *Configure) All() *Configure {
	return this
}

// Get 获取配置
func (this *Configure) Get(key string) interface{} {
	paths := strings.Split(key, ".")
	if len(paths) > 1 {
		return parse(this, key)
	}
	if value, ok := this.Value[key]; ok {
		return value
	} else {
		panic(key + " non-existent")
	}
}

func (this *Configure) Remove(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) > 1 {
		previous := make([]string, 0)
		var prev *Configure
		current := this
		currKey := ""
		for _, key := range keys {
			currKey = key
			previous = append(previous, currKey)
			if _, ok := current.Value[currKey]; ok {
				if cur, ok := current.Value[currKey].(*Configure); ok {
					prev = current
					current = cur
				} else {
					panic(strings.Join(previous, ".") + " isn't *Configure type")
				}
			} else {
				return false
			}
		}
		return remove(prev, currKey)
	}
	return remove(this, key)
}

func (this *Configure) Exist(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) > 1 {
		previous := make([]string, 0)
		current := this
		currKey := ""
		for i, key := range keys {
			currKey = key
			previous = append(previous, key)
			if _, ok := current.Value[key]; ok {
				if i < len(keys)-1 {
					if cur, ok := current.Value[key].(*Configure); ok {
						current = cur
					} else {
						panic(strings.Join(previous, ".") + " isn't *Configure type")
					}
				}
			} else {
				return false
			}
		}
		key = currKey
		this = current
	}
	return exist(this, key)
}

func (this *Configure) Empty() bool {
	if len(this.Value) == 0 {
		return true
	}
	return false
}

func (this *Configure) Len() int {
	return len(this.Value)
}

// 第一个字符串元素
func (this *Configure) FirstString() string {
	key := this.Sort[0]
	return this.Value[key].(string)
}

// 最后字符串元素
func (this *Configure) EndString() string {
	key := this.Sort[len(this.Sort)-1]
	return this.Value[key].(string)
}

// 第一个数字元素
func (this *Configure) FirstInt() int {
	key := this.Sort[0]
	return this.Value[key].(int)
}

// 最后数字元素
func (this *Configure) EndInt() int {
	key := this.Sort[len(this.Sort)-1]
	return this.Value[key].(int)
}

// 获取索引字符串元素
func (this *Configure) IndexString(index int) string {
	key := this.Sort[index]
	return this.Value[key].(string)
}

// 获取索引字符串元素
func (this *Configure) IndexInt(index int) int {
	key := this.Sort[index]
	return this.Value[key].(int)
}
