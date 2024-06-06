import { defHttp } from '/@/utils/http/axios';
import { getTaskList } from '/@/api/task';

const PATH = {
  list: '/getDynamicModel',
  onOff: '/onoffDynamicModel',
  delete: '/deleteDynamicModel',
  upload: '/dynamic/uploadDynamicModel',
};

const AbilitiesPATH = {
  prefix: '/abilities',
  list: '/list',
  token_get: '/token/get',
  set: '/set',
  unset: '/unset',
  token_release: '/token/release',
};

const abilitiesList = () => {
  return defHttp
    .post(
      { url: AbilitiesPATH.list },
      { apiUrl: AbilitiesPATH.prefix, isTransformResponse: false, joinParamsToUrl: true },
    )
    .then((res) => {
      return res.Abilities;
    });
};

const abilitiesGetToken = () => {
  return defHttp
    .post(
      { url: AbilitiesPATH.token_get },
      { apiUrl: AbilitiesPATH.prefix, isTransformResponse: false, joinParamsToUrl: true },
    )
    .then((res) => {
      console.log('abilities token get', res);
      return res.Token;
    });
};

const abilitiesReleaseToken = (params) => {
  return defHttp
    .post(
      { url: AbilitiesPATH.token_release, params: params },
      { apiUrl: AbilitiesPATH.prefix, isTransformResponse: false, joinParamsToUrl: false },
    )
    .then((res) => {
      console.log('abilities token release', res);
      return res;
    });
};

/**
 * { token: xxxxx, full_structure: 16 }
 */
const abilitiesSet = (params) => {
  return defHttp
    .post(
      { url: AbilitiesPATH.set, params: params },
      { apiUrl: AbilitiesPATH.prefix, isTransformResponse: false, joinParamsToUrl: false },
    )
    .then((res) => {
      console.log('abilities set', res);
      return res;
    });
};

const abilitiesUnset = (params) => {
  return defHttp
    .post(
      { url: AbilitiesPATH.unset, params: params },
      { apiUrl: AbilitiesPATH.prefix, isTransformResponse: false, joinParamsToUrl: false },
    )
    .then((res) => {
      console.log('abilities unset', res);
      return res;
    });
};

const dynamicList = () => {
  return defHttp
    .get({ url: PATH.list }, { apiUrl: '/dynamic', isTransformResponse: false })
    .then((res) => {
      return res.data;
    });
};

const list = () => {
  return Promise.all([abilitiesList(), dynamicList()]).then((res) => {
    let abilities = res[0];
    let data = res[1];

    if (abilities == null) {
      abilities = {};
    }

    data.forEach((item) => {
      const name = item.annotator_name;
      item.sts = Object.values(abilities).includes(name) ? 1 : 0;
    });

    console.log('abilities, dynamic', res);
    return data;
  });
};

const uploadFile = (params, onUploadProgress) => {
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

const upload = async (params, onUploadProgress) => {
  const uploadRes = await uploadFile(params, onUploadProgress);
  const annotatorName = uploadRes.data.data.annotator_name; // also is algorithm name

  const taskList = await getTaskList({ pageNo: -1, pageSize: -1 });

  const needStop =
    taskList.items.find((item) => {
      return item.abilities.includes(annotatorName) && item.status === 1;
    }) !== undefined;
  if (needStop) {
    throw 'ERROR: please stop the tasks include [' + annotatorName + '] first.';
  }

  const token = await abilitiesGetToken();
  let aList = await abilitiesList();
  if (aList == null) {
    aList = {};
  }
  const keys = Object.keys(aList);
  const index = keys.length > 0 ? Math.max(...keys.map(Number)) + 1 : 1;

  await abilitiesSet({ token: token, [annotatorName]: '' + index });
  await abilitiesReleaseToken({ token: token });

  return true;
};

const dynamicStart = (name) => {
  return defHttp.post(
    { url: PATH.onOff, params: { annotatorName: name, onoff: true } },
    { apiUrl: '/dynamic', isTransformResponse: false, joinParamsToUrl: true },
  );
};

const dynamicStop = (name) => {
  return defHttp.post(
    { url: PATH.onOff, params: { annotatorName: name, onoff: false } },
    { apiUrl: '/dynamic', isTransformResponse: false, joinParamsToUrl: true },
  );
};

const start = (name) => {
  return dynamicStart(name);
};

const stop = async (name) => {
  const taskList = await getTaskList({ pageNo: -1, pageSize: -1 });

  const needStop =
    taskList.items.find((item) => {
      return item.abilities.includes(name) && item.status === 1;
    }) !== undefined;
  if (needStop) {
    throw 'ERROR: please stop the tasks include [' + name + '] first.';
  }

  await dynamicStop(name);
  return true;
};

const deleteFile = (name) => {
  return defHttp
    .delete(
      { url: PATH.delete, params: { annotatorName: name } },
      { apiUrl: '/dynamic', isTransformResponse: false, joinParamsToUrl: true },
    )
    .then((res) => {
      console.log(res);
      return res;
    });
};

const unload = async (name) => {
  const taskList = await getTaskList({ pageNo: -1, pageSize: -1 });

  const needStop =
    taskList.items.find((item) => {
      return item.abilities.includes(name) && item.status === 1;
    }) !== undefined;
  if (needStop) {
    throw 'ERROR: please stop the tasks include [' + name + '] first.';
  }

  await deleteFile(name);

  const token = await abilitiesGetToken();
  const aList = await abilitiesList();

  const index = Object.keys(aList).find((key) => aList[key] === name);
  await abilitiesUnset({ token: token, [name]: '' + index });
  await abilitiesReleaseToken({ token: token });

  return true;
};

const apis = {
  list,
  upload,
  start,
  stop,
  unload,
};

export default apis;
