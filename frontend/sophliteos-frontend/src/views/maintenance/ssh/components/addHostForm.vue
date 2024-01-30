<template>
  <BasicModal
    v-bind="$attrs"
    @register="register"
    :title="isAdd ? t('maintenance.ssh.addHost') : t('maintenance.ssh.editHost')"
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
  import { addHost, editHost } from '/@/api/maintenance/index';
  import { defineEmits } from 'vue';

  const { t } = useI18n();
  const emits = defineEmits(['form-success']);
  const schemas: FormSchema[] = [
    {
      field: 'name',
      component: 'Input',
      label: t('maintenance.ssh.hostName'),
      colProps: {
        span: 24,
      },
      componentProps: {
        placeholder: t('maintenance.ssh.hostNamePlaceholder'),
      },
      rules: [
        { required: true, message: t('maintenance.ssh.hostNamePlaceholder'), trigger: 'change' },
      ],
    },
    {
      field: 'addr',
      component: 'Input',
      label: t('maintenance.ssh.addr'),
      colProps: {
        span: 24,
      },
      componentProps: {
        placeholder: t('maintenance.ssh.addrPlaceholder'),
      },
      rules: [
        {
          required: true,
          pattern:
            /^((25[0-5]|2[0-4][0-9]|[0-1]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[0-1]?[0-9][0-9]?)$/,
          message: t('maintenance.ssh.addrvVlidate'),
          trigger: 'blur',
        },
      ],
      dynamicDisabled: () => {
        return isAdd.value ? false : true;
      },
    },
    {
      field: 'user',
      component: 'Input',
      label: t('maintenance.ssh.user'),
      colProps: {
        span: 24,
      },
      componentProps: {
        placeholder: t('maintenance.ssh.userPlaceholder'),
      },
      rules: [{ required: true, message: t('maintenance.ssh.userPlaceholder'), trigger: 'change' }],
    },
    {
      field: 'password',
      component: 'InputPassword',
      label: t('maintenance.ssh.password'),
      colProps: {
        span: 24,
      },
      componentProps: {
        placeholder: t('maintenance.ssh.passwordPlaceholder'),
      },
      rules: [
        { required: true, message: t('maintenance.ssh.passwordPlaceholder'), trigger: 'change' },
      ],
    },
    {
      field: 'port',
      component: 'InputNumber',
      label: t('maintenance.ssh.port'),
      colProps: {
        span: 24,
      },
      componentProps: {
        placeholder: t('maintenance.ssh.portPlaceholder'),
      },
      rules: [
        {
          required: true,
          pattern: /^[1-9]\d*$/,
          message: t('maintenance.ssh.portVlidate'),
          trigger: 'blur',
        },
      ],
    },
    {
      field: 'description',
      component: 'InputTextArea',
      label: t('maintenance.ssh.description'),
      colProps: {
        span: 24,
      },
      componentProps: {
        placeholder: t('maintenance.ssh.descriptionPlaceholder'),
      },
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

  const [register, { closeModal, changeOkLoading }] = useModalInner((data) => {
    onDataReceive(data);
  });
  function onDataReceive(data) {
    if (data?.add) {
      isAdd.value = true;
      resetFields(); //添加主机初始化
    } else {
      isAdd.value = false;
      setFieldsValue(data); //编辑主机初始化
    }
  }

  function submit() {
    let params = getFieldsValue();
    params.password = btoa(params.password); //密码编码成64位的格式
    validate().then(() => {
      changeOkLoading(true);
      const actionPromise = isAdd.value ? addHost(params) : editHost(params);
      actionPromise
        .then(() => {
          emits('form-sucess', true);
          changeOkLoading(false);
          closeModal();
        })
        .catch(() => {
          changeOkLoading(false);
        });
    });
  }
</script>
