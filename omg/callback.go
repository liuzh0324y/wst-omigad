package omg

import (
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/wst-libs/wst-sdk/sdk/manager"
)

type callbackParam struct {
	CallbackURL  string `json:"callbackUrl"`
	CallbackBody string `json:"callbackBody"`
}

type callbackBody struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type resCallback struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

// callbackhandler handle the callback
func callbackhandler(body []byte) []byte {
	log.Println("callback body: ", string(body))
	var req callbackBody
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err.Error())
	}
	fileURL := "https://" + config.Server.Bucket + "." + config.Server.Endpoint + "/" + req.Name
	fileSize := getObjectSizeHandler(config.Server.Bucket, req.Name)
	manager.Update(config.Server.ManagerURL+"/"+req.ID, req.ID, req.Name, fileURL, fileSize, 4)
	t := resCallback{
		Code: 0,
		Msg:  "success",
	}
	res, _ := json.Marshal(t)
	return res
}

func callbackParamEncode(id, name string) string {
	body := callbackBody{
		ID:   id,
		Name: name,
	}
	b, err := json.Marshal(body)
	if err != nil {
		log.Println("callback param error: ", err.Error())
	}

	param := callbackParam{
		CallbackURL:  config.Server.CallbackURL,
		CallbackBody: string(b),
	}
	data, err := json.Marshal(param)
	if err != nil {
		log.Println(err.Error())
	}
	out := base64.StdEncoding.EncodeToString(data)

	return out
}
