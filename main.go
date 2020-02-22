package main

import (
	"fmt"
	"net/http"
	"wcjs/access_token/common"
	"wcjs/access_token/configs"
	"wcjs/access_token/routers"
	"wcjs/lib/log"
)

func main() {

	baseconfig := configs.IniBaseGet()
	logconfig := configs.IniLogGet()
	collectconfig := configs.IniCollectGet()
	redisconfig := configs.IniSessionGet()

	log.LogRun(logconfig)
	common.CmInit(collectconfig)
	common.RdsInit(redisconfig)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", baseconfig.Host, baseconfig.Port), routers.Getrouter())
	if err != nil {
		fmt.Printf("Can't start the server: %s", err)
	}

}
