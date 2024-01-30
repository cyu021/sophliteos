export interface MediaServerParams {
  name: string;
  pageNo: number;
  pageSize: number;
}

export interface AlgorithmsItemType {
  Type: number;
  TrackInterval: number;
  DetectInterval: number;
  TargetSize: {
    MinDetect: number;
    MaxDetect: number;
  };
  DetectInfos: any;
  Extend: any;
}
export interface AlgoTaskInfoType {
  device: {
    codec: string; //编码类型
    deviceId: string; //设备id
    name: string; //设备名称
    protocol: number; //协议，2表示rtsp协议
    ptzType: number;
    resolution: string; //分辨率
    status: string; //设备状态
    type: string; //设备类型
    url: string; //设备地址
    mediaServer: string; //设备接入的流媒体
    mediaPull: number; //是否拉流成功
  };
  algorithms: AlgorithmsItemType[];
}

export interface algoTaskPanrams {
  taskId: string;
}

export interface algoTaskEditParams {
  taskId: string;
  Algorithm: {
    Type: number;
    TrackInterval: number;
    DetectInterval: number;
    TargetSize: {
      MinDetect: number;
      MaxDetect: number;
    };
    DetectInfos: [
      {
        TripWire: any;
        HotArea: [
          {
            X: number;
            Y: number;
          },
          {
            X: number;
            Y: number;
          },
          {
            X: number;
            Y: number;
          },
          {
            X: number;
            Y: number;
          },
        ];
      },
    ];
    Extend: any;
  };
}
