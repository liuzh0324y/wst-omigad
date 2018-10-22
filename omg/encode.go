package omg

import (
	"encoding/json"
)

func PutFileResponse() []byte {
	s := ResponseCommon{
		Version: Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "success",
		Code:    0,
	}
	//  ResPutFile{
	// 	utils.ResponseCommon{
	// 		Version: utils.Version,
	// 		SeqNum:  1,
	// 		From:    "imigad",
	// 		To:      "client",
	// 		Type:    "imigad",
	// 		Number:  "XXXX-XXXX-XXXX-XXXX",
	// 		Message: "success",
	// 		Code:    0,
	// 	},
	// 	utils.ID{
	// 		Id: "",
	// 	},
	// }
	out, _ := json.Marshal(s)
	return out
}

func GetFileResponse(url string) []byte {
	s := ResGetFile{
		ResponseCommon{
			Version: Version,
			SeqNum:  1,
			From:    "omigad",
			To:      "client",
			Type:    "omigad",
			Number:  "XXXX-XXXX-XXXX-XXXX",
		},
		GetFileData{
			Id:     "",
			Name:   "",
			Type:   "",
			Path:   "",
			Url:    url,
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
		ResponseCommon{
			Version: Version,
			SeqNum:  1,
			From:    "omigad",
			To:      "client",
			Type:    "omigad",
			Number:  "XXXX-XXXX-XXXX-XXXX",
		},
	}
	out, _ := json.Marshal(s)
	return out
}

func DeleteFileResponse() []byte {
	s := ResponseCommon{
		Version: Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "success",
		Code:    0,
	}
	out, _ := json.Marshal(s)
	return out
}

func GetUrlForFileResponse(url, t string) []byte {
	s := resGetURLForFile{
		ResponseCommon{
			Version: Version,
			SeqNum:  1,
			From:    "omigad",
			To:      "client",
			Type:    "omigad",
			Number:  "success",
			Code:    0,
		},
		putFileInfo{
			URL:  url,
			Type: t,
		},
	}

	out, _ := json.Marshal(s)
	return out
}
