import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const overview: AppRouteModule = {
  path: '/overview',
  name: 'Overview',
  component: LAYOUT,
  redirect: '/overview/index',
  meta: {
    orderNo: 1,
    icon: 'bx:grid-alt',
    title: t('dashboard.basicInfor'),
    // hideChildrenInMenu: true,
    // hideMenu: true,
  },
  children: [
    {
      path: 'index',
      name: 'OverviewPage',
      component: () => import('/@/views/overview/index.vue'),
      meta: {
        title: t('dashboard.ovewview'),
        hideMenu: false,
        // icon: 'simple-icons:about-dot-me',
        // hideBreadcrumb: true, // 在隐藏面包屑中，隐藏当前菜单
      },
    },
    {
      path: 'detail',
      name: 'OverviewDetail',
      component: () => import('/@/views/overview/detail.vue'),
      meta: {
        title: t('dashboard.boardDetail'),
        hideMenu: true,
        // icon: 'simple-icons:about-dot-me',
        // hideBreadcrumb: true, // 在隐藏面包屑中，隐藏当前菜单
      },
    },
    {
      path: 'detail/:sn',
      name: 'OverviewDetailDny',
      component: () => import('/@/views/overview/detail.vue'),
      meta: {
        currentActiveMenu: '/overview/detail',
        hideMenu: true,
        title: t('dashboard.boardDetail'),
        // icon: 'simple-icons:about-dot-me',
        // hideBreadcrumb: true, // 在隐藏面包屑中，隐藏当前菜单
        realPath: '/overview/detail',
        dynamicLevel: 1,
      },
    },
    // {
    //   path: 'threshold',
    //   name: 'Threshold',
    //   component: () => import('/@/views/maintenance/threshold/index.vue'),
    //   meta: {
    //     title: t('dashboard.alarmThreshold'),
    //   },
    // },
  ],
};

export default overview;
