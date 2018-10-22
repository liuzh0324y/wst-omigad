package omg

import (
	"log"
	"net/http"
	"os"

	"github.com/wst-libs/wst-sdk/conf"
)

var config httpconfig

var (
	configurl = os.Getenv("WST_OMIGAD_CONFIG")
)

// Run is start function
func Run() {
	// go uploadChan()

	var addr string
	err := getconfig()
	if err != nil {
		addr = ":18012"
	} else {
		addr = ":" + config.Server.HTTPPort
	}

	log.Fatal(http.ListenAndServe(addr, router()))
}

func getconfig() error {
	err := conf.GetConf(configurl, &config)
	if err != nil {
		log.Println("confserver: ", err.Error())
		return err
	}
	log.Printf("%v\n", config)
	return nil
}

// // uploadinfo get a upload url of oss and create the record to manager.
// func uploadinfo(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	// Parse url

// 	m, _ := url.ParseQuery(r.URL.RawQuery)
// 	log.Println(m)
// 	bucket := m.Get("bucket")
// 	object := m.Get("object")
// 	if len(bucket) == 0 {
// 		bucket = defaultBucket
// 	}
// 	if len(object) == 0 {
// 		w.Write(InvalidParams())
// 		return
// 	}
// 	// New object for oss
// 	obj, err := ost.NewOst(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
// 	if err != nil {
// 		log.Println("GetUrlFromFileHandler")
// 		w.Write(BucketNotFound())
// 		return
// 	}

// 	// check is file exist
// 	isExist, err := obj.IsObjectExist(object)
// 	if err != nil {
// 		log.Println("GetUrlForFileHandler error: ", err.Error())
// 		w.Write(InternalError())
// 		return
// 	}
// 	if isExist != false {
// 		log.Println("file not exist")
// 		w.Write(FileAlreadyExist())
// 		return
// 	}

// 	// get id for manager
// 	res := manager.Add(config.Server.ManagerURL, object)
// 	if res.Code != 0 {
// 		w.Write(CreateRecordFailed())
// 		return
// 	}
// 	url, err := obj.PutFileWithURL(object)
// 	if err != nil {
// 		log.Println("GetUrlForFileHandler error: ", err.Error())
// 		w.Write(InternalError())
// 		return
// 	}

// 	w.Write(GetUrlForFileResponse(res.Id, url, "", ""))
// }

func setHeader(h http.Header) {
	h.Add("Content-Type", "application/json")
	h.Add("Connection", "close")
}
