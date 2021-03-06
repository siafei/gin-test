package redis

import (
	"github.com/gomodule/redigo/redis"
	"github/siafei/gin-test/pkg/setting"
)

func NewRedis(redisSetting *setting.RedisSettings) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			c, err := redis.Dial("tcp", redisSetting.Host+":"+redisSetting.Port)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", redisSetting.Password); err != nil {
				_ = c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT", redisSetting.DB); err != nil {
				_ = c.Close()
				return nil, err
			}
			return c, nil
		},
	}
	return pool
}


