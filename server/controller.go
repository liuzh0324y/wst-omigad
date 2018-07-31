package server

import (
	"log"

	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

func (t *Controller) PutFile() {
	log.Printf("params: %v\n", t.Ctx.Input.Params())
	log.Println("=========================")
	t.Ctx.WriteString("omigad")
}

func (t *Controller) GetFile() {
	t.Ctx.WriteString("GetFile")
}

func (t *Controller) UpdateFile() {
	t.Ctx.WriteString("UpdateFile")
}

func (t *Controller) DeleteFile() {
	t.Ctx.WriteString("DeleteFile")
}
