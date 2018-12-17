package dm

// OpType 操作类型
type OpType int

const (
	// RespTopic dm返回给ac的response队列名
	RespTopic = "dm_response_topic"

	// RecognizeTopic 识别消息的队列名
	RecognizeTopic = "recognize_pub"

	addFaceType = "add_face"
	delFaceType = "del_face"
)

const (
	_ OpType = iota
	// OpTypeAddFace operation for add face
	OpTypeAddFace
	// OpTypeDelFace operation for delete face
	OpTypeDelFace
)

// OpTypeMap 映射OpType对应的字符串
var OpTypeMap = map[OpType]string{
	OpTypeAddFace: addFaceType,
	OpTypeDelFace: delFaceType,
}

// StrOpTypeMap 字符串到OpType的反向查询
var StrOpTypeMap = map[string]OpType{
	addFaceType: OpTypeAddFace,
	delFaceType: OpTypeDelFace,
}

// --- Device ---

// AddDevReq is request of add device
type AddDevReq struct {
	ReqID    string `json:"req_id"`
	DevURL   string `json:"dev_url"`
	DevTopic string `json:"dev_topic"`
}

// AddDevResp is response of add device
type AddDevResp struct {
	ReqID  string `json:"req_id"`
	DevURL string `json:"dev_url"`
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

// DelDevReq 删除设备请求
type DelDevReq struct {
	ReqID  string `json:"req_id"`
	DevURL string `json:"dev_url"`
}

// DelDevResp 删除设备响应
type DelDevResp struct {
	ReqID  string `json:"req_id"`
	DevURL string `json:"dev_url"`
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

// --- Group ---

// AddGroupReq is request for add group
type AddGroupReq struct {
	ReqID     string  `json:"req_id"`
	DevURL    string  `json:"dev_url"`
	Group     string  `json:"group"`
	Threshold float32 `json:"threshold"`
	TopN      int32   `json:"top_n"`
}

// AddGroupResp is response for add group
type AddGroupResp struct {
	ReqID  string `json:"req_id"`
	DevURL string `json:"dev_url"`
	Group  string `json:"group"`
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

// DelGroupReq is request for add group
type DelGroupReq struct {
	ReqID  string `json:"req_id"`
	Group  string `json:"group"`
	DevURL string `json:"dev_url"`
}

// DelGroupResp is response for add group
type DelGroupResp struct {
	ReqID  string `json:"req_id"`
	DevURL string `json:"dev_url"`
	Group  string `json:"group"`
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

// AddFaceReq 添加人脸请求
type AddFaceReq struct {
	ReqID   string `json:"req_id"`
	Group   string `json:"group"`
	Face    string `json:"face"`
	FaceURL string `json:"face_url"`
}

// AddFaceResp 增加人脸响应
type AddFaceResp struct {
	ReqID  string `json:"req_id"`
	DevURL string `json:"dev_url"`
	Group  string `json:"group"`
	Face   string `json:"face"`
}

// DelFaceReq 删除人脸请求
type DelFaceReq struct {
	ReqID string `json:"req_id"`
	Group string `json:"group"`
	Face  string `json:"face"`
}

// DelFaceResp 增加人脸响应
type DelFaceResp struct {
	ReqID  string `json:"req_id"`
	Group  string `json:"group"`
	Face   string `json:"face"`
	DevURL string `json:"dev_url"`
}

// SyncReq 同步请求
// OpType 操作类型
// OpData 操作数据。json格式化。
type SyncReq struct {
	OpType string `json:"op_type,omitempty"`
	OpData []byte `json:"op_data,omitempty"`
	Vers   int    `json:"vers,omitempty"`
}

// SyncResp 同步响应
type SyncResp struct {
	Result bool   `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
	OpType string `json:"op_type,omitempty"`
	OpData []byte `json:"op_data,omitempty"`
	Vers   int    `json:"vers,omitempty"`
}
