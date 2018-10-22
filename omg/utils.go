package omg

const (
	defaultBucket = "glxsslive-test"
	Version       = "V1.0"
)

type ResNotFound struct {
	Status int64  `json:"status"`
	Msg    string `json:"message"`
}

///
/// Common request structure
///
type RequestCommon struct {
	Version string `json:"version"`
	SeqNum  int64  `json:"seqnum"`
	From    string `json:"from"`
	To      string `json:"to"`
	Type    string `json:"type"`
	Number  string `json:"number"`
	Uid     string `json:"uid"`
}

///
/// Common response structure
///
type ResponseCommon struct {
	Version string `json:"version"`
	SeqNum  int64  `json:"seqnum"`
	From    string `json:"from"`
	To      string `json:"to"`
	Type    string `json:"type"`
	Number  string `json:"number"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

// PutFile request structure
type ReqPutFile struct {
	Data struct {
		Id     string `json:"id"`
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
type ID struct {
	Id string `json:"id"`
}

// PutFile response structure
type ResPutFile struct {
	ResponseCommon
	ID `json:"data"`
}

// GetFile response structure
type GetFileData struct {
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
type ResGetFile struct {
	ResponseCommon
	GetFileData `json:"data"`
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
	ResponseCommon
}

type ResDeleteFile struct {
	ResponseCommon
}

type ReqDeleteFileList struct {
}

type ResDeleteFileList struct {
	ResponseCommon
}

type PostInfo struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Key      string `json:"key"`
	Secret   string `json:"secret"`
	Bucket   string `json:"bucket"`
	Object   string `json:"object"`
	Region   string `json:"region"`
	Endpoint string `json:"endpoint"`
	Desc     string `json:"description"`
}

type ResUpdateFileToP struct {
	RequestCommon
	PostInfo `json:"data"`
}
type fileParam struct {
	Type string `json:"Content-Type"`
	Data string `json:"X-Oss-Callback"`
}
type putFileInfo struct {
	URL  string `json:"url"`
	Type string `json:"Content-Type"`
}

type resGetURLForFile struct {
	ResponseCommon
	putFileInfo `json:"data"`
}

type fileInfo struct {
	ID       string
	FilePath string
	FileName string
	FileType string
	Channel  string
	SeqNum   int64
	TmpPath  string
	Bucket   string
	Object   string
}

type resObjectSize struct {
	Size string `json:"size"`
}
