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
    message: 'positive integer',
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
    message: 'positive float',
    trigger: 'blur',
  },
];

export const algoFormSchema: FormSchema[] = [
  {
    field: 'Threshold',
    label: t('paramConfig.param.algoThreshold'),
    component: 'Input',
    // helpMessage: '0 as default threshold, 0.85, 0.9, ...',
    componentProps: {
      suffix: '',
      allowClear: false,
      placeholder: '0 as default threshold, 0.85, 0.9, ...',
    },
    rules: [
      {
        type: 'number',
        min: 0,
        transform(val) {
          if(val.length == 0) {
            return -1
          }
          return +val;
        },
        message: '0 ~ 1.0',
        trigger: 'blur',
      },
    ],
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
  {
    field: 'ExpansionRatio',
    label: t('paramConfig.param.expansionRatio'),
    component: 'Input',
    // helpMessage: '1.1, 1.2, ...',
    componentProps: {
      suffix: '',
      allowClear: false,
      placeholder: '≥1.0, 1.1, 1.2, ...',
    },
    rules: [
      {
        type: 'number',
        min: 1.0,
        transform(val) {
          return +val;
        },
        message: '≥1.0',
        trigger: 'blur',
      },
    ],
  },
  // {
  //   field: 'RetentionTime',
  //   label: t('paramConfig.param.retentionTime'),
  //   component: 'Input',
  //   componentProps: {
  //     suffix: 's',
  //     allowClear: false,
  //   },
  //   rules: [
  //     {
  //       type: 'integer',
  //       min: 2,
  //       max: 5,
  //       transform(val) {
  //         return +val;
  //       },
  //       message: '2 ~ 5',
  //       trigger: 'blur',
  //     },
  //   ],
  // },
  {
    field: 'NotSaveImage',
    label: t('paramConfig.param.notSaveImage'),
    component: 'Checkbox',
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
    MinDetect: 15,
    MaxDetect: 2500,
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
