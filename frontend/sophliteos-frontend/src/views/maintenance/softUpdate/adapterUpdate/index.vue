<template>
  <PageWrapper title="Adapter" content="">
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
                  currentVersion
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
                    :file-list="fileList"
                    @remove="handleRemove"
                    :before-upload="beforeUpload"
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
                    @click="handleUpload"
                    :loading="fileLoading"
                    :disabled="fileLoading || fileList.length < 1"
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
                @click="handleUpgrade"
                :disabled="!fileUploaded"
              >
                {{ t('maintenance.systemUpdate.startUpdate') }}
              </a-button>
              <a-button
                type="primary"
                danger
                style="margin-right: 4.5vw"
                @click="handleRestart"
              >{{
                t('overview.restart')
              }}</a-button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </PageWrapper>
</template>
<script lang="ts" setup> 

import { PageWrapper } from '/@/components/Page';
 import { reactive, ref, onMounted } from 'vue';
import type { UnwrapRef } from 'vue';
import { Upload } from 'ant-design-vue';
import { UploadOutlined } from '@ant-design/icons-vue';

import apis from './apis.js';

import { useMessage } from '/@/hooks/web/useMessage';

import { useI18n } from '/@/hooks/web/useI18n';
const { t } = useI18n();

const currentVersion = ref<string>('');

onMounted(async () => {
  console.log('on mounted')
  await fetchVersion();

  fileUploaded.value = false;
});

const fetchVersion = async () => {
  const version = await apis.currentVersion();
  currentVersion.value = version;
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

const fileLoading = ref(false);
const fileUploaded = ref(false);

const fileList = ref<any>([]);

const handleRemove = (file) => {
  const index = fileList.value.indexOf(file);
  const newFileList = fileList.value.slice();
  newFileList.splice(index, 1);
  fileList.value = newFileList;

  fileUploaded.value = false;
};

const beforeUpload = (file) => {
  fileList.value = [file];
  formState.file = file;
  return false;
};

async function handleUpload() {
  if (fileList.value.length <= 0) {
    createMessage.error('Select file first');
  } else {
    await uploadFile();
  }
}

const handleRestart = async () => {
  const res = await apis.restart();
  if (res.code === 0) {
    createMessage.success('Restart done');
  } else {
    createMessage.error(res.msg);
  }
}

const handleUpgrade = async () => {
  console.log('handle upgrade')

  const res = await apis.upgrade();
  if (res.code === 0) {
    createMessage.success('Upgrade done');
    await fetchVersion();
  } else {
    createMessage.error(res.msg);
  }
}

const uploadFile = async () => {
  const item = formState.file
  try {
    fileUploaded.value = false;
    fileLoading.value = true;

    return apis
      .upload({ name: 'adapterZip', file: item, },
        function onUploadProgress(progressEvent: ProgressEvent) {
          const complete = ((progressEvent.loaded / progressEvent.total) * 99) | 0;
          item.percent = complete;
        }
      )
      .then(async (res) => {
        fileLoading.value = false;
        item.percent = 100;
        if (res.data.code !== 0) {
          createMessage.error(res.data.msg, 3);
        } else {
          fileUploaded.value = true;
          createMessage.success(t('component.upload.uploadSuccess'));
        }
      })
      .catch(() => {
        fileLoading.value = false;
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

</style>
