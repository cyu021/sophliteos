<template>
  <div class="p-4">
    <BasicTable @register="hostComputerTable" :title="t('dashboard.hostComputer')">
      <template #toolbar>
        <a-button type="primary" @click="addHostForm" class="addHostButt">{{
          t('maintenance.ssh.addHost')
        }}</a-button>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                label: t('maintenance.ssh.delete'),
                icon: 'ic:outline-delete-outline',
                popConfirm: {
                  title: t('maintenance.ssh.deleteConfirm'),
                  confirm: handleDelete.bind(null, record),
                },
              },
              {
                label: t('maintenance.ssh.edit'),
                icon: 'clarity:edit-line',
                onClick: editHostForm.bind(null, record),
              },
            ]"
          />
        </template>
      </template>
    </BasicTable>
    <AddHostForm @register="register" @form-sucess="formSucess" />
  </div>
</template>
<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { getHostColumns } from './hostTableData';
  import { getHostList, delHost } from '/@/api/maintenance/index';
  import { useModal } from '/@/components/Modal';
  import AddHostForm from './components/addHostForm.vue';
  import { message } from 'ant-design-vue';

  const { t } = useI18n();
  const [register, { openModal: FormAddHost }] = useModal();

  const [hostComputerTable, { reload }] = useTable({
    api: getHostList,
    columns: getHostColumns(),
    actionColumn: {
      width: 160,
      title: t('maintenance.ssh.operation'),
      dataIndex: 'action',
    },
  });

  function addHostForm() {
    FormAddHost(true, { add: true });
  }
  function editHostForm(record: Recordable) {
    FormAddHost(true, record);
  }
  function handleDelete(record: Recordable) {
    let params = record;
    delete params.ID;
    delHost(params);
    reload();
  }
  const formSucess = (msg: boolean) => {
    if (msg) {
      message.success(t('sys.api.success'));
      reload();
    }
  };
</script>
<style scoped lang="less">
  .addHostButt {
    margin: 10px 50px;
  }
</style>
