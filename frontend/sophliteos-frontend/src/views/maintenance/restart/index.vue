<template>
  <PageWrapper :title="t('dashboard.restart')" >
    <a-tabs v-model:activeKey="activeKey" class="!m-2 !p-2 bg-white" animated>
      <a-tab-pane key="hardware" :tab="t('maintenance.restart.hardware')">
        <div style="background-color: #F0F2F5; height: 240px;">
          <div style="padding: 14px;">
            {{ t('maintenance.restart.hardwareTip') }}
          </div>

          <div style="padding: 14px;">
            {{ hardwareStep }}
          </div>

        </div>
        <div>
          <a-button :loading="hardwareRestartButtonLoading" type="primary" style="margin: 14px;" @click="onHardwareRestart"> {{ t('maintenance.restart.restart') }} </a-button>
        </div>
      </a-tab-pane>
      <a-tab-pane key="services" :tab="t('maintenance.restart.services')">
        <div style="background-color: #F0F2F5; height: 240px;">
          <div style="padding: 14px;">
            {{ t('maintenance.restart.servicesTip') }}
          </div>

          <div style="padding: 14px; white-space: pre-wrap;">
            {{ serviceSteps }}
          </div>

        </div>
        <div>
          <a-button :loading="serviceRestartButtonLoading" type="primary" style="margin: 14px;" @click="onServicesRestart"> {{ t('maintenance.restart.restart') }} </a-button>
        </div>

      </a-tab-pane>
    </a-tabs>
</PageWrapper>
</template>

<script setup>
import { ref } from 'vue';
import { PageWrapper } from '/@/components/Page';
import { Tabs, Button, } from 'ant-design-vue';

import { useI18n } from '/@/hooks/web/useI18n';

import apis from './api.js'

import apisCore from '../softUpdate/coreUpdate/apis.js';
import apisAdapter from '../softUpdate/adapterUpdate/apis.js';
import apisOtaDaemon from '../softUpdate/otaDaemonUpdate/apis.js';

const { t } = useI18n();

const ATabs = Tabs;
const ATabPane = Tabs.TabPane;
const AButton = Button

const activeKey = ref('hardware');
const serviceSteps = ref('');

const hardwareStep = ref('')
const hardwareRestartButtonLoading = ref(false);
const serviceRestartButtonLoading = ref(false);

const onHardwareRestart = async () => {
  hardwareRestartButtonLoading.value = true
  try {
    const res = await apis.hardwareRestart()
    if (res.code === 0) {
      hardwareStep.value = t('maintenance.restart.restartDone')
    } else {
      hardwareStep.value = t('maintenance.restart.failed')
    }
  } catch (error) {
    hardwareStep.value = t('maintenance.restart.failed')
  }
  hardwareRestartButtonLoading.value = false
}

const onServicesRestart = async () => {
  serviceRestartButtonLoading.value = true 

  // 1. 插件
  serviceSteps.value = t('maintenance.restart.plugin') + "..." + t('maintenance.restart.restarting') + "..."
  try {
    const res = await apisCore.restart();
    if (res.code === 0) {
      serviceSteps.value = serviceSteps.value + ' ' + t('maintenance.restart.restartDone')
    } else {
      serviceSteps.value = serviceSteps.value + ' ' + t('maintenance.restart.failed')
    }
  } catch {
    serviceSteps.value = serviceSteps.value + ' ' + t('maintenance.restart.failed')
  }

  // 2. Biz
  serviceSteps.value = serviceSteps.value + "\n"
  serviceSteps.value = t('maintenance.restart.biz') + "..." + t('maintenance.restart.restarting') + "..."
  try {
    const r = await apisAdapter.restart()
    if (r.code === 0) {
      serviceSteps.value = serviceSteps.value + ' ' + t('maintenance.restart.restartDone')
    } else {
      serviceSteps.value = serviceSteps.value + ' ' + t('maintenance.restart.failed')
    }
  } catch {
    serviceSteps.value = serviceSteps.value + ' ' + t('maintenance.restart.failed')
  }

  // 3. Daemon
  serviceSteps.value = serviceSteps.value + "\n"
  serviceSteps.value = t('maintenance.restart.daemon') + "..." + t('maintenance.restart.restarting') + "..."
  try {
    const r3 = await apisOtaDaemon.restart()
    if (r3.code === 0) {
      serviceSteps.value = serviceSteps.value + ' ' + t('maintenance.restart.restartDone')
    } else {
      serviceSteps.value = serviceSteps.value + ' ' + t('maintenance.restart.failed')
    }
  } catch {
    serviceSteps.value = serviceSteps.value + ' ' + t('maintenance.restart.failed')
  }

  serviceRestartButtonLoading.value = false
}

</script>