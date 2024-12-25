import type { AppRouteModule } from '/@/router/types';
import { t } from '/@/hooks/web/useI18n';
import { LAYOUT } from '/@/router/constant';

const privilegeMgmt: AppRouteModule = {
  path: '/privilegeMgmt',
  name: 'privilegeMgmt',
  component: LAYOUT,
  redirect: '/privilegeMgmt/privileges',
  meta: {
    orderNo: 7,
    // hideChildrenInMenu: true,
    icon: 'tdesign:task',
    title: t('dashboard.privilegeMgmt'),
  },
  children: [
    {
      path: 'acctManage',
      name: 'acctManage',
      component: () => import('/@/views/privilegeMgmt/acctMgmt/index.vue'),
      meta: {
        title: t('dashboard.acctManage'),
      },
    },
    {
      path: 'privilegeConfig',
      name: 'privilegeConfig',
      component: () => import('/@/views/privilegeMgmt/privilegeConfig/index.vue'),
      meta: {
        title: t('dashboard.privilegeConfig'),
      },
    }
  ],
};

export default privilegeMgmt;
