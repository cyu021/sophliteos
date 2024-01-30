<template>
  <BasicDrawer
    v-bind="$attrs"
    @register="register"
    :isDetail="true"
    :title="rowData?.name"
    @close="closeBasicDrawer"
  >
    <Terminal isContainer="true" :containerId="rowData?.id" :isCloseDrawer="isCloseDrawer" />
  </BasicDrawer>
</template>
<script lang="ts" setup>
  import { ref, watchEffect } from 'vue';
  import { BasicDrawer, useDrawerInner } from '/@/components/Drawer';
  // import Terminal from './termial.vue';
  import Terminal from '../../ssh/components/termial.vue';
  import { useRoute } from 'vue-router';

  const rowData = ref();
  const [register, { closeDrawer }] = useDrawerInner((data) => {
    isCloseDrawer.value = false;
    rowData.value = data;
  });

  const isCloseDrawer = ref<Boolean>(false);
  function closeBasicDrawer() {
    isCloseDrawer.value = true;
  }
  // 监控路由变化，若路由有变化时，关闭drawer
  const route = useRoute();
  const currentRoute = ref(route.fullPath);
  watchEffect(() => {
    if (currentRoute.value !== route.fullPath) {
      closeDrawer();
    }
  });
</script>
