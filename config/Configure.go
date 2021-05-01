package config

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Configure struct {
	Sort  []string
	Value map[string]interface{}
}

func New() *Configure {
	return &Configure{Sort: make([]string, 0), Value: make(map[string]interface{})}
}

func (this *Configure) set(key string, level, length int, value interface{}) *Configure {
	if level < length-1 {
		exist, ok := this.Value[key];
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

// 设置
func (this *Configure) Set(key string, value interface{}) *Configure {
	paths := strings.Split(key, ".")
	current := this
	for level, key := range paths {
		current = current.set(key, level, len(paths), value)
	}
	return this
}

// 获取所有配置
func (this *Configure) All() *Configure {
	return this
}

// 获取配置
func (this *Configure) Get(key string) interface{} {
	paths := strings.Split(key, ".")
	if len(paths) > 1 {
		return parse(key)
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
