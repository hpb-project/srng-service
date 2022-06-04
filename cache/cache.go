package cache

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	red "github.com/hpb-project/srng-service/cache/redis"
	"time"
)

var Redis red.IRedis
var pool *redis.Pool

func init() {
	conn := beego.AppConfig.String("cache::conn")
	dbNum := beego.AppConfig.String("cache::dbNum")
	password := beego.AppConfig.String("cache::password")
	if err := NewPool(conn, dbNum, password); err != nil {
		logs.Error(err)
		return
	}
	Redis = red.NewStoreRedis(pool)
}

func NewPool(conn, dbNum, password string) error {
	fmt.Println("redis连接池里的连接为空,重新创建连接池,starting...")
	pool = &redis.Pool{
		MaxIdle:     50, //最大空闲连接数
		MaxActive:   0,  //若为0，则活跃数没有限制
		Wait:        true,
		IdleTimeout: 30 * time.Second, //最大空闲连接时间
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conn)
			if err != nil {
				panic(err)
				return nil, err
			}
			// 设置密码
			if _, err := c.Do("AUTH", password); err != nil {
				panic(err)
				return nil, err
			}
			// 选择db
			c.Do("SELECT", dbNum)
			return c, nil
		},
	}
	return nil
}
