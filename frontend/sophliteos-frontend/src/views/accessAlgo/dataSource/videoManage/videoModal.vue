<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    :showOkBtn="false"
    width="40vw"
    :height="600"
    @cancel="player.destroy()"
  >
    <video ref="video" controls></video>
  </BasicModal>
</template>
<script ts setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { ref } from 'vue';
  const title = ref('');
  const url = ref();
  const video = ref();
  let player = null;
  const [registerModal, { setModalProps }] = useModalInner((data) => {
    title.value = data.record.name || data.record.deviceName;
    url.value = data.res;

    player = new WebRTCPlayer({
      video: video.value,
      type: 'whep',
    });
    player.load(new URL(url.value))

    setModalProps({ confirmLoading: false });
  });
</script>
<style scoped>
  video {
    height: 100%;
    width: 100%;
  }
</style>
