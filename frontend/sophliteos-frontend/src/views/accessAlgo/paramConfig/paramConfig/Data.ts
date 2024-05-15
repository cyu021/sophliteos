import { BasicColumn, FormSchema } from '/@/components/Table';
import { useI18n } from '/@/hooks/web/useI18n';
const { t } = useI18n();
export function getBasicColumns(): BasicColumn[] {
  return [
    {
      title: t('paramConfig.param.algoType'),
      dataIndex: 'ability',
      width: 150,
      align: 'left',
    },
    {
      title: t('paramConfig.param.minDetectRange'),
      dataIndex: 'minBox',
      width: 150,
      align: 'left',
    },
    {
      title: t('paramConfig.param.alarmInterval'),
      dataIndex: 'interval',
      width: 150,
      align: 'left',
    },
    {
      title: t('paramConfig.param.algoThreshold'),
      dataIndex: 'threshold',
      width: 150,
      align: 'left',
    },
  ];
}

export const paramFormSchema: FormSchema[] = [
  {
    field: 'ability',
    label: t('paramConfig.param.algoType'),
    component: 'Input',
    componentProps: {
      disabled: true,
    },
  },
  // {
  //   field: 'minBox',
  //   label: t('paramConfig.param.minDetectRange'),
  //   component: 'Input',
  //   slot: 'minBox',
  // },
  {
    field: 'width',
    label: t('paramConfig.param.minDetectRange'),
    component: 'Input',
    slot: 'width',
  },
  {
    field: 'height',
    label: ' ',
    component: 'Input',
    slot: 'height',
  },
  {
    field: 'interval',
    label: t('paramConfig.param.alarmInterval'),
    component: 'Input',
    slot: 'interval',
  },
  {
    field: 'threshold',
    label: t('paramConfig.param.algoThreshold'),
    component: 'Input',
    slot: 'threshold',
  },
];

const rule: any[] = [
  {
    type: 'integer',
    min: 1,
    transform(val) {
      return +val;
    },
    message: '请输入正整数',
    trigger: 'blur',
  },
];

const floatRule: any[] = [
  {
    type: 'number',
    min: -999,
    transform(val) {
      return +val;
    },
    message: '可输入小数',
    trigger: 'blur',
  },
];

export const algoFormSchema: FormSchema[] = [
  {
    field: 'Threshold',
    label: t('paramConfig.param.algoThreshold'),
    component: 'Input',
    componentProps: { },
    rules: floatRule,
  },
  {
    field: 'MinDetect',
    label: t('paramConfig.param.minRatio'),
    component: 'Input',
    componentProps: {
      suffix: 'pix',
      allowClear: false,
    },
    rules: rule,
  },
  {
    field: 'MaxDetect',
    label: t('paramConfig.param.maxRatio'),
    component: 'Input',
    componentProps: {
      suffix: 'pix',
      allowClear: false,
    },
    rules: rule,
  },
];

export const AlgoTaskInfoDefault = {
  device: {
    codec: '', //编码类型
    deviceId: '', //设备id
    name: '', //设备名称
    protocol: Number(), //协议，2表示rtsp协议
    ptzType: Number(),
    resolution: '', //分辨率
    status: '', //设备状态
    type: '', //设备类型
    url: '', //设备地址
    mediaServer: '', //设备接入的流媒体
    mediaPull: Number(), //是否拉流成功
  },
  algorithms: [
    {
      Type: Number(),
      TrackInterval: Number(),
      DetectInterval: Number(),
      TargetSize: {
        MinDetect: Number(),
        MaxDetect: Number(),
      },
      DetectInfos: null,
      Extend: null,
    },
  ],
};

export const defaultEditConfigParams = {
  Type: 1,
  TrackInterval: 2,
  DetectInterval: 5,
  TargetSize: {
    MinDetect: 30,
    MaxDetect: 250,
  },
  DetectInfos: [],
  Extend: null,
};

export const statusMap = new Map([
  ['ON', t('common.ON')],
  ['OFF', t('common.OFF')],
]);

export const protocolMap = new Map([
  [0, '未知'],
  [1, '国标'],
  [2, 'RTSP'],
  [3, '海康SDK'],
  [4, '大华SDK'],
]);
