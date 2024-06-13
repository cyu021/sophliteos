<template>
  <div>
    <Card ref="topSearch">
      <Descriptions>
        <DescriptionsItem :label="t('paramConfig.param.selectTask')">
          <FormItem style="width: 30%">
            <Select
              v-model:value="taskId"
              :options="taskIdOptions"
              @change="handleSelectChange(false)"
            />
          </FormItem>
        </DescriptionsItem>
      </Descriptions>
    </Card>
    <Card class="!my-4" ref="cardRef">
      <Descriptions
        :title="t('paramConfig.param.channelDetail')"
        class="channel"
        :column="4"
        ref="descRef"
      >
        <DescriptionsItem :label="t('paramConfig.param.channel')">{{
          algoTaskInfo.device.name
        }}</DescriptionsItem>
        <DescriptionsItem :label="t('paramConfig.param.media')">{{
          algoTaskInfo.device.mediaServer
        }}</DescriptionsItem>
        <DescriptionsItem :label="t('paramConfig.param.channelStatus')">{{
          statusMap.get(algoTaskInfo.device.status)
        }}</DescriptionsItem>
        <DescriptionsItem :label="t('paramConfig.param.ratio')">{{
          algoTaskInfo.device.resolution
        }}</DescriptionsItem>
        <DescriptionsItem :label="t('paramConfig.param.protocol')">{{
          protocolMap.get(algoTaskInfo.device.protocol)
        }}</DescriptionsItem>
        <DescriptionsItem :label="t('paramConfig.param.codeType')">{{
          algoTaskInfo.device.codec
        }}</DescriptionsItem>
      </Descriptions>
      <!-- <Descriptions :title="t('paramConfig.param.algoList')" /> -->
      <Card :title="t('paramConfig.param.algoList')" :bordered="false" class="set">
        <template #extra>
          <Space>
            <Button @click="reset" size="small">{{ t('common.resetText') }}</Button>
            <Button type="primary" @click="submit" size="small">{{ t('common.save') }}</Button>
          </Space>
        </template>
        <div class="flex">
          <Tabs
            v-model:activeKey="activeKey"
            @change="handleTabChange"
            tab-position="left"
            ref="tabsRef"
            style="min-width: 130px"
          >
            <TabPane
              v-for="item in algoTaskInfo.algorithms"
              :key="item.Type"
              :tab="abilitiesMap.get(Number(item.Type))"
            />
          </Tabs>
          <div class="flex flexDerecton" style="flex: 1" ref="divRef">
            <div>
              <div style="margin: 20px 0; font-weight: 600">
                <span style="padding-left: 10px; border-left: 2px solid #0960bdb8"></span>
                {{ t('paramConfig.param.setAlgoParams') }}
              </div>
              <BasicForm @register="registerForm" style="width: 450px;" />
              <div style="width: 450px; display: flex; flex-direction: row; align-items: center;">
                <div style="width: 150px;">{{ t('paramConfig.param.filterRule') }}</div>
                <Button style="max-width: 250px; text-overflow: ellipsis; overflow: hidden; white-space: nowrap;" @click="showFilterRule">{{ filterName }}</Button>
                <Button v-if="hasExtend" type="text" danger @click="deleteExtend"><DeleteOutlined /></Button>
              </div>
            </div>
            <div :style="{ width: `${videoWidth}px`, height: `${(9 * videoWidth) / 16 + 30}px` }">
              <div
                class="video"
                :style="{ width: `${videoWidth}px`, height: `${(9 * videoWidth) / 16}px` }"
                style="position: relative; border: 1px solid #f7faff"
              >
                <video autoplay class="w-full h-full" ref="video" controls></video>
                <canvas id="canvasId" v-if="drawRegion || drawLine"></canvas>
              </div>
              <div
                class="flex justify-between pb-1 video"
                style="background-color: #f7faff; border: 1px solid #f7faff; width: 100%"
              >
                <Space>
                  <Button :type="drawRegion ? 'primary' : 'default'" @click="drawPolygon">{{
                    t('paramConfig.draw.region')
                  }}</Button>
                  <Button :type="drawLine ? 'primary' : 'default'" @click="drawLineDerect">{{
                    t('paramConfig.draw.line')
                  }}</Button>
                </Space>
                <Space v-if="drawRegion || drawLine">
                  <Tooltip
                    :title="
                      drawRegion ? t('paramConfig.draw.regionTip') : t('paramConfig.draw.lineTip')
                    "
                    color="#0960bd"
                  >
                    <Button
                      @click="changeDrawMode"
                      :icon="h(EditOutlined)"
                      :style="{
                        color: drawMode ? '#0960bd' : 'black',
                        backgroundColor: drawMode ? '#e4effd' : 'white',
                      }"
                    >{{ drawText }}</Button
                    >
                  </Tooltip>
                  <Tooltip :title="t('paramConfig.draw.rectTip')" color="#0960bd">
                    <Button
                      v-if="drawRegion"
                      @click="drawRect"
                      :icon="h(EditOutlined)"
                      :style="{
                        color: rectMode ? '#0960bd' : 'black',
                        backgroundColor: rectMode ? '#e4effd' : 'white',
                      }"
                    >{{ t('paramConfig.draw.drawRect') }}</Button>
                  </Tooltip>

                  <Select
                    v-if="drawRegion"
                    v-model:value="roiMode"
                    :options="modeOptions"
                    @change="handleRoiModeChange"
                  />

                  <Select
                    v-else
                    v-model:value="crossLineMode"
                    :options="modeOptions"
                    @change="handleCrossModeChange"
                  />

                  <Button @click="clearDraw" :icon="h(ClearOutlined)">{{
                    t('paramConfig.draw.clear')
                  }}</Button>
                </Space>
              </div>
            </div>
          </div>
        </div>
      </Card>
    </Card>
    <FilterRuleModal :extend="extend" @register="RegisterFilterRuleModal" :taskName="taskId" :algorithmName="algorithmName" @success="updateFilterRule" />
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted, nextTick, h, onUnmounted, watch, computed } from 'vue';
  import {
    Button,
    Card,
    Select,
    FormItem,
    Descriptions,
    DescriptionsItem,
    Tabs,
    TabPane,
    Space,
    Tooltip,
  } from 'ant-design-vue';
  import { EditOutlined, ClearOutlined, DeleteOutlined, } from '@ant-design/icons-vue';
  import { fabric } from 'fabric';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { getTaskList } from '/@/api/task';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { useAbilitesStore } from '/@/store/modules/abilities';
  import { AlgoTaskInfoType } from '/@/api/paramConfig/model';
  import { editAlgoTaskConfig, getAlgoTaskConfig } from '/@/api/paramConfig';
  import { ResultEnum } from '/@/enums/httpEnum';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { LivePreview } from '/@/api/dataSource';
  import {
    algoFormSchema,
    AlgoTaskInfoDefault,
    defaultEditConfigParams,
    statusMap,
    protocolMap,
  } from './Data';

  import { WebRTCPlayer } from '@eyevinn/webrtc-player';

  import FilterRuleModal from './FilterRuleModal.vue';
  import { useModal } from '/@/components/Modal';
  import apis from './api';

  const [RegisterFilterRuleModal, { openModal: openFilterRuleModal }] = useModal();

  const { createMessage } = useMessage();
  const { t } = useI18n();

  // video 宽高计算所需值
  const cardRef = ref();
  const tabsRef = ref();
  const videoWidth = ref();
  const descRef = ref();
  const divRef = ref();
  const pixRatio = ref(); // 当前canvas宽高与标准宽高的缩放比，兼容不同大小的canvas提交的绘制

  const topSearch = ref();
  const taskId = ref('');
  const activeKey = ref(1);
  const taskIdOptions = ref([]);
  const algoTaskInfo = ref<AlgoTaskInfoType>(AlgoTaskInfoDefault);
  const algoAbilityStore = useAbilitesStore();
  const abilitiesMap = ref(new Map());
  const currentAlgoParams = ref({
    DetectInterval: 0,
    TrackInterval: 0,
    MinDetect: 0,
    MaxDetect: 0,
  });
  // 视频播放
  const video = ref();
  const player = ref();
  // 绘制
  const canvas = ref();
  const drawRegion = ref(false); // 绘制区域，打开在视频上显示从接口获取的区域
  const drawLine = ref(false); // 绘制检测线，打开在视频上显示从接口获取的检测线
  const drawMode = ref(false); // 打开绘制区域或者绘制检测线之后，打开绘制，可在视频上绘制对应图形
  const rectMode = ref(false); // 快速绘制矩形模式
  const pylogonPoints = ref<any>([]); //当前绘制了的所有点数据
  const lastPoint = ref(); // 上一个点
  const activeLine = ref(); //绘制草稿，移动时跟随鼠标移动的线对象
  const withArrow = ref(false);
  const draftLines = ref([{}]); //绘制草稿，线对象
  const draftPoints = ref([{}]); //绘制草稿，点对象
  const pylogonObject = ref(); //当前区域对象
  const pylogonSubmitPoint = ref<any>([]); //当前区域的点数组
  const lineObject = ref(); //当前检测线对象
  const lineSubmitPoint = ref(); //当前检测线的点
  const originPaintData = ref();
  const rectDownPoint = ref(); //绘制矩形-按下的点
  const rectUpPoint = ref(); //绘制矩形-抬起的点
  const activeRect = ref(); //按下之后移动的临时矩形对象
  const drawText = ref('');

  const roiMode = ref(0);
  const crossLineMode = ref(0);
  const modeOptions = ref([
  { value: 0, label: t('paramConfig.param.top'), },
  { value: 1, label: t('paramConfig.param.bottom'), },
  { value: 2, label: t('paramConfig.param.center'), }
  ]);

  onMounted(() => {
    algoAbilityStore.setAbilities().then((res) => {
      abilitiesMap.value = res;
    });
    getTaskOptions();
    window.addEventListener('resize', computeVideoWidth);
  });

  onUnmounted(() => {
    clearCanvas();
    if (player.value) {
      player.value.unload()
      player.value.destroy()
    }
    window.removeEventListener('resize', computeVideoWidth);
  });

  watch([drawRegion, drawLine], () => {
    if (drawRegion.value) {
      drawText.value = t('paramConfig.draw.drawRegion');
    } else if (drawLine.value) {
      drawText.value = t('paramConfig.draw.drawLineDetect');
    }
  });

  const filterName = computed(() => {
    if (algoTaskInfo.value && activeKey.value) {
      const algorithms = algoTaskInfo.value.algorithms.filter((v) => v.Type === activeKey.value)
      if (algorithms.length > 0) {
        const extend = algorithms[0].Extend;
        if (extend && extend.FilterName) {
          return extend.FilterName;
        }
      }
    }

    return t('paramConfig.param.createFilterRule')
  })

  const hasExtend = computed(() => {    
    if (algoTaskInfo.value && activeKey.value) {
      const algorithms = algoTaskInfo.value.algorithms.filter((v) => v.Type === activeKey.value)
      if (algorithms.length > 0) {
        const extend = algorithms[0].Extend;
        if (extend && extend.FilterName) {
          return true;
        }
      }
    }

    return false;
  })

  const algorithmName = computed(() => {
    if (algoTaskInfo.value && activeKey.value) {
      const algorithms = algoTaskInfo.value.algorithms.filter((v) => v.Type === activeKey.value)
      if (algorithms.length > 0) {
        const algorithm = algorithms[0];
        return abilitiesMap.value.get(Number(algorithm.Type))
      }
    }

    return ''
  })

  const extend = computed(async () => {
    if (algoTaskInfo.value && activeKey.value) {
      const algorithms = algoTaskInfo.value.algorithms.filter((v) => v.Type === activeKey.value)
      if (algorithms.length > 0) {
        return algorithms[0].Extend || {};
      }
    } 

    return {};
  })

  function updateFilterRule(value) {
    if (algoTaskInfo.value && activeKey.value) {
      const algorithms = algoTaskInfo.value.algorithms.filter((v) => v.Type === activeKey.value)
      if (algorithms.length > 0) {
        algorithms[0].Extend = value;
      }
    } 
  }

  async function getCurrentExtend() {
    const rtsp = await apis.videoUrl(algoTaskInfo.value.device.name);

    if (algoTaskInfo.value && activeKey.value) {
      const algorithms = algoTaskInfo.value.algorithms.filter((v) => v.Type === activeKey.value)
      if (algorithms.length > 0) {
        const extend = algorithms[0].Extend;

        if (extend) {
          if (extend.Filter) {
            extend.Filter.forEach(element => {
              element.Rtsp = rtsp;
            });
          }
        }

        return extend;
      }
    } 

    return null;
  }

  function computeVideoWidth() {
    // 只要开始缩放，就关闭canvas
    clearCanvas();
    initDrawSelect();
    clearDraftDraw();
    // 用于计算video宽高的元素或宽高
    const content = document.getElementsByClassName('sophgo-layout-content')[0];
    const body = document.getElementsByTagName('body')[0];
    const topSearchHeight = topSearch.value.$el.clientHeight;
    const cardWidth = cardRef.value.$el.clientWidth;
    const tabWidth = Math.max(tabsRef.value.$el.clientWidth, 605); // 放大时tab宽度被视频覆盖，设置最小宽度防止tab撑不开
    // 计算留给视频区域的最大宽度和高度
    const videoWidthMAx = cardWidth - tabWidth - 50;
    const videoHeightMax =
      content.clientHeight - topSearchHeight - descRef.value.$el.clientHeight - 100;
    // 按照宽高比选择以宽度还是高度为准，页面宽度小于1440时采用上下布局
    if (body.clientWidth <= 1440) {
      videoWidth.value = divRef.value.clientWidth - 100;
    } else if (videoWidthMAx / videoHeightMax > 16 / 9) {
      videoWidth.value = (videoHeightMax * 16) / 9;
    } else {
      videoWidth.value = videoWidthMAx;
    }
    // 计算当前canvas与标准1600*900的尺寸比例
    pixRatio.value = 1600 / videoWidth.value;
    // 缩放之后更新正确的object
    if (originPaintData.value) {
      const [points, line] = originPaintData.value;
      initObjectByApi(points, line);
    }
  }

  const [registerForm, { setFieldsValue, validate }] = useForm({
    labelAlign: 'left',
    labelWidth: 150,
    schemas: algoFormSchema,
    showActionButtonGroup: false,
  });

  async function getTaskOptions() {
    const res = await getTaskList({ pageNo: -1, pageSize: -1 });
    taskIdOptions.value = res.items.map((item) => ({ value: item.taskId, label: item.taskId }));
    taskId.value = res.items[0].taskId;
    handleSelectChange(false);
  }

  async function handleTabChange(val) {
    if (!val) return;
    clearCanvas();
    const algo = algoTaskInfo.value.algorithms.filter((v) => v.Type === activeKey.value)[0];
    const formData = {
      Threshold: algo.Extend?.Threshold || 0,
      // DetectInterval: algo.DetectInterval,
      // TrackInterval: algo.TrackInterval,
      MinDetect: algo.TargetSize.MinDetect,
      MaxDetect: algo.TargetSize.MaxDetect,
      ExpansionRatio: algo.Extend?.roiExtendRatio || 1.0,
      RetentionTime: algo.Extend?.dwellTimeSec || 3,
    };

    roiMode.value = algo.Extend?.detectModeRoi || 0;
    crossLineMode.value = algo.Extend?.detectModeCrossline || 0;

    nextTick(() => {
      setFieldsValue(formData);
      computeVideoWidth();
      initDrawSelect();
      clearDraftDraw();
      // 绘制的点坐标适配当前canvas
      const points = algo.DetectInfos[0].HotArea;
      const line = algo.DetectInfos[0].TripWire;
      originPaintData.value = [points, line];
      initObjectByApi(points, line);
    });
    currentAlgoParams.value = formData;
  }

  async function handleSelectChange(noChangeTab) {
    player.value && player.value.destroy();
    if (!taskId.value) return;
    const res = await getAlgoTaskConfig({ taskId: taskId.value });
    algoTaskInfo.value = res;
    activeKey.value = noChangeTab ? activeKey.value : res.algorithms[0].Type;
    handleTabChange(activeKey.value);
    const url = await LivePreview({ deviceId: algoTaskInfo.value.device.deviceId });
    play(url);
  }

  function handleRoiModeChange(val) {
    roiMode.value = val;
  }

  function handleCrossModeChange(val) {
    crossLineMode.value = val;
  }

  async function submit() {
    const values = await validate();

    lineSubmitPoint.value.forEach((item) => {
      item.x = Math.round(item.x * pixRatio.value) || 0;
      item.y = Math.round(item.y * pixRatio.value) || 0;
    });

    const extend = await getCurrentExtend();
    extend.Threshold = Number(values.Threshold);

    extend.roiExtendRatio = Number(values.ExpansionRatio);
    extend.dwellTimeSec = Number(values.RetentionTime);
    
    extend.detectModeRoi = Number(roiMode.value);
    extend.detectModeCrossline = Number(crossLineMode.value);

    const params: any = {
      taskId: taskId.value,
      Algorithm: {
        ...defaultEditConfigParams,
        Type: activeKey.value,
        // TrackInterval: Number(values.TrackInterval),
        // DetectInterval: Number(values.DetectInterval),
        TargetSize: {
          MinDetect: Number(values.MinDetect),
          MaxDetect: Number(values.MaxDetect),
        },
        DetectInfos: [
          {
            TripWire: {
              LineStart: {
                X: lineSubmitPoint.value?.[0].x,
                Y: lineSubmitPoint.value?.[0].y,
              },
              LineEnd: {
                X: lineSubmitPoint.value?.[1].x,
                Y: lineSubmitPoint.value?.[1].y,
              },
              DirectStart: {
                X: lineSubmitPoint.value?.[3].x,
                Y: lineSubmitPoint.value?.[3].y,
              },
              DirectEnd: {
                X: lineSubmitPoint.value?.[2].x,
                Y: lineSubmitPoint.value?.[2].y,
              },
            },
            HotArea: pylogonSubmitPoint.value.map((item) => ({
              X: Math.round(item.x * pixRatio.value),
              Y: Math.round(item.y * pixRatio.value),
            })),
          },
        ],
        Extend: extend,
      },
    };
    const res = await editAlgoTaskConfig(params);
    if (res.code === ResultEnum.SUCCESS) {
      createMessage.success(t('common.saveSuccess'));
    }
    handleSelectChange(true);
  }

  // 重置，重置表单和绘制
  function reset() {
    setFieldsValue(currentAlgoParams.value);
    clearCanvas();
    initDrawSelect();
    clearDraftDraw();
    const [points, line] = originPaintData.value;
    initObjectByApi(points, line);
  }

  async function play(url: any) {
    console.log('play', url)
    if (player.value) {
      await player.value.unload();
      player.value.destroy();
    }

    player.value = new WebRTCPlayer({
      video: video.value,
      type: 'whep',
    });
    await player.value.load(new URL(url))
  }

  function initObjectByApi(points, line) {
    if (points.length) {
      const convertPoints = points.map((item) => ({
        x: item.X / pixRatio.value,
        y: item.Y / pixRatio.value,
      }));
      pylogonObject.value = makePylogon(convertPoints);
      pylogonSubmitPoint.value = convertPoints;
    }
    if (line) {
      const p1 = { x: line.LineStart.X / pixRatio.value, y: line.LineStart.Y / pixRatio.value };
      const p2 = { x: line.LineEnd.X / pixRatio.value, y: line.LineEnd.Y / pixRatio.value };
      const p3 = { x: line.DirectStart.X / pixRatio.value, y: line.DirectStart.Y / pixRatio.value };
      const p4 = { x: line.DirectEnd.X / pixRatio.value, y: line.DirectEnd.Y / pixRatio.value };
      const line1 = makeLine([p1, p2]);
      const line2 = makeLine([p3, p4], {}, true);
      lineObject.value = makeGroup([line1, line2]);
      lineSubmitPoint.value = [p1, p2, p4, p3];
    }
  }

  function initDrawSelect() {
    drawRegion.value = false;
    drawLine.value = false;
    drawMode.value = false;
    rectMode.value = false;
  }

  function initCanvas() {
    const width = video.value.clientWidth;
    const height = video.value.clientHeight;
    canvas.value = new fabric.Canvas('canvasId', {
      width: width,
      height: height,
      hoverCursor: 'auto',
    });
  }

  function initDraft() {
    draftLines.value = [];
    draftPoints.value = [];
    pylogonPoints.value = [];
    activeLine.value = null;
    lastPoint.value = null;
    withArrow.value = false;
    activeRect.value = null;
    rectDownPoint.value = null;
    rectUpPoint.value = null;
  }

  function clearDraftDraw() {
    if (canvas.value) {
      canvas.value.remove(...draftLines.value);
      canvas.value.remove(...draftPoints.value);
      canvas.value.remove(activeLine.value);
      canvas.value.remove(activeRect.value);
    }
    initDraft();
    // changeDrawMode();
  }

  function clearDraw() {
    if (canvas.value) {
      canvas.value.clear();
      initDraft();
      if (drawLine.value) {
        lineObject.value = null;
        lineSubmitPoint.value = [];
      }
      if (drawRegion.value) {
        pylogonObject.value = null;
        pylogonSubmitPoint.value = [];
      }
    }
  }

  function deleteExtend() {
    if (algoTaskInfo.value && activeKey.value) {
      const algorithms = algoTaskInfo.value.algorithms.filter((v) => v.Type === activeKey.value)
      if (algorithms.length > 0) {
        algorithms[0].Extend = {};
      }
    }
  }

  function showFilterRule() {
    openFilterRuleModal(true);
  }

  function clearCanvas() {
    if (canvas.value) {
      canvas.value.dispose();
      canvas.value = null;
    }
    drawMode.value = false;
  }

  function onMouseDown(event) {
    if (drawMode.value) {
      let opts = {};
      if (pylogonPoints.value.length && lastPoint.value) {
        // 绘制直线
        const thisLine = makeLine([lastPoint.value, event.pointer]);
        canvas.value.add(thisLine);
        draftLines.value.push(thisLine);
      }
      if (drawRegion.value && pylogonPoints.value.length === 0)
        opts = { radius: 15, fill: 'red', stroke: 'red' };
      pylogonPoints.value.push(event.pointer);
      const thisPoint = makePoint(event.pointer, opts);
      canvas.value.add(thisPoint);
      draftPoints.value.push(thisPoint);
      // 完成多边形绘制
      if (
        pylogonPoints.value.length > 3 &&
        checkPointIntersect(pylogonPoints.value[0], event.pointer)
      ) {
        pylogonPoints.value.pop();
        const thisPylogon = makePylogon(pylogonPoints.value);
        canvas.value.add(thisPylogon);
        pylogonObject.value && canvas.value.remove(pylogonObject.value);
        pylogonObject.value = thisPylogon;
        pylogonSubmitPoint.value = pylogonPoints.value;
        clearDraftDraw();
        changeDrawMode();
        return;
      }
      // 完成检测线绘制
      if (drawLine.value && withArrow.value) {
        const lines = makeGroup([...draftLines.value]);
        lineObject.value && canvas.value.remove(lineObject.value);
        lineObject.value = lines;
        lineSubmitPoint.value = [...pylogonPoints.value, lastPoint.value];
        canvas.value.add(lines);
        clearDraftDraw();
        changeDrawMode();
        return;
      }
      lastPoint.value = event.pointer;
    }
  }
  function onMousemove(event) {
    if (drawMode.value) {
      const movePoint = event.pointer;
      if (drawLine.value && pylogonPoints.value.length === 2) {
        lastPoint.value = getCenterPoint(pylogonPoints.value[0], pylogonPoints.value[1]);
        withArrow.value = true;
      }
      if (activeLine.value) {
        canvas.value.remove(activeLine.value);
      }
      if (lastPoint.value) {
        activeLine.value = makeLine([lastPoint.value, movePoint]);
        canvas.value.add(activeLine.value);
      }
    }
  }

  function onDbclick() {
    clearDraftDraw();
  }

  function onMouseDownRect(event) {
    if (rectMode.value) {
      rectDownPoint.value = event.pointer;
    }
  }

  function onMouseMoveRect(event) {
    if (rectMode.value) {
      if (activeRect.value) {
        canvas.value.remove(activeRect.value);
      }
      if (rectDownPoint.value) {
        getRectPoints(rectDownPoint.value, event.pointer);
        const thisPylogon = makePylogon(pylogonPoints.value);
        activeRect.value = thisPylogon;
        canvas.value.add(thisPylogon);
      }
    }
  }

  // up的时候绘制矩形
  function onMouseUpRect(event) {
    if (JSON.stringify(rectDownPoint.value) === JSON.stringify(event.pointer)) {
      clearDraftDraw();
      return;
    }
    if (rectMode.value) {
      rectUpPoint.value = event.pointer;
      getRectPoints(rectDownPoint.value, rectUpPoint.value);
      pylogonSubmitPoint.value = pylogonPoints.value;
      pylogonObject.value && canvas.value.remove(pylogonObject.value);
      const thisPylogon = makePylogon(pylogonPoints.value);
      pylogonObject.value = thisPylogon;
      canvas.value.add(thisPylogon);
      clearDraftDraw();
    }
  }

  function getRectPoints(downPoint, upPoint) {
    const ponit2 = { x: downPoint.x, y: upPoint.y };
    const point4 = { x: upPoint.x, y: downPoint.y };
    pylogonPoints.value = [downPoint, ponit2, upPoint, point4];
  }

  // 开启/关闭绘制
  function changeDrawMode() {
    if (!canvas.value) return;
    drawMode.value = !drawMode.value;
    if (drawMode.value) {
      rectMode.value && drawRect();
      // 开启绘制
      clearDraftDraw(); //初始化草稿
      canvas.value.defaultCursor = 'crosshair'; // 画布光标样式设置为十字
      canvas.value
        .on('mouse:down', onMouseDown)
        .on('mouse:move', onMousemove)
        .on('mouse:dblclick', onDbclick); //开启鼠标监听
    } else {
      canvas.value
        .off('mouse:down', onMouseDown)
        .off('mouse:move', onMousemove)
        .off('mouse:dblclick', onDbclick);
      canvas.value.defaultCursor = 'auto';
    }
  }

  // 快捷绘制-矩形
  function drawRect() {
    rectMode.value = !rectMode.value;
    if (rectMode.value) {
      clearDraftDraw();
      drawMode.value && changeDrawMode();
      //监听鼠标down和up
      canvas.value
        .on('mouse:down', onMouseDownRect)
        .on('mouse:up', onMouseUpRect)
        .on('mouse:move', onMouseMoveRect);
      canvas.value.defaultCursor = 'crosshair'; // 画布光标样式设置为十字
    } else {
      canvas.value
        .off('mouse:down', onMouseDownRect)
        .off('mouse:up', onMouseUpRect)
        .off('mouse:move', onMouseMoveRect);
      canvas.value.defaultCursor = 'auto'; // 画布光标样式设置为十字
    }
  }

  /**
   * 开启关闭区域
   * 先清除删除当前canvas，设置drawmode=false、设置多边形可选、更新darwregion值
   * 若drawregion=true:
   *    置drawline绘制检测线为false，只能绘制一个东西
   *    初始化canvas
   *    如果有多边形object，就把它绘制到canvas上
   */
  function drawPolygon() {
    // pylogonObject.value?.set({ selectable: true });
    clearCanvas();
    drawRegion.value = !drawRegion.value;
    if (drawRegion.value) {
      drawLine.value = false;
      nextTick(() => {
        initCanvas();
        if (pylogonObject.value && canvas.value) {
          canvas.value.add(pylogonObject.value);
        }
      });
    }
  }

  function drawLineDerect() {
    // lineObject.value?.set({ selectable: true });
    clearCanvas();
    drawLine.value = !drawLine.value;
    if (drawLine.value) {
      drawRegion.value = false;
      nextTick(() => {
        initCanvas();
        if (lineObject.value && canvas.value) {
          canvas.value.add(lineObject.value);
        }
      });
    }
  }

  function makePoint(point, opts = {}) {
    const obj = new fabric.Circle({
      originY: 'center',
      originX: 'center',
      left: point.x,
      top: point.y,
      strokeWidth: 1,
      radius: 4,
      fill: 'green',
      stroke: 'green',
      selectable: false,
      hasControls: false,
      hasBorders: false,
      strokeUniform: true,
      ...opts,
    });
    return obj;
  }

  function makeLine(points, opts = {}, needArrow = false) {
    const pointArray: number[] = [];
    points.forEach((v) => {
      pointArray.push(v.x, v.y);
    });
    let color = 'green';
    let arrow;
    if (withArrow.value || needArrow) {
      color = 'red';
      arrow = makeTriangle(points[0], points[1]);
    }
    const obj = new fabric.Line(pointArray, {
      fill: color,
      stroke: color,
      strokeWidth: 2,
      evented: false,
      selectable: false,
      hasControls: false,
      hasBorders: false,
      strokeUniform: true,
      ...opts,
    });
    // 返回带箭头的线
    if (arrow) return makeGroup([obj, arrow]);
    return obj;
  }

  function makePylogon(points, opts = {}) {
    const obj = new fabric.Polygon(points, {
      fill: 'transparent',
      strokeWidth: 2,
      stroke: 'green',
      objectCaching: false,
      transparentCorners: false,
      selectable: false,
      strokeUniform: true,
      ...opts,
    });
    return obj;
  }

  function makeGroup(objs, opts = {}) {
    const group = new fabric.Group(objs, {
      selectable: false,
      hasBorders: true,
      strokeUniform: true,
      ...opts,
    });
    return group;
  }

  // 绘制箭头
  function makeTriangle(point1, point2, opts = {}) {
    let angle = 0;
    const X = point2.x - point1.x;
    const Y = point2.y - point1.y;
    if (X === 0 && Y === 0) return null;
    else if (X === 0 && Y > 0) {
      angle = 180;
    } else if (Y === 0) {
      if (X > 0) angle = 90;
      else angle = -90;
    } else if (X > 0 && Y > 0) {
      angle = getDegree(Math.atan(Y / X)) + 90;
    } else if (X > 0 && Y < 0) {
      angle = getDegree(Math.atan(-X / Y));
    } else if (X < 0 && Y > 0) {
      angle = getDegree(Math.atan(Y / X)) - 90;
    } else if (X < 0 && Y < 0) {
      angle = getDegree(Math.atan(-X / Y));
    }
    const obj = new fabric.Triangle({
      originY: 'center',
      originX: 'center',
      left: point2.x + 1,
      top: point2.y + 1,
      width: 16, // 底边长度
      height: 16, // 底边到对角的距离（三角形的高）
      fill: 'red',
      angle: angle,
      ...opts,
    });
    return obj;
  }

  function getDegree(angle) {
    return angle * (180 / Math.PI);
  }

  // 检测两个点是否重叠，
  function checkPointIntersect(p1, p2) {
    return Math.abs(p1.x - p2.x) <= 15 && Math.abs(p1.y - p2.y) <= 15;
  }

  // 获取两个点的中点
  function getCenterPoint(p1, p2) {
    return { x: (p1.x + p2.x) / 2, y: (p1.y + p2.y) / 2 };
  }
</script>

<style lang="less">
  .channel.ant-descriptions {
    background-color: #f7faff;
    border-radius: 10px;
    padding: 10px;
  }

  .ant-row.ant-form-item {
    margin-bottom: 0;
  }

  .set.ant-card {
    .ant-card-head {
      padding: 0;

      .ant-card-head-title {
        font-weight: 700;
      }
    }

    .ant-card-body {
      padding: 0;
    }
  }

  .ant-tabs-nav-wrap {
    background-color: #f7faff;
    border-radius: 5px;
  }

  .ant-tabs-tab.ant-tabs-tab-active {
    background-color: #e4effd;
  }

  .ant-form-item-control {
    width: 250px !important;
  }

  .flexDerecton {
    flex-direction: row;
  }
  @media screen and (max-width: 1440px) {
    .flexDerecton {
      flex-direction: column;
    }
  }

  .canvas-container {
    position: absolute !important;
    bottom: 0;
  }
</style>
