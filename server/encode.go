package server

import (
	"encoding/json"

	"github.com/wst-libs/wst-sdk/utils"
)

func PutFileResponse() []byte {
	s := ResPutFile{
		utils.ResponseCommon{
			Version: utils.Version,
			SeqNum:  1,
			From:    "",
			To:      "",
			Type:    "",
			Number:  "",
			Uid:     "",
		},
		utils.ID{
			Id: "",
		},
	}
	out, _ := json.Marshal(s)
	return out
}

func GetFileResponse() []byte {
	s := ResGetFile{
		utils.ResponseCommon{
			Version: utils.Version,
			SeqNum:  1,
			From:    "",
			To:      "",
			Type:    "",
			Number:  "",
			Uid:     "",
		},
		GetFileData{
			Id:     "",
			Name:   "",
			Type:   "",
			Path:   "",
			Url:    "",
			Size:   "",
			Bucket: "",
			Object: "",
			Desc:   "",
		},
	}
	out, _ := json.Marshal(s)
	return out
}

func UpdateFileResponse() []byte {
	s := ResUpdateFile{
		utils.ResponseCommon{
			Version: utils.Version,
			SeqNum:  1,
			From:    "",
			To:      "",
			Type:    "",
			Number:  "",
			Uid:     "",
		},
	}
	out, _ := json.Marshal(s)
	return out
}

func DeleteFileResponse() []byte {
	s := ResDeleteFile{
		utils.ResponseCommon{
			Version: utils.Version,
			SeqNum:  1,
			From:    "",
			To:      "",
			Type:    "",
			Number:  "",
			Uid:     "",
		},
	}
	out, _ := json.Marshal(s)
	return out
}
