import type { AppRouteModule } from '/@/router/types';
import { t } from '/@/hooks/web/useI18n';
import { LAYOUT } from '/@/router/constant';

const algoModel: AppRouteModule = {
  path: '/algoModel',
  name: 'algoModel',
  component: LAYOUT,
  meta: {
    orderNo: 4,
    // hideChildrenInMenu: true,
    icon: 'tdesign:task',
    title: t('dashboard.algoModel'),
  },
  children: [
    {
      path: 'algorithmsList',
      name: 'AlgorithmsList',
      component: () => import('/@/views/accessAlgo/algorithms/index.vue'),
      meta: {
        title: t('dashboard.algorithm'),
      },
    },
    {
      path: 'prepSpringZip',
      name: 'PrepSpringZip',
      component: () => import('/@/views/maintenance/prepSpringZip/index.vue'),
      meta: {
        title: t('dashboard.prepSpringZip'),
      },
    },
  ],
};

export default algoModel;
