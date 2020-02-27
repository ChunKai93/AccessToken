package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
	"time"
	"wcjs/access_token/common"
	"wcjs/lib/session"
)

const API_URL_PREFIX = "https://api.weixin.qq.com/cgi-bin"
const AUTH_URL = "/token?grant_type=client_credential&"

type WeixinController struct {
	Mu *sync.RWMutex
	JsonController
}

type WeixinStruct struct {
	errorCode   int64  `json:"errorcode"`
	errMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func NewWeixinController() *WeixinController {
	ac := &WeixinController{
		&sync.RWMutex{},
		JsonController{
			Sm:       session.NewSession(),
			HostName: "angelandy_pc",
		},
	}
	return ac
}

func (weixin *WeixinController) GetToken(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	wx := p.ByName("weixin")
	d, err := common.ChemmMysqlclient.GetRead().QueryOne("select * from sp_weixin_account where weixin='%s'", wx)
	if err != nil {
		panic(err)
	}
	appid := d["appid"]
	appsecret := d["appsecret"]

	//读redis
	weixin.Mu.RLock()
	cache_name := "CMMAPI_wechat_access_token" + appid
	access_token, _ := common.RedisDB.Get(cache_name).Result()
	weixin.Mu.RUnlock()

	//写redis
	if access_token == "" {
		weixin.Mu.Lock()
		url := API_URL_PREFIX + AUTH_URL + "appid=" + appid + "&secret=" + appsecret
		res := common.Get(url)
		ws := WeixinStruct{}
		err = json.Unmarshal(res, &ws)
		if err != nil {
			panic(err)
		}
		if ws.errorCode == 0 {
			access_token = ws.AccessToken

			expire := ws.ExpiresIn - 1000
			common.RedisDB.Set(cache_name, ws.AccessToken, time.Duration(expire) * time.Second)
		}
		weixin.Mu.Unlock()
	}

	fmt.Fprintf(w, access_token)
}
