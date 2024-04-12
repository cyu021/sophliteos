<template>
  <div style="margin: 20px; display: flex; flex-direction: column">
    <div style="font-size: 20px; margin-bottom: 20px"> 上传新算法包 </div>
    <Space>
      <Upload :file-list="fileList" :before-upload="beforeUpload" @remove="handleRemove">
        <Button>
          <UploadOutlined />
          选择文件
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
          <p v-if="record.sts === 1">启用中</p>
          <p v-else>未启用</p>
        </Space>
      </template>
      <template v-else-if="column.key === 'action'">
        <Space>
          <Button v-if="record.sts === 0" type="primary" @click="handleStart(record)">启用</Button>
          <Button v-if="record.sts === 1" type="primary" @click="handleStop(record)">暂停</Button>
          <Button type="danger" @click="handleDelete(record)">删除</Button>
        </Space>
      </template>
    </template>
  </Table>
</template>
<script setup>
  import { ref, onMounted } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Table, Divider, Button, Space, Upload } from 'ant-design-vue';
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
  const handleUpload = () => {
    const formData = new FormData();
    fileList.value.forEach((file) => {
      formData.append('files[]', file);
    });
    uploading.value = true;

    // You can use any AJAX library you like
    // request('https://www.mocky.io/v2/5cc8019d300000980a055e76', {
    //   method: 'post',
    //   data: formData,
    // })
    //   .then(() => {
    //     fileList.value = [];
    //     uploading.value = false;
    //     message.success('upload successfully.');
    //   })
    //   .catch(() => {
    //     uploading.value = false;
    //     message.error('upload failed.');
    //   });
  };

  const { t } = useI18n();

  const handleStart = (record) => {
    console.log('handle start', record);
  };

  const handleStop = (record) => {
    console.log('handle stop', record);
  };

  const handleDelete = (record) => {
    console.log('handle delete', record);
  };

  const columns = ref([
    {
      title: '算法包名称',
      dataIndex: 'annotator_name',
      key: 'annotator_name',
    },
    {
      title: '当前状态',
      dataIndex: 'sts',
      key: 'sts',
    },
    {
      title: '操作',
      dataIndex: 'action',
      key: 'action',
    },
  ]);

  // 声明表格的数据源
  const dataSource = ref([]);

  onMounted(() => {
    apis.list().then((res) => {
      dataSource.value = res;
    });
  });
</script>
