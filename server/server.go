package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var FileChan chan FileInfo = make(chan FileInfo, 1024)

func Run() {
	go UploadChan()
	beego.Run()
}

func UploadChan() {
	for {
		f := <-FileChan
		log.Println("upload file: ", f.FileName)

		filepath := f.FilePath + "/" + f.FileName
		fileinfo, err := os.Stat(filepath)
		if err != nil {
			logs.Error("PutFile error: ", err.Error())
			continue
			// return FileNotExist(request.SeqNum)
		}
		if fileinfo.IsDir() != false {
			logs.Error("PutFile error: This is a dir.")
			continue
		}

		aliyun, err := NewAliyunObject(beego.AppConfig.String("endpoint"), beego.AppConfig.String("accesskey"), beego.AppConfig.String("secretkey"), f.Bucket)
		if err != nil {
			logs.Error(err.Error())
			continue
		}

		err = aliyun.PutFile(f.Object, filepath)
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		err = os.Remove(filepath)

		if err != nil {
			logs.Error(err.Error())
			continue
		}

		res := ResUpdateFileToP{}
		out, _ := json.Marshal(res)
		posturl := beego.AppConfig.String("puthost") + "/" + beego.AppConfig.String("putport") + "/" + beego.AppConfig.String("putpath") + "/" + f.Id
		resp, err := http.Post(posturl, "application/json", strings.NewReader(string(out)))
		if err != nil {
			logs.Error(err.Error())
		}
		defer resp.Body.Close()
	}
}
