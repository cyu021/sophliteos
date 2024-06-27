import type { AppRouteModule } from '/@/router/types';
import { t } from '/@/hooks/web/useI18n';
import { LAYOUT } from '/@/router/constant';

const accessAlgo: AppRouteModule = {
  path: '/accessAlgo',
  name: 'accessAlgo',
  component: LAYOUT,
  redirect: '/accessAlgo/videos',
  meta: {
    orderNo: 2,
    // hideChildrenInMenu: true,
    icon: 'tdesign:task',
    title: t('routes.dashboard.accessAlgo'),
  },
  children: [
    {
      path: 'videoManage',
      name: 'videoManage',
      component: () => import('/@/views/accessAlgo/dataSource/videoManage/videos.vue'),
      meta: {
        title: t('routes.dashboard.videoManage'),
      },
    },
    // {
    //   path: 'algorithmsList',
    //   name: 'AlgorithmsList',
    //   component: () => import('/@/views/accessAlgo/algorithms/index.vue'),
    //   meta: {
    //     title: t('routes.dashboard.algorithm'),
    //   },
    // },
    {
      path: 'taskList',
      name: 'TaskList',
      component: () => import('/@/views/accessAlgo/task/taskList/index.vue'),
      meta: {
        title: t('routes.dashboard.task'),
      },
    },
    {
      path: 'paramList',
      name: 'paramList',
      component: () => import('/@/views/accessAlgo/paramConfig/paramConfig/index.vue'),
      meta: {
        title: t('routes.dashboard.AlgoParamConfig'),
      },
    }
  ],
};

export default accessAlgo;
