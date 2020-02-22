package common

import (
	"github.com/go-redis/redis"
	//"wcjs/access_token/configs"
	"wcjs/lib/client/mysql"
)

type CollectConfig struct {
	FilePath         string
	ChemmMysqlConfig []*mysql.MultiMysqlOptions
}

type SessionConfig struct {
	Host       string
	Select     int
	Authstring string
}

var ChemmMysqlclient *mysql.MultiMysqlClient
var RedisDB *redis.Client

func CmInit(config *CollectConfig) {
	ChemmMysqlclient = mysql.NewMultiMsqlClient(config.ChemmMysqlConfig)
}

func RdsInit(config *SessionConfig) {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     config.Host,
		DB:       config.Select,
		Password: config.Authstring,
	})
}
