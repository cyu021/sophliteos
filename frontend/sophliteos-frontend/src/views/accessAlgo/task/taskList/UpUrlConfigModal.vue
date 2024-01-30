<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    @ok="submit"
    :min-height="100"
  >
    <div class="config">
      <BasicForm @register="registerForm" style="padding-top: 30px" />
    </div>
  </BasicModal>
</template>

<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { addUpUrl } from '/@/api/task/index';
  import { upUrlConfigFormSchema } from './taskData';
  const { t } = useI18n();
  const title = t('taskList.taskList.upUrl');
  const emit = defineEmits(['configsuccess', 'register']);
  const [registerForm, { resetFields, validate, setFieldsValue }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 21 },
    schemas: upUrlConfigFormSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 20,
    },
  });
  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    resetFields();
    setFieldsValue({
      ...data,
    });
    setModalProps({ confirmLoading: false });
  });
  async function submit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });
      await addUpUrl({ ip: values.ip });
      emit('configsuccess');
      closeModal();
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
<style lang="less">
  .config {
    vertical-align: middle;
  }
</style>
