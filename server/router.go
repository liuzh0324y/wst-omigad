package server

import (
	"log"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/v1/cloudstorage/file", &Controller{}, "put:PutFile;get:GetFile;post:UpdateFile;delete:DeleteFile")
	log.Println("Initialize Router.")
}
