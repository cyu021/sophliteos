<template>
  <BasicTable
    @register="registerTable"
    :title="t('maintenance.prepSpringZip.customAlgoDownloadLinks')"
    :noloading="true"
  >
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'action'">
        <Space>
          <Button type="primary" :href="record.url">download</Button>
          <Button danger type="primary">purge</Button>
        </Space>
      </template>
    </template>
  </BasicTable>
</template>
<script lang="ts" setup>
  import { BasicTable, useTable } from '/@/components/Table';
  import { getBasicColumns } from './tableData';
  import apis from '../apis.js';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { onBeforeUnmount, onMounted } from 'vue';
  import { Space, Button } from 'ant-design-vue';
  const { t } = useI18n();
  const [registerTable, { reload }] = useTable({
    title: t('maintenance.prepSpringZip.customAlgoDownloadLinks'),
    api: apis.list,
    columns: getBasicColumns(),
    showTableSetting: true,
    tableSetting: { fullScreen: true },
    showIndexColumn: false,
    indexColumnProps: {
      width: 60,
    },
    rowKey: 'name',
  });
  // 停留此页面每秒刷新一次表格数据
  onMounted(() => {
    // const intervalId = setInterval(reload, 1000);

    // // 在组件销毁前清理定时器
    // onBeforeUnmount(() => {
    //   clearInterval(intervalId);
    // });
  });
</script>
