<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button type="primary" @click="addTask">{{ t('taskList.taskList.addTask') }}</a-button>
        <!-- <a-button type="primary" @click="algoConfig">{{
          t('taskList.taskList.algoConfigTitle')
        }}</a-button>
        <a-button type="primary" @click="upUrlConfig">{{
          t('taskList.taskList.upUrlConfig')
        }}</a-button> -->
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'mdi:success-circle-outline',
                tooltip: t('component.table.start'),
                color: 'success',
                onClick: handle.bind(null, record),
                ifShow: !Boolean(record.status),
              },
              {
                icon: 'nimbus:stop',
                tooltip: t('component.table.stop'),
                color: 'warning',
                onClick: handle.bind(null, record),
                ifShow: Boolean(record.status),
              },
              {
                icon: 'carbon:play-filled',
                tooltip: t('component.cropper.btn_play'),
                onClick: VideoPreview.bind(null, record),
              },
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
        <template v-if="column.key === 'abilities'">
          <a-tooltip placement="bottom">
            <span>{{ (record.abilities && record.abilities.join('; ')) || '-' }}</span>
          </a-tooltip>
        </template>
      </template>
    </BasicTable>
    <TaskFormModal @register="registerModal" @success="handleSuccess" />
    <AddTaskModal @register="addModal" @addsuccess="addSuccess" />
    <AlgoConfigModal @register="algoConfigModal" @configsuccess="configSuccess" />
    <UpUrlConfigModal @register="upUrlConfigModal" @configsuccess="configSuccess" />
    <videoModal @register="videoPreviewModal" @success="videoSuccess" />
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
  import TaskFormModal from './TaskFormModal.vue';
  import AddTaskModal from './AddTaskModal.vue';
  import AlgoConfigModal from './AlgoConfigModal.vue';
  import UpUrlConfigModal from './UpUrlConfigModal.vue';
  import videoModal from '../../dataSource/videoManage/videoModal.vue';
  import { LivePreview } from '/@/api/dataSource';
  // import { useRedo } from '/@/hooks/web/usePage';
  const { createMessage } = useMessage();
  // const redo = useRedo();
  const { t } = useI18n();
  const ATooltip = Tooltip;
  const [registerTable, { reload, getDataSource, setLoading }] = useTable({
    api: getTaskList,
    rowKey: 'taskId',
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

  const [registerModal, { openModal: editOpenModal }] = useModal();
  async function handleEdit(record: Recordable) {
    const usedUrl = getDataSource().map((item) => item.deviceName) || [];
    const algoOptions = await getAbilityOption();
    editOpenModal(true, {
      record,
      usedUrl,
      algoOptions: algoOptions,
    });
  }
  async function handleDelete(id: string) {
    await PostDeleteTask({ taskId: id });
    reload();
  }
  async function handle(record: any) {
    try {
      setLoading(true);
      if (record.status === 0) {
        await StartTask({ taskId: record.taskId }).then();
        createMessage.success(t('maintenance.containers.activated'));
        record.status = 1;
      } else {
        await StopTask({ taskId: record.taskId, deviceName: record.taskId }).then();
        createMessage.success(t('maintenance.containers.stop'));
        record.status = 0;
      }
      reload();
      setLoading(false);
    } catch {
      setLoading(false);
    }
  }
  function handleSuccess() {
    createMessage.success(t('sys.api.success'));
    reload();
  }

  const [addModal, { openModal: addOpenModal }] = useModal();
  async function addTask() {
    const usedUrl = getDataSource() ? getDataSource().map((item) => item.deviceName) : [];
    console.info("usedUrl = " + JSON.stringify(usedUrl))
    const algoOptions = await getAbilityOption();
    addOpenModal(true, {
      usedUrl,
      algoOptions: algoOptions,
    });
  }
  function addSuccess() {
    createMessage.success('Add success');
    reload();
  }
  const [algoConfigModal, { openModal: algoConfigOpenModal }] = useModal();
  async function algoConfig() {
    const res = await getAlgorithm();
    algoConfigOpenModal(true, res);
  }
  function configSuccess() {
    createMessage.success('Config success');
    reload();
  }
  const [upUrlConfigModal, { openModal: upUrlConfigOpenModal }] = useModal();
  async function upUrlConfig() {
    const res = await getUpUrl();
    upUrlConfigOpenModal(true, { ip: res });
  }

  async function getAbilityOption() {
    const ret = await getAbilites();
    return (
      ret
        .map((item) => ({ label: item.name, value: item.type }))
        .sort((a, b) => a.value - b.value) || []
    );
  }

  const [videoPreviewModal, { openModal: openvideoModal }] = useModal();
  async function VideoPreview(record: any) {
    const res = await LivePreview({ deviceId: record.deviceId });
    openvideoModal(true, {
      record,
      res,
    });
  }
  function videoSuccess() {
    reload();
  }
</script>
