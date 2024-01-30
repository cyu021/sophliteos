<template>
  <Card class="!m-4" v-bind="$attrs" :active-tab-key="activeKey" :loading="loading">
    <a-menu
      mode="horizontal"
      :forceSubMenuRender="true"
      :inlineIndent="28"
      @select="selectedContent"
      :selectedKeys="selectedKeys"
      class="detail-menu"
      :subMenuCloseDelay="0.3"
    >
      <a-menu-item :key="tabListTitle.CONTROL.key">
        {{ tabListTitle.CONTROL.tab }}
      </a-menu-item>
      <a-sub-menu :key="tabListTitle.CORE.key">
        <template #title>{{ coreTabName }}</template>
        <a-menu-item v-for="plain in plainOptions" :key="plain.value">{{
          plain.label
        }}</a-menu-item>
      </a-sub-menu>
    </a-menu>
    <CircleGrid :grid-list="gridList" class="circle-grid" />
    <a-divider />
    <div v-if="activeKey === tabListTitle.CONTROL.key">
      <a-descriptions :title="t('overview.basicInfor')" bordered :column="controlColumn">
        <a-descriptions-item :label="t('overview.device.name')">{{
          originData.deviceName
        }}</a-descriptions-item>
        <a-descriptions-item :label="t('overview.device.type')">{{
          originData.deviceType
        }}</a-descriptions-item>
        <a-descriptions-item :label="t('overview.device.ctrSN')">{{
          originData.ctrlBoardSn
        }}</a-descriptions-item>
        <a-descriptions-item :label="t('overview.device.sdkVersion')">{{
          originData.sdkVersion
        }}</a-descriptions-item>
        <a-descriptions-item :label="t('overview.buildTime')">{{
          originData.bmssmVersion
        }}</a-descriptions-item>
        <a-descriptions-item :label="t('overview.device.ip')">{{
          originData.deviceIp
        }}</a-descriptions-item>
        <a-descriptions-item label="WAN IP">{{ originData.wanIp }}</a-descriptions-item>
        <a-descriptions-item label="LAN IP">
          {{ originData.lanIp.split(',')[0] }}
          <br />
          {{ originData.lanIp.split(',')[1] }}
        </a-descriptions-item>
        <a-descriptions-item :label="t('overview.operatingSystem')">{{
          originData.operatingSystem
        }}</a-descriptions-item>
        <a-descriptions-item
          v-for="item in originData.netCard"
          :label="t('overview.netCard') + item.name"
          :key="item.ip"
        >
          {{ t('overview.bandwidth') + '：' + item.bandwidth + t('overview.bandwidthUnit') }}
          <br />
          {{ t('overview.ip') + '：' + item.ip }}
          <br />
          {{ t('overview.mac') + '：' + item.mac }}
        </a-descriptions-item>
      </a-descriptions>
    </div>
    <a-row v-if="activeKey === tabListTitle.CORE.key">
      <a-col :xs="24" :lg="12">
        <a-descriptions :title="t('overview.basicInfor')" bordered :column="1">
          <a-descriptions-item :label="t('overview.device.sn')">{{
            currentBoardInfor.boardSn
          }}</a-descriptions-item>
          <a-descriptions-item :label="t('overview.device.sdkVersion')">{{
            currentBoardInfor.sdkVersion
          }}</a-descriptions-item>
          <a-descriptions-item :label="t('overview.device.type')">{{
            currentBoardInfor.boardType
          }}</a-descriptions-item>
          <a-descriptions-item :label="t('overview.chip') + t('overview.health')">{{
            chipStatus[currentBoardInfor.chip[0].health]
          }}</a-descriptions-item>
          <a-descriptions-item :label="t('overview.ip')">{{
            currentBoardInfor.netCard[0]?.ip
          }}</a-descriptions-item>
          <a-descriptions-item :label="t('overview.bandwidth')">{{
            currentBoardInfor.netCard[0].bandwidth + t('overview.bandwidthUnit')
          }}</a-descriptions-item>
          <a-descriptions-item :label="t('overview.mac')">{{
            currentBoardInfor.netCard[0].mac
          }}</a-descriptions-item>
        </a-descriptions>
      </a-col>
      <a-col :xs="24" :lg="12" class="!flex items-center">
        <a-row class="w-full">
          <a-col :xs="12" :md="12">
            <GaugeChart
              :value="currentBoardInfor.chip[0].temperature"
              :unit="t('overview.coreTemperature') + '（℃）'"
            />
          </a-col>
          <a-col :xs="12" :md="12">
            <GaugeChart
              :value="currentBoardInfor.temperature"
              :unit="t('overview.boardTemperature') + '（℃）'"
            />
          </a-col>
          <!-- <a-col :xs="24" :md="12">
            <GaugeChart
              :value="+(Math.max(0, currentBoardInfor.fanSpeed) / 1000).toFixed(0)"
              :colors="['#80B1F9', '#0C33F5']"
              :max="20"
              :unit="t('overview.fanSpeed') + '（x1000r/min）'"
            />
          </a-col> -->
        </a-row>
      </a-col>
    </a-row>
    <!-- <template #tabBarExtraContent>
      <a-select
        v-show="activeKey === 'core'"
        v-model:value="currentCore"
        style="width: 140px"
        :options="plainOptions"
      />
    </template> -->
  </Card>
</template>
<script lang="ts" setup>
  import { ref, computed, watch } from 'vue';
  import { Card, Menu, Descriptions, Divider, Row, Col } from 'ant-design-vue';
  import { storeToRefs } from 'pinia';
  import { useDeviceInfo } from '/@/store/modules/overview';
  import { useRouter } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import CircleGrid from './components/CircleGrid.vue';
  import GaugeChart from './components/Gauge.vue';
  import { useWindowSize } from '@vueuse/core';
  import { Board, chipTypeMap, tabListTitle } from './data';

  const { t } = useI18n();
  const { currentRoute, push } = useRouter();

  const AMenu = Menu;
  const AMenuItem = Menu.Item;
  const ASubMenu = Menu.SubMenu;
  const ADescriptions = Descriptions;
  const ADescriptionsItem = Descriptions.Item;
  const ARow = Row;
  const ACol = Col;
  const ADivider = Divider;

  const loading = ref(false);
  const deviceInfoStore = useDeviceInfo();
  const { deviceStatus, originData, deviceInfo } = storeToRefs(deviceInfoStore);
  const selectedKeys = ref([tabListTitle.CONTROL.key]);
  if (!deviceInfo.value.deviceSn) {
    loading.value = true;
    deviceInfoStore.getDeviceInfo().then(() => {
      loading.value = false;
    });
  }
  const activeKey = ref(tabListTitle.CONTROL.key);
  const chipStatus = {
    0: t('overview.healthy'),
    1: t('overview.hitch'),
  };

  watch(selectedKeys, (select) => {
    // console.log('+++', currentRoute.value.params.sn, select[0]);
    if (select[0] === tabListTitle.CONTROL.key) {
      push({ name: 'OverviewDetail' });
    } else {
      push({ name: 'OverviewDetailDny', params: { sn: select[0] } });
    }
  });
  function selectedContent({ key }) {
    selectedKeys.value = [key];
    if (key === tabListTitle.CONTROL.key) {
      activeKey.value = key;
    } else {
      activeKey.value = tabListTitle.CORE.key;
      currentCore.value = key;
    }
  }
  // 核心板选择器
  const plainOptions = computed(() => {
    return deviceStatus.value.map((board: Board) => ({
      value: board.sn,
      label: `${t('overview.coreBoard')}-${board.number}`,
    }));
  });

  const coreTabName = computed(() => {
    return (
      plainOptions.value.find((v) => v.value === selectedKeys.value[0])?.label ||
      tabListTitle.CORE.tab
    );
  });

  // 当前核心板
  const currentCore = ref('');
  //   watch(plainOptions, () => {
  //     currentCore.value = plainOptions.value[0]?.value;
  //   });
  // 当前核心板信息
  const currentBoardInfor = computed(() => {
    try {
      return originData.value.coreComputingUnit.board.find((board) => {
        return board.boardSn === currentCore.value;
      });
    } catch (error) {
      console.log(error);
      return {};
    }
  });
  if (currentRoute.value.params.sn) {
    activeKey.value = tabListTitle.CORE.key;
    const sn = currentRoute.value.params.sn as string;
    selectedKeys.value = [sn];
    currentCore.value = sn;
  }
  const gridList = computed(() => {
    let result: any[] = [];
    if (activeKey.value === tabListTitle.CONTROL.key && originData.value.cpu) {
      result = [
        {
          title: t('overview.cpu'),
          usage: originData.value.cpu.usage,
          text: [
            `${originData.value.cpu.cores}${t('overview.core')}${
              originData.value.cpu.frequency / 1000
            }GHz`,
          ],
        },
        {
          title: t('overview.memory'),
          usage: originData.value.memory.usage,
          total: originData.value.memory.total,
        },
        {
          title: t('overview.disk'),
          usage: originData.value.disk[0].usage,
          total: originData.value.disk[0].total,
        },
      ];
    } else if (activeKey.value === tabListTitle.CORE.key && currentBoardInfor.value.cpu) {
      result = [
        {
          title: t('overview.cpu'),
          usage: currentBoardInfor.value.cpu.usage,
          text: [
            `${currentBoardInfor.value.cpu.cores}${t('overview.core')}${
              currentBoardInfor.value.cpu.frequency / 1000
            }GHz`,
          ],
        },
        {
          title: t('overview.memory'),
          usage: currentBoardInfor.value.memory.usage,
          total: currentBoardInfor.value.memory.total,
        },
        {
          title: t('overview.disk'),
          usage: currentBoardInfor.value.disk[0].usage,
          total: currentBoardInfor.value.disk[0].total,
        },
        {
          title: t('overview.tpu'),
          usage: currentBoardInfor.value.chip[0].tpuUtililizationRate,
          text: [
            chipTypeMap.get(currentBoardInfor.value?.chip?.[0]?.chipType),
            `INT8 ${currentBoardInfor.value?.chip?.[0]?.theoretialCalculationCapacity || 0}TOPS`,
          ],
        },
        {
          title: t('overview.tpuMemory'),
          used: currentBoardInfor.value?.chip?.[0]?.memoryUsedBytes || 0,
          total: currentBoardInfor.value?.chip?.[0]?.memoryTotalBytes || 0,
        },
      ];
    }
    return result;
  });
  const { width } = useWindowSize();
  const controlColumn = computed(() => {
    return width.value > 768 ? 2 : 1;
  });
</script>
<style scoped lang="less">
  .circle-grid {
    margin-top: 30px;
  }

  .detail-menu {
    font-size: 20px;
  }
</style>
