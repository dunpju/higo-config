package config

import "sync"

var (
	container      Configure
	onceConfig     sync.Once
)

func init() {
	onceConfig.Do(func() {
		container = make(Configure)
	})
}
