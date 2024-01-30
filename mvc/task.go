package mvc

// 任务下发参数
type Task struct {
	TaskId    string      `json:"TaskID"`
	InputSrc  InputSource `json:"InputSrc"`
	Algorithm []Algorithm `json:"Algorithm"`
	Reporting Reporting   `json:"Reporting"`
}

type InputSource struct {
	SrcID     string       `json:"SrcID"`
	StreamSrc StreamSource `json:"StreamSrc"`
}

type StreamSource struct {
	Address string `json:"Address"`
}

type Algorithm struct {
	Type           int          `json:"Type"`
	TrackInterval  int          `json:"TrackInterval"`
	DetectInterval int          `json:"DetectInterval"`
	TargetSize     TargetSize   `json:"TargetSize,omitempty"`
	DetectInfos    []DetectInfo `json:"DetectInfos,omitempty"`
	Extend         interface{}  `json:"Extend,omitempty"` // 使用 interface{} 以支持任何类型的扩展字段
}

type TargetSize struct {
	MinDetect int `json:"MinDetect"`
	MaxDetect int `json:"MaxDetect"`
}

type DetectInfo struct {
	TripWire TripWire  `json:"TripWire,omitempty"`
	HotArea  []Point2D `json:"HotArea",omitempty`
}

type TripWire struct {
	LineStart   Point2D `json:"LineStart"`
	LineEnd     Point2D `json:"LineEnd"`
	DirectStart Point2D `json:"DirectStart"`
	DirectEnd   Point2D `json:"DirectEnd"`
}

type Point2D struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

type Reporting struct {
	ReportUrlList []string `json:"ReportUrlList"`
}

// 任务参数修改
type TaskConf struct {
	TaskId    string    `json:"TaskID"`
	Algorithm Algorithm `json:"Algorithm"`
}

// 查询算法服务的任务状态列表
type TaskList struct {
	Code   int        `json:"Code"`
	Msg    string     `json:"Msg"`
	Result []TaskInfo `json:"Result"`
}

type TaskInfo struct {
	TaskId string `json:"TaskID"`
	Status int    `json:"Status"`
}

// 返回给前端的任务列表
type TaskListRes struct {
	Total     int        `json:"total"`
	PageSize  int        `json:"pageSize"`
	PageCount int        `json:"pageCount"`
	PageNo    int        `json:"pageNo"`
	Items     []TaskItem `json:"items"`
}

type TaskItem struct {
	TaskId      string   `json:"taskId"`
	DeviceName  string   `json:"deviceName"`
	DeviceId    string   `json:"deviceId"`
	Status      int      `json:"status"`
	ErrorReason string   `json:"errorReason"`
	Abilities   []string `json:"abilities"`
	Types       []int    `json:"types"`
}

type AbilityList struct {
	Type int    `json:"type"`
	Name string `json:"name"`
}

// 算法配置信息，包括视频通道信息和能力列表
type AlgorithmConf struct {
	Device     Device      `json:"device"`
	Algorithms []Algorithm `json:"algorithms"`
}
