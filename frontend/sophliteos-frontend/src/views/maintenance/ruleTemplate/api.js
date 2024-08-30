import { getVideosList } from '/@/api/dataSource';
import { getTaskList } from '/@/api/task';
import { defHttp } from '/@/utils/http/axios';

const attrList = () => {
  return defHttp
    .post({ url: '/list' }, { apiUrl: '/attr', isTransformResponse: false, joinParamsToUrl: false })
    .then((res) => {
      console.log('attr list', res);
      return res.AlgoAttrDict;
    });
};

const modelList = () => {
  return defHttp
    .get({ url: '/getDynamicModel' }, { apiUrl: '/dynamic', isTransformResponse: false })
    .then((res) => {
      return res.data;
    });
};

const defaultExtend = (annotator) => {
  return attrList().then((aRes) => {
    const attr = aRes[annotator];
    if (!attr) {
      throw 'no annotator found in attr list, annoatorname' + annotator;
    }

    const label = Object.keys(attr)[0];
    const kvp = Object.keys(attr[label]['kvp'])[0];
    const rule = {
      Grp: 1,
      Op: '=',
      K: kvp,
      V: attr[label]['kvp'][kvp][0] || 0,
    };

    const value = {
      FilterName: '规则名称1',
      FilterPeriod: [
        {
          DateStart: '20240101',
          DateEnd: '20241230',
          TimeStart: '0800',
          TimeEnd: '2000',
          Weekday: '0,1,2,3,4,5,6',
        },
      ],
      Filter: [
        {
          Annotator: annotator,
          Label: '1420',
          Rtsp: 'rtsp://192.168.0.158:8554/sample',
          Rule: [rule],
        },
      ],
    };

    return value;
  });
};

const ruleTemplateList = () => {
  return defHttp
    .get(
      { url: '/template/get' },
      { apiUrl: '/filter', isTransformResponse: false, joinParamsToUrl: false },
    )
    .then((res) => {
      return res.FilterTemplate;
    });
};

const deleteRuleTemplate = (modalName, filterName) => {
  const params = {
    anno: modalName,
    filterName: filterName,
  };

  return defHttp.post(
    { url: '/template/del', params },
    { apiUrl: '/filter', isTransformResponse: false },
  );
};

const updateRuleTemplate = (modelName, extend) => {
  const params = { FilterTemplate: {} };
  params.FilterTemplate[modelName] = [extend];

  return defHttp.post(
    { url: '/template/upsert', params },
    { apiUrl: '/filter', isTransformResponse: false },
  );
};

const updateRuleTemplateJson = (params) => {
  return defHttp.post(
    { url: '/template/upsert', params },
    { apiUrl: '/filter', isTransformResponse: false },
  );
};

const apis = {
  attrList,
  modelList,
  ruleTemplateList,
  deleteRuleTemplate,
  updateRuleTemplate,
  updateRuleTemplateJson,
};

export default apis;
