<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    :width="1000"
    :height="600"
    :footer="null"
  >
    <Row :gutter="20" class="info">
      <Col :span="17">
        <div class="left-box" ref="imgWrap">
          <img
            :src="image"
            alt=""
            class="img"
            :style="{
              width: `${imageWidth}px`,
              height: `${imageHeight}px`,
              marginTop: `${imageMarginTop}px`,
            }"
          />
        </div>
      </Col>
      <Col :span="7">
        <div class="right-box">
          <Descriptions
            :column="1"
            layout="horizontal"
            :title="t('alarmRetrieval.alarm.detailInfo')"
          >
            <DescriptionsItem :label="t('alarmRetrieval.alarm.taskId')">{{
              iamgeInfo.taskId
            }}</DescriptionsItem>
            <DescriptionsItem :label="t('alarmRetrieval.alarm.alarmType')">{{
              abilitiesMap.get(Number(iamgeInfo.type))
            }}</DescriptionsItem>
            <DescriptionsItem :label="t('alarmRetrieval.alarm.srcId')">{{
              iamgeInfo.srcId || t('alarmRetrieval.alarm.blank')
            }}</DescriptionsItem>
            <DescriptionsItem :label="t('alarmRetrieval.alarm.captureTime')">{{
              dayjs(iamgeInfo.time * 1000).format('YYYY-MM-DD HH:mm:ss')
            }}</DescriptionsItem>

            <DescriptionsItem
              v-for="item in Object.keys(iamgeInfo.Extend)"
              :label="item"
              :key="item"
              >{{ iamgeInfo.Extend[item] }}</DescriptionsItem
            >
          </Descriptions>
        </div>
      </Col>
    </Row>
  </BasicModal>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Descriptions, DescriptionsItem, Row, Col } from 'ant-design-vue';
  import dayjs from 'dayjs';
  import { storeToRefs } from 'pinia';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { getAlarmImage } from '/@/api/alrmRetrieval/index';
  import { drawRectOnImage } from '/@/utils/image/index';
  // import { option } from '/@/components/Data/algoData';
  import { alarmInfo } from '/@/store/modules/alrmRetrieval';
  import { useAbilitesStore } from '/@/store/modules/abilities';

  const { t } = useI18n();
  const title = t('alarmRetrieval.alarm.alarmDetail');
  const store = alarmInfo();
  const algoAbilityStore = useAbilitesStore();
  const { abilitiesMap } = storeToRefs(algoAbilityStore);
  const imgWrap = ref();
  const imageWidth = ref();
  const imageHeight = ref();
  const imageMarginTop = ref(0);

  const { iamgeInfo } = storeToRefs(store);
  const image = ref();

  const [registerModal, { setModalProps }] = useModalInner(async (data) => {
    await getImage(data.image);
    setModalProps({ confirmLoading: false });
  });

  async function getImage(url: string) {
    await getAlarmImage(url).then((res) => {
      const myBlob = new window.Blob([res.data], { type: 'image/jpg' });
      drawRectOnImage(window.URL.createObjectURL(myBlob), iamgeInfo.value.box).then((res) => {
        computedWidthHeight(res);
        image.value = res;
      });
    });
  }

  function computedWidthHeight(image) {
    const img = new Image();
    img.src = image;
    img.onload = () => {
      const { width, height } = img;
      const rate = width / height;
      const wrapWidth = imgWrap.value.clientWidth;
      const wrapHeight = imgWrap.value.clientHeight;
      const rateStand = wrapWidth / wrapHeight;
      if (rate >= rateStand) {
        imageWidth.value = wrapWidth;
        imageHeight.value = (rateStand * wrapWidth) / rate;
      } else {
        imageWidth.value = wrapHeight * rate;
        imageHeight.value = wrapHeight;
      }
    };
  }
</script>
<style lang="less" scoped>
  .info {
    padding: 0 10px;
  }

  .left-box {
    // flex: 8; /* 8:2 宽度比例 */
    // width: 650px;
    // padding: 20px; /* 可以根据需要添加内边距 */
    border: 2px dotted #d4d4d4;
    height: 600px;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .right-box {
    // width: 250px;
    // border-left: 1px solid; /* 可以根据需要添加边框样式 */
    height: 600px;
    padding: 0 0 0 10px;
  }

  .img {
    margin: 0 auto;
  }
</style>
