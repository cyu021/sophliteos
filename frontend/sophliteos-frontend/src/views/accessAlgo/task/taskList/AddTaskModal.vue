<template>
  <BasicModal v-bind="$attrs" @register="registerModal" @ok="submit" width="80vw">
    <BasicForm @register="registerForm">
      <template #urlList="{ model, field }">
        <a-select v-model:value="model[field]" :options="urlList" />
      </template>
      <template #abilitie="{ model, field }">
        <div class="task">
          <a-checkbox-group v-model:value="model[field]" :options="algoOptions" />
        </div>
      </template>
    </BasicForm>
    <template #title>
      <div style="display: flex; flex-direction: row">
        <div>{{ message1 }}</div>
        <Alert :message="message2" show-icon class="alert" />
      </div>
    </template>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { Alert, CheckboxGroup, Select } from 'ant-design-vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { addTask } from '/@/api/task/index';
  import { getVideosList } from '/@/api/dataSource/index';
  import { addTaskSchema } from './taskData';
  import { ref } from 'vue';
  const { t } = useI18n();
  const message1 = t('taskList.taskList.addTask');
  const message2 = t('taskList.taskList.message');
  const ACheckboxGroup = CheckboxGroup;
  const ASelect = Select;
  const urlList = ref();
  const algoOptions = ref();
  const emit = defineEmits(['addsuccess', 'register']);
  const [registerForm, { resetFields, validate }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 10 },
    schemas: addTaskSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 20,
    },
  });
  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    resetFields();
    algoOptions.value = data.algoOptions;
    console.info("add task, algoOptions.value = " + JSON.stringify(algoOptions))
    const res = await getVideosList();
    setModalProps({ confirmLoading: false });
    urlList.value = res.filter((item) => !data.usedUrl.includes(item.name));
    urlList.value = urlList.value.map((item) => {
      //@ts-ignore
      const { name, deviceId, url } = item;
      return { label: `${name}-${deviceId}-${url}`, value: JSON.stringify(item) };
    });
  });

  async function submit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });
      const { name: deviceName, deviceId, url } = JSON.parse(values.urlList);
      await addTask({
        taskId: values.taskId,
        abilities: values.abilities,
        deviceName: deviceName,
        deviceId: deviceId,
        url: url,
      });
      emit('addsuccess');
      resetFields();
      closeModal();
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
<style lang="less" scoped>
  .alert {
    position: relative;
    left: 22%;
    top: 50%;
    width: 50%;
    height: 50px;
    font-size: 14px;
  }

  .ant-checkbox-group {
    margin-top: 30px;
  }
</style>
<style lang="less">
  .task {
    display: flex;
    flex-direction: row;

    .ant-checkbox-group-item {
      width: 190px;
    }
  }
</style>
