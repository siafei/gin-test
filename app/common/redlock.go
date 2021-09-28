package common

import (
	"github.com/gomodule/redigo/redis"
)

const (
	// key, field, num
	RED_LOCK_LUA = `
if redis.call('get',KEYS[1]) == ARGV[1] then 
	return redis.call('del',KEYS[1]) 
else
   return 0 
end
`
)

type RedLock struct {
	Key string
	Value string
}

/**
加锁
 */
func (l RedLock) Lock() bool  {
	redisCon := RedisGetCon()
	defer redisCon.Close()
	_, err := redis.String(redisCon.Do("SET", l.Key, l.Value, "EX", 5, "NX"))
	if err != nil {
		return false
	}
	return true
}

/**
解锁
 */
func (l RedLock) Unlock() bool {
	redisCon := RedisGetCon()
	script := redis.NewScript(1,RED_LOCK_LUA)
	ok, err := redis.Int(script.Do(redisCon, l.Key, l.Value))
	if err != nil || ok == 0 {
		return false
	}
	return true
}
