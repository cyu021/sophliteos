<template>
  <div class="step1">
    <div class="step1-form">
      <BasicForm @register="register">
        <template #resetBefore>
          <div class="place-hold"></div>
        </template>
      </BasicForm>
    </div>
  </div>
</template>
<script lang="ts" setup>
  import { BasicForm, useForm, FormSchema } from '/@/components/Form';
  import { slectImagelist } from '/@/api/maintenance/index';
  import { getStoreList } from '/@/api/maintenance/index';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { onMounted } from 'vue';
  import { mirrorFromDockerHub } from '../MirrorTableData';
  const { t } = useI18n();

  const emit = defineEmits(['next', 'image', 'change']);
  const step1Schemas: FormSchema[] = [
    {
      field: 'insecureRegistry',
      component: 'Select',
      label: t('maintenance.containers.insecureRegistry'),
      required: true,
      componentProps: {
        placeholder: t('maintenance.containers.insecureRegistryPlaceholder'),
        onChange: (val) => {
          emit('change', val);
        },
      },
      colProps: {
        span: 24,
      },
    },
  ];

  const [register, { getFieldsValue, updateSchema, validate, setProps }] = useForm({
    labelWidth: 100,
    schemas: step1Schemas,
    actionColOptions: {
      span: 14,
    },
    showResetButton: false,
    submitButtonOptions: {
      text: t('maintenance.containers.next'),
    },
    submitFunc: customSubmitFunc,
  });
  onMounted(async () => {
    const formattedArray = [{ ...mirrorFromDockerHub, key: '-1' }];
    const store = await getStoreList();
    const storeArr =
      store?.map((item, index) => ({
        label: item.addr,
        value: item.addr,
        key: `${index + 1}`,
      })) || [];
    updateSchema({
      field: 'insecureRegistry',
      componentProps: {
        options: [...formattedArray, ...storeArr],
      },
    });
  });
  async function customSubmitFunc() {
    let params = getFieldsValue();

    setProps({
      submitButtonOptions: {
        loading: true,
      },
    });
    validate().then(async () => {
      try {
        let formattedRegistry = null;
        if (params.insecureRegistry !== mirrorFromDockerHub.value) {
          const insecureRegistry = await slectImagelist(params);
          formattedRegistry = insecureRegistry.repositories.map((item, index) => ({
            label: item,
            value: item,
            key: `${index + 1}`,
          }));
        }
        setProps({
          submitButtonOptions: {
            loading: false,
          },
        });
        emit('image', params);
        emit('next', formattedRegistry);
      } catch (error) {
        setProps({
          submitButtonOptions: {
            loading: false,
          },
        });
      }
    });
  }
</script>
<style lang="less" scoped>
  .step1 {
    &-form {
      width: 450px;
      margin: 0 auto;
    }
  }

  .place-hold {
    height: 130px;
  }
</style>
