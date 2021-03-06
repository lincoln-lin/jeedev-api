package myredis

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

// Conn return redis connection.
func Conn() redis.Conn {
	rs := pool.Get()
	rs.Do("SELECT", beego.AppConfig.DefaultInt("cache::dbno",0))
	return rs
}


func Close() {
	pool.Close()
}


func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				_, err := c.Do("AUTH", password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	server := beego.AppConfig.String("cache::server")
	password := beego.AppConfig.String("cache::password")

	pool = newPool(server, password)
}
func GetString(key string) string{
	rs := Conn()
	n, _ := redis.String(rs.Do("GET", key))
	return n
}

func SetString(key string, value string, expire string) bool{
	rs := Conn()
	_, err := rs.Do("SET", key,  value , "EX", expire)
	if err != nil {
		return false
	}else{
		return true
	}
}