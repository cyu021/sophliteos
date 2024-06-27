import { defHttp } from '/@/utils/http/axios';

const currentVersion = () => {
  return defHttp
    .get({ url: PATH.currentVersion }, { apiUrl: PATH.prefix, isTransformResponse: false })
    .then((res) => {
      if (res.code == 0) {
        return res.data.vps;
      }

      return '-';
    });
};

const upload = (params, onUploadProgress) => {
  return defHttp.uploadFile(
    {
      url: PATH.upload,
      baseURL: PATH.prefix,
      onUploadProgress,
      timeout: 1000 * 60 * 10,
      requestOptions: {
        ignoreCancelToken: false,
        isTransformResponse: false,
      },
    },
    params,
  );
};

const upgrade = () => {
  return defHttp.post({ url: PATH.upgrade }, { apiUrl: PATH.prefix, isTransformResponse: false });
};

const restart = () => {
  return defHttp.post({ url: PATH.restart }, { apiUrl: PATH.prefix, isTransformResponse: false });
};

const deleteCache = () => {};

const PATH = {
  prefix: '/ota/core',
  upload: '/upload',
  upgrade: '/upgrade',
  currentVersion: '/currver',
  restart: '/restart',
};

const apisCore = {
  upload,
  upgrade,
  restart,
  currentVersion,
  deleteCache,
};

export default apisCore;
