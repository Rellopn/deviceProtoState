/**
 * @Author: Rellopn
 * @Description:
 * @File:  option_test.go
 * @Version: 0.1.1
 * @Date: 2020/7/3 10:12
 */
package deviceProtoState

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	host = "12.0.0.1"
	port = "6379"
	pwd  = "123456"
	db   = 0
)

func TestOption_RedisStateCfg(t *testing.T) {
	option := Option{}
	option.RedisStateCfg(host, port, pwd)
	assert.Equal(t, host, option.Host)
	assert.Equal(t, port, option.Port)
	assert.Equal(t, pwd, option.Pwd)
	assert.Equal(t, db, option.Db)
}
