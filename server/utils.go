package server

import "github.com/wst-libs/wst-sdk/utils"

// PutFile request structure
type ReqPutFile struct {
	Data struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Path   string `json:"path"`
		Url    string `json:"url"`
		Size   int64  `json:"size"`
		Bucket string `json:"bucket"`
		Object string `json:"object"`
		Desc   string `json:"description"`
	} `json:"data"`
}

// PutFile response structure
type ResPutFile struct {
	utils.ResponseCommon
	Data struct {
		Id string `json:"id"`
	}
}

// GetFile response structure
type ResGetFile struct {
	utils.ResponseCommon
	Data struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Path   string `json:"path"`
		Url    string `json:"url"`
		Size   string `json:"size"`
		Bucket string `json:"bucket"`
		Object string `json:"object"`
		Desc   string `json:"description"`
	}
}

// UpdateFile request structure
type ReqUpdateFile struct {
	Data struct {
		Object string `json:"object"`
		Desc   string `json:"description"`
	}
}

// UpdateFile response structure
type ResUpdateFile struct {
	utils.ResponseCommon
}

type ResDeleteFile struct {
	utils.ResponseCommon
}

type ReqDeleteFileList struct {

}

type ResDeleteFileList struct {
	utils.ResponseCommon
}
