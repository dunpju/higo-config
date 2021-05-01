package config

import (
	"sync"
)

const (
	EnvConf      = "env"
	ServeConf    = "serve"
	AppConf      = "app"
	DatabaseConf = "database"
	AuthConf     = "auth"
)

var (
	container  *Configure
	onceConfig sync.Once
)

func init() {
	onceConfig.Do(func() {
		container = New()
	})
}
