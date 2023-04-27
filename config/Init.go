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
	AnnoConf     = "anno"
)

var (
	container   *Configure
	onceConfig  sync.Once
	EnvPrefix   string
	ServePrefix string
	AppPrefix   string
	DbPrefix    string
	AuthPrefix  string
	AnnoPrefix  string
)

func init() {
	onceConfig.Do(func() {
		container = New()
	})
}
