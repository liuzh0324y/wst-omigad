package main

import (
	"log"
	"net/url"

	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

func init() {
	beego.Router("/api/v1/storage/file", &Controller{}, "put:PutFile;get:GetFile;post:UpdateFile;delete:DeleteFile")
	log.Println("Initialize Router.")
}

func main() {
	beego.Run()
}

func (t *Controller) PutFile() {
	// Parse uri query
	u, err := url.Parse(t.Ctx.Input.URI())
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	log.Println("raw query: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	log.Println(m)

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
