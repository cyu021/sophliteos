<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button type="primary" @click="editHandle('add')">{{
          t('dataSource.videoManage.addDevice')
        }}</a-button>
        <a-button type="primary" @click="oneStart">{{
          t('dataSource.videoManage.allCheck')
        }}</a-button>
        <!-- <a-button type="primary" @click="mediaServer">{{
          t('dataSource.mediaServers.configService')
        }}</a-button> -->
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              // {
              //   icon: 'nimbus:stop',
              //   tooltip: t('component.table.stop'),
              //   color: 'warning',
              //   onClick: handle.bind(null, record),
              //   ifShow: record.status === 'ON',
              // },
              // {
              //   icon: 'mdi:success-circle-outline',
              //   tooltip: t('component.table.start'),
              //   color: 'success',
              //   onClick: handle.bind(null, record),
              //   ifShow: record.status !== 'ON',
              // },
              {
                icon: 'typcn:edit',
                color: 'success',
                tooltip: t('component.cropper.btn_edit'),
                onClick: editHandle.bind(null, record),
              },
              {
                icon: 'ant-design:delete-outlined',
                color: 'error',
                tooltip: t('component.cropper.btn_delete'),
                popConfirm: {
                  title: t('component.cropper.btn_delete') + '?',
                  placement: 'left',
                  confirm: handleDelete.bind(null, record.deviceId),
                },
              },
              {
                icon: 'carbon:inspection',
                color: 'error',
                tooltip: t('component.cropper.btn_inspect'),
                onClick: PatrolInsepect.bind(null, record),
              },
              {
                icon: 'carbon:play-filled',
                tooltip: t('component.cropper.btn_play'),
                onClick: VideoPreview.bind(null, record),
              },
            ]"
          />
        </template>
        <template v-if="column.key === 'ptzType'"> {{ optType[record.ptzType] }} </template>
        <template v-if="column.key === 'protocol'"> {{ optProtocol[record.protocol] }} </template>
        <template v-if="column.key === 'mediaPull'">
          {{ optMediaPull[record.mediaPull] }}
        </template>
        <template v-if="column.key === 'isNeedDetect'">
          {{ optIsMediaPull[record.mediaPull] }}
        </template>
      </template>
      ></BasicTable
    >
    <videoModal @register="registerModal" @success="handleSuccess" />
    <MediaServers @register="mediaModal" @success="configSuccess" />
    <editVideo @register="DeviceModal" @success="VideoSuccess" />
  </div>
</template>

<script lang="ts" setup>
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { getBasicColumns } from './tableData';
  import {
    LivePreview,
    getVideosList,
    getMediaServer,
    DelDevice,
    DeviceCheck,
  } from '/@/api/dataSource/index';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useModal } from '/@/components/Modal';
  import videoModal from './videoModal.vue';
  import MediaServers from './MediaServers.vue';
  import editVideo from './editVideo.vue';
  import { useMessage } from '/@/hooks/web/useMessage';
  const { createMessage } = useMessage();
  const { t } = useI18n();
  const optProtocol = {
    0: '未知',
    1: '国标',
    2: 'RTSP',
    3: '海康SDK',
    4: '大华SDK',
  };
  const optType = {
    0: '未知',
    1: '球机',
    2: '半球',
    3: '枪机',
    4: '遥控枪机',
  };
  const optMediaPull = {
    0: '失败',
    1: '成功',
  };
  const optIsMediaPull = {
    1: '巡检完成',
    0: '需要巡检',
  };
  const [registerTable, { getDataSource, reload, setLoading, }] = useTable({
    columns: getBasicColumns(),
    api: getVideosList,
    showTableSetting: true,
    tableSetting: { fullScreen: true },
    showIndexColumn: true,
    indexColumnProps: {
      width: 120,
    },
    actionColumn: {
      width: 150,
      title: t('component.cropper.btn_action'),
      dataIndex: 'action',
    },
    rowKey: 'deviceName',
  });
  const [registerModal, { openModal }] = useModal();
  async function VideoPreview(record: any) {
    try {
      setLoading(true);
      const res = await LivePreview({ deviceId: record.deviceId });
      setLoading(false);
      openModal(true, {
        record,
        res,
      });
    } catch {
      setLoading(false);
    }
  }
  function handleSuccess() {
    reload();
  }
  async function handleDelete(id) {
    const device = [id];
    await DelDevice({ device: device });
    createMessage.success('Delete stream success');
    reload();
  }
  const [DeviceModal, { openModal: OpenVideoModal }] = useModal();
  function editHandle(record: any) {
    OpenVideoModal(true, { record });
  }
  function VideoSuccess() {
    createMessage.success('Create stream success');
    reload();
  }
  const [mediaModal, { openModal: OpenMediaModal }] = useModal();
  async function mediaServer() {
    const res = await getMediaServer();
    OpenMediaModal(true, { res });
  }
  function configSuccess() {
    createMessage.success('Config stream sucess');
    reload();
  }
  async function oneStart() {
    const deviceIds = getDataSource().map((item) => item.deviceId);
    setLoading(true);
    await DeviceCheck({ deviceIds: deviceIds });
    setLoading(false);
    createMessage.success('Check stream success');
    reload();
  }
  async function PatrolInsepect(record) {
    const deviceId = [record.deviceId];
    setLoading(true);
    await DeviceCheck({ deviceIds: deviceId });
    setLoading(false);
    createMessage.success('Check stream success');
    reload();
  }
</script>
