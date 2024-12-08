package data

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var Cache *redis.Pool

func InitRedis(cacheHost, cachePort string) error {
	Cache = &redis.Pool{
		MaxIdle:     20,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cacheHost+":"+cachePort)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Wait: true,
	}

	return nil
}

func CacheGet(key string) ([]byte, error) {
	c := Cache.Get()

	v, err := c.Do("GET", key)
	if err != nil || v == nil {
		return nil, err
	}
	defer c.Close()
	return v.([]byte), err
}

func CacheSet(key string, value []byte) error {
	c := Cache.Get()
	_, err := c.Do("SET", key, value)
	defer c.Close()
	return err
}

func CacheSetExp(key string, value []byte, expireSeconds int) error {
	c := Cache.Get()
	_, err := c.Do("SETEX", key, expireSeconds, value)
	defer c.Close()
	return err
}

func CacheExists(key string) (bool, error) {
	c := Cache.Get()
	ext, err := redis.Bool(c.Do("EXISTS", key))
	defer c.Close()
	return ext, err
}
