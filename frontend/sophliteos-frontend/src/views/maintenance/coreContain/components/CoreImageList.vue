<template>
  <div class="p-4">
    <BasicTable
      @register="register"
      :title="`${t('overview.coreBoard')} ${$props.number} ${t('dashboard.mirror')}`"
    >
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
  </div>
</template>
<script lang="ts" setup>
  import { watch } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { getCoreImageList, deleteCoreImage } from '/@/api/maintenance/index';
  import { getMirrorColumns } from '../../contain/MirrorTableData';

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

  const [register, { reload }] = useTable({
    api: getCoreImageList,
    columns: getMirrorColumns(),
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
  });

  async function handleDelete(record: Recordable) {
    let params = { image: record.ID, ip: props.ip };
    await deleteCoreImage(params);
    reload();
  }
</script>
<style scoped lang="less"></style>
