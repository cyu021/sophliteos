<template>
  <div ref="chartRef" :style="{ width, height }"></div>
</template>
<script lang="ts" setup>
  // @ts-nocheck
  import { Ref, ref, watch } from 'vue';
  import { storeToRefs } from 'pinia';

  // import { Card } from 'ant-design-vue';
  import { useECharts } from '/@/hooks/web/useECharts';
  import { useDeviceInfo } from '/@/store/modules/overview';
  import { useI18n } from '/@/hooks/web/useI18n';
  const { t } = useI18n();

  const props = defineProps({
    loading: Boolean,
    width: {
      type: String as PropType<string>,
      default: '100%',
    },
    height: {
      type: String as PropType<string>,
      default: '300px',
    },
    colors: {
      type: Array,
      default: () => ['#f9e284', '#D89A68', '#B75224', '#E03E4C'],
    },
  });

  const chartRef = ref<HTMLDivElement | null>(null);
  const { setOptions, echarts } = useECharts(chartRef as Ref<HTMLDivElement>);

  // 设备基础信息
  const deviceInfoStore = useDeviceInfo();
  const { chipTemperature: cpuchipTemperatureData, deviceStatus } = storeToRefs(deviceInfoStore);

  watch(
    () => props.loading,
    () => {
      if (props.loading) {
        return;
      }
      setOptions({
        grid: {
          containLabel: true,
          left: '16px',
          right: '16px',
          top: '32px',
          bottom: '0',
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross',
            crossStyle: {
              color: '#999',
            },
          },
        },
        // legend: {
        //   bottom: '0',
        //   data: ['CPU使用率', '内存使用率', '温度'],
        // },
        xAxis: [
          {
            type: 'category',
            data: deviceStatus.value.map((board) => `${t('overview.coreBoard')}-${board.number}`),
            axisPointer: {
              type: 'shadow',
            },
          },
        ],
        yAxis: [
          {
            type: 'value',
            name: '',
            min: 0,
            max: 100,
            interval: 20,
            axisLabel: {
              formatter: '{value} °C',
            },
          },
        ],
        series: [
          {
            name: t('overview.coreTemperature'),
            type: 'line',
            smooth: true,
            tooltip: {
              valueFormatter: function (value) {
                return value + ' °C';
              },
            },
            showSymbol: false,
            lineStyle: {
              width: 0,
            },
            areaStyle: {
              opacity: 0.8,
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: props.colors[3],
                },
                {
                  offset: 0.3,
                  color: props.colors[2],
                },
                {
                  offset: 0.7,
                  color: props.colors[1],
                },
                {
                  offset: 1,
                  color: props.colors[0],
                },
              ]),
            },
            markPoint: {
              data: [
                {
                  type: 'max',
                  name: 'Max',
                  itemStyle: { color: '#B36666' },
                  label: { color: '#FFFFFF' },
                },
                {
                  type: 'min',
                  name: 'Min',
                  itemStyle: { color: '#9999CC' },
                  label: { color: '#FFFFFF' },
                },
              ],
            },
            data: cpuchipTemperatureData,
          },
        ],
      });
    },
    { immediate: true },
  );
</script>
