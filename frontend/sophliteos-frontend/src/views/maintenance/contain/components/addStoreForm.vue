<template>
  <BasicModal
    v-bind="$attrs"
    @register="register"
    :title="t('maintenance.containers.addStore')"
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
  import { addStore, addCoreRegistry } from '/@/api/maintenance/index';
  import { defineEmits } from 'vue';

  const { t } = useI18n();

  const emits = defineEmits(['form-success']);

  const props = defineProps({
    ip: {
      type: String,
      default: '',
    },
  });

  const schemas: FormSchema[] = [
    {
      field: 'insecureRegistry',
      component: 'Input',
      label: t('maintenance.containers.addr'),
      colProps: {
        span: 24,
      },
      componentProps: {
        placeholder: t('maintenance.containers.addrPlaceholder'),
      },
      rules: [
        {
          required: true,
          pattern:
            /^(?:(?:25[0-5]|2[0-4][0-9]|[0-1]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[0-1]?[0-9][0-9]?):\d{1,5}$/,
          message: t('maintenance.containers.addrVlidate') + '，格式为192.168.1.100:9500',
          trigger: 'blur',
        },
      ],
    },
  ];
  const modelRef = ref({});
  const [registerForm, { getFieldsValue, validate, resetFields }] = useForm({
    labelWidth: 120,
    schemas,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 24,
    },
  });

  const [register, { closeModal }] = useModalInner(() => {
    resetFields();
  });

  function submit() {
    let params = getFieldsValue();
    validate().then(() => {
      if (props.ip) {
        params.ip = props.ip;
        addCoreRegistry({ ip: props.ip, insecureRegistry: params.insecureRegistry }).then(() => {
          emits('form-sucess');
          closeModal();
        });
        return;
      }
      addStore(params).then(() => {
        emits('form-sucess', true);
        closeModal();
      });
    });
  }
</script>
