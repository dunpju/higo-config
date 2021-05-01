package config

import (
	"sync"
)

const (
	EnvPath      = "env"
	ServePath    = "serve"
	AppPath      = "app"
	DatabasePath = "database"
	AuthPath     = "auth"
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
