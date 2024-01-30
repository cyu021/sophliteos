<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    width="80vw"
    @cancel="handleClose"
    :footer="null"
  >
    <BasicTable @register="pullTable" :title="t('maintenance.containers.pullImageList')">
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'status'">
          <Tag
            v-if="pullMirrorStatusMap.get(record.status)"
            :color="pullMirrorStatusMap.get(record.status)?.color"
            >{{ pullMirrorStatusMap.get(record.status)?.label }}</Tag
          >
        </template>
      </template>
    </BasicTable>
  </BasicModal>
</template>

<script lang="ts" setup>
  import { ref, onUnmounted } from 'vue';
  import { Tag } from 'ant-design-vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicTable, useTable } from '/@/components/Table';
  import { getPulledMirrorList } from '/@/api/maintenance';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { getPullMirrorColumns, pullMirrorStatusMap } from '../MirrorTableData';

  const { t } = useI18n();
  const interval = ref();

  const emit = defineEmits(['close']);

  const [pullTable, { reload }] = useTable({
    api: getPulledMirrorList,
    columns: getPullMirrorColumns(),
    maxHeight: 500,
  });

  const [registerModal] = useModalInner(() => {
    intervalGetList();
  });

  function intervalGetList() {
    interval.value = setInterval(() => {
      reload();
    }, 2000);
  }

  onUnmounted(() => {
    handleClose();
  });

  function handleClose() {
    interval.value && clearInterval(interval.value);
    interval.value = null;
    emit('close');
  }
</script>
