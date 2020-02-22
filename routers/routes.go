package routers

import (
	"github.com/julienschmidt/httprouter"
	"wcjs/access_token/controllers"
)

/*
命名规范：
路由：小驼峰
方法：大驼峰
静态文件：全小写
*/
func Getrouter() *httprouter.Router {
	wx := controllers.NewWeixinController()

	router := httprouter.New()
	router.GET("/weixin/access_token/weixin/:weixin", wx.GetToken)

	return router
}
