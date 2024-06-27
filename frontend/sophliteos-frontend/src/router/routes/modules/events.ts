import type { AppRouteModule } from '/@/router/types';
import { t } from '/@/hooks/web/useI18n';
import { LAYOUT } from '/@/router/constant';

const events: AppRouteModule = {
  path: '/event',
  name: 'events',
  component: LAYOUT,
  redirect: '/events/events',
  meta: {
    orderNo: 3,
    // hideChildrenInMenu: true,
    icon: 'tdesign:task',
    title: t('routes.dashboard.eventManage'),
  },
  children: [
    {
      path: 'eventList',
      name: 'EvnetList',
      component: () => import('/@/views/events/events/index.vue'),
      meta: {
        title: t('routes.dashboard.eventList'),
      },
    },
    {
      path: 'AlarmRetrieval',
      name: 'AlarmRetrieval',
      component: () => import('/@/views/events/alarmRetrieval/index.vue'),
      meta: {
        title: t('routes.dashboard.alarmRetrieval'),
        hideChildrenInMenu: true,
      },
      children: [
        {
          path: 'AlarmDetail/:image',
          name: 'AlarmDetail',
          component: () => import('/@/views/events/alarmRetrieval/alarmDetail.vue'),
          meta: {
            title: t('routes.dashboard.alarmDetail'),
            // hideMenu: true,
            // hideBreadcrumb: true, // 在隐藏面包屑中，隐藏当前菜单
            // hideTab: true,
          },
        },
      ],
    },
  ],
};

export default events;
