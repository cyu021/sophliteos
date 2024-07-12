import type { AppRouteModule } from '/@/router/types';
import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';
// import { useDeviceInfo } from '/@/store/modules/overview';
// const deviceStore = useDeviceInfo();
// console.log(deviceStore.isSingleBoard);
// setTimeout(() => {
//   console.log(deviceStore.isSingleBoard);
// }, 3000);
const maintenance: AppRouteModule = {
  path: '/maintenance',
  name: 'Maintenance',
  component: LAYOUT,
  redirect: '/maintenance/sysSoft',
  meta: {
    orderNo: 5,
    icon: 'bx:server',
    title: t('dashboard.maintenance'),
    // hideMenu: true,
  },
  children: [
    // {
    //   path: 'prepSpringZip',
    //   name: 'PrepSpringZip',
    //   component: () => import('/@/views/maintenance/prepSpringZip/index.vue'),
    //   meta: {
    //     title: t('dashboard.prepSpringZip'),
    //   },
    // },
    {
      path: 'sysSoft',
      name: 'SysSoft',
      component: () => import('/@/views/maintenance/sysSoft/index.vue'),
      meta: {
        title: t('dashboard.sysSoft'),
      },
    },
    {
      path: 'softUpdate',
      name: 'softUpdate',
      // redirect: '/maintenance/softUpdate/LiteOSUpdate',
      component: () => import('/@/views/maintenance/softUpdate/index.vue'),
      meta: {
        title: t('dashboard.softUpdate'),
      },
      // children: [
        // {
        //   path: 'ssmUpdate',
        //   name: 'ssmUpdate',
        //   component: () => import('/@/views/maintenance/softUpdate/ssmUpdate/index.vue'),
        //   meta: {
        //     title: t('dashboard.ssmUpdate'),
        //   },
        // },
        // {
        //   path: 'LiteOSUpdate',
        //   name: 'LiteOSUpdate',
        //   component: () => import('/@/views/maintenance/softUpdate/index.vue'),
        //   meta: {
        //     title: t('dashboard.liteOsUpdate'),
        //   },
        // },
        // {
        //   path: 'OTADaemon',
        //   name: 'OTADaemon',
        //   component: () => import('/@/views/maintenance/softUpdate/otaDaemonUpdate/index.vue'),
        //   meta: {
        //     title: t('dashboard.otaDaemonUpdate'),
        //   },
        // },
        // {
        //   path: 'Adapter',
        //   name: 'Adapter',
        //   component: () => import('/@/views/maintenance/softUpdate/adapterUpdate/index.vue'),
        //   meta: {
        //     title: t('dashboard.adapterUpdate'),
        //   },
        // },
        // {
        //   path: 'VPS',
        //   name: 'VPS',
        //   component: () => import('/@/views/maintenance/softUpdate/vpsUpdate/index.vue'),
        //   meta: {
        //     title: t('dashboard.vpsUpdate'),
        //   },
        // },
        // {
        //   path: 'BMS',
        //   name: 'BMS',
        //   component: () => import('/@/views/maintenance/softUpdate/bmsUpdate/index.vue'),
        //   meta: {
        //     title: t('dashboard.bmsUpdate'),
        //   },
        // },
      // ],
    },
    {
      path: 'coreBoardMap',
      name: 'coreBoardMap',
      component: () => import('/@/views/maintenance/coreBoardMap/index.vue'),
      meta: {
        title: t('dashboard.coreBoardMap'),
        hideMenu: true,
      },
    },
    {
      path: 'networkSetting',
      name: 'NetworkSetting',
      component: () => import('../../../views/maintenance/networkSetting/index.vue'),
      meta: {
        title: t('dashboard.sysSetting'),
        // title: t('dashboard.networkSetting'),
      },
    },
    // {
    //   path: 'ssh',
    //   name: 'ssh',
    //   redirect: '/maintenance/ssh',
    //   meta: {
    //     title: t('dashboard.sshTerminal'),
    //   },
    //   children: [
    //     {
    //       path: 'sshTerminal',
    //       name: 'sshTerminal',
    //       component: () => import('/@/views/maintenance/ssh/sshTerminal.vue'),
    //       meta: {
    //         title: t('dashboard.sshTerminal'),
    //         ignoreKeepAlive: true,
    //       },
    //     },
    //     {
    //       path: 'hostComputer',
    //       name: 'hostComputer',
    //       component: () => import('/@/views/maintenance/ssh/hostComputer.vue'),
    //       meta: {
    //         title: t('dashboard.hostComputer'),
    //         ignoreKeepAlive: true,
    //       },
    //     },
    //     {
    //       path: 'quickCommands',
    //       name: 'quickCommands',
    //       component: () => import('/@/views/maintenance/ssh/quickCommands.vue'),
    //       meta: {
    //         title: t('dashboard.quickCommands'),
    //         ignoreKeepAlive: true,
    //       },
    //     },
    //   ],
    // },
    // {
    //   path: 'contain',
    //   name: 'contain',
    //   redirect: '/maintenance/contain',
    //   meta: {
    //     title: t('dashboard.DockerContainers'),
    //   },
    //   children: [
    //     {
    //       path: 'containers',
    //       name: 'containers',
    //       component: () => import('/@/views/maintenance/contain/containers.vue'),
    //       meta: {
    //         title: t('dashboard.containers'),
    //       },
    //     },
    //     {
    //       path: 'mirror',
    //       name: 'mirror',
    //       component: () => import('/@/views/maintenance/contain/mirror.vue'),
    //       meta: {
    //         title: t('dashboard.mirror'),
    //       },
    //     },
    //     {
    //       path: 'store',
    //       name: 'store',
    //       component: () => import('/@/views/maintenance/contain/store.vue'),
    //       meta: {
    //         title: t('dashboard.store'),
    //       },
    //     },
    //   ],
    // },
    {
      path: 'coreContain/:number/:type',
      name: 'coreContain',
      component: () => import('/@/views/maintenance/coreContain/index.vue'),
      meta: {
        title: t('dashboard.coreDockerContainers'),
        hideMenu: true,
        dynamicLevel: 1,
        currentActiveMenu: '/maintenance',
        realPath: '/maintenance/coreContain',
      },
    },
    // {
    //   path: 'threshold',
    //   name: 'Threshold',
    //   component: () => import('/@/views/maintenance/threshold/index.vue'),
    //   meta: {
    //     title: t('dashboard.sysSetting'),
    //   },
    // },
  ],
};

export default maintenance;
