/************************************************************
 * Author:        jackey
 * Date:        2022/11/25
 * Description:
 * Version:    V1.0.0
 **********************************************************/

package common

import (
	"fmt"
	"github.com/go-redis/redis"
	"liteIm/pkg/config"
	"strconv"
	"time"
)

var (
	RedisClient *redis.Client
)

func InitRedis() {
	getRedisClient()
}

type redisConn struct {
	Host     string
	Password string
	DB       int
}

func (r *redisConn) connect() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         r.Host,          // Redis地址
		Password:     r.Password,      // Redis账号
		DB:           r.DB,            // Redis库
		PoolSize:     50,              // Redis连接池大小
		MaxRetries:   3,               // 最大重试次数
		IdleTimeout:  5 * time.Second, // 空闲链接超时时间
		MinIdleConns: 20,              // 空闲连接数量
	})
	pong, err := client.Ping().Result()
	if err == redis.Nil {
		fmt.Println("Redis 链接池异常-db", r.Host, r.DB)
	} else if err != nil {
		fmt.Println("redis 链接池失败:", err, r.Host, r.DB)
	} else {
		fmt.Println("redis " + pong + " " + r.Host + " db" + strconv.Itoa(r.DB))
	}
	return client, nil
}

func getRedisClient() {
	var err error
	addr := config.Config.Section(ConfigSectionRedisClient).Key("addr").String()
	password := config.Config.Section(ConfigSectionRedisClient).Key("password").String()
	db, _ := config.Config.Section(ConfigSectionRedisClient).Key("db").Int()
	redisConfig := &redisConn{
		Host:     addr,
		Password: password,
		DB:       db,
	}
	RedisClient, err = redisConfig.connect()
	if err != nil {
		fmt.Println("getRedisClient", err)
	}
}
