<template>
  <BasicForm @register="register">
    <template #resetBefore>
      <div class="place-hold"></div>
    </template>
  </BasicForm>
</template>
<script lang="ts" setup>
  import { BasicForm, useForm } from '/@/components/Form';
  import { FormSchema } from '/@/components/Form';
  import { pullImage } from '/@/api/maintenance/index';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { onMounted, watch } from 'vue';
  import { useMessage } from '/@/hooks/web/useMessage';
  const { t } = useI18n();
  const { createMessage } = useMessage();
  const props = defineProps(['image', 'imageChoice', 'tags']);

  const emit = defineEmits(['redo', 'prev']);

  const step2Schemas: FormSchema[] = [
    {
      field: 'tag',
      component: 'Select',
      label: t('maintenance.containers.tag'),
      required: true,
      componentProps: {
        placeholder: t('maintenance.containers.tagPlaceholder'),
      },
      colProps: {
        span: 24,
      },
    },
  ];
  const [register, { validate, getFieldsValue, updateSchema, setProps }] = useForm({
    labelWidth: 80,
    schemas: step2Schemas,
    actionColOptions: {
      span: 14,
    },
    resetButtonOptions: {
      text: t('maintenance.containers.prev'),
    },
    submitButtonOptions: {
      text: t('maintenance.containers.pullMirror'),
    },
    resetFunc: customResetFunc,
    submitFunc: customSubmitFunc,
  });
  const updateTags = async (tags) => {
    updateSchema({
      field: 'tag',
      componentProps: {
        options: tags,
      },
    });
  };
  onMounted(() => updateTags(props.tags));
  watch(() => props.tags, updateTags);

  async function customResetFunc() {
    emit('prev');
  }

  async function customSubmitFunc() {
    setProps({
      submitButtonOptions: {
        loading: true,
      },
    });
    validate().then(async () => {
      try {
        const params = { image: `${props.image}/${props.imageChoice}:${getFieldsValue().tag}` };
        await pullImage(params);
        setProps({
          submitButtonOptions: {
            loading: false,
          },
        });
        createMessage.success(t('maintenance.containers.pullSuccess'));
        emit('redo');
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
  .place-hold {
    height: 130px;
  }
</style>
