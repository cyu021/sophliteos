export default {
  dashboard: 'Dashboard',
  about: '关于',
  workbench: '工作台',
  analysis: '分析页',
  basicInfor: '基础信息',
  ovewview: '设备概览',
  boardDetail: '板卡详情',
  logs: '日志管理',
  warning: '告警日志',
  operate: '操作日志',
  threshold: '告警阈值',
  maintenance: '设备运维',
  appSoft: '应用升级',
  sysSoft: '系统升级',
  sysSetting: '系统配置',
  softUpdate: '软件升级',
  ssmUpdate: 'SSM升级',
  liteOsUpdate: 'LiteOS升级',
  networkSetting: '网络设置',
  sshTerminal: '终端',
  hostComputer: '主机',
  quickCommands: '快速命令',
  DockerContainers: 'DOCKER容器',
  coreDockerContainers: '核心板DOCKER容器',
  containers: '容器',
  mirror: '镜像',
  store: '仓库',
  timeSetting: '时间设置',
  task: '任务管理',
  taskList: '任务查询',
  accessAlgo: '算法业务',
  taskSubscribe: '任务订阅',
  dataSource: '数据源维护',
  mediaServers: '流媒体服务',
  videoManage: '视频资源管理',
  AlgoParamConfig: '算法参数配置',
  alarmRetrieval: '告警检索',
  alarmDetail: '告警详情',
  coreBoardMap: '核心板端口映射',
  logDownload: '日志下载',
  content: {
    sysContent: '系统OTA升级,适用于版本升级,不支持降级和跨版本升级,否则可能导致系统损坏。',
    Content: `系统SDK OTA升级，不支持跨3.0版本升级,否则可能导致系统损坏。  
     1.  控制板升级请上传卡刷包文件，控制板升级完成后，再进行核心板升级，选择需要升级的核心板，无需上传文件，直接点击升级。
     2.  核心板升级，上传tftp刷机包，选择核心板进行升级。传错升级包可能导致系统损坏，请谨慎操作。`,
    ssmContent:
      'SSM软件是部署于AI算力设备上的管理服务,提供系统管理、系统升级、算力设备信息获取、算法应用部署等功能。',
    liteContent:
      'LiteOS软件基于SSM设备管理服务,提供了WEB页面的设备监控、告警查询软件和系统升级等功能',
  },
};
