<template>
  <div style="margin: 20px; display: flex; flex-direction: column">
    <div style="font-size: 20px; margin-bottom: 20px"> {{ t('algorithm.upload') }} </div>
    <Space>
      <Upload :file-list="fileList" :before-upload="beforeUpload" @remove="handleRemove">
        <Button>
          <UploadOutlined />
          {{ t('algorithm.select') }}
        </Button>
      </Upload>
      <Button
        type="primary"
        :disabled="fileList.length === 0"
        :loading="uploading"
        style="margin-left: 16px"
        @click="handleUpload"
      >
        {{ uploading ? 'Uploading' : 'Start Upload' }}
      </Button>
    </Space>
  </div>
  <Divider />
  <Table :columns="columns" :data-source="dataSource">
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'sts'">
        <Space>
          <p v-if="record.sts === 1">{{ t('algorithm.state_on') }}</p>
          <p v-else>{{ t('algorithm.state_off') }}</p>
        </Space>
      </template>
      <template v-else-if="column.key === 'action'">
        <Space>
          <Button v-if="record.sts === 0" type="primary" @click="handleStart(record)">{{ t('algorithm.enable') }}</Button>
          <Button v-if="record.sts === 1" type="primary" @click="handleStop(record)">{{ t('algorithm.disable') }}</Button>
          <Button type="danger" @click="handleDelete(record)">{{ t('algorithm.delete') }}</Button>
        </Space>
      </template>
    </template>
  </Table>
</template>
<script setup>
  import { ref, onMounted } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Table, Divider, Button, Space, Upload, message } from 'ant-design-vue';
  import { UploadOutlined } from '@ant-design/icons-vue';
  import apis from './apis.js';

  const fileList = ref([]);
  const uploading = ref(false);
  const handleRemove = (file) => {
    const index = fileList.value.indexOf(file);
    const newFileList = fileList.value.slice();
    newFileList.splice(index, 1);
    fileList.value = newFileList;
  };

  const beforeUpload = (file) => {
    fileList.value = [...(fileList.value || []), file];
    return false;
  };

  const onUploadProgress = (progressEvent) => {
    console.log('on upload progress:', progressEvent)
  }

  const handleUpload = () => {
    console.log('handle upload', fileList.value)

    if (!fileList.value.length) {
      console.log('fileList value length null')
      return;
    }

    uploading.value = true;

    const file = fileList.value[0];
    apis.upload({
      name: 'modelFile',
      file: file,
      onUploadProgress,
    }).then((res) => {      
      uploading.value = false;
      message.success('upload successfully')

      refresh()

    }).catch((e) => {
      uploading.value = false;
      message.error(e.toString())
    })
  };

  const { t } = useI18n();

  const handleStart = (record) => {
    console.log('handle start', record);
    apis
      .start(record.annotator_name)
      .then((res) => {
        message.success('started')
        refresh()
      }).catch((e) => {
        message.error(e.toString())
      })
  };

  const handleStop = (record) => {
    console.log('handle stop', record);
    apis
      .stop(record.annotator_name)
      .then((res) => {
        message.success('stopped')
        refresh()
      }).catch((e) => {
        message.error(e.toString())
      })
  };

  const handleDelete = (record) => {
    console.log('handle delete', record);
    apis
      .unload(record.annotator_name)
      .then((res) => {
        message.success('deleted')
        refresh()
      }).catch((e) => {
        message.error(e.toString())
      })
  };

  const columns = ref([
    {
      title: t('algorithm.name'),
      dataIndex: 'annotator_name',
      key: 'annotator_name',
    },
    {
      title: t('algorithm.state'),
      dataIndex: 'sts',
      key: 'sts',
    },
    {
      title: t('algorithm.action'),
      dataIndex: 'action',
      key: 'action',
    },
  ]);

  const refresh = () => {
    fileList.value = []

    apis.list().then((res) => {
      dataSource.value = res;
    });
  }

  // 声明表格的数据源
  const dataSource = ref([]);

  onMounted(() => {
    refresh()
  });
</script>
