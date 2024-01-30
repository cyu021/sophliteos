<template>
  <div class="p-4">
    <BasicTable
      @register="storeTable"
      :title="`${t('overview.coreBoard')} ${$props.number} ${t('routes.dashboard.store')}`"
    >
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
    <AddStoreForm v-if="$props.ip" @register="register" @form-sucess="formSucess" :ip="$props.ip" />
  </div>
</template>
<script lang="ts" setup>
  import { watch } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { getStoreColumns } from '../../contain/StoreTableData';
  import { getCoreRegistryList, deleteCoreRegistry } from '/@/api/maintenance/index';
  import { useModal } from '/@/components/Modal';
  import AddStoreForm from '../../contain/components/addStoreForm.vue';

  const { t } = useI18n();

  const props = defineProps({
    ip: {
      type: String,
      default: '',
    },
    number: {
      type: Number,
      default: 0,
    },
  });

  watch(
    () => props.ip,
    () => {
      reload();
    },
  );

  const [register, { openModal: FormAddStore }] = useModal();

  const [storeTable, { reload }] = useTable({
    api: getCoreRegistryList,
    columns: getStoreColumns(),
    actionColumn: {
      width: 160,
      title: t('maintenance.ssh.operation'),
      dataIndex: 'action',
    },
    beforeFetch: (params) => {
      return {
        ...params,
        ip: props.ip,
      };
    },
    afterFetch: (data) => {
      return data.map((item) => ({ addr: item }));
    },
  });

  function addStoreForm() {
    FormAddStore(true, { add: true });
  }

  async function handleDelete(record: Recordable) {
    let params = { insecureRegistry: record.addr, ip: props.ip };
    await deleteCoreRegistry(params);
    reload();
  }

  const formSucess = () => {
    reload();
  };
</script>
<style scoped lang="less"></style>
