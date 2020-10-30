/*
@Time : 2020/10/29 4:59 下午
@Author : mmy83
@File : envconfig_test.go
@Software: GoLand
*/

package envconfig

import (
	"github.com/wonderivan/logger"
	"testing"
)


type config struct {
	Key1 string `envconfig:"ENV1"`
	Key2 int32 `envconfig:"ENV2"`
	Key3 string `envconfig:"ENV3"`
	Key4 bool `envconfig:"ENV4"`
	Key5 string `envconfig:"ENV5"`
	Key6 string `envconfig:"ENV6"`
	Key7 string `envconfig:"ENV7"`
}


func TestGetConfig(t *testing.T) {
	obj := new(config)
	GetConfig(obj)
	logger.Debug(obj)
}