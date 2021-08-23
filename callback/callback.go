package callback

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"qiniu/httptool"
	"time"
	qlog "qiniu/log"
)

type Callback struct {}

type callbackItem struct {
	time 			string
	source          string
	remoteip 		string
	requestHeader   string
	requestBody     string
}

func (c *Callback) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Println("only support POST")
		http.NotFound(w, r)
		return
	}

	url := r.URL
	switch url.Path {
	case "/callback/kodo":
		handleKodoCallback(w, r)
	case "/callback/qcdn":
		handleQcdnCallback(w, r)
	case "/callback/dora":
		handleDoraCallback(w, r)
	case "/callback/pili":
		handlePiliCallback(w, r)
	default:
		fmt.Println("no such service")
		http.NotFound(w, r)
	}
}

func handleKodoCallback(w http.ResponseWriter, r *http.Request) {
	handleCallback("KODO", w, r)
}

func handleQcdnCallback(w http.ResponseWriter, r *http.Request) {
	handleCallback("QCDN", w, r)
}

func handleDoraCallback(w http.ResponseWriter, r *http.Request) {
	handleCallback("DORA", w, r)
}

func handlePiliCallback(w http.ResponseWriter, r *http.Request) {
	handleCallback("PILI", w, r)
}

func handleCallback(source string, w http.ResponseWriter, r *http.Request) {
	// 由于中间件要对request.body再次读取，所以这里读取一次之后，将bytes再copy回去
	requestBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		qlog.AppLog.Println("request body read failed:", err)
		return
	}
	r.Body = ioutil.NopCloser(bytes.NewReader(requestBodyBytes))

	header, err := json.Marshal(r.Header)
	if err != nil {
		qlog.AppLog.Println("header decode failed:", err)
		return
	}

	responseBody := map[string]string {
		"ReqId":httptool.GetReqId(),
	}
	responseBodyJson, err := json.Marshal(responseBody)
	if err != nil {
		qlog.AppLog.Println("response body encode failed:", err)
	}

	// 返回response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBodyJson)

	// 保存数据库
	callback := callbackItem{
		time: time.Now().Format("2006-01-02 15:04:05"),
		source: source,
		remoteip: r.RemoteAddr,
		requestHeader: string(header),
		requestBody: string(requestBodyBytes),
	}
	insertCallback(callback)
}
