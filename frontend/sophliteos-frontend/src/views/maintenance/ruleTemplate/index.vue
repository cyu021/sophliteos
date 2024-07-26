<template>
  <div style="padding: 24px; padding-right: 64px; justify-content: flex-end; display: flex; flex-direction: row;">
    <Button type="primary" @click="onAddNew">Add Template</Button>
  </div>
  <Table :columns="columns" :data-source="data" style="padding: 14px;" :pagination="paginationConfig">
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
    @success="onModalSuccess" />
</template>
<script setup>
import { Table, Space, Button, } from 'ant-design-vue';
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons-vue'
import { onMounted, ref } from 'vue'
import apis from './api.js'

import RuleTemplateModal from './RuleTemplateModal.vue'
import { useModal } from '/@/components/Modal';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

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
    title: t('maintenance.ruleTemplate.action'),
    key: 'action',
  },
];
const data = ref([])
</script>