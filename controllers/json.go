package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"wcjs/lib/log"
	"wcjs/lib/session"
)

type ResultJson struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

type JsonController struct {
	Sm        *session.SessionManager
	HostName  string
}

func NewJsonController() *JsonController {
	jc := &JsonController{
		Sm:        session.NewSession(),
		HostName:  "angelandy_pc",
	}
	return jc
}

func (c *JsonController) request_after(w http.ResponseWriter, r *http.Request) {
	if e := recover(); e != nil {
		var msg, logmsg string
		var code int

		switch err := e.(type) {
		case error:
			code = 500
			msg = fmt.Sprint(err)
			logmsg = fmt.Sprint(err)
		case string:
			code = 500
			msg = err
			logmsg = err
		case int:
			code = err
			if sarr, ok := ErrorCode[code]; ok {
				msg = sarr[0]
				logmsg = sarr[1]
			} else {
				msg = fmt.Sprint("服务器繁忙,请稍后在试")
				logmsg = fmt.Sprintf("未定义服务器错误:%d", code)
			}
		case int64:
			code = int(err)
			if sarr, ok := ErrorCode[code]; ok {
				msg = sarr[0]
				logmsg = sarr[1]
			} else {
				msg = fmt.Sprint("服务器繁忙,请稍后在试")
				logmsg = fmt.Sprintf("未定义服务器错误:%d", code)
			}
		case int32:
			code = int(err)
			if sarr, ok := ErrorCode[code]; ok {
				msg = sarr[0]
				logmsg = sarr[1]
			} else {
				msg = fmt.Sprint("服务器繁忙,请稍后在试")
				logmsg = fmt.Sprintf("未定义服务器错误:%d", code)
			}
		default:
			code = 500
			msg = fmt.Sprint(err)
			logmsg = fmt.Sprint(err)
		}

		re := &ResultJson{
			Code: code,
			Msg:  msg,
			Data: map[string]interface{}{},
		}

		log.FileError(logmsg)
		fulluri := r.RequestURI
		uri_arr := strings.Split(r.RequestURI, "?")
		urilen := len(uri_arr)
		var uri string
		if urilen != 1 {
			uri = uri_arr[0]
		} else {
			uri = fulluri
		}
		log.FileError(logmsg, c.HostName, uri, fulluri)
		jstr, _ := json.Marshal(re)
		fmt.Fprintf(w, string(jstr))
	}
}

func (c *JsonController) Display(w http.ResponseWriter,code int,d map[string]interface{}) {
	var msg string = ""
	if sarr, ok := ErrorCode[code]; ok {
		msg = sarr[0]
	} else {
		msg = "服务器繁忙,请稍后在试"
	}
	re := &ResultJson{
		Code: code,
		Msg:  msg,
		Data: d,
	}

	jstr, _ := json.Marshal(re)
	fmt.Fprintf(w, string(jstr))
}