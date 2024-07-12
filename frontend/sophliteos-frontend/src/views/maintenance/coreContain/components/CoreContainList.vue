<template>
  <div class="p-4">
    <BasicTable
      @register="containerTable"
      :title="`${t('overview.coreBoard')} ${$props.number} ${t('dashboard.containers')}`"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                label: record.running
                  ? t('maintenance.containers.stop')
                  : t('maintenance.containers.run'),
                icon: record.running
                  ? 'fluent:subtract-circle-12-regular'
                  : 'mdi:success-circle-outline',
                onClick: stopOrStart.bind(null, record),
              },
              {
                label: t('maintenance.ssh.delete'),
                icon: 'ic:outline-delete-outline',
                popConfirm: {
                  title: t('maintenance.containers.containerDelConfirm'),
                  confirm: handleDelete.bind(null, record),
                },
              },
            ]"
          />
        </template>
        <template v-if="column.key === 'running'">
          <Tag color="green" v-if="record.running">
            {{ t('maintenance.containers.activated') }}
          </Tag>
          <Tag color="blue" v-else>
            {{ t('maintenance.containers.stop') }}
          </Tag>
        </template>
        <template v-else-if="column.key === 'ports'">
          <div v-for="(port, index) in record.ports" :key="index">
            <Tag color="blue">
              {{ port }}
            </Tag>
          </div>
        </template>
      </template>
    </BasicTable>
  </div>
</template>
<script lang="ts" setup>
  import { watch } from 'vue';
  import { Tag, message } from 'ant-design-vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import {
    getCoreContainList,
    coreContainStop,
    coreContainStart,
    deleteCoreContain,
  } from '/@/api/maintenance/index';
  import { getContainersColumns } from '../../contain/ContainersTableData';

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

  const [containerTable, { reload, setLoading }] = useTable({
    api: getCoreContainList,
    columns: getContainersColumns(),
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

  async function stopOrStart(record: Recordable) {
    const { name, running } = record;
    setLoading(true);
    const success = () => {
      setLoading(false);
      message.success(t('sys.api.success'));
      reload();
    };
    if (running) {
      coreContainStop({ name: name, ip: props.ip })
        .then(() => {
          success();
        })
        .catch(() => {
          setLoading(false);
        });
    } else {
      coreContainStart({ name: name, ip: props.ip })
        .then(() => {
          success();
        })
        .catch(() => {
          setLoading(false);
        });
    }
  }

  async function handleDelete(record: Recordable) {
    let params = { name: record.name, ip: props.ip };
    await deleteCoreContain(params);
    reload();
  }
</script>
<style scoped lang="less"></style>
