# 这是一个go重写的获取微信access_token的一个web服务项目
---
## 目录

- common:公共函数包 
    - db.go:初始化redis和mysql客户端连接
    - fun.go:公共函数
- conf:配置文件(git忽略，需要自行创建) 
    - app.conf:项目配置信息
    - chemm_mysql.json:MySQL信息
- configs：
    - config.go 
    - ini.go:初始化ini配置文件
    - json.go:初始化json配置文件
- controllers:控制器
    - errorcode.go:定义map报错码
    - json.go:返回自定义json数据结构处理
    - weixin.go:获取access_token
- routers
    - routes.go:项目路由
- main.go:入口文件
- access_token:go build生成的可执行二进制文件

## app.conf
```
[base]
host = "0.0.0.0"
port = 7091

[session]
host = "**********"
select = 2
authstring = ""

[log]
path = "/tmp/log"
distingtype = 1
level = 0
debug = false

[collect]
filepath = "/"
```

## chemm_mysql.json
```
[
  {
    "optionstype":"readwrite",
    "options":{
      "username":"root",
      "password":"*********",
      "host":"*********",
      "port":3306,
      "Name":"***",
      "maxopenconns":2000,
      "maxidleconns":200
    }
  }
]

```

## 生产环境
登陆测试环境后，*** 命令登陆生产机；

项目工作目录为 ***$GOPATH/src/wcjs/access_token*** ，监听 ***7091*** 端口;

项目更新记得进到工作目录，***go build*** 重新编译生成可执行文件 ***access_token*** ;

守护进程文件为 ***/etc/systemd/system/accesstoken.service*** ;

开启服务命令 ***systemctl start accesstoken*** ;

关闭服务命令 ***systemctl stop accesstoken*** ;

Rsync同步命令 ***rsync -avrtzopgPW --delete --exclude=conf --progress --port=873 
--password-file=/etc/rsync/rsync.pass  test@*::access_token /files/go/src/wcjs/access_token***

### accesstoken.service
```
[Unit]
Description=Go Web Service for Weixin Access Token
#After=network.target

[Service]
Type=simple
ExecStart=/files/go/src/wcjs/access_token/access_token
WorkingDirectory=/files/go/src/wcjs/access_token/
Restart=on-failure
RestartSec=5s
RestartPreventExitStatus=23

[Install]
WantedBy=multi-user.target
```

## 请求测试
服务只能在已构成局域网的测试和生产几台机器上请求
``` 
curl ***:7091/weixin/access_token/weixin/gzwcjs 

curl ***:7091/weixin/access_token/weixin/wx-wcjs 

```
