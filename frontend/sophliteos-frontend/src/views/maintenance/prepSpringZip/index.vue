<template>
  <PageWrapper :title="t('routes.dashboard.prepSpringZip')" content="">
    <div style="display: flex; flex-direction: row; margin-bottom: 20px">
      <div style="background-color: white; width: 100%">
        <a-skeleton :loading="loading" active>
          <div style="font-size: 16px; margin: 20px 0 0 20px; font-weight: 550">{{
            t('maintenance.prepSpringZip.pack')
          }}</div>
          <div class="formStyle">
            <div style="width: 20vw; margin-left: 20px; min-width: 550px; margin-right: 20px">
              <a-form
                :model="formState"
                v-bind="formItemLayout"
                class="!my-4"
                labelAlign="left"
              >
                <a-form-item
                  name="file"
                  required
                  :label="t('maintenance.prepSpringZip.selectModel')"
                  v-model:value="formState.file"
                >
                  <a-upload
                    :file-list="fileList"
                    @remove="handleRemove"
                    :before-upload="beforeUpload"
                    name="model"
                  >
                    <a-button>
                      <a-upload-outlined />
                      {{ t('maintenance.systemUpdate.selectFile') }}
                    </a-button>
                    <br />
                  </a-upload>
                </a-form-item>

                <a-form-item
                  :label="t('maintenance.prepSpringZip.customAlgoName')"
                  :name="algoName"
                  required
                  v-model:value="formState.algoName"
                >
                  <a-input v-model:value="formState.algoName" />
                </a-form-item>

                <!-- <a-form-item :wrapper-col="{ span: 14, offset: 4 }" /> -->
                <a-button
                  type="primary"
                  style="margin-right: 4.5vw"
                  @click="handleUpload"
                  :loading="fileLoading"
                  :disabled="uploading"
                  :class="{ 'is-disabled': uploading }"
                >{{
                  fileLoading
                    ? t('maintenance.systemUpdate.filesUploading')
                    : t('maintenance.prepSpringZip.pack')
                }}</a-button
                >
              </a-form>
            </div>
          </div>
        </a-skeleton>
      </div>
    </div>
    <BasicTable
      @register="registerTable"
      :title="t('maintenance.prepSpringZip.customAlgoDownloadLinks')"
      :noloading="true"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <Space>
            <Button type="primary" :icon="h(CloudDownloadOutlined)" :href="record.url" />
            <Button danger type="primary" :icon="h(DeleteOutlined)" @click="handleDelete(record.name)" />
          </Space>
        </template>
      </template>
    </BasicTable>
  </PageWrapper>
</template>
<script lang="ts" setup>
import { PageWrapper } from '/@/components/Page';
import { useI18n } from '/@/hooks/web/useI18n';

import { reactive, ref, onMounted, watch, h } from 'vue';
import type { UnwrapRef } from 'vue';
import { storeToRefs } from 'pinia';
import { Upload, Progress } from 'ant-design-vue';
import { UploadOutlined } from '@ant-design/icons-vue';

import apis from './apis.js';

import { useDeviceInfo } from '/@/store/modules/overview';
import { useMessage } from '/@/hooks/web/useMessage';
import { AxiosCanceler } from '/@/utils/http/axios/axiosCancel';
import { useGlobSetting } from '/@/hooks/setting';
import { Loading } from '/@/components/Loading';

import { BasicTable, useTable } from '/@/components/Table';
import { getBasicColumns } from './tableData';
import { Space, Button } from 'ant-design-vue';
import { CloudDownloadOutlined, DeleteOutlined } from '@ant-design/icons-vue';

const { t } = useI18n();
const { createMessage } = useMessage();

const [registerTable, { reload }] = useTable({
  title: t('maintenance.prepSpringZip.customAlgoDownloadLinks'),
  api: apis.list,
  columns: getBasicColumns(),
  showTableSetting: true,
  tableSetting: { fullScreen: true },
  showIndexColumn: false,
  indexColumnProps: {
    width: 60,
  },
  rowKey: 'name',
});

const handleDelete = (name) => {
  console.log('handle delete', name);
  apis.deleteFile(name).then((res) => {
    if (res) {
      createMessage.success('Delete successfully');
    } else {
      createMessage.error('Delete failed');
    }

    reload();
  })
}

onMounted(async () => {

});
const { uploadUrl = '' } = useGlobSetting();

const axiosCanceler = new AxiosCanceler();
const AUpload = Upload;
const AProgress = Progress;
const AUploadOutlined = UploadOutlined;

const deviceInfoStore = useDeviceInfo();
const { deviceInfo } = storeToRefs(deviceInfoStore);

const loading = ref(false);
if (!deviceInfo.value.deviceSn) {
  loading.value = true;
  deviceInfoStore.getDeviceInfo().then(() => {
    loading.value = false;
  });
}

interface FormState {
  file: any;
  algoName: string,
}

const formState: UnwrapRef<FormState> = reactive({
  file: '',
  algoName: '',
});

const formItemLayout = {
  labelCol: { span: 10 },
  wrapperCol: { span: 14 },
};
const fileLoading = ref(false);
const tip = ref('');
const show = ref(false);
// 上传逻辑
const fileList = ref<any>([]);
const uploading = ref(false);
const progress = {
  strokeColor: {
    '0%': '#108ee9',
    '100%': '#87d068',
  },
  strokeWidth: 3,
  format: (percent) => `${parseFloat(percent.toFixed(2))}%`,
  class: 'test',
};
const progressStatus = ref('normal');
const handleRemove = (file) => {
  const index = fileList.value.indexOf(file);
  const newFileList = fileList.value.slice();
  newFileList.splice(index, 1);
  fileList.value = newFileList;
  // 取消请求
  axiosCanceler.removePending({ method: 'post', url: uploadUrl });
  uploading.value = false;
};

const beforeUpload = (file) => {
  fileList.value = [...(fileList.value || []), file];
  formState.file = file;
  return false;
};


async function handleUpload() {
  if (fileList.value.length <= 0) {
    createMessage.error('请先选择文件');
  } else {
    return uploadFile()
  }
}
// 定义超时时间为 15 分钟
const timeoutDuration = 60 * 15 * 1000;
// 计时器
let timerId;
async function uploadFile() {
  const item = formState.file
  const algoName = formState.algoName

  try {
    const dataParams = { annoName: algoName };
    fileLoading.value = true;
    progressStatus.value = 'normal';
    tip.value = '上传文件中...';

    return apis.upload({
      name: 'model',
      data: dataParams,
      file: item,
    },
    function onUploadProgress(progressEvent: ProgressEvent) {
        const complete = ((progressEvent.loaded / progressEvent.total) * 99) | 0;
        item.percent = complete;
      },)
      .then(async (res) => {
        fileLoading.value = false;
        item.percent = 100;
        progressStatus.value = 'success';
        if (res.data.code !== 0) {
          createMessage.error(res.data.msg, 3);
        } else {
          createMessage.success(t('component.upload.uploadSuccess'));

          reload();
        }
      })
      .catch(() => {
        fileLoading.value = false;
        progressStatus.value = 'exception';
        createMessage.error(t('component.upload.uploadError'));
        clearTimeout(timerId);
      });
    return {
      success: true,
      error: null,
    };
  } catch (e) {
    item.status = 'error';
    return {
      success: false,
      error: e,
    };
  }
}

//以下是文件上传超时逻辑
// 监听 fileList[0]?.percent 变化
watch(
  () => fileList.value[0]?.percent,
  (newVal, oldVal) => {
    if (oldVal === 0) {
      // 开始计时
      // 设置计时器，在 timeoutDuration 后执行超时逻辑
      timerId = setTimeout(() => {
        handleTimeout(newVal);
      }, timeoutDuration);
    } else if (newVal === 100) {
      // 清空计时器
      clearTimeout(timerId);
    }
  },
  { deep: true }, // 开启深层监听
);
// 超时逻辑
const handleTimeout = (percent) => {
  if (percent < 100) {
    fileLoading.value = false;
    progressStatus.value = 'exception';
    createMessage.error('文件传输超时');
  }
};

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

  .is-disabled,
  .is-disabled:hover,
  .is-disabled:focus,
  .is-disabled:active {
    border-color: #0960bd;
    background: #0960bd;
    color: #fff;
  }
</style>
