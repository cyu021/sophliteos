import { getVideosList, LivePreview } from '/@/api/dataSource/index';
import { getAbilites } from '/@/api/task/index';

const videosList = () => {
  return getVideosList()
};

const preview = (param) => {
  return LivePreview(param)
}

const getAbilitesTypes = () => {
  return getAbilites()
}

const apis = {
  videosList,
  preview,
  getAbilitesTypes,
};

/**
 * TODO: 
 * 1. 获取视频源列表
 * 2. 
 */

export default apis;
