package server

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego/context"
)

type CallBack struct {
}

type ReqBody struct {
	Code int64 `json:"code"`
}
type ResSuccess struct {
	Code int64 `json:"code"`
}
type ResFailed struct {
	Code int64 `json:"code"`
}

func ResponseSuccess() []byte {
	out := ResSuccess{
		Code: 0,
	}
	res, _ := json.Marshal(out)
	return res
}
func ResponseFailed() []byte {
	out := ResFailed{
		Code: 404,
	}
	res, _ := json.Marshal(out)
	return res
}

func CallBackHandler(ctx *context.Context) []byte {
	log.Println("callback: ", ctx.Input.RequestBody)
	var req ReqBody
	if json.Unmarshal(ctx.Input.RequestBody, req) != nil {
		log.Println("Failed to json unmarshal.")
		return ResponseFailed()
	}
	return ResponseSuccess()
}
