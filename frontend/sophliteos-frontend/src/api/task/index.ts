import { defHttp } from '/@/utils/http/axios';
import { BasicApiResponse } from '../model/baseModel';
import {
  MediaServerParams,
  AbilityListGetResultModel,
  UpUrlGetResultModal,
  upUrlParams,
  UpUrlAddResultModal,
} from './model/index';
enum Api {
  taskList = '/task/list',
  deleteTask = '/task/delete',
  startTask = '/task/start',
  addTask = '/task/add',
  modTask = '/task/modify',
  stopTask = '/task/stop',
  addAlgorithm = '/server/add',
  getAlgorithm = '/server/get',
  getAbilites = '/task/abilities',
  getUpUrl = '/alarm/upload/get',
  addUpUrl = '/alarm/upload/mod',
  getRotateCfg = '/rotate/config/get',
  modRotateCfg = '/rotate/config/mod'
}

export function getTaskList(params?: MediaServerParams) {
  return defHttp.post({ url: Api.taskList, params }, { apiUrl: 'algorithm' });
}
export function PostDeleteTask(params: any) {
  return defHttp.post({ url: Api.deleteTask, params }, { apiUrl: 'algorithm' });
}

export function StartTask(params: any) {
  return defHttp.post({ url: Api.startTask, params }, { apiUrl: 'algorithm' });
}

export function StopTask(params: any) {
  return defHttp.post({ url: Api.stopTask, params }, { apiUrl: 'algorithm' });
}

export function getAlgorithm() {
  return defHttp.get({ url: Api.getAlgorithm }, { apiUrl: 'algorithm' });
}
export function addAlgorithm(params: any) {
  return defHttp.post({ url: Api.addAlgorithm, params }, { apiUrl: '/algorithm' });
}
export function addTask(params: any) {
  return defHttp.post<BasicApiResponse>({ url: Api.addTask, params }, { apiUrl: '/algorithm' });
}
export function modTask(params: any) {
  return defHttp.post<BasicApiResponse>({ url: Api.modTask, params }, { apiUrl: '/algorithm' });
}
export const getAbilites = () =>
  defHttp.get<AbilityListGetResultModel>(
    {
      url: Api.getAbilites,
    },
    {
      apiUrl: 'algorithm',
    },
  );

export const getUpUrl = () =>
  defHttp.get<UpUrlGetResultModal>(
    {
      url: Api.getUpUrl,
    },
    {
      apiUrl: 'algorithm', isTransformResponse: false
    },
  ).then((res) => {
    if (res.Code == 0) {
      return res.result;
    }
    return 'http://127.0.0.1:8081';
  }
  );

export const addUpUrl = (params: upUrlParams) =>
  defHttp.post<UpUrlAddResultModal>(
    {
      url: Api.addUpUrl,
      params,
    },
    {
      apiUrl: 'algorithm', isTransformResponse: false
    },
  );

  export function getRotateCfg() {
    return defHttp.get({ url: Api.getRotateCfg }, { apiUrl: '/algorithm', isTransformResponse: false });
  }
  export function modRotateCfg(params: any) {
    return defHttp.post({ url: Api.modRotateCfg, params }, { apiUrl: '/algorithm', isTransformResponse: false });
  }