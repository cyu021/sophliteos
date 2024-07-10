<template>
  <div class="p-4">
    <BasicTable @register="CommandTable" :title="t('dashboard.quickCommands')">
      <template #toolbar>
        <a-button type="primary" @click="addCommandForm" class="addCommandButt">{{
          t('maintenance.ssh.addCommand')
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
                  title: t('maintenance.ssh.deleteCommandConfirm'),
                  confirm: handleDelete.bind(null, record),
                },
              },
              {
                label: t('maintenance.ssh.edit'),
                icon: 'clarity:edit-line',
                onClick: editCommandForm.bind(null, record),
              },
            ]"
          />
        </template>
      </template>
    </BasicTable>
    <CommandForm @register="register" @form-sucess="formSucess" />
  </div>
</template>
<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { getCommandColumns } from './CommandTableData';
  import { getCommandList, delCommand } from '/@/api/maintenance/index';
  import { useModal } from '/@/components/Modal';
  import CommandForm from './components/CommandForm.vue';

  const { t } = useI18n();
  const [register, { openModal: FormCommand }] = useModal();

  const [CommandTable, { reload }] = useTable({
    api: getCommandList,
    columns: getCommandColumns(),
    actionColumn: {
      width: 160,
      title: t('maintenance.ssh.operation'),
      dataIndex: 'action',
    },
  });

  function addCommandForm() {
    FormCommand(true, { add: true });
  }
  function editCommandForm(record: Recordable) {
    FormCommand(true, record);
  }
  function handleDelete(record: Recordable) {
    let params = record;
    delete params.ID;
    delCommand(params);
    reload();
  }
  const formSucess = (message: boolean) => {
    if (message) {
      reload();
    }
  };
</script>
<style scoped lang="less">
  .addCommandButt {
    margin: 10px 50px;
  }
</style>
