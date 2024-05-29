<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="submit" width="80vw">
    <BasicForm @register="registerForm">
      <template #urlList="{ model, field }">
        <a-select v-model:value="model[field]" :options="urlList" />
      </template>
      <template #abilitie="{ model, field }">
        <div class="editTask">
          <a-checkbox-group v-model:value="model[field]" :options="algoOptions" />
        </div>
      </template>
    </BasicForm>
  </BasicModal>
</template>

<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { taskFormSchema } from './taskData';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { modTask } from '/@/api/task/index';
  import { CheckboxGroup } from 'ant-design-vue';
  import { getVideosList } from '/@/api/dataSource/index';
  import { ref } from 'vue';
  const urlList = ref();
  const algoOptions = ref();
  const { t } = useI18n();
  const ACheckboxGroup = CheckboxGroup;
  const title = t('taskList.taskList.editTask');
  const emit = defineEmits(['success', 'register']);
  const [registerForm, { setFieldsValue, resetFields, validate }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 10 },
    schemas: taskFormSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 20,
    },
  });
  // onMounted(async () => {
  //   const res = await getVideosList();
  //   urlList.value = res.map((item) => {
  //     //@ts-ignore
  //     const { name, deviceId, url } = item;
  //     return { label: `${name}-${deviceId}-${url}`, value: `${name}-${deviceId}-${url}` };
  //   });
  // });
  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {    
    resetFields();
    algoOptions.value = data.algoOptions;
    setModalProps({ confirmLoading: false });
    let video = '';
    const res = await getVideosList();
    data.usedUrl = data.usedUrl.filter((name) => name !== data.record.deviceName);
    urlList.value = res.filter((item) => !data.usedUrl.includes(item.name));
    urlList.value = urlList.value.map((item) => {
      //@ts-ignore
      const { name, deviceId, url } = item;
      if (name === data.record.deviceName) video = JSON.stringify(item);
      return { label: `${name}-${deviceId}-${url}`, value: JSON.stringify(item) };
    });

    setFieldsValue({
      ...data.record,
      urlList: video,
      abilities: data.record.types,
    });
  });
  async function submit() {
    try {
      const values = await validate();
      const { deviceId } = JSON.parse(values.urlList);
      setModalProps({ confirmLoading: true });
      // TODO custom api
      await modTask({ taskId: values.taskId, deviceId: deviceId, abilities: values.abilities });
      emit('success');
      closeModal();
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>

<style lang="less" scoped>
  .ant-checkbox-group {
    margin-top: 30px;
  }
</style>
<style lang="less">
  .editTask {
    .ant-checkbox-group-item {
      width: 190px;
    }
  }
</style>
