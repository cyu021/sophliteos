<template>
  <div class="p-4">
    <BasicTable @register="containerTable" :title="t('dashboard.containers')">
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
                label: t('maintenance.containers.terminal'),
                icon: 'material-symbols-light:terminal',
                disabled: !record.running,
                onClick: Terminal.bind(null, record),
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
    <terminalDrawer @register="register" />
  </div>
</template>
<script lang="ts" setup>
  import { Tag, message } from 'ant-design-vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import {
    containerList,
    delContainer,
    startContainer,
    stopContainer,
  } from '/@/api/maintenance/index';
  import { useDrawer } from '/@/components/Drawer';
  import terminalDrawer from './components/terminalDrawer.vue';
  import { getContainersColumns } from './ContainersTableData';

  const { t } = useI18n();

  const [containerTable, { reload, setLoading }] = useTable({
    api: containerList,
    columns: getContainersColumns(),
    actionColumn: {
      width: 160,
      title: t('maintenance.ssh.operation'),
      dataIndex: 'action',
    },
  });
  const [register, { openDrawer: openTerminal }] = useDrawer();

  async function stopOrStart(record: Recordable) {
    const { name, running } = record;
    setLoading(true);
    const success = () => {
      setLoading(false);
      message.success(t('sys.api.success'));
      reload();
    };
    if (running) {
      stopContainer({ name: name })
        .then(() => {
          success();
        })
        .catch(() => {
          setLoading(false);
        });
    } else {
      startContainer({ name: name })
        .then(() => {
          success();
        })
        .catch(() => {
          setLoading(false);
        });
    }
  }

  function Terminal(record: Recordable) {
    openTerminal(true, record);
  }
  async function handleDelete(record: Recordable) {
    let params = { name: record.name };
    await delContainer(params);
    reload();
  }
</script>
<style scoped lang="less"></style>
