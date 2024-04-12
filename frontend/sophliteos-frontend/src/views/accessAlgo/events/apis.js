import { getVideosList, LivePreview } from '/@/api/dataSource/index';

const videosList = () => {
  return getVideosList()
};

const preview = (param) => {
  return LivePreview(param)
}

const apis = {
  videosList,
  preview,
};

/**
 * TODO: 
 * 1. 获取视频源列表
 * 2. 
 */

export default apis;
