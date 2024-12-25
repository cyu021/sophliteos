<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button type="primary" @click="editHandle('add')">{{
          t('dataSource.privilegeCfg.addRole')
        }}</a-button>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'clarity:note-edit-line',
                tooltip: t('component.cropper.btn_edit'),
                onClick: handleEdit.bind(null, record),
              },
              {
                icon: 'ant-design:delete-outlined',
                color: 'error',
                tooltip: t('component.cropper.btn_delete'),
                popConfirm: {
                  title: t('component.cropper.btn_delete') + '?',
                  placement: 'left',
                  confirm: handleDelete.bind(null, record.taskId),
                },
              },
            ]"
          />
        </template>
      </template>
    </BasicTable>
  </div>
</template>

<script lang="ts" setup>
  import { Tooltip } from 'ant-design-vue';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { useI18n } from '/@/hooks/web/useI18n';
  import {
    getTaskList,
    PostDeleteTask,
    StartTask,
    StopTask,
    getAlgorithm,
    getUpUrl,
    getAbilites,
  } from '/@/api/task/index';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { useModal } from '/@/components/Modal';
  import { columns } from './taskData';
  // import { useRedo } from '/@/hooks/web/usePage';
  const { createMessage } = useMessage();
  // const redo = useRedo();
  const { t } = useI18n();
  const ATooltip = Tooltip;
  const [registerTable, { reload, getDataSource, setLoading }] = useTable({
    api: getTaskList,
    rowKey: 'roleName',
    columns,
    showTableSetting: true,
    bordered: true,
    tableSetting: { fullScreen: true },
    showIndexColumn: true,
    indexColumnProps: {
      width: 160,
    },
    actionColumn: {
      width: 120,
      title: t('component.cropper.btn_action'),
      dataIndex: 'action',
    },
  });
</script>
