<template>
  <PageWrapper>
    <a-menu
      mode="horizontal"
      :forceSubMenuRender="true"
      @select="newTerminal"
      :selectedKeys="selectedKeys"
    >
      <a-menu-item v-for="ip in linkIpList" :key="extractIpFromString(ip)">
        {{ ip }}
        <a-button type="text" preIcon="mdi:remove" @click.stop="removeTermianl(ip)" />
      </a-menu-item>
      <a-sub-menu key="ipList">
        <a-sub-menu key="core" v-if="needCore">
          <template #title>{{ t('maintenance.ssh.coreIpMenu') }}</template>
          <a-menu-item v-for="ip in coreBoardList" :key="ip.key">
            {{ ip.label }}
          </a-menu-item>
        </a-sub-menu>
        <template #title>{{ t('maintenance.ssh.createShh') }}</template>
        <a-menu-item v-for="ip in ips" :key="ip.addr">{{ ip.addr }}</a-menu-item>
      </a-sub-menu>
    </a-menu>

    <Terminal :ip="currentIp" :core="core" />
    <Loading :loading="sshState.loading" :tip="sshState.tip" />
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { onMounted, reactive, ref, watchEffect } from 'vue';
  import { Menu, message } from 'ant-design-vue';
  import { useRoute } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { PageWrapper } from '/@/components/Page';
  import { getHostList, getCoreBoardList } from '/@/api/maintenance/index';
  import { Loading } from '/@/components/Loading';
  import { historyTermial } from '/@/store/modules/terimal';
  import { useDeviceInfo } from '/@/store/modules/overview';
  import Terminal from './components/termial.vue';

  const AMenu = Menu;
  const AMenuItem = Menu.Item;
  const ASubMenu = Menu.SubMenu;

  const storeHstory = historyTermial();

  const { t } = useI18n();

  const ips = ref<any[]>([]);
  const linkIpList = ref<string[]>([]);
  const currentIp = ref('');
  const selectedKeys = ref<any[]>([]);
  const sshState = reactive({
    loading: false,
    tip: '连接建立中...',
  });
  const initIP = ref('127.0.0.1'); // 初始化页面中显示的终端
  const coreBoardList = ref<any[]>([]);
  const core = ref(0);
  const needCore = ref(false);

  const deviceStore = useDeviceInfo();
  if (!deviceStore.deviceType) {
    deviceStore.getDeviceInfo();
  }

  onMounted(async () => {
    needCore.value = deviceStore.needCoreBoard;
    initTerminal();
  });

  // 监控路由变化，若路由在此页面，固定滚动条，若不在，取消固定滚动条
  const route = useRoute();
  watchEffect(() => {
    if (route.fullPath === '/maintenance/ssh/sshTerminal') {
      document.body.style.height = '100vh';
      document.body.style.overflow = 'hidden';
    } else {
      document.body.style.height = '';
      document.body.style.overflow = '';
    }
  });

  async function getCoreBoradOption() {
    const res = await getCoreBoardList();
    coreBoardList.value = res.map((item) => {
      const label = `${t('overview.coreBoard')}-${item.number}: ${item.ip}`;
      return {
        ...item,
        label: label,
        key: `core-${item.ip}`,
      };
    });
  }

  const initLinkIpList = () => {
    const ipList = storeHstory.getIpList();
    if (ipList.length > 0) {
      ipList.forEach((ip) => {
        const formatAddr = formatAddress(ip, checkCore(ip));
        formatAddr && linkIpList.value.push(formatAddr);
      });
      initIP.value = ipList[0];
    } else {
      const formatAddr = formatAddress(initIP.value, false);
      if (formatAddr && !linkIpList.value.includes(formatAddr)) {
        linkIpList.value.push(formatAddr);
      }
    }
  };

  const newTerminal = async ({ key }) => {
    const isCore = checkCore(key);
    let formatAddr = formatAddress(key, isCore);
    if (!formatAddr) {
      message.warning(t('maintenance.ssh.notIp'));
      return;
    }
    selectedKeys.value = [key];
    if (!linkIpList.value.includes(formatAddr)) {
      // 最多5个标签
      if (linkIpList.value.length >= 5) {
        linkIpList.value.shift();
      }
      linkIpList.value.push(formatAddr);
    }
    await updateTerminal(key, isCore);
  };

  const fetchHostList = async () => {
    const data = await getHostList();
    ips.value = data.map((item) => ({ ...item, label: `${item.user}@${item.addr}:${item.port}` }));
  };

  const updateTerminal = async (addr, isCore) => {
    sshState.loading = true;
    // TODO
    // const result = await terminalsTest({ addr });
    // if (result) {
    currentIp.value = addr;
    core.value = isCore;
    // } else {
    //   currentIp.value = '';
    // }
    sshState.loading = false;
  };

  const initTerminal = async () => {
    await fetchHostList();
    needCore.value && (await getCoreBoradOption());
    initLinkIpList();
    let isAddrPresent =
      ips.value?.some((item) => item.addr === initIP.value) ||
      coreBoardList.value?.some((item) => item.key === initIP.value);
    if (isAddrPresent) {
      // 初始化终端存在，建立连接
      const isCore = checkCore(linkIpList.value[0]);
      await updateTerminal(initIP.value, isCore);
      selectedKeys.value = [initIP.value];
    } else {
      currentIp.value = 'null';
    }
  };

  async function removeTermianl(key: string) {
    const ip = extractIpFromString(key);
    const indexToDeleteLinkIpList = linkIpList.value.indexOf(key);
    if (indexToDeleteLinkIpList !== -1) {
      linkIpList.value.splice(indexToDeleteLinkIpList, 1);
    }
    storeHstory.deleteSocket(ip);
    // 删完显示哪个页面的逻辑
    if (linkIpList.value.length === 0) {
      currentIp.value = 'null'; //没得显示，特殊处理
      selectedKeys.value = [];
    } else if (ip === selectedKeys.value?.[0]) {
      // 删的是当前终端,切换后一个，没后一个，切换第一个
      if (indexToDeleteLinkIpList < linkIpList.value.length) {
        // 显示后一个
        const link = linkIpList.value[indexToDeleteLinkIpList];
        const nextIp = extractIpFromString(link);
        selectedKeys.value[0] = nextIp;
        await updateTerminal(nextIp, checkCore(link));
      } else {
        // 显示第一个
        const firstIp = extractIpFromString(linkIpList.value[0]);
        selectedKeys.value[0] = firstIp;
        await updateTerminal(firstIp, checkCore(linkIpList.value[0]));
      }
    }
  }

  function formatAddress(addr: string, isCore) {
    if (isCore) {
      return coreBoardList.value?.find((ip) => ip.key === addr)?.label || null;
    }
    return ips.value.find((ip) => ip.addr === addr)?.label || null;
  }

  function extractIpFromString(inputString) {
    if (checkCore(inputString)) {
      return coreBoardList.value?.find((ip) => ip.label === inputString)?.key;
    }
    const ipRegex = /\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b/;
    const match = inputString.match(ipRegex);

    return match ? match[0] : null;
  }

  function checkCore(label) {
    return Number(
      label?.split('-')[0] === t('overview.coreBoard') || label?.split('-')[0] === 'core',
    );
  }
</script>
<style scoped lang="less">
  .background-terminal {
    background-color: black;
  }
</style>
