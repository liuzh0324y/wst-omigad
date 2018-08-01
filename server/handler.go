package server

import (
	"log"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/wst-libs/wst/logs"
)

func PutFileHandler(ctx *context.Context) []byte {
	s, err := PutFileRequest(ctx.Input.RequestBody)
	if err != nil {
		return []byte("")
	}

	log.Printf("%v\n", s)
	return PutFileResponse()
}

func GetFileHandler(ctx *context.Context) []byte {
	// Parse uri query
	u, err := url.Parse(ctx.Input.URI())
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	log.Println("raw query: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	log.Println(m)

	aliyun, err := NewAliyunObject(beego.AppConfig.String("endpoint"), beego.AppConfig.String("accesskey"), beego.AppConfig.String("secretkey"), beego.AppConfig.String("bucket"))
	if err != nil {
		logs.Error("PostFileWithUrl error: ", err.Error())
		return []byte{}
	}

	isExist, err := aliyun.IsFileExist("filename")
	if err != nil {
		logs.Error("PostFileWithUrl Error: ", err.Error())
	}
	if isExist != true {

	}
	returl, err := aliyun.GetFileWithURL("filename", 3600)
	log.Println("url: ", returl)
	return GetFileResponse()
}

func UpdateFileHandler(ctx *context.Context) []byte {
	// Parse uri query
	u, err := url.Parse(ctx.Input.URI())
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	log.Println("raw query: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	log.Println(m)

	return UpdateFileResponse()
}

func DeleteFileHandler(ctx *context.Context) []byte {
	// Parse uri query
	u, err := url.Parse(ctx.Input.URI())
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	log.Println("raw query: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	log.Println(m)

	return DeleteFileResponse()
}
