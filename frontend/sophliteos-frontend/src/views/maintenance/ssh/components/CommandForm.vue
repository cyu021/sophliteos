<template>
  <BasicModal
    v-bind="$attrs"
    @register="register"
    :title="isAdd ? t('maintenance.ssh.addCommand') : t('maintenance.ssh.editCommand')"
    @ok="submit"
  >
    <div class="pt-3px pr-3px">
      <BasicForm @register="registerForm" :model="modelRef" />
    </div>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { ref } from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, FormSchema, useForm } from '/@/components/Form';
  import { addCommand, editCommand } from '/@/api/maintenance/index';
  import { defineEmits } from 'vue';

  const { t } = useI18n();
  const emits = defineEmits(['form-success']);
  const schemas: FormSchema[] = [
    {
      field: 'name',
      component: 'Input',
      label: t('maintenance.ssh.commandName'),
      colProps: {
        span: 24,
      },
      componentProps: {
        placeholder: t('maintenance.ssh.commandNamePlaceholder'),
      },
      rules: [
        { required: true, message: t('maintenance.ssh.commandNamePlaceholder'), trigger: 'change' },
      ],
      dynamicDisabled: () => {
        return isAdd.value ? false : true;
      },
    },
    {
      field: 'command',
      component: 'Input',
      label: t('maintenance.ssh.command'),
      colProps: {
        span: 24,
      },
      componentProps: {
        placeholder: t('maintenance.ssh.commandPlaceholder'),
      },
      rules: [
        {
          required: true,
          message: t('maintenance.ssh.commandPlaceholder'),
          trigger: 'change',
        },
      ],
    },
  ];
  const modelRef = ref({});
  let isAdd = ref(true);
  const [registerForm, { getFieldsValue, setFieldsValue, validate, resetFields }] = useForm({
    labelWidth: 120,
    schemas,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 24,
    },
  });

  const [register, { closeModal }] = useModalInner((data) => {
    onDataReceive(data);
  });
  function onDataReceive(data) {
    if (data?.add) {
      isAdd.value = true;
      resetFields(); //添加快速命令初始化
    } else {
      isAdd.value = false;
      setFieldsValue(data); //编辑快速命令初始化
    }
  }

  function submit() {
    let params = getFieldsValue();
    validate().then(() => {
      const actionPromise = isAdd.value ? addCommand(params) : editCommand(params);

      actionPromise.then(() => {
        emits('form-sucess', true);
        closeModal();
      });
    });
  }
</script>
