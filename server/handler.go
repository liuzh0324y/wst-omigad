package server

import (
	"log"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func PutFileHandler(ctx *context.Context) []byte {
	s, err := PutFileRequest(ctx.Input.RequestBody)
	if err != nil {
		return JsonFormatErr()
	}
	info := FileInfo{
		FilePath: s.Data.Path,
		FileName: s.Data.Name,
		FileType: s.Data.Type,
		Bucket:   s.Data.Bucket,
		Object:   s.Data.Object,
	}
	FileChan <- info

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
	bucket := m.Get("bucket")
	object := m.Get("object")
	if len(bucket) <= 0 {
		bucket = "llvision"
	}
	if len(object) <= 0 {
		return UnknownReq()
	}

	obj, err := NewAliyunObject(beego.AppConfig.String("endpoint"), beego.AppConfig.String("accesskey"), beego.AppConfig.String("secretkey"), bucket)
	if err != nil {
		log.Println("GetFileHandler error: ", err.Error())
		return BucketNotFound()
	}

	isExist, err := obj.IsFileExist(m.Get("object"))
	if err != nil {
		log.Println("GetFileHandler Error: ", err.Error())
		return InternalError()
	}
	if isExist != true {
		log.Println("file not exist.")
		return FileNotFound()
	}
	returl, err := obj.GetFileWithURL(object, 3600)
	if err != nil {
		log.Println("GetFileHandler error: ", err.Error())
		return InternalError()
	}
	log.Println("url: ", returl)
	return GetFileResponse(returl)
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
	bucket := m.Get("bucket")
	object := m.Get("object")

	s, err := UpdateFileRequest(ctx.Input.RequestBody)
	if err != nil {
		return JsonFormatErr()
	}

	obj, err := NewAliyunObject(beego.AppConfig.String("endpoint"), beego.AppConfig.String("accesskey"), beego.AppConfig.String("secretkey"), bucket)
	if err != nil {
		log.Println("UpdateFileHandler error: ", err.Error())
		return BucketNotFound()
	}

	isExist, err := obj.IsFileExist(object)
	if err != nil {
		log.Println("UpdateFileHandler error: ", err.Error())
		return InternalError()
	}
	if isExist != true {
		log.Println("file not exist.")
		return FileNotFound()
	}

	obj.UpdateFile(object, "description", s.Data.Desc)

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
	bucket := m.Get("bucket")
	object := m.Get("object")

	obj, err := NewAliyunObject(beego.AppConfig.String("endpoint"), beego.AppConfig.String("accesskey"), beego.AppConfig.String("secretkey"), bucket)
	if err != nil {
		log.Println("UpdateFileHandler error: ", err.Error())
		return BucketNotFound()
	}

	isExist, err := obj.IsFileExist(object)
	if err != nil {
		log.Println("UpdateFileHandler error: ", err.Error())
		return InternalError()
	}
	if isExist != true {
		log.Println("file not exist.")
		return FileNotFound()
	}

	obj.DeleteFile(object)

	return DeleteFileResponse()
}
