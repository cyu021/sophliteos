import { defHttp } from '/@/utils/http/axios';
import { getTaskList } from '/@/api/task';

// 接口:
// //用model+annoName打包spring zip, lock直到完成, 强制一次只有一个操作
// curl -H 'Content-Type: multipart/form-data'
//      -F"annoName=detect"
//      -F "model=@./KM_harpy_FASTER-RCNN_nart_sophgo-bm1684-fp32-sdkV221201_b1_1.0.2.model;type=application/x-zip-compressed"
//      -X POST http://10.162.18.26:8080/spring/pack

// //以file server页面形式显示可下载的spring zip http://10.162.18.26:8080/springpacks/

// //删除指定spring filename的zip
// curl -H 'Content-Type: multipart/form-data' -F"springZip=spring_detect.zip" -X POST http://10.162.18.26:8080/spring/purge

const PATH = {
  prefix: '/spring',
  list: '/zips',
  purge: '/purge',
  upload: '/spring/pack',
};

const list = () => {
  return defHttp
    .get({ url: PATH.list }, { apiUrl: PATH.prefix, isTransformResponse: false })
    .then((res) => {
      console.log('list', res);

      return res.data;
    });
};

const upload = (params, onUploadProgress) => {
  console.log('upload', params);
  return defHttp.uploadFile(
    {
      url: PATH.upload,
      onUploadProgress,
      timeout: 1000 * 60 * 5,
      requestOptions: {
        ignoreCancelToken: false,
        isTransformResponse: false,
      },
    },
    params,
  );
};

const purge = (name) => {
  return defHttp
    .delete(
      { url: PATH.delete, params: { annotatorName: name } },
      { apiUrl: PATH.prefix, isTransformResponse: false, joinParamsToUrl: true },
    )
    .then((res) => {
      console.log(res);
      return res;
    });
};

const download = (name) => {
  return {};
};

const apis = {
  list,
  upload,
  purge,
  download,
};

export default apis;
