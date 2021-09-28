package common

import (
	"github.com/gomodule/redigo/redis"
	"github/siafei/gin-test/global"
)

func RedisGetCon() redis.Conn  {
	return global.Redis.Get()
}

