import { defHttp } from '/@/utils/http/axios';

const PATH = {
  list: '/getDynamicModel',
  onOff: '/onoffDynamicModel',
  delete: '/deleteDynamicModel',
  upload: '/dynamic/uploadDynamicModel',
}

const list = () => {
  return defHttp
    .get({ url: PATH.list, }, { apiUrl: '/dynamic', isTransformResponse: false })
    .then((res) => {
      return res.data
    })
};

const upload = (params, onUploadProgress) => {
  return defHttp.uploadFile(
  {
    url: PATH.upload,
    onUploadProgress,
    timeout: 1000 * 60 * 5,
    requestOptions: {
      ignoreCancelToken: false,
      isTransformResponse: false,
    }
  },
  params,
  );
};

const start = (name) => {
  return defHttp
    .post({ url: PATH.onOff, params: {annotatorName: name, onoff: true} }, { apiUrl: '/dynamic', isTransformResponse: false, joinParamsToUrl: true, })
    .then((res) => {
      console.log(res)
      // return res.data
    })
};

const stop = (name) => {
  return defHttp
    .post({ url: PATH.onOff, params: {annotatorName: name, onoff: false} }, { apiUrl: '/dynamic', isTransformResponse: false, joinParamsToUrl: true, })
    .then((res) => {
      console.log(res)
      // return res.data
    })
};

const deleteFile = (name) => {
  return defHttp
    .delete({ url: PATH.delete, params: { annotatorName: name } }, { apiUrl: '/dynamic', isTransformResponse: false, joinParamsToUrl: true, })
    .then((res) => {
      console.log(res)
      // return res.data
    })
};

const apis = {
  list,
  upload,
  start,
  stop,
  deleteFile,
};

export default apis;
