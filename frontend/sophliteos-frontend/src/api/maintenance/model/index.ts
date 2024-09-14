export interface IpSetParams {
  device: string;
  ipType: number; // ip类型 1静态ip 2动态ip
  ip: string; // ip地址
  subnetMask: string; // 子网掩码
  gateway: string; // 网关
  dns: string; // dns
}

export interface AlgoConfigSetParams {
  ip: string; // ip地址
  port: string; // port
}

export interface UpUrlConfigSetParams {
  ip: string; // ip地址
  port: string;
  protocol: string; //http|https
  endpoint: string;
  retryRows: number; //rows, 1000-5000
  enableRetry: boolean;
}

export interface RotateCfgSetParams {
  record: number; //days, 1-180
  serviceLog: number; //days, 1-180
  retryRows: number; //rows, 1000-5000
  enableRetry: boolean;
}

export interface AlarmParams {
  fanSpeed: number; // 风扇转速
  boardTemperature: number; // 主板温度
  coreTemperature: number; // 芯片结温
  cpuRate: number; // cpu使用率
  totalMemoryScale: number; // 内存使用率
  // systemScale: number; // system内存使用率
  // videoScale: number; // video内存使用率
  tpuScale: number; // tpu内存使用率
  // externalHardDiskRate: number; // 外挂存储使用率
  diskRate: number; // 存储使用率
  tpuRate: number;
}

export interface RollbackParams {
  workflowId: number;
}

// multipart/form-data: upload file
export interface UploadFileParamsSys {
  // Other parameters
  module: string;
  // file name
  file: File | Blob;
  // file name
  ip?: string;
}
export interface UploadApiResult {
  msg: string;
  code: number;
  result: string;
}

export interface BasicPageParams {
  pageNo: number;
  pageSize: number;
}

export interface stopOrStartContainerParams {
  name: string;
}

export interface pullMirrorItem {
  name: string;
  cmd: string;
  status: number;
}

// -----S  核心板docker容器页面   S-----
export interface CoreDockerGetParam {
  ip: string;
}

export interface CoreImageListItem {
  Repository: string; //镜像名称
  Tag: string; //标签（显示镜像的时候，那一列是名称+：+标签）
  ID: string; //镜像ID
  Created: number; //秒为单位时间戳，表示该镜像被创建的时间点，不是创建了多久
  Size: number; //镜像大小，定位MB
}

export interface CoreImageDeleteParam {
  ip: string;
  image: string;
}

export interface CoreContainListItem {
  name: string; //容器名称
  image: string; //镜像
  id: string; //容器ID
  createdAt: string; //创建时间
  status: string; //运行时长
  ports: string[]; //端口列表
  running: boolean; //是否正在运行，对于停止状态的容器，前端页面不允许进入容器终端
}

export interface CoreContainParam {
  //删除、启动、停止的参数
  ip: string;
  name: string;
}

export interface CoreRegistryParam {
  ip: string;
  insecureRegistry: string;
}

export interface CoreBoardListItem {
  ip: string; //核心板ip，ssh连接核心板，传入该地址
  deviceSn: string; //核心板SN
  number: number; //核心板编号
}
// -----E  核心板docker容器页面   E-----
