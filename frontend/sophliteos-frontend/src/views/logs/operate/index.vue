<template>
  <BasicTable @register="registerTable">
    <template #form-operationFunc="{ model, field }">
      <Select v-model:value="model[field]" :options="options.operationFunc" />
    </template>
    <template #form-username="{ model, field }">
      <Select v-model:value="model[field]" :options="options.username" />
    </template>
    <template #form-operationIp="{ model, field }">
      <Select v-model:value="model[field]" :options="options.operationIp" />
    </template>
  </BasicTable>
</template>
<script lang="ts" setup>
  import { onMounted, reactive } from 'vue';
  import { Select } from 'ant-design-vue';
  import { BasicTable, useTable } from '/@/components/Table';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { getOperRecord, getLogFunctionInfo } from '/@/api/logs/index';
  import { useDeviceInfo } from '/@/store/modules/overview';
  import { getBasicColumns, getFormConfig } from './tableData';

  const deviceStore = useDeviceInfo();
  if (!deviceStore.deviceType) {
    deviceStore.getDeviceInfo();
  }
  const { t } = useI18n();

  const options = reactive({
    operationFunc: new Array(0),
    username: new Array(0),
    operationIp: new Array(0),
  });

  onMounted(() => {
    getOptions();
  });

  async function getOptions() {
    const res = await getLogFunctionInfo();
    options.operationFunc = res.operationFunc.map((item) => ({ value: item, label: item }));
    options.operationIp = res.operationIp.map((item) => ({ value: item, label: item }));
    options.username = res.userName.map((item) => ({ value: item, label: item }));
  }

  const [registerTable] = useTable({
    title: t('logs.operate.title'),
    api: getOperRecord,
    columns: getBasicColumns(),
    useSearchForm: true,
    formConfig: getFormConfig(options),
    showTableSetting: true,
    tableSetting: { fullScreen: true },
    showIndexColumn: true,
    indexColumnProps: {
      width: 60,
    },
    rowKey: 'recordId',
  });
</script>
