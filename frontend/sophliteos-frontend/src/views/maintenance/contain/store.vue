<template>
  <div class="p-4">
    <BasicTable @register="storeTable" :title="t('routes.dashboard.store')">
      <template #toolbar>
        <a-button type="primary" @click="addStoreForm" class="addStoreButt">{{
          t('maintenance.containers.addStore')
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
                  title: t('maintenance.containers.deleteConfirm'),
                  confirm: handleDelete.bind(null, record),
                },
              },
            ]"
          />
        </template>
      </template>
    </BasicTable>
    <AddStoreForm @register="register" @form-sucess="formSucess" />
  </div>
</template>
<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { getStoreColumns } from './StoreTableData';
  import { getStoreList, delStore } from '/@/api/maintenance/index';
  import { useModal } from '/@/components/Modal';
  import AddStoreForm from './components/addStoreForm.vue';

  const { t } = useI18n();
  const [register, { openModal: FormAddStore }] = useModal();

  const [storeTable, { reload }] = useTable({
    api: getStoreList,
    columns: getStoreColumns(),
    actionColumn: {
      width: 160,
      title: t('maintenance.ssh.operation'),
      dataIndex: 'action',
    },
  });

  function addStoreForm() {
    FormAddStore(true, { add: true });
  }
  async function handleDelete(record: Recordable) {
    let params = { insecureRegistry: record.addr };
    await delStore(params);
    reload();
  }
  const formSucess = (message: boolean) => {
    if (message) {
      reload();
    }
  };
</script>
<style scoped lang="less">
  .addStoreButt {
    margin: 10px 50px;
  }
</style>
