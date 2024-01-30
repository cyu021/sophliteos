package mvc

// 接受告警信息
type AlarmDate struct {
	TaskID           string         `json:"TaskID" validate:"required"`
	SceneImageBase64 string         `json:"SceneImageBase64" validate:"required"`
	SrcID            string         `json:"SrcID"`
	FrameIndex       int            `json:"FrameIndex"`
	AnalyzeEvents    []AnalyzeEvent `json:"AnalyzeEvents" validate:"required"`
}

type AnalyzeEvent struct {
	ImageBase64 string      `json:"ImageBase64"`
	Type        int         `json:"Type"`
	Box         Box         `json:"Box"`
	Extend      interface{} `json:"Extend"`
}

type Box struct {
	LeftTopY  int `json:"LeftTopY"`  //左上⻆ y 坐标
	RightBtmY int `json:"RightBtmY"` //右下⻆ y 坐标
	LeftTopX  int `json:"LeftTopX"`  //左上⻆ x 坐标
	RightBtmX int `json:"RightBtmX"` //右下⻆ x 坐标
}
type AlgorithmHost struct {
	Ip   string `json:"ip"  validate:"required"`
	Port int    `json:"port"  validate:"required"`
}

// 告警图片数据库存储
type Record struct {
	ID                   uint `gorm:"primary_key"`
	TaskId               string
	SrcID                string
	FrameIndex           int
	Type                 int
	Date                 int64
	SamllPictureFilename string
	BigPictureFilename   string
	LeftTopY             int
	RightBtmY            int
	LeftTopX             int
	RightBtmX            int
	Extend               string `gorm:"type:text"`
	Number               int
}

type ItemInBox struct {
	Confidence float64 `json:"confidence"`
	Type       string  `json:"type"`
	Content    struct {
		NotSatisfy int `json:"notSatisfy"`
	} `json:"content"`
	FaceId string `json:"faceId"`
}

type Extra struct {
	Cei        interface{} `json:"cei"`
	ItemsInBox []ItemInBox `json:"itemsInBox"`
}

type AlarmInfo struct {
	Boxes []Box `json:"boxes"`
	Extra Extra `json:"extra"`
}

// 新增算法任务接收结构体
type AlgoTask struct {
	TaskId     string `json:"taskId" validate:"required"`
	DeviceName string `json:"deviceName" validate:"required"`
	DeviceId   string `json:"deviceId" validate:"required"`
	Url        string `json:"url" validate:"required"`
	Abilities  []int  `json:"abilities" validate:"required"`
}

// 修改算法任务接收结构体
type AlgoTaskMod struct {
	TaskId    string `json:"taskId" validate:"required"`
	DeviceId  string `json:"deviceId"`
	Abilities []int  `json:"abilities"`
}

// 算法任务数据库存储
type AlgoTaskSql struct {
	ID          uint   `gorm:"primary_key"`
	TaskId      string `json:"taskId"`
	Status      int    `json:"status"`
	DeviceName  string `json:"deviceName"`
	DeviceId    string `json:"deviceId"`
	Url         string `json:"url"`
	VideoWidth  int    `json:"VideoWidth"`
	VideoHeight int    `json:"VideoHeight"`
	Abilities   string `json:"abilities" gorm:"type:text"`
}

// 算法应用响应
type MediaReq struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// 算法应用响应
type AlgoReq struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}

// 告警信息查询请求
type AlgoQuery struct {
	BeginTime int64  `json:"beginTime"`
	EndTime   int64  `json:"endTime"`
	PageNo    int    `json:"pageNo"`
	PageSize  int    `json:"pageSize"`
	TaskID    string `json:"taskID"`
	SrcID     string `json:"srcID"`
	Types     []int  `json:"types"`
}

// 告警信息查询响应
type AlarmReq struct {
	Total     int            `json:"total"`
	PageNo    int            `json:"pageNo"`
	PageSize  int            `json:"pageSize"`
	PageCount int            `json:"pageCount"`
	Items     []AlarmReqItem `json:"items"`
	UsedSize  string         `json:"usedSize"`
	MaxSize   string         `json:"maxSize"`
}

type QueryInfo struct {
	Types   []int    `json:"types"`
	TaskIds []string `json:"taskIds"`
	SrcIds  []string `json:"srcIds"`
}
type AlarmReqItem struct {
	Id         int    `json:"id"`
	TaskId     string `json:"taskId"`
	SrcID      string `json:"srcId"`
	FrameIndex int    `json:"frameIndex"`
	Type       int    `json:"type"`
	SmallImage string `json:"smallImage"`
	BigImage   string `json:"bigImage"`
	Time       int64  `json:"time"`
	Box        Box    `json:"box"`
	Extend     string `json:"Extend"`
}

// 下发算法任务body
type AlgoStart struct {
	CameraID   string        `json:"cameraId"`
	URL        string        `json:"url"`
	ImageOut   string        `json:"imageOut"`
	InputType  string        `json:"inputType"`
	DecodeType string        `json:"decodeType"`
	NotifyURL  string        `json:"notifyUrl"`
	SkipFrame  int           `json:"skipFrame"`
	ROI        []interface{} `json:"roi"`
	AreaBoxes  []interface{} `json:"areaBoxes"`
	Abilities  []Ability     `json:"abilities"`
}

type Ability struct {
	Name  string       `json:"name"`
	Value AbilityValue `json:"value"`
}

// 算法能力结构体
type AbilityValue struct {
	Interval      float64 `json:"interval"`
	MinTarry      int     `json:"minTarry"`
	AlarmInterval float64 `json:"alarmInterval"`
	Threshold     float64 `json:"threshold"`
	AreaIsReverse bool    `json:"areaIsReverse"`
	PointType     int     `json:"pointType"`
	ZoomFactor    float64 `json:"zoomFactor"`
	ConfirmCount  int     `json:"confirmCount"`
	MinBox        MinRect `json:"minBox"`
	AreaBoxes     [][]struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"areaBoxes"`
}

type MinRect struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// 接受算法参数配置
type AlgoConfig struct {
	Ability   string  `json:"ability"`
	Interval  float64 `json:"interval"`
	MinBox    MinRect `json:"minBox"`
	Threshold float64 `json:"threshold"`
}

// 接受前端请求内容
type TaskReq struct {
	TaskId string `json:"taskId"`
}

// 向算法应用发送请求
type TaskRes struct {
	TaskId string `json:"TaskID"`
}

type UploadIp struct {
	Ip string `json:"ip"`
}
