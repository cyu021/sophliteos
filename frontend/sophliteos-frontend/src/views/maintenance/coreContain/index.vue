<template>
  <Tabs v-model:activeKey="activeKey" class="core" type="card">
    <TabPane key="contain" :tab="t('dashboard.containers')">
      <CoreContainList :ip="ip" :number="slectBoard[0]" v-if="ip && activeKey === 'contain'" />
    </TabPane>
    <TabPane key="image" :tab="t('dashboard.mirror')">
      <CoreImageListVue :ip="ip" :number="slectBoard[0]" v-if="ip && activeKey === 'image'" />
    </TabPane>
    <TabPane key="regisrty" :tab="t('dashboard.store')">
      <CoreRegistryListVue :ip="ip" :number="slectBoard[0]" v-if="ip && activeKey === 'regisrty'" />
    </TabPane>
    <template #leftExtra>
      <Menu
        mode="horizontal"
        :forceSubMenuRender="true"
        @select="handleSelect"
        :selectedKeys="slectBoard"
        :subMenuCloseDelay="0.3"
      >
        <SubMenu>
          <template #title>{{ `${t('overview.coreBoard')}-${slectBoard[0]}` }}</template>
          <MenuItem v-for="board in boardList" :key="board.number">
            {{ `${t('overview.coreBoard')}-${board.number}` }}
          </MenuItem>
        </SubMenu>
      </Menu>
    </template>
  </Tabs>
</template>

<script lang="ts" setup>
  import { ref, onMounted, watch } from 'vue';
  import { Menu, SubMenu, MenuItem, Tabs, TabPane } from 'ant-design-vue';
  import { useRouter } from 'vue-router';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { getCoreBoardList } from '/@/api/maintenance';
  import CoreContainList from './components/CoreContainList.vue';
  import CoreImageListVue from './components/CoreImageList.vue';
  import CoreRegistryListVue from './components/CoreRegistryList.vue';

  const { t } = useI18n();
  const { push, currentRoute } = useRouter();
  const types = ['contain', 'image', 'regisrty'];
  const type = types.includes(currentRoute.value.params.type as string)
    ? currentRoute.value.params.type
    : 'contain';
  const activeKey = ref(type as string);
  const boardList = ref();
  const boardMap = ref();
  const slectBoard = ref([1]);
  const ip = ref();

  onMounted(() => {
    getCoreBoardMenu();
  });

  watch([activeKey, slectBoard], () => {
    push({ name: 'coreContain', params: { number: slectBoard.value[0], type: activeKey.value } });
  });

  async function getCoreBoardMenu() {
    const res = await getCoreBoardList();
    res.sort((a, b) => a.number - b.number);
    boardList.value = res;
    boardMap.value = new Map(res.map((item) => [item.number, item]));
    // 初始化，如果router上有，就用router上的query,如果没有，就push
    const { number } = currentRoute.value.params;
    if (Number(number)) {
      slectBoard.value = [Number(number)];
      ip.value = boardMap.value.get(Number(number)).ip;
    } else {
      ip.value = res[0].ip;
      push({ name: 'coreContain', params: { number: slectBoard.value[0], type: activeKey.value } });
    }
  }

  function handleSelect(val) {
    slectBoard.value = val.selectedKeys;
    ip.value = boardMap.value.get(slectBoard.value[0]).ip;
  }
</script>

<style lang="less">
  .core {
    .ant-tabs-nav {
      margin: 10px 0 5px 0;
      padding: 0 16px;

      .ant-tabs-extra-content {
        width: 150px;

        .ant-menu-submenu {
          flex: 1;

          .ant-menu-submenu-title {
            text-align: center;
            font-size: 16px;
          }
        }
      }

      .ant-tabs-nav-wrap {
        flex: 1;
        justify-content: flex-end;
        // background-color: white;

        .ant-tabs-nav-list {
          width: 50%;

          .ant-tabs-tab {
            flex: 1;

            .ant-tabs-tab-btn {
              width: 100%;
              text-align: center;
              font-size: 16px;
              letter-spacing: 20px;
            }
          }

          .ant-tabs-tab-active {
            border: solid #0960bd;
          }
        }
      }
    }
  }
</style>
