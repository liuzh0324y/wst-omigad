package omg

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// setResHeader set response header
func setResHeader(header http.Header) {
	header.Add("Content-Type", "application/json")
	header.Add("Connection", "close")
}

// uploadLocalFileToCloud upload the local file to cloud
func uploadLocalFileToCloud(w http.ResponseWriter, r *http.Request) {
	setResHeader(w.Header())
	// u, err := url.ParseQuery(r.URL.RawQuery)
	// if err != nil {

	// }

	// bucket := u.Get("bucket")
	// object := u.Get("object")
	// log.Println("bucket: ", bucket)
	// log.Println("object: ", object)
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {

	// }

	// w.Write(uploadLocalFileToCloudHandler(body))

	w.Write(notFound())
}

// getURLOfFileForCloud get the url of file for cloud
func getURLOfFileForCloud(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setResHeader(w.Header())

	u, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.Write(InvalidParams())
		return
	}
	bucket := u.Get("bucket")
	object := u.Get("object")
	if len(bucket) == 0 {
		bucket = defaultBucket
	}
	if len(object) == 0 {

	}
	w.Write(getURLOfFileForCloudHandler(bucket, object))
}

// updateFileInfoOfCloud  update the file info of cloud
func updateFileInfoOfCloud(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setResHeader(w.Header())
	u, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {

	}

	bucket := u.Get("bucket")
	object := u.Get("object")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	}

	w.Write(updateFileInfoOfCloudHandler(bucket, object, body))
}

// deleteFileOfCloud delete the file of cloud
func deleteFileOfCloud(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setResHeader(w.Header())
	u, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
	}
	bucket := u.Get("bucket")
	object := u.Get("object")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	}

	w.Write(deleteFileOfCloudHandler(bucket, object, body))
}

// getURLOfUploadFile get the url of the upload file
func getURLOfUploadFile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setResHeader(w.Header())
	// Parse url

	m, _ := url.ParseQuery(r.URL.RawQuery)

	bucket := m.Get("bucket")
	if len(bucket) == 0 {
		bucket = defaultBucket
	}
	object := m.Get("object")
	if len(object) == 0 {
		w.Write(InvalidParams())
		return
	}
	t := m.Get("type")
	if len(t) == 0 {

	}

	w.Write(getURLOfUploadFileHandler(bucket, object, t))
}

func callback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setResHeader(w.Header())

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("failed to oss callback")
	}

	w.Write(callbackhandler(body))

}

// getStsAuth get aliyuncs sts auth
func getStsAuth(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setResHeader(w.Header())

	var outBody []byte
	outBody = getStsAuthHandler()
	w.Write(outBody)
}

func getObjectSize(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	setResHeader(w.Header())
	// vars := mux.Vars(r)
	// w.Write(getObjectSizeHandler(vars["bucket"], vars["object"]))
}
