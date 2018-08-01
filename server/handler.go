package server

import (
	"log"
	"net/url"

	"github.com/astaxie/beego/context"
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
