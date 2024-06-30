export default {
  appSoftUpdate: 'Application Upgrade',
  sysSoftUpdate: 'System Upgrade',
  updatedFiles: 'Uploaded File',
  upgradePackage: 'Upgrade Package',
  prepSpringZip: {
    selectModel: 'Select model file',
    customAlgoName: 'Custom algo name',
    pack: 'Pack',
    customAlgoDownloadLinks: 'Custom algo download links:',
    createDate: 'Create date',
    download: 'Download',
    purge: 'Purge',
  },
  systemUpdate: {
    uploadFiles: 'Upload Files',
    systemUpdate: 'Control Board',
    coreBoardUpdate: 'Core Board',
    localUpdate1: 'Local Upgrade',
    updateList: 'Upgrade List',
    softwareUpdate: 'Software Upgrade',
    updateType: 'Upgrade Way:',
    localUpdate: 'Local Upload',
    currentVersion: 'System Version',
    currentSoftVersion: 'Software Version',
    currentSoVersion: 'SO version',
    currentVpsVersion: 'VPS version',
    currentBmsVersion: 'BMS version',
    currentSsmVersion: 'SSM Version',
    selectCoreBoard: 'Select Core Board:',
    selectUpgradeCoreBoard: 'Please Select Core Board',
    updateStatus: 'Upgrade Status',
    controlStartUpload: 'Start upgrding..., please wait five minites.',
    softwareUpdateStatus: 'Start upgrding..., please wait a minite！.',
    selectNeedFile: 'Please select file to upload',
    fileFormat: 'Only .tgz files are supported',
    startUpdate: 'Start Upgrade',
    updaing: 'Upgrading',
    selectFile: 'Select File',
    taskList: {
      name: 'Name',
      product: 'Device Type',
      moduleName: 'Module',
      status: 'Status',
      step: 'Step',
      strategy: 'Strategy',
      type: 'Operate Type',
      fileName: 'File Name',
      createTime: 'Create Time',
    },
    status: {
      '1': 'Submit',
      '2': 'Success',
      '3': 'Failure',
      '4': 'Expired',
    },
    type: {
      '1': 'Upgrade',
      '2': 'Install',
      '3': 'Rollback',
    },
    strategy: {
      flash: 'Flash',
      reboot: 'Reboot',
    },
    portal: 'Portal',
    otaDaemon: 'Daemon',
    adapter: 'Biz Module',
    core: 'Plugin Module'
  },
  ssmUpdate: {
    serverUsername: 'Server Username',
    serverPassword: 'Server Password',
    ssmUpdate: 'SSM Update',
  },
  coreBoardMap: {
    sysMap: 'system Map',
    UserDefineMap: 'User Define Map',
    number: 'Number',
    mode: 'Mode',
    protocol: 'protocol',
    ControlBoardIP: 'Control Board IP',
    ControlPort: 'Control Board Port',
    CoreBoardIP: 'Core Board IP',
    CorePort: 'Core Board Port',
    add: 'add',
    addMap: 'Add Mapping',
    msg: 'After the system restarts, the user-defined mapping will become invalid.',
  },
  containers: {
    // TODO 要启用中英文版本的时候搞这些
    // addStore: '添加仓库',
    // deleteConfirm: '确认删除此仓库？',
    // mirrorDelConfirm: '确认删除此镜像？',
    // storeName: '仓库名称',
    // addr: '仓库地址',
    // addrPlaceholder: '192.168.1.100:9500',
    // addrVlidate: '仓库地址填写有误',
    // mirror: '镜像',
    // mirrorName: '镜像名',
    // ID: 'ID',
    // mirrorSize: '镜像大小',
    // tag: '标签',
    // creat: '时间',
    // pullMirror: '拉取镜像',
    // insecureRegistry: '私仓地址',
    // containerName: '容器名',
    // containerPlaceholder: '请选择镜像名',
    // tagPlaceholder: '请选择镜像标签',
    // insecureRegistryPlaceholder: '请选择仓库',
    // terminal: '终端',
    // containerState: '状态',
    // containerPort: '端口',
    // containerStarttime: '开始时间',
    // activated: '已启动',
    // stop: 'Stop',
    // containerDelConfirm: '确认删除此容器？',
    runningTime: 'Running time',
    run: 'Run',
  },
  newworkSettings: {
    ipType: 'Connection Mode',
    ip: 'IP Address',
    subnetMask: 'Subnet Mask',
    netCard: 'Netcard',
    gateway: 'Gateway',
    dns: 'DNS',
    staticIP: 'static IP',
    dynmicIP: 'Dynmic IP',
    core1: 'IP1',
    core2: 'IP2',
    wan: 'WAN Setting',
    lan: 'LAN Setting',
    lan2: 'LAN2 Setting',
    core: 'Core board setting',
    inputRightIp: 'IP address is invalid',
    inputIp: 'Please enter IP address',
    inpuSubnetMask: 'Please enter subnet mask',
    inputGateway: 'Please enter gateway',
    inputDns: 'Please enter DNS',
    inputTwoIps: 'Please enter two IP address separated by commas',
    validateWrong: 'format is invalid, please re-enter',
    ipRightOrNotNetword:
      'Please confirm that the IP setting information is filled in correctly, otherwise the device may not be able to connect to the Internet.',
    restartAfterSubmit: 'After submission, the device will restart',
    useAddress: 'After the modification is successful, please use the address',
    copyAddress: 'Copy address',
    reOpenPage: ', Revisit this page',
    dynIpDialogContent:
      'After submission, the device will restart. After the modification is successful, please determine the device IP according to the routing system.',
    msg: 'This SDK version does not support modifying network settings',
    inputPort: 'Please enter port',
    serverPort: 'ServerPort',
    protocol: "Protocol",
    inputProtocol: "Please enter one of: http, https"
  },
  threshold: {
    fanSpeed: 'Fan Speed',
    boardTemperature: 'Board Temperature',
    coreTemperature: 'Chip Temperature',
    cpuRate: 'CPU Usage',
    totalMemoryScale: 'Memory Usage',
    // memoryScale: '内存使用率',
    // videoScale: 'video内存使用率',
    tpuScale: 'TPU Memory Usage',
    // externalHardDiskRate: '外挂存储使用率',
    diskRate: 'Disk Usage',
    title: 'Resource Alarm Threshold',
    cloud: 'Cloud Platform Config',
  },
};
