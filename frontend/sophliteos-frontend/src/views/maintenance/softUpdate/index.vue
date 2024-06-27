<template>
  <!-- <PageWrapper title="LiteOS" :content="t('routes.dashboard.content.liteContent')"> -->
  <PageWrapper title="System Upgrade" >
    <a-tabs v-model:activeKey="activeKey" class="!m-2 !p-2 bg-white" animated>
      <a-tab-pane key="control" :tab="controlText">
        <ControlForm :isSoftware="true" />
      </a-tab-pane>
      <a-tab-pane key="otadaemon" :tab="t('maintenance.systemUpdate.otadaemon')">
        <div style="display: flex; flex-direction: row; margin-bottom: 20px">
          <div style="background-color: white">
            <div style="font-size: 16px; margin: 20px 0 0 20px; font-weight: 550">{{
              t('maintenance.systemUpdate.localUpdate1')
            }}</div>
            <div class="formStyle">
              <div style="margin-left: 20px; min-width: 650px; margin-right: 20px">
                <a-form
                  :model="formState"
                  v-bind="formItemLayout"
                  class="!my-4"
                  labelAlign="left"
                >
                  <a-form-item
                    :label="t('maintenance.systemUpdate.updateType')"
                    :rules="[{ required: true }]"
                  >
                    {{ t('maintenance.systemUpdate.localUpdate') }}
                  </a-form-item>

                  <a-form-item
                    :label="t('maintenance.systemUpdate.currentSoftVersion')"
                    :rules="[{ required: true }]"
                  >
                    <span>{{
                      currentVersionOtaDamon
                    }}</span>
                  </a-form-item>

                  <a-form-item
                    name="file"
                    required
                    :label="t('maintenance.systemUpdate.selectFile')"
                    v-model:value="formState.file"
                  >
                    <div style="display: flex; flex-direction: row; justify-content: space-between;">
                      <a-upload
                        :file-list="fileListOtaDaemon"
                        @remove="handleRemoveOtaDaemon"
                        :before-upload="beforeUploadOtaDaemon"
                        name="otadaemonZip"
                      >
                        <a-button>
                          <a-upload-outlined />
                          {{ t('maintenance.systemUpdate.selectFile') }}
                        </a-button>
                      </a-upload>

                      <a-button
                        type="primary"
                        style="margin-right: 4.5vw"
                        @click="handleUploadOtaDaemon"
                        :loading="fileLoadingOtaDaemon"
                        :disabled="fileLoadingOtaDaemon || fileListOtaDaemon.length < 1"
                      >{{
                        t('component.upload.startUpload')
                      }}</a-button>
                    </div>

                  </a-form-item>
                </a-form>

                <div style="padding: 14px;">
                  <a-button
                    type="primary"
                    style="margin-right: 4.5vw"
                    @click="handleUpgradeDaemon"
                    :disabled="!fileUploadedOtaDaemon"
                  >
                    {{ t('maintenance.systemUpdate.startUpdate') }}
                  </a-button>
                  <a-button
                    type="primary"
                    danger
                    style="margin-right: 4.5vw"
                    @click="handleRestartDaemon"
                  >{{
                    t('overview.restart')
                  }}</a-button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </a-tab-pane>
      <a-tab-pane key="adapter" :tab="t('maintenance.systemUpdate.adapter')">
        <div style="display: flex; flex-direction: row; margin-bottom: 20px">
          <div style="background-color: white">
            <div style="font-size: 16px; margin: 20px 0 0 20px; font-weight: 550">{{
              t('maintenance.systemUpdate.localUpdate1')
            }}</div>
            <div class="formStyle">
              <div style="margin-left: 20px; min-width: 650px; margin-right: 20px">
                <a-form
                  :model="formState"
                  v-bind="formItemLayout"
                  class="!my-4"
                  labelAlign="left"
                >
                  <a-form-item
                    :label="t('maintenance.systemUpdate.updateType')"
                    :rules="[{ required: true }]"
                  >
                    {{ t('maintenance.systemUpdate.localUpdate') }}
                  </a-form-item>

                  <a-form-item
                    :label="t('maintenance.systemUpdate.currentSoftVersion')"
                    :rules="[{ required: true }]"
                  >
                    <span>{{
                      currentVersionAdapter
                    }}</span>
                  </a-form-item>

                  <a-form-item
                    name="file"
                    required
                    :label="t('maintenance.systemUpdate.selectFile')"
                    v-model:value="formState.file"
                  >
                    <div style="display: flex; flex-direction: row; justify-content: space-between;">
                      <a-upload
                        :file-list="fileListAdapter"
                        @remove="handleRemoveAdapter"
                        :before-upload="beforeUploadAdapter"
                        name="adapterZip"
                      >
                        <a-button>
                          <a-upload-outlined />
                          {{ t('maintenance.systemUpdate.selectFile') }}
                        </a-button>
                      </a-upload>

                      <a-button
                        type="primary"
                        style="margin-right: 4.5vw"
                        @click="handleUploadAdapter"
                        :loading="fileLoadingAdapter"
                        :disabled="fileLoadingAdapter || fileListAdapter.length < 1"
                      >{{
                        t('component.upload.startUpload')
                      }}</a-button>
                    </div>

                  </a-form-item>
                </a-form>

                <div style="padding: 14px;">
                  <a-button
                    type="primary"
                    style="margin-right: 4.5vw"
                    @click="handleUpgradeAdapter"
                    :disabled="!fileUploadedAdapter"
                  >
                    {{ t('maintenance.systemUpdate.startUpdate') }}
                  </a-button>
                  <a-button
                    type="primary"
                    danger
                    style="margin-right: 4.5vw"
                    @click="handleRestartAdapter"
                  >{{
                    t('overview.restart')
                  }}</a-button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </a-tab-pane>
      <a-tab-pane key="core" :tab="t('maintenance.systemUpdate.core')">
      </a-tab-pane>
    </a-tabs>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import { ref } from 'vue';
  import { PageWrapper } from '/@/components/Page';
  import { Tabs } from 'ant-design-vue';
  import ControlForm from '../sysSoft/components/controlForm.vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  // import { useDeviceInfo } from '/@/store/modules/overview';
  // const deviceStore = useDeviceInfo();

  const { t } = useI18n();

  const ATabs = Tabs;
  const ATabPane = Tabs.TabPane;

  const activeKey = ref('control');

  // const controlText = t('maintenance.systemUpdate.softwareUpdate');
  const controlText = t('maintenance.systemUpdate.portal');

  /////////////////////////////////////////////////////////////////////////
  
 import { reactive, onMounted } from 'vue';
import type { UnwrapRef } from 'vue';
import { Upload } from 'ant-design-vue';
import { UploadOutlined } from '@ant-design/icons-vue';

import apisOtaDaemon from './otaDaemonUpdate/apis.js';

import { useMessage } from '/@/hooks/web/useMessage';


const currentVersionOtaDamon = ref<string>('');

onMounted(async () => {
  console.log('on mounted')
  await fetchVersionOtaDaemon();

  fileUploadedOtaDaemon.value = false;
});

const fetchVersionOtaDaemon = async () => {
  const version = await apisOtaDaemon.currentVersion();
  currentVersionOtaDamon.value = version;
}

const { createMessage } = useMessage();
const AUpload = Upload;
const AUploadOutlined = UploadOutlined;

interface FormState {
  file: any;
}

const formState: UnwrapRef<FormState> = reactive({
  file: '',
});

const formItemLayout = {
  labelCol: { span: 10 },
  wrapperCol: { span: 14 },
};

const fileLoadingOtaDaemon = ref(false);
const fileUploadedOtaDaemon = ref(false);

const fileListOtaDaemon = ref<any>([]);

const handleRemoveOtaDaemon = (file) => {
  const index = fileListOtaDaemon.value.indexOf(file);
  const newFileList = fileListOtaDaemon.value.slice();
  newFileList.splice(index, 1);
  fileListOtaDaemon.value = newFileList;

  fileUploadedOtaDaemon.value = false;
};

const beforeUploadOtaDaemon = (file) => {
  fileListOtaDaemon.value = [file];
  formState.file = file;
  return false;
};

async function handleUploadOtaDaemon() {
  if (fileListOtaDaemon.value.length <= 0) {
    createMessage.error('请先选择文件');
  } else {
    await uploadFileDaemon();
  }
}

const handleRestartDaemon = async () => {
  const res = await apisOtaDaemon.restart();
  if (res.code === 0) {
    createMessage.success('重启成功');
  } else {
    createMessage.error(res.msg);
  }
}

const handleUpgradeDaemon = async () => {
  console.log('handle upgrade')

  const res = await apisOtaDaemon.upgrade();
  if (res.code === 0) {
    createMessage.success('升级成功');
    await fetchVersionOtaDaemon();
  } else {
    createMessage.error(res.msg);
  }
}

const uploadFileDaemon = async () => {
  const item = formState.file
  try {
    fileUploadedOtaDaemon.value = false;
    fileLoadingOtaDaemon.value = true;

    return apisOtaDaemon
      .upload({ name: 'otadaemonZip', file: item, },
        function onUploadProgress(progressEvent: ProgressEvent) {
          const complete = ((progressEvent.loaded / progressEvent.total) * 99) | 0;
          item.percent = complete;
        }
      )
      .then(async (res) => {
        fileLoadingOtaDaemon.value = false;
        item.percent = 100;
        if (res.data.code !== 0) {
          createMessage.error(res.data.msg, 3);
        } else {
          fileUploadedOtaDaemon.value = true;
          createMessage.success(t('component.upload.uploadSuccess'));
        }
      })
      .catch(() => {
        fileLoadingOtaDaemon.value = false;
        createMessage.error(t('component.upload.uploadError'));
      });
  } catch (e) {
    item.status = 'error';
    return {
      success: false,
      error: e,
    };
  }
}

////////////////////////////////////////////////////////

import apisAdapter from './adapterUpdate/apis.js';

const currentVersionAdapter = ref<string>('');

onMounted(async () => {
  console.log('on mounted')
  await fetchVersionAdapter();

  fileUploadedAdapter.value = false;
});

const fetchVersionAdapter = async () => {
  const version = await apisAdapter.currentVersion();
  currentVersionAdapter.value = version;
}

interface FormState {
  file: any;
}

const fileLoadingAdapter = ref(false);
const fileUploadedAdapter = ref(false);

const fileListAdapter = ref<any>([]);

const handleRemoveAdapter = (file) => {
  const index = fileListAdapter.value.indexOf(file);
  const newFileList = fileListAdapter.value.slice();
  newFileList.splice(index, 1);
  fileListAdapter.value = newFileList;

  fileUploadedAdapter.value = false;
};

const beforeUploadAdapter = (file) => {
  fileListAdapter.value = [file];
  formState.file = file;
  return false;
};

async function handleUploadAdapter() {
  if (fileListAdapter.value.length <= 0) {
    createMessage.error('请先选择文件');
  } else {
    await uploadFileAdapter();
  }
}

const handleRestartAdapter = async () => {
  const res = await apisAdapter.restart();
  if (res.code === 0) {
    createMessage.success('重启成功');
  } else {
    createMessage.error(res.msg);
  }
}

const handleUpgradeAdapter = async () => {
  console.log('handle upgrade')

  const res = await apisAdapter.upgrade();
  if (res.code === 0) {
    createMessage.success('升级成功');
    await fetchVersionAdapter();
  } else {
    createMessage.error(res.msg);
  }
}

const uploadFileAdapter = async () => {
  const item = formState.file
  try {
    fileUploadedAdapter.value = false;
    fileLoadingAdapter.value = true;

    return apisAdapter
      .upload({ name: 'adapterZip', file: item, },
        function onUploadProgress(progressEvent: ProgressEvent) {
          const complete = ((progressEvent.loaded / progressEvent.total) * 99) | 0;
          item.percent = complete;
        }
      )
      .then(async (res) => {
        fileLoadingAdapter.value = false;
        item.percent = 100;
        if (res.data.code !== 0) {
          createMessage.error(res.data.msg, 3);
        } else {
          fileUploadedAdapter.value = true;
          createMessage.success(t('component.upload.uploadSuccess'));
        }
      })
      .catch(() => {
        fileLoadingAdapter.value = false;
        createMessage.error(t('component.upload.uploadError'));
      });
  } catch (e) {
    item.status = 'error';
    return {
      success: false,
      error: e,
    };
  }
}

</script>

<style scoped lang="less">
  .ant-form-item .ant-upload .tips {
    font-size: 12px;
    color: red;
    position: relative;
    top: 4px;
    opacity: 0.8;
  }

  .formStyle {
    display: flex;
    flex-direction: row;

    .tips {
      font-size: 12px;
      color: red;
      position: relative;
      top: 4px;
      opacity: 0.8;
    }
  }

  .isSsmbtn {
    margin-top: 105px !important;
  }

  :deep(.ant-upload-list-item-progress) {
    bottom: -20px;
  }

  .md5Style {
    width: 18vw;
    display: flex;
    flex-direction: column;
    position: relative;

    .spanStyle {
      margin-top: 20px;
      margin-bottom: 38px;
    }
  }

  .isClass {
    margin-top: 65px !important;
    position: absolute;
    bottom: 1rem;
  }

</style>./otaDaemonUpdate/apisOtaDaemon.js/index.js
