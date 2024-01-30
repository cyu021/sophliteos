import { defHttp } from '/@/utils/http/axios';
import { AlarmQueryInfo } from './model';

enum Api {
  alarmList = '/alarm/list', //告警图片列表
  alarmImage = '/image?',
  modSize = '/alarm/modSize',
  alarmQueryInfo = '/alarm/info',
}
export function AlarmList(params: any) {
  return defHttp.post({ url: Api.alarmList, params }, { apiUrl: 'algorithm' });
}
export function getAlarmImage(params: any) {
  return defHttp.get(
    { url: Api.alarmImage, params, responseType: 'blob' },
    { apiUrl: 'algorithm', joinTime: false, isReturnNativeResponse: true },
  );
}

export function modSize(params: any) {
  return defHttp.post({ url: Api.modSize, params }, { apiUrl: 'algorithm' });
}

export function getAlarmQueryInfo() {
  return defHttp.get<AlarmQueryInfo>(
    {
      url: Api.alarmQueryInfo,
    },
    {
      apiUrl: 'algorithm',
    },
  );
}
