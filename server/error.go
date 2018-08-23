package server

import (
	"encoding/json"

	"github.com/wst-libs/wst-sdk/utils"
)

const (
	NotFound     = 404
	InternalErr  = 503
	AlreadyExist = 504
)

func JsonFormatErr() []byte {
	s := utils.ResponseCommon{
		Version: utils.Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "json format error",
		Code:    InternalErr,
	}
	out, _ := json.Marshal(s)
	return out
}
func InternalError() []byte {
	s := utils.ResponseCommon{
		Version: utils.Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "Internal error",
		Code:    InternalErr,
	}
	out, _ := json.Marshal(s)
	return out
}

func UnknownReq() []byte {
	s := utils.ResponseCommon{
		Version: utils.Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "Unknown request params",
		Code:    NotFound,
	}
	out, _ := json.Marshal(s)
	return out
}

func BucketNotFound() []byte {
	s := utils.ResponseCommon{
		Version: utils.Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "Bucket not found",
		Code:    NotFound,
	}
	out, _ := json.Marshal(s)
	return out
}

// File not exist
func FileNotFound() []byte {
	s := utils.ResponseCommon{
		Version: utils.Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "File not found",
		Code:    NotFound,
	}

	out, _ := json.Marshal(s)
	return out
}

// File already exist
func FileAlreadyExist() []byte {
	s := utils.ResponseCommon{
		Version: utils.Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "File already exist",
		Code:    AlreadyExist,
	}

	out, _ := json.Marshal(s)
	return out
}
