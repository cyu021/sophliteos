<template>
  <Table :columns="columns" :data-source="data" style="padding: 14px;" :pagination="paginationConfig">
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'action'">
        <Space>
          <Button type="primary" @click="onEdit(record)">Edit</Button>
          <Button type="primary" danger @click="onDelete(record.aiModelName, record.templateName)">Delete</Button>
        </Space>
      </template>
    </template>
  </Table>

  <Button shape="circle" size="large" style="position: absolute; right: 24px; bottom: 24px; width: 60px; height: 60px;" type="primary" @click="onAddNew">
    <template #icon>
      <PlusOutlined />
    </template>
  </Button>

  <RuleTemplateModal 
    :extend="extend" 
    @register="RegisterFilterRuleModal" 
    :algorithmName="algorithmName"
    @success="onModalSuccess" />
</template>
<script setup>
import { Table, Space, Button, } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue'
import { onMounted, ref } from 'vue'
import apis from './api.js'

import RuleTemplateModal from './RuleTemplateModal.vue'
import { useModal } from '/@/components/Modal';

const [RegisterFilterRuleModal, { openModal: openRuleTemplateModal }] = useModal();

const extend = ref({})
const algorithmName = ref('')

const paginationConfig = {
  position: ['bottomCenter'],
}

const onAddNew = () => {
  extend.value = {}
  algorithmName.value = ''

  openRuleTemplateModal(true)
}

const onModalSuccess = () => {
  refresh()
}

const onDelete = (modelName, filterName) => {
  apis.deleteRuleTemplate(modelName, filterName).then((res) => {
    refresh()
  })
}

const cacheRuleTemplate = ref({})

const refresh = () => {
  apis.ruleTemplateList().then((res) => {
    cacheRuleTemplate.value = res

    let templates = []
    for (let key in res) {
      if (res.hasOwnProperty(key)) {
        if (Array.isArray(res[key])) {
          res[key].forEach((item, index) => {
            templates.push({
              aiModelName: key,
              templateName: item.FilterName,
              value: item,
            })
          });
        }
      }
    }

    data.value = templates
  })

}

onMounted(() => {
  refresh()
})

const onEdit = (e) => {
  extend.value = e.value
  algorithmName.value = e.aiModelName

  openRuleTemplateModal(true)
}

const columns = [
  {
    title: 'AI Model Name',
    dataIndex: 'aiModelName',
    key: 'aiModelName',
  },
  {
    title: 'Template Name',
    dataIndex: 'templateName',
    key: 'templateName',
  },
  {
    title: 'Action',
    key: 'action',
  },
];
const data = ref([])
</script>