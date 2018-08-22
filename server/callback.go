package server

import (
	"log"

	"github.com/astaxie/beego/context"
)

func CallBackHandler(ctx *context.Context) []byte {
	log.Println("callback: ", ctx.Input.RequestBody)
	return []byte{}
}
