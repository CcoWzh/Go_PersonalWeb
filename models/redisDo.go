package models

import (
	"fmt"
	"time"
	"github.com/garyburd/redigo/redis"
)

const (
	RedisURL            = "redis://127.0.0.1:6379"
	redisMaxIdle        = 3   //最大空闲连接数
	redisIdleTimeoutSec = 240 //最大空闲连接时间
	RedisPassword       = "test000"
)

// NewRedisPool 返回redis连接池
func NewRedisPool(redisURL string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			//验证redis密码
			if _, authErr := c.Do("AUTH", RedisPassword); authErr != nil {
				return nil, fmt.Errorf("redis auth password error: %s", authErr)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}

//string 键值对
func Set(k, v string) {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

//得到键值 set方法
func GetStringValue(k string) string {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	username, err := redis.String(c.Do("GET", k))
	if err != nil {
		fmt.Println("Get Error: ", err.Error())
		return ""
	}
	return username
}

//为哈希表中的字段赋值
//可以设置多值
//可追加，可修改
func Hset(myhash ,field ,v string)  {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	_, err := c.Do("HSET", myhash ,field ,v)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

/**
	用于为哈希表中不存在的的字段赋值 。
	如果哈希表不存在，一个新的哈希表被创建并进行 HSET 操作。
	如果字段已经存在于哈希表中，操作无效。
	如果 key 不存在，一个新哈希表被创建并执行 HSETNX 命令。
 */
func Hsetnx(myhash ,field ,v string)  {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	_, err := c.Do("HSET", myhash ,field ,v)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

//返回哈希表中指定字段的值
func Hget(myhash ,field string) string  {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	value, err := redis.String(c.Do("HGET", myhash,field))
	if err != nil {
		fmt.Println("Get Error: ", err.Error())
		return ""
	}
	return value
}

//用于删除哈希表 key 中的一个或多个指定字段，不存在的字段将被忽略。
//成功为1，不成功为0
func Hdel(myhash ,field string) string {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	value, err := redis.String(c.Do("HDEL", myhash,field))
	if err != nil {
		fmt.Println("Get Error: ", err.Error())
		return "0"
	}
	return value
}

