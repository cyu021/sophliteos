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
  import { pullImage, slectImageTaglist } from '/@/api/maintenance/index';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { onMounted, watch } from 'vue';
  import { useMessage } from '/@/hooks/web/useMessage';

  const { t } = useI18n();
  const { createMessage } = useMessage();

  const props = defineProps(['image', 'insecureRegistry']);

  const emit = defineEmits(['next', 'prev', 'imageChoice', 'redo']);

  const step2Schemas: FormSchema[] = [
    {
      field: 'insecureRegistry',
      component: 'Select',
      label: t('maintenance.containers.mirrorName'),
      required: true,
      componentProps: {
        placeholder: t('maintenance.containers.containerPlaceholder'),
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
      text: props.insecureRegistry
        ? t('maintenance.containers.next')
        : t('maintenance.containers.pullMirror'),
    },
    resetFunc: customResetFunc,
    submitFunc: customSubmitFunc,
  });
  const updateInsecureRegistry = async (Registry) => {
    let updateProps = {};
    if (Registry) {
      updateProps = {
        component: 'Select',
        componentProps: {
          options: Registry,
          placeholder: t('maintenance.containers.containerPlaceholder'),
        },
      };
    } else {
      updateProps = {
        component: 'Input',
        componentProps: {
          placeholder: t('maintenance.containers.mirrorInput'),
        },
      };
    }
    updateSchema({
      field: 'insecureRegistry',
      ...updateProps,
    });
    setProps({
      submitButtonOptions: {
        text: props.insecureRegistry
          ? t('maintenance.containers.next')
          : t('maintenance.containers.pullMirror'),
      },
    });
  };

  watch(
    () => props.insecureRegistry,
    () => {
      updateInsecureRegistry(props.insecureRegistry);
    },
  );
  onMounted(() => updateInsecureRegistry(props.insecureRegistry));
  async function customResetFunc() {
    emit('prev');
  }

  async function customSubmitFunc() {
    validate().then(async () => {
      try {
        setProps({
          submitButtonOptions: {
            loading: true,
          },
        });
        if (!props.insecureRegistry) {
          const imageParams = {
            image: `${props.image}/${getFieldsValue().insecureRegistry}`,
          };
          await pullImage(imageParams);
          setProps({
            submitButtonOptions: {
              loading: false,
            },
          });
          createMessage.success(t('maintenance.containers.pullSuccess'));
          emit('redo');
          return;
        }
        const params = { image: getFieldsValue().insecureRegistry, insecureRegistry: props.image };
        const tag = await slectImageTaglist(params);
        const formattedTag = tag.tags.map((item, index) => {
          return {
            label: item,
            value: item,
            key: `${index + 1}`,
          };
        });
        setProps({
          submitButtonOptions: {
            loading: false,
          },
        });
        emit('imageChoice', getFieldsValue());
        emit('next', formattedTag);
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
