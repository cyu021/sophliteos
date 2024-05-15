<template>
  <div style="margin: 10px; display: flex; flex-direction: column">
    <div class="p-4 mb-2 bg-white">
      <BasicForm @register="registerForm">
        <template #srcId="{ model, field }">
          <a-select v-model:value="model[field]" :options="srcIdOptions" allowClear />
        </template>
        <template #type="{ model, field }">
          <a-select v-model:value="model[field]" :options="typeOptions" mode="multiple" />
        </template>
        <template #taskId="{ model, field }">
          <a-select v-model:value="model[field]" :options="taskIdOptions" allowClear />
        </template>
      </BasicForm>
    </div>
    <div class="p-2 bg-white">
      <div>
        <a-p v-if="total"
          >{{
            t('alarmRetrieval.alarm.totalNumber') + ': ' + total + t('alarmRetrieval.alarm.images')
          }}
        </a-p>
        <a-p v-else
          >{{ t('alarmRetrieval.alarm.totalNumber') + ': 0' + t('alarmRetrieval.alarm.images') }}
        </a-p>
        <a-p style="margin-left: 20px" v-if="maxSize"
          >{{
            t('alarmRetrieval.alarm.useSpace') +
            ' : ' +
            usedSize +
            '  ' +
            t('alarmRetrieval.alarm.maxSpace') +
            ' : ' +
            maxSize
          }}
        </a-p>
        <a-button
          v-if="usedSize"
          style="margin-left: 20px; color: blue; outline: none; border: none"
          @click="settingSpace()"
          >[{{ t('alarmRetrieval.alarm.setSpace') }}]
        </a-button>
      </div>
      <div style="width: 100%">
        <div style="min-height: calc(100vh - 325px); max-height: max-content">
          <a-card
            hoverable
            class="myCard"
            v-for="(item, index) in list"
            :key="index"
            @click="alarmDetail(item)"
          >
            <template #cover>
              <div>
                <img :src="item.smallImage" alt="" style="width: 310px; height: 220px" />
              </div>
            </template>
            <a-card-meta>
              <template #description>
                <div class="des_card">
                  <p style="color: blue">
                    <Icon icon="ri:time-line" />{{
                      ' ' + dayjs(item.time * 1000).format('YYYY-MM-DD HH:mm:ss')
                    }}</p
                  >
                  <p style="color: blue">
                    <Icon icon="bx:detail" />{{ item.taskId || t('alarmRetrieval.alarm.blank') }}</p
                  >
                </div>
                <div class="des_card">
                  <p style="color: blue">
                    <Icon icon="material-symbols:where-to-vote-outline" />{{
                      item.srcId || t('alarmRetrieval.alarm.blank')
                    }}</p
                  >
                  <p style="color: blue">
                    <Icon icon="pajamas:issue-type-incident" />{{
                      item.type
                        ? abilitiesMap.get(Number(item.type)) || t('alarmRetrieval.alarm.notKown')
                        : t('alarmRetrieval.alarm.blank')
                    }}</p
                  >
                </div>
              </template>
            </a-card-meta>
          </a-card>
        </div>
        <a-pagination
          v-model:current="pageNo"
          :total="total"
          :pageSize="pageSize"
          @change="onChange"
          size="small"
          show-quick-jumper
          :show-total="
            (total) =>
              ` ${t('alarmRetrieval.alarm.total')} ${total} ${t('alarmRetrieval.alarm.images')}`
          "
          style="text-align: right"
          class="bg-white"
        />
      </div>
    </div>
    <AlarmDetailModal @register="alarmDetailModal" />
    <SettingSpace @register="registerModal" @success="handleSuccess" />
  </div>
</template>

<script setup>
  // import { useMessage } from '/@/hooks/web/useMessage';
  import { onMounted, ref, watchEffect } from 'vue';
  import { Card, CardMeta, Pagination, Select } from 'ant-design-vue';
  import dayjs from 'dayjs';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useModal } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form';
  import { Icon } from '/@/components/Icon';
  import { AlarmList, getAlarmQueryInfo } from '/@/api/alrmRetrieval/index';
  import { alarmInfo } from '/@/store/modules/alrmRetrieval';
  import { cropperImageByRect } from '/@/utils/image/index';
  import { useAbilitesStore } from '/@/store/modules/abilities';
  import { schemas } from './Data';
  import AlarmDetailModal from './component/AlarmDetailModal.vue';
  import SettingSpace from './component/SettingSpace.vue';

  const { t } = useI18n();
  let params = {
    beginTime: 0,
    endTime: 0,
    pageNo: 1,
    pageSize: 20,
    srcId: '',
    types: [],
    taskId: '',
  };
  const ACard = Card;
  const ACardMeta = CardMeta;
  const ASelect = Select;
  const APagination = Pagination;
  const list = ref(); //卡片列表
  const total = ref();
  const pageNo = ref();
  const pageSize = ref();
  let maxSize = ref('');
  let usedSize = ref('');
  // const ableList = ref(); //检索列表
  const srcIdOptions = ref([]);
  const typeOptions = ref([]);
  const taskIdOptions = ref([]);
  const algoAbilityStore = useAbilitesStore();
  const abilitiesMap = ref(new Map());
  onMounted(() => {
    algoAbilityStore.setAbilities().then((res) => {
      abilitiesMap.value = res;
    });
    getQueryOptions();
    getList(params);
    // ablelist();
  });
  async function getQueryOptions() {
    const res = await getAlarmQueryInfo();
    const { srcIds, types, taskIds } = res;
    srcIdOptions.value = srcIds
      .filter((v) => v)
      .map((item) => ({
        value: item,
        label: item,
      }));
    typeOptions.value = types
      .filter((v) => v)
      .map((item) => ({
        value: item,
        label: abilitiesMap.value.get(Number(item)),
      }));
    taskIdOptions.value = taskIds
      .filter((v) => v)
      .map((item) => ({
        value: item,
        label: item,
      }));
  }
  async function getList(params) {
    const res = await AlarmList(params);
    list.value = res.items;
    total.value = res.total;
    pageNo.value = res.pageNo;
    pageSize.value = res.pageSize;
    maxSize.value = convert(res.maxSize);
    usedSize.value = res.usedSize;
  }
  const [registerModal, { openModal }] = useModal();
  function settingSpace() {
    openModal(true, { max, usedSize });
  }
  function handleSuccess() {
    getList(params);
  }
  const [registerForm, { validate }] = useForm({
    schemas: schemas,
    labelWidth: 100,
    submitOnChange: true,
    labelCol: {
      sm: { span: 24 }, // 小型平板电脑
      md: { span: 12 }, // 中型平板电脑
      lg: { span: 7 }, // 大屏幕桌面
    },
    baseColProps: { span: 6 },
    actionColOptions: { span: 24 },
    autoSubmitOnEnter: true,
    submitFunc: handleSubmit,
  });
  async function handleSubmit() {
    const data = await validate();
    data.beginTime = data.beginTime ? `${data.beginTime} 00:00:00` : 0;
    data.endTime = data.endTime ? `${data.endTime} 23:59:59` : 0;
    params = {
      beginTime: dayjs(data.beginTime).unix(),
      endTime: dayjs(data.endTime).unix(),
      pageNo: 1,
      pageSize: 20,
      srcId: data.srcId,
      types: data.types,
      taskId: data.taskId,
    };
    getList(params);
  }
  async function onChange(p, size) {
    const data = await validate();
    data.beginTime = data.beginTime ? `${data.beginTime} 00:00:00` : data.beginTime;
    data.endTime = data.endTime ? `${data.endTime} 23:59:59` : data.endTime;
    params = {
      beginTime: dayjs(data.beginTime || 0).unix(),
      endTime: dayjs(data.endTime || 0).unix(),
      pageNo: p,
      pageSize: size,
      srcId: data.srcId,
      types: data.types,
      taskId: data.taskId,
    };
    getList(params);
  }

  const [alarmDetailModal, { openModal: openAlarmDetailModal }] = useModal();

  function alarmDetail(item) {
    const store = alarmInfo();
    const { LeftTopX, LeftTopY, RightBtmX, RightBtmY } = item.box;
    item.box = {
      ...item.box,
      x: LeftTopX,
      y: LeftTopY,
      width: RightBtmX - LeftTopX,
      height: RightBtmY - LeftTopY,
    };
    try {
      item.Extend && (item.Extend = JSON.parse(item.Extend));
    } catch {
      item.Extend = {};
    }
    store.setInfo(item);
    openAlarmDetailModal(true, { image: item.bigImage?.split('?')[1] || '' });
  }
  watchEffect(() => {
    if (list.value) {
      list.value.map((item) => {
        cropperImageByRect(item.image, { width: '300', height: '220' }, item.box[0]).then((res) => {
          item.canvas = res;
        });
      });
    }
  });
  const max = ref(0);
  function convert(sizeInMB) {
    max.value = sizeInMB;
    if (sizeInMB < 1024) {
      return sizeInMB + ' MB';
    } else if (sizeInMB < 1024 * 1024) {
      const sizeInGB = sizeInMB / 1024;
      return sizeInGB.toFixed(2) + ' GB';
    } else {
      const sizeInTB = sizeInMB / (1024 * 1024);
      return sizeInTB.toFixed(2) + ' TB';
    }
  }
</script>

<style lang="less">
  .des_card {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    height: 20px;
  }

  .myCard {
    display: inline-block;
    margin-left: 15px !important;
    width: 310px;

    .ant-card-body {
      padding: 15px !important;
    }
  }
</style>
