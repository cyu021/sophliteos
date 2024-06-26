import { defHttp } from '/@/utils/http/axios';
import axios from 'axios';

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
  delete: '/purge',
  upload: '/spring/pack',
};

const list = () => {
  const host = 'http://' + `${window.location.host}`;
  return axios
    .get(host + PATH.prefix + PATH.list, { headers: { 'Cache-Control': 'no-cache' } })
    .then((res) => {
      let tempPre = document.createElement('pre');
      tempPre.innerHTML = res.data;

      const links = tempPre.querySelectorAll('a[href]');
      const data = Array.from(links).map((item) => {
        return {
          name: item.text,
          url: host + PATH.prefix + PATH.list + '/' + item.text,
        };
      });

      return data;
    });
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
      },
    },
    params,
  );
};

const deleteFile = (name) => {
  const formData = new window.FormData();
  formData.append('springZip', name);

  return defHttp
    .post(
      {
        url: PATH.delete,
        data: formData,
        headers: { 'Content-type': 'multipart/form-data;charset=UTF-8' },
      },
      { apiUrl: PATH.prefix, isTransformResponse: false },
    )
    .then((res) => {
      return res;
    });
};

const apis = {
  list,
  upload,
  deleteFile,
};

export default apis;
