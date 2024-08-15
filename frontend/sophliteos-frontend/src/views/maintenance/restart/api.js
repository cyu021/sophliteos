import { defHttp } from '/@/utils/http/axios';

const hardwareRestart = () => {
  return defHttp.post({ url: '/box/restart' }, { apiUrl: '/ota', isTransformResponse: false })
}

const apis = {
  hardwareRestart,
};

export default apis;
