import { defHttp } from '/@/utils/http/axios';
import { useGlobSetting } from '/@/hooks/setting';

import {
  IpSetParams,
  AlarmParams,
  RollbackParams,
  // UploadFileParamsSys,
  UploadApiResult,
  stopOrStartContainerParams,
  pullMirrorItem,
  CoreBoardListItem,
  CoreImageListItem,
  CoreContainListItem,
  CoreImageDeleteParam,
  CoreContainParam,
  CoreRegistryParam,
  CoreDockerGetParam,
} from './model/index';
import { BasicApiResponse } from '../model/baseModel';

const { uploadUrl = '' } = useGlobSetting();

enum Api {
  IpSet = '/device/ip',
  Alarm = '/device/configure/alarm',
  Upgrade = '/device/ota/upgrade',
  // Upload = '/api/device/ota/upgrade',
  Rollback = '/device/ota/rollback',
  SsmList = '/ssm/list',
  SysTables = '/device/iptable/get',
  deleteUserIp = '/device/iptable/delete',
  addUserIp = '/device/iptable/add',
  getComIP = '/device/basic',
  modComIP = '/device/mod',
  getHostList = '/terminals/host/list',
  addHost = '/terminals/host/add',
  delHost = '/terminals/host/del',
  editHost = '/terminals/host/update',
  getCommandList = '/terminals/command',
  addCommand = '/terminals/command',
  delCommand = '/terminals/command/del',
  editCommand = '/terminals/command/update',
  terminalsTest = '/terminals/test',
  getStoreList = '/container/insecure/list',
  addStore = '/container/insecure/add',
  delStore = '/container/insecure/delete',
  getMirrorList = '/container/image/list',
  delMirror = '/container/image/delete',
  slectImagelist = '/container/insecure/imagelist',
  slectImageTaglist = '/container/insecure/taglist',
  pullImage = '/container/image/pull',
  containerList = '/container/list',
  delContainer = '/container/delete',
  stopContainer = '/container/stop',
  startContainer = '/container/start',
  getPulledMirrorList = '/container/image/listpulling',
}
enum CoreDockerApi {
  coreBoardList = '/terminals/cores/list',
  coreImageList = '/container/core/image/list',
  coreImageDelete = '/container/core/image/delete',
  coreRegistryList = '/container/core/insecure/list',
  coreRegistryAdd = '/container/core/insecure/add',
  coreRegistryDelete = '/container/core/insecure/delete',
  coreContainList = '/container/core/list',
  coreContainDelete = '/container/core/delete',
  coreContainStop = '/container/core/stop',
  coreContainStart = '/container/core/start',
}

// IP地址设置
export function ipSet(params: IpSetParams) {
  return defHttp.post<BasicApiResponse>({ url: Api.IpSet, params }, { isTransformResponse: false });
}

// IP地址查询
export function ipGet() {
  return defHttp.get({ url: Api.IpSet });
}

// 告警阈值设置
export function setAlarm(params: AlarmParams) {
  return defHttp.post<BasicApiResponse>({ url: Api.Alarm, params });
}

// 告警阈值查询
export function getAlarm() {
  return defHttp.get({ url: Api.Alarm });
}

// OTA一键升级
export function upgradeApi(params, onUploadProgress: (progressEvent: ProgressEvent) => void) {
  return defHttp.uploadFile<UploadApiResult>(
    {
      url: uploadUrl,
      onUploadProgress,
      // timeout: 1000 * 60 * 60 * 24,   超时时间改成15分钟
      timeout: 1000 * 60 * 15,
      // @ts-ignore
      requestOptions: {
        ignoreCancelToken: false,
        isReturnNativeResponse: true,
      },
    },
    params,
  );
}
export function checkFileList() {
  return defHttp.get({
    url: '/device/ota/list',
  });
}

//上传文件
export function checkFile(params, onUploadProgress: (progressEvent: ProgressEvent) => void) {
  return defHttp.uploadFile<UploadApiResult>(
    {
      url: '/api/device/ota/file',
      onUploadProgress,
      timeout: 1000 * 60 * 60 * 24,
      // timeout: 1000 * 60 * 15,超时时间改成15分钟
      // @ts-ignore
      requestOptions: {
        ignoreCancelToken: false,
        isTransformResponse: false,
      },
    },
    params,
  );
}

// 软件升级
export function upgradeSoftApi(params, onUploadProgress: (progressEvent: ProgressEvent) => void) {
  return defHttp.uploadFile<UploadApiResult>(
    {
      url: '/api/upgrade',
      onUploadProgress,
      timeout: 1000 * 60 * 15,
      // @ts-ignore
      requestOptions: {
        ignoreCancelToken: false,
        isReturnNativeResponse: true,
      },
    },
    params,
  );
}

// ssm升级
export function upgradeSsmApi(params, onUploadProgress: (progressEvent: ProgressEvent) => void) {
  return defHttp.uploadFile<UploadApiResult>(
    {
      url: '/api/ssm/upgrade',
      onUploadProgress,
      timeout: 1000 * 60 * 15,
      // @ts-ignore
      requestOptions: {
        ignoreCancelToken: false,
        isReturnNativeResponse: true,
      },
    },
    params,
  );
}

// OTA升级状态
export function upgradeStatusApi() {
  return defHttp.get({ url: Api.Upgrade });
}

// OTA一键回滚
export function rollbackApi(params: RollbackParams) {
  return defHttp.post({ url: Api.Rollback, params });
}

// OTA升级状态
export function ssmStatusApi() {
  return defHttp.get({ url: Api.SsmList });
}

export function getSysTables() {
  const res = defHttp.get({ url: Api.SysTables }).then((res) => {
    return res.sysTables;
  });

  return res;
}
export function getUserTables() {
  const res = defHttp.get({ url: Api.SysTables }).then((res) => {
    return res.userTables;
  });

  return res;
}

export function DeleteUserTables(params: any) {
  return defHttp.post({ url: Api.deleteUserIp, params });
}

export function addUserMap(params: any) {
  return defHttp.post({ url: Api.addUserIp, params });
}

export function getComIP() {
  return defHttp.get({ url: Api.getComIP }).then((res) => {
    return res.configure;
  });
}

export function modComIP(params: any) {
  return defHttp.post({ url: Api.modComIP, params });
}

// 获取主机列表
export function getHostList() {
  return defHttp.get({ url: Api.getHostList });
}
// 增加主机
export function addHost(params: any) {
  return defHttp.post({ url: Api.addHost, params });
}
//删除主机
export function delHost(params: any) {
  return defHttp.post({ url: Api.delHost, params });
}
//修改主机
export function editHost(params: any) {
  return defHttp.post({ url: Api.editHost, params });
}
// 获取快速命令列表
export function getCommandList() {
  return defHttp.get({ url: Api.getCommandList });
}
// 增加快速命令
export function addCommand(params: any) {
  return defHttp.post({ url: Api.addCommand, params });
}
//删除快速命令
export function delCommand(params: any) {
  return defHttp.post({ url: Api.delCommand, params });
}
//修改快速命令
export function editCommand(params: any) {
  return defHttp.post({ url: Api.editCommand, params });
}
//ssh连通性测试
export function terminalsTest(params: any) {
  return defHttp.post({ url: Api.terminalsTest, params });
}
// 获取容器仓库列表
export function getStoreList() {
  return defHttp.get({ url: Api.getStoreList });
}
// 添加私仓
export function addStore(params: any) {
  return defHttp.post({ url: Api.addStore, params });
}
// 删除私仓
export function delStore(params: any) {
  return defHttp.post({ url: Api.delStore, params });
}
// 获取容器镜像列表
export function getMirrorList() {
  return defHttp.get({ url: Api.getMirrorList });
}
// 容器删除镜像
export function delMirror(params: any) {
  return defHttp.post({ url: Api.delMirror, params });
}
// 获取算法仓容器名列表
export function slectImagelist(params: any) {
  return defHttp.post({ url: Api.slectImagelist, params });
}
// 获取算法仓特定容器tag列表
export function slectImageTaglist(params: any) {
  return defHttp.post({ url: Api.slectImageTaglist, params });
}
// 拉取镜像
export function pullImage(params: any) {
  return defHttp.post({ url: Api.pullImage, params });
}
// 获取容器列表
export function containerList() {
  return defHttp.get({ url: Api.containerList });
}
// 删除容器
export function delContainer(params: any) {
  return defHttp.post({ url: Api.delContainer, params });
}
// 停止容器
export function stopContainer(params: stopOrStartContainerParams) {
  return defHttp.post({ url: Api.stopContainer, params });
}

// 启动容器
export function startContainer(params: stopOrStartContainerParams) {
  return defHttp.post({ url: Api.startContainer, params });
}

// 获取已拉取从镜像列表
export function getPulledMirrorList() {
  return defHttp.get<pullMirrorItem[]>({ url: Api.getPulledMirrorList });
}

// -----S  核心板docker容器页面   S-----

// 核心板列表
export function getCoreBoardList() {
  return defHttp.get<CoreBoardListItem[]>({ url: CoreDockerApi.coreBoardList });
}

// 镜像列表
export function getCoreImageList(params: CoreDockerGetParam) {
  return defHttp.get<CoreImageListItem[]>({ url: CoreDockerApi.coreImageList, params });
}
// 镜像删除
export function deleteCoreImage(params: CoreImageDeleteParam) {
  return defHttp.post({ url: CoreDockerApi.coreImageDelete, params });
}

// 容器列表
export function getCoreContainList(params: CoreDockerGetParam) {
  return defHttp.get<CoreContainListItem[]>({ url: CoreDockerApi.coreContainList, params });
}
// 容器删除
export function deleteCoreContain(params: CoreContainParam) {
  return defHttp.post({ url: CoreDockerApi.coreContainDelete, params });
}
// 容器停止
export function coreContainStop(params: CoreContainParam) {
  return defHttp.post({ url: CoreDockerApi.coreContainStop, params });
}
// 容器启动
export function coreContainStart(params: CoreContainParam) {
  return defHttp.post({ url: CoreDockerApi.coreContainStart, params });
}

// 仓库列表
export function getCoreRegistryList(params: CoreDockerGetParam) {
  return defHttp.get<string[]>({ url: CoreDockerApi.coreRegistryList, params });
}
// 仓库删除
export function deleteCoreRegistry(params: CoreRegistryParam) {
  return defHttp.post({ url: CoreDockerApi.coreRegistryDelete, params });
}
// 添加仓库
export function addCoreRegistry(params: CoreRegistryParam) {
  return defHttp.post({ url: CoreDockerApi.coreRegistryAdd, params });
}

// -----E  核心板docker容器页面   E-----
