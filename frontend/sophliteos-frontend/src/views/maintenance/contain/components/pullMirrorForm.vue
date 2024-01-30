<template>
  <BasicModal
    :footer="null"
    @register="register"
    v-bind="$attrs"
    :title="t('maintenance.containers.pullMirror')"
    :minHeight="350"
  >
    <div class="pt-3px pr-3px">
      <div class="step-form-form">
        <Steps :current="current">
          <Steps.Step :title="t('maintenance.containers.selectStoreAddr')" />
          <Steps.Step
            :title="
              image === mirrorFromDockerHub.value
                ? t('maintenance.containers.inputMirror')
                : t('maintenance.containers.selectMirror')
            "
          />
          <Steps.Step
            :title="t('maintenance.containers.selectMirrorTag')"
            v-if="image !== mirrorFromDockerHub.value"
          />
        </Steps>
      </div>
      <div class="mt-5 pull-mirror-form">
        <Step1
          @next="handleStep1Next"
          @change="hangleStep1Change"
          v-show="current === 0"
          @image="handleImage"
          :image="image"
        />
        <Step2
          @prev="handleStepPrev"
          @next="handleStep2Next"
          @imageChoice="handleImageChoice"
          v-show="current === 1"
          :image="image"
          :insecureRegistry="insecureRegistry"
          @redo="handleRedo"
          v-if="state.initStep2"
        />
        <Step3
          @prev="handleStepPrev"
          v-show="current === 2"
          :image="image"
          :imageChoice="imageChoice"
          :tags="tags"
          @redo="handleRedo"
          v-if="state.initStep3"
        />
      </div>
    </div>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { reactive, ref } from 'vue';
  import Step1 from './pullMirrorStep1.vue';
  import Step2 from './pullMirrorStep2.vue';
  import Step3 from './pullMirrorStep3.vue';
  import { Steps } from 'ant-design-vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import {} from '/@/api/maintenance/index';
  import { mirrorFromDockerHub } from '../MirrorTableData';
  // import { defineEmits } from 'vue';

  const { t } = useI18n();
  // const emits = defineEmits();

  const state = reactive({
    initStep2: false,
    initStep3: false,
  });
  const current = ref(0);
  const insecureRegistry = ref(null);
  const image = ref('');
  const tags = ref([]);
  const imageChoice = ref('');
  const [register, { closeModal }] = useModalInner(() => {
    current.value = 0;
    state.initStep2 = false;
    state.initStep3 = false;
  });

  function handleImage(params: any) {
    image.value = params?.insecureRegistry;
  }
  function handleImageChoice(params: any) {
    imageChoice.value = params?.insecureRegistry;
  }
  function hangleStep1Change(value: string) {
    image.value = value;
  }
  function handleStep1Next(step1Values: any) {
    current.value++;
    state.initStep2 = true;
    insecureRegistry.value = step1Values;
  }

  function handleStepPrev() {
    current.value--;
  }

  function handleStep2Next(step2Values: any) {
    current.value++;
    state.initStep3 = true;
    tags.value = step2Values;
  }

  function handleRedo() {
    current.value = 0;
    state.initStep2 = false;
    state.initStep3 = false;
    closeModal();
  }
</script>
<style lang="less" scoped>
  .pull-mirror-form {
    margin-top: 50px;
  }
</style>
