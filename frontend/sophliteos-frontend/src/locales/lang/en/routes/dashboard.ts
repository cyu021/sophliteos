export default {
  dashboard: 'Dashboard',
  about: 'About',
  workbench: 'Workbench',
  analysis: 'Analysis',
  basicInfor: 'Information',
  ovewview: 'Overview',
  boardDetail: 'Board Detail',
  logs: 'Logs',
  warning: 'Alarm',
  operate: 'Operation',
  threshold: 'Threshold',
  maintenance: 'Maintenance',
  appSoft: 'Application Upgrade',
  sysSoft: 'System Upgrade',
  softUpdate: 'Software Upgrade',
  ssmUpdate: 'SSM Upgrade',
  sysSetting: 'System Setting',
  liteOsUpdate: 'LiteOS Upgrade',
  networkSetting: 'Network Setting',
  sshTerminal: 'SSH Terminal',
  containers: 'Containers',
  mirror: 'Mirror',
  store: 'Repository',
  hostComputer: 'Host Computer',
  timeSetting: 'Time Setting',
  algorithm: 'Algorithm Manage',
  eventList: 'Event List',
  algorithmList: 'Algorithm List',
  task: 'Task Manage',
  taskList: 'Task List',
  taskSubscribe: 'Task Subscribe',
  accessAlgo: 'Algorithm business',
  dataSource: 'DataSource',
  mediaServers: 'Streaming Services',
  videoManage: 'Video Resource',
  AlgoParamConfig: 'Parameter config',
  alarmRetrieval: 'Alarm Retrieval',
  alarmDetail: 'Alarm Detail',
  coreBoardMap: 'Core Board port Map',
  logDownload: 'logs Download',
  content: {
    sysContent:
      'System OTA upgrade, suitable for version upgrade, does not support downgrade and cross version upgrade, otherwise it may cause system damage.',
    Content: `The system SDK OTA upgrade does not support cross 3.0 version upgrades, otherwise it may cause system damage.
    1. To upgrade the control board, please upload the card swiping file. After the control board upgrade is completed, proceed with the core board upgrade. Select the core board that needs to be upgraded, and there is no need to upload the file. Click Upgrade directly.
    2. Upgrade the core board, upload the tftp flash package, and select the core board for upgrade. Transferring the wrong upgrade package may cause system damage, please operate with caution.`,
    ssmContent:
      'SSM software is a management service deployed on AI computing devices, providing functions such as system management, system upgrade, computing device information acquisition, and algorithm application deployment.',
    liteContent:
      'LiteOS software is based on SSM device management services and provides functions such as device monitoring, alarm query software, and system upgrade on WEB pages',
  },
};
