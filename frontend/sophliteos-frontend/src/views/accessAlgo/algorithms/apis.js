import { defHttp } from '/@/utils/http/axios';
import { getTaskList } from '/@/api/task';

const tasks = () => {
  return getTaskList({ pageNo: -1, pageSize: -1 });
};

const PATH = {
  list: '/getDynamicModel',
  onOff: '/onoffDynamicModel',
  delete: '/deleteDynamicModel',
  upload: '/dynamic/uploadDynamicModel',
}

const AbilitiesPATH = {
  prefix: '/abilities',
  list: '/list',
  token_get: '/token/get',
  set: '/set',
  unset: '/unset',
  token_release: '/token/release'
}

const abilitiesList = () => {
  return defHttp
    .post({ url: AbilitiesPATH.list, }, { apiUrl: AbilitiesPATH.prefix, isTransformResponse: false, joinParamsToUrl: true, })
    .then((res) => {
      return res.Abilities
    })
}

const abilitiesGetToken = () => {
  return defHttp
  .post({ url: AbilitiesPATH.token_get, }, { apiUrl: AbilitiesPATH.prefix, isTransformResponse: false, joinParamsToUrl: true, })
  .then((res) => {
    console.log('abilities token get', res)
    return res.Token
  })
}

const abilitiesReleaseToken = (params) => {
  return defHttp
  .post({ url: AbilitiesPATH.token_release, params: params }, { apiUrl: AbilitiesPATH.prefix, isTransformResponse: false, joinParamsToUrl: false, })
  .then((res) => {
    console.log('abilities token release', res)
    return res
  })
}

/**
 * { token: xxxxx, full_structure: 16 }
 */
const abilitiesSet = (params) => {
  return defHttp
  .post({ url: AbilitiesPATH.set, params: params }, { apiUrl: AbilitiesPATH.prefix, isTransformResponse: false, joinParamsToUrl: false, })
  .then((res) => {
    console.log('abilities set', res)
    return res
  })
}

const abilitiesUnset = (params) => {
  return defHttp
  .post({ url: AbilitiesPATH.unset, params: params }, { apiUrl: AbilitiesPATH.prefix, isTransformResponse: false, joinParamsToUrl: false, })
  .then((res) => {
    console.log('abilities unset', res)
    return res
  })
}

const dynamicList = () => {
  return defHttp
    .get({ url: PATH.list, }, { apiUrl: '/dynamic', isTransformResponse: false })
    .then((res) => {
      return res.data
    })
}

const list = () => {
  return Promise.all([abilitiesList(), dynamicList()])
    .then((res) => {
      const abilities = res[0]
      let data = res[1]

      data.forEach(item => {
        const name = item.annotator_name;
        item.sts = Object.values(abilities).includes(name) ? 1 : 0
      });

      console.log('abilities, dynamic', res);
      return data;
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

const dynamicStart = (name) => {
  return defHttp
    .post({ url: PATH.onOff, params: { annotatorName: name, onoff: true } }, { apiUrl: '/dynamic', isTransformResponse: false, joinParamsToUrl: true, })
}

const dynamicStop = (name) => {
  return defHttp
    .post({ url: PATH.onOff, params: { annotatorName: name, onoff: false } }, { apiUrl: '/dynamic', isTransformResponse: false, joinParamsToUrl: true, })
}

const start = (name) => {
  // 1. get token 
  // 2. get abilities list -> get newest id 
  // 3. dynamic start 
  // 4. abilities set 
  // 5. token release
  return Promise.all([abilitiesGetToken(), abilitiesList()])
  .then((res) => {
    const token = res[0]
    const list = res[1]
    const keys = Object.keys(list);
    const index = keys.length > 0 ? (Math.max(...keys.map(Number)) + 1) : 1;

    return dynamicStart(name)
    .then((res) => {
      return abilitiesSet({ token: token, [name]: '' + index })
    })
    .then((res) => {
      return abilitiesReleaseToken({ token: token })
    })
  })
};

const stop = (name) => {
  // prefix: check if task exist 
  // 1. get token 
  // 2. get abilities list -> get ability index 
  // 3. dynamic stop 
  // 4. abilities unset 
  // 5. token release

  return tasks().then((res) => {
    let existNames = []

    if (res.items) {
      res.items.forEach((item) => {
        existNames = existNames.concat(item.abilities)
      })  
    }

    if (existNames.includes(name)) {
      throw new Error('当前算法包存在任务使用，请先删除对应任务');
    }

    return true;
  })
  .then((res) => {
    return Promise.all([abilitiesGetToken(), abilitiesList()])
    .then((res) => {
      const token = res[0]
      const list = res[1]
      const index = Object.keys(list).find(key => list[key] === name);
  
      return dynamicStop(name)
      .then((res) => {
        return abilitiesUnset({ token: token, [name]: '' + index })
      })
      .then((res) => {
        return abilitiesReleaseToken({ token: token })
      })
    })  
  })
};

const deleteFile = (name) => {
  return defHttp
    .delete({ url: PATH.delete, params: { annotatorName: name } }, { apiUrl: '/dynamic', isTransformResponse: false, joinParamsToUrl: true, })
    .then((res) => {
      console.log(res)
      return res
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
