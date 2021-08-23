package main

import (
	"fmt"
	"net/http"
	"qiniu/callback"
	"qiniu/db"
	"qiniu/httptool"
	qlog "qiniu/log"
	"qiniu/setting"
)

func init() {
	setting.InitSetting()
	qlog.InitLogger()
}

func main() {

	// 链接数据库
	if err := db.ConnectMysql(); err != nil {
		qlog.AppLog.Fatalln("connect mysql failed:", err)
	}

	// 创建表
	if err := callback.CreateTable(); err != nil {
		qlog.AppLog.Fatalln("create table 'callbacks' failed:", err)
	}

	// 服务中间件
	router := httptool.NewRouter()
	router.Use(&httptool.Mid_timer{})
	router.Use(&httptool.Mid_logger{})
	router.Add("/callback/", &callback.Callback{})
	router.Add("/", http.HandlerFunc(index))

	// 监听http服务
	mux := http.NewServeMux()

	// router绑定到mux
	router.BindMux(mux)

	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", setting.Setting().Server.HttpPort),
		Handler:      mux,
		ReadTimeout:  setting.Setting().Server.ReadTimeout,
		WriteTimeout: setting.Setting().Server.WriteTimeout,
		ErrorLog:     qlog.HttpError,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello!")
}

