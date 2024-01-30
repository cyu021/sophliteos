<template>
  <div class="p-4">
    <BasicTable @register="mirrorTable" :title="t('routes.dashboard.mirror')">
      <template #toolbar>
        <a-button type="primary" @click="pullMirrorForm">{{
          t('maintenance.containers.pullMirror')
        }}</a-button>
        <a-button type="primary" @click="pullMirrorList">
          {{ t('maintenance.containers.pullImageList') }}
        </a-button>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                label: t('maintenance.ssh.delete'),
                icon: 'ic:outline-delete-outline',
                popConfirm: {
                  title: t('maintenance.containers.mirrorDelConfirm'),
                  confirm: handleDelete.bind(null, record),
                },
              },
            ]"
          />
        </template>
      </template>
    </BasicTable>
    <PullMirrorForm @register="register" @form-sucess="formSucess" />
    <PullMirrorList @register="pullMirrorListRegister" @close="handelClose" />
  </div>
</template>
<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { getMirrorColumns } from './MirrorTableData';
  import { delMirror, getMirrorList } from '/@/api/maintenance/index';
  import { useModal } from '/@/components/Modal';
  import PullMirrorForm from './components/pullMirrorForm.vue';
  import PullMirrorList from './components/PullMirrorList.vue';

  const { t } = useI18n();
  const [register, { openModal: FormPullMirror }] = useModal();
  const [pullMirrorListRegister, { openModal: openPullMirrorListModal }] = useModal();

  const [mirrorTable, { reload }] = useTable({
    api: getMirrorList,
    columns: getMirrorColumns(),
    actionColumn: {
      width: 160,
      title: t('maintenance.ssh.operation'),
      dataIndex: 'action',
    },
  });

  function pullMirrorForm() {
    FormPullMirror(true, { add: true });
  }

  function pullMirrorList() {
    openPullMirrorListModal(true, { open: true });
  }

  async function handleDelete(record: Recordable) {
    let params = { image: record.ID };
    await delMirror(params);
    reload();
  }
  const formSucess = (message: boolean) => {
    if (message) {
      reload();
    }
  };

  function handelClose() {
    reload();
  }
</script>
<style scoped lang="less">
  .pullMirrorButt {
    margin: 10px 50px;
  }
</style>
