package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}

// CounterHandler 计数器接口
func LawyerHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		body := make(map[string]interface{})
		if err := decoder.Decode(&body); err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			fmt.Println(body)
			resp := ResponseMsg{
				ToUserName:   fmt.Sprint(body["FromUserName"]),
				FromUserName: fmt.Sprint(body["ToUserName"]),
				CreateTime:   time.Now().Unix(),
				MsgType:      fmt.Sprint(body["MsgType"]),
				Content:      fmt.Sprint(body["Content"]),
			}
			res.Data = resp
		}

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res.Data)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
