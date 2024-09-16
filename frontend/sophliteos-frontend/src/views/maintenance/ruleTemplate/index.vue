<template>
  <div
    style="
      padding: 24px;
      padding-right: 64px;
      justify-content: flex-end;
      display: flex;
      flex-direction: row;
    "
  >
    <Button type="primary" @click="onAddNew">Add</Button>
    <div style="width: 14px"></div>

    <Upload :file-list="[]" accept=".json" :beforeUpload="handleBeforeUpload">
      <Button type="primary">Import</Button>
    </Upload>
    <div style="width: 14px"></div>
    <Button type="primary" @click="onExport">{{ exportButtonTitle }}</Button>
  </div>
  <Table
    :row-selection="rowSelection"
    :columns="columns"
    :data-source="data"
    style="padding: 14px"
    :pagination="paginationConfig"
  >
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'action'">
        <Space>
          <Button type="primary" @click="onEdit(record)">
            <template #icon>
              <EditOutlined />
            </template>
          </Button>

          <Button type="primary" danger @click="onDelete(record.aiModelName, record.templateName)">
            <template #icon>
              <DeleteOutlined />
            </template>
          </Button>
        </Space>
      </template>
    </template>
  </Table>

  <RuleTemplateModal
    :extend="extend"
    @register="RegisterFilterRuleModal"
    :algorithmName="algorithmName"
    @success="onModalSuccess"
  />
</template>
<script setup>
  import { Table, Space, Button, Upload, message } from 'ant-design-vue';
  import {
    PlusOutlined,
    EditOutlined,
    DeleteOutlined,
    DownloadOutlined,
  } from '@ant-design/icons-vue';
  import { computed, onMounted, ref, unref } from 'vue';
  import apis from './api.js';

  import RuleTemplateModal from './RuleTemplateModal.vue';
  import { useModal } from '/@/components/Modal';
  import { useI18n } from 'vue-i18n';
  import { useMessage } from '/@/hooks/web/useMessage';

  const { t } = useI18n();

  const [RegisterFilterRuleModal, { openModal: openRuleTemplateModal }] = useModal();

  const extend = ref({});
  const algorithmName = ref('');

  const paginationConfig = {
    position: ['bottomCenter'],
  };

  const onAddNew = () => {
    extend.value = {};
    algorithmName.value = '';

    openRuleTemplateModal(true);
  };

  const onModalSuccess = () => {
    refresh();
  };

  const { createMessage } = useMessage();
  const handleBeforeUpload = (file) => {
    const reader = new FileReader();
    reader.onload = function (e) {
      const content = JSON.parse(e.target.result);

      apis.updateRuleTemplateJson(content).then(() => {
        refresh();

        createMessage.success('ok');
      });
    };

    reader.readAsText(file);
    return false;
  };

  const selectedRows = ref([]);
  const exportButtonTitle = ref('Export');
  const onSelectChange = (changableRowKeys, rows) => {
    selectedRows.value = rows;
  };
  const selectedRowsKeys = computed(() => {
    return selectedRows.value.map((i) => i.key);
  });
  const rowSelection = ref({
    fixed: true,
    selectedRowKeys: selectedRowsKeys,
    onChange: onSelectChange,
  });

  const onExport = () => {
    const data = selectedRows.value;

    const content = { FilterTemplate: {} };
    data.forEach((d) => {
      const modelName = d.aiModelName;
      let v = content.FilterTemplate[modelName];
      if (v) {
        v.push(d.value);
      } else {
        v = [d.value];
      }
      content.FilterTemplate[modelName] = v;
    });

    saveObjectAsJson(content, 'template.json');
    createMessage.success('ok');

    selectedRows.value = [];
  };

  const onDelete = (modelName, filterName) => {
    apis.deleteRuleTemplate(modelName, filterName).then((res) => {
      refresh();
    });
  };

  const saveObjectAsJson = (object, filename = 'data.json') => {
    const jsonString = JSON.stringify(object, null, 2);
    const blob = new Blob([jsonString], { type: 'application/json' });
    const link = document.createElement('a');
    link.href = URL.createObjectURL(blob);
    link.download = filename;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  };

  const cacheRuleTemplate = ref({});

  const refresh = () => {
    apis.ruleTemplateList().then((res) => {
      cacheRuleTemplate.value = res;

      let templates = [];
      for (let key in res) {
        if (res.hasOwnProperty(key)) {
          if (Array.isArray(res[key])) {
            res[key].forEach((item, index) => {
              templates.push({
                key: key + index,
                aiModelName: key,
                templateName: item.FilterName,
                templateDesc: item.FilterDescription,
                value: item,
              });
            });
          }
        }
      }

      data.value = templates;
    });
  };

  onMounted(() => {
    refresh();
  });

  const onEdit = (e) => {
    extend.value = e.value;
    algorithmName.value = e.aiModelName;

    openRuleTemplateModal(true);
  };

  const columns = [
    {
      title: t('maintenance.ruleTemplate.modelName'),
      dataIndex: 'aiModelName',
      key: 'aiModelName',
    },
    {
      title: t('maintenance.ruleTemplate.templateName'),
      dataIndex: 'templateName',
      key: 'templateName',
    },
    {
      title: t('maintenance.ruleTemplate.templateDesc'),
      dataIndex: 'templateDesc',
      key: 'templateDesc',
    },
    {
      title: t('maintenance.ruleTemplate.action'),
      key: 'action',
    },
  ];
  const data = ref([]);
</script>
