package omg

// var filechan = make(chan fileInfo, 1024)

// Put file
// func uploadChan() {
// 	for {
// 		f := <-filechan
// 		log.Println("upload file: ", f.FileName)

// 		filepath := f.FilePath + "/" + f.FileName
// 		fileinfo, err := os.Stat(filepath)
// 		if err != nil {

// 			continue
// 			// return FileNotExist(request.SeqNum)
// 		}
// 		if fileinfo.IsDir() != false {

// 			continue
// 		}

// 		obj, err := ost.NewOst(config.Server.Endpoint, config.Server.AccessKey, config.Server.SecretKey, f.Bucket)
// 		if err != nil {

// 			continue
// 		}

// 		err = obj.PutObjectFromFile(f.Object, filepath, callbackParamEncode())
// 		if err != nil {

// 			continue
// 		}

// 		if config.Server.IsRemove == true {
// 			err = os.Remove(filepath)

// 			if err != nil {

// 				continue
// 			}
// 		}

// 		res := ResUpdateFileToP{
// 			utils.RequestCommon{
// 				Version: utils.Version,
// 				SeqNum:  1,
// 				From:    "omigad",
// 				To:      "file manager",
// 				Type:    "omigad",
// 				Number:  "XXXX-XXXX-XXXX-XXXX",
// 				Uid:     f.ID,
// 			},
// 			PostInfo{
// 				Name:     f.FileName,
// 				Type:     f.FileType,
// 				Url:      "https://sample.com/sample.mp4",
// 				Key:      "",
// 				Secret:   "",
// 				Bucket:   f.Bucket,
// 				Object:   f.Object,
// 				Region:   "",
// 				Endpoint: config.Server.Endpoint,
// 				Desc:     "",
// 			},
// 		}
// 		out, _ := json.Marshal(res)
// 		posturl := "http://" + config.Server.ManagerURL + "/" + f.ID

// 		resp, err := http.Post(posturl, "application/json", strings.NewReader(string(out)))
// 		if err != nil {
// 			log.Println("Upload file error: ", err.Error())
// 		} else {
// 			resp.Body.Close()
// 		}
// 	}
// }
