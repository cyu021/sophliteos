package mvc

type DeviceId struct {
	DeviceID string `json:"deviceId"`
}

type DeviceInfo struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
	Url      string `json:"url"`
	PtzType  int    `json:"ptzType"`
}

type DeviceMod struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
	Url      string `json:"url"`
	PtzType  int    `json:"ptzType"`
	DeviceId string `json:"deviceId"`
}

type DeviceDetectBody struct {
	DetectReportURL string     `json:"detectReportUrl"`
	DevList         []DeviceId `json:"devList"`
}

type DevicesList struct {
	Code      int      `json:"code"`
	Msg       string   `json:"msg"`
	TotalSize int      `json:"totalSize"`
	PageCount int      `json:"pageCount"`
	PageSize  int      `json:"pageSize"`
	PageNo    int      `json:"pageNo"`
	Devices   []Device `json:"device"`
}

// 实时预览请求接收
type LiveUrl struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	TsUrl string `json:"tsUrl"`
}

type DeviceIP struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result string `json:"result"`
}

// 流媒体请求接收
type MediaRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type MediaHost struct {
	Ip   string `json:"ip"  validate:"required"`
	Port int    `json:"port"  validate:"required"`
}

// 流媒体添加设备结构体
type Device struct {
	Codec       string `json:"codec"`
	DeviceId    string `json:"deviceId"  gorm:"primary_key"`
	DeviceName  string `json:"name"`
	Protocol    int    `json:"protocol"` //协议类型，0-未知，1-国标，2-RTSP，3-海康SDK，4-大华SDK
	PtzType     int    `json:"ptzType"`  //摄像机类型，0-未知，1-球机，2-半球，3-枪机，4-遥控枪机
	Resolution  string `json:"resolution"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	MediaServer string `json:"mediaServer"`
	MediaPull   int    `json:"mediaPull"`
}

// 接收巡检信息
type Detect struct {
	FaultSource int                `json:"faultSource"`
	List        []DetectDeviceInfo `json:"list"`
	OperationID string             `json:"operationId"`
}

type DetectDeviceInfo struct {
	DeviceID     string `json:"deviceId"`
	EncodeFormat string `json:"encodeFormat"`
	FaultType    int    `json:"faultType"`
	MediaPull    int    `json:"mediaPull"`
	OperationID  string `json:"operationId"`
	Resolution   string `json:"resolution"`
}
