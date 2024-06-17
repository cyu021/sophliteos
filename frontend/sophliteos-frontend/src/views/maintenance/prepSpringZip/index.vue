<template>
  <PageWrapper :title="t('routes.dashboard.prepSpringZip')" content="">
    <div style="display: flex; flex-direction: row; margin-bottom: 20px">
      <div style="background-color: white; width: 100%">
        <ControlForm :isSysSoft="true" />
      </div>
    </div>
    <List />
  </PageWrapper>
</template>
<script setup>
import { PageWrapper } from '/@/components/Page';
import ControlForm from './components/controlForm.vue';
import List from './components/list.vue';
import { useI18n } from '/@/hooks/web/useI18n';
import { ref } from 'vue';
const { t } = useI18n();
</script>
