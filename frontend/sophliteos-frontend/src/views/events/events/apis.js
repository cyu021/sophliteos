import { LivePreview } from '/@/api/dataSource/index';
import { getAbilites } from '/@/api/task/index';
import { getAlgoTaskConfig } from '/@/api/paramConfig/index';
import { getTaskList } from '/@/api/task';

const tasks = () => {
  return getTaskList({ pageNo: -1, pageSize: -1 });
};

const preview = (param) => {
  return LivePreview(param)
}

const getAbilitesTypes = () => {
  return getAbilites()
}

const getRois = (param) => {
  return getAlgoTaskConfig(param)
}

const apis = {
  tasks,
  preview,
  getAbilitesTypes,
  getRois,
};

export default apis;
