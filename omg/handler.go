package omg

import (
	"log"

	"github.com/wst-libs/wst-sdk/sdk/ost"
)

// func uploadLocalFileToCloudHandler(body []byte) []byte {
// 	s, err := PutFileRequest(body)
// 	if err != nil {
// 		return JsonFormatErr()
// 	}
// 	info := fileInfo{
// 		ID:       s.Data.Id,
// 		FilePath: s.Data.Path,
// 		FileName: s.Data.Name,
// 		FileType: s.Data.Type,
// 		Bucket:   s.Data.Bucket,
// 		Object:   s.Data.Object,
// 	}
// 	filechan <- info

// 	return PutFileResponse()
// }

func getURLOfFileForCloudHandler(bucket, object string) []byte {

	// obj, err := NewAliyunObject(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
	// if err != nil {
	// 	log.Println("GetFileHandler error: ", err.Error())
	// 	return BucketNotFound()
	// }
	obj, err := ost.NewOst(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
	if err != nil {
		log.Println("sts err: ", err.Error())
		return BucketNotFound()
	}
	isExist, err := obj.IsObjectExist(object)
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

func updateFileInfoOfCloudHandler(bucket, object string, body []byte) []byte {

	s, err := UpdateFileRequest(body)
	if err != nil {
		return JsonFormatErr()
	}

	obj, err := ost.NewOst(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
	if err != nil {
		log.Println("UpdateFileHandler error: ", err.Error())
		return BucketNotFound()
	}

	isExist, err := obj.IsObjectExist(object)
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

func deleteFileOfCloudHandler(bucket, object string, body []byte) []byte {
	obj, err := ost.NewOst(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
	if err != nil {
		log.Println("UpdateFileHandler error: ", err.Error())
		return BucketNotFound()
	}

	isExist, err := obj.IsObjectExist(object)
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

// getURLOfUploadFileHandler
func getURLOfUploadFileHandler(bucket, object, t string) []byte {
	obj, err := ost.NewOst(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
	if err != nil {
		log.Println(err.Error())
		return BucketNotFound()
	}
	isE, err := obj.IsObjectExist(object)
	if err != nil {
		log.Println(err.Error())
		return InternalError()
	}
	if isE != false {
		return FileAlreadyExist()
	}
	// Create the record for file manager
	// res := manager.Add(config.Server.ManagerURL, object)
	// if res.Code != 0 {
	// 	return CreateRecordFailed()
	// }

	// Create upload url with callback param
	// param := callbackParamEncode(res.Id, object)
	url := obj.SignPutURL(object, t, 3600)

	// X-Oss-Callback
	return GetUrlForFileResponse(url, t)
}

func getStsAuthHandler() []byte {
	return []byte{}
}

func getObjectSizeHandler(bucket, object string) string {
	obj, err := ost.NewOst(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, bucket)
	if err != nil {
		log.Println(err.Error())
		// return BucketNotFound()
	}
	isE, err := obj.IsObjectExist(object)
	if err != nil {
		log.Println(err.Error())
		// return InternalError()
	}
	if isE != true {
		// return FileNotFound()
	}

	return obj.GetObjectSize(object)
	// resO := resObjectSize{
	// 	Size: size,
	// }
	// res, _ := json.Marshal(resO)
	// return res
}
