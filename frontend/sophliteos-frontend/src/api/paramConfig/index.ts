import { defHttp } from '/@/utils/http/axios';
import {
  AlgoTaskInfoType,
  MediaServerParams,
  algoTaskEditParams,
  algoTaskPanrams,
} from './model/index';

enum Api {
  algoConfigMod = '/config/mod',
  getalgoConfig = '/config/get',
  AlgoTaskConfig = '/config/get',
  algoTaskEdit = '/config/mod',
}

export function getalgoConfig(params: MediaServerParams) {
  return defHttp.get({ url: Api.getalgoConfig, params }, { apiUrl: 'algorithm' });
}

export function algoConfigMod(params: any) {
  return defHttp.post({ url: Api.algoConfigMod, params }, { apiUrl: 'algorithm' });
}

export function getAlgoTaskConfig(params: algoTaskPanrams) {
  return defHttp.post<AlgoTaskInfoType>(
    { url: Api.AlgoTaskConfig, params },
    { apiUrl: 'algorithm' },
  );
}

export function editAlgoTaskConfig(params: algoTaskEditParams) {
  return defHttp.post(
    { url: Api.algoTaskEdit, params },
    { apiUrl: 'algorithm', isTransformResponse: false },
  );
}
