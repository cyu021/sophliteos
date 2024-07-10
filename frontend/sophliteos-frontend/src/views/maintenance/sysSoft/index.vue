<template>
  <PageWrapper :title="t('dashboard.sysSoft')" :content="content">
    <div style="display: flex; flex-direction: row; margin-bottom: 20px">
      <div style="background-color: white; padding-bottom: 10px" v-if="!deviceStore.isSingleBoard">
        <ControlForm :isSysSoft="true" />
      </div>
      <div style="background-color: white; width: 100%" v-else>
        <ControlForm :isSysSoft="true" />
      </div>
      <div
        style="background-color: white; margin-left: 20px; width: 100%"
        v-if="!deviceStore.isSingleBoard"
      >
        <CoreForm ref="core" :isSysSoft="true" />
      </div>
    </div>
    <List />
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { PageWrapper } from '/@/components/Page';
  import CoreForm from './components/coreForm.vue';
  import ControlForm from './components/controlForm.vue';
  import List from './components/list.vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useDeviceInfo } from '/@/store/modules/overview';
  import { ref } from 'vue';
  const deviceStore = useDeviceInfo();
  const content = ref('');
  const { t } = useI18n();
  if (deviceStore.isSingleBoard) {
    content.value = t('dashboard.content.sysContent');
  } else {
    content.value = t('dashboard.content.Content');
  }
</script>
