import { BasicColumn, FormSchema } from '/@/components/Table';
import { useI18n } from '/@/hooks/web/useI18n';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';
// import { options } from '/@/components/Data/algoData';
const { t } = useI18n();

export const columns: BasicColumn[] = [
  {
    title: t('dataSource.privilegeCfg.roleName'),
    dataIndex: 'taskId',
    width: 120,
  },
];

export const taskFormSchema: FormSchema[] = [
  {
    field: 'taskId',
    label: t('taskList.taskList.id'),
    component: 'Input',
    componentProps: {
      disabled: true,
    },
    required: true,
  },
  {
    field: 'urlList',
    label: t('taskList.taskList.videoSource'),
    component: 'Select',
    slot: 'urlList',
    required: true,
  },
  {
    field: 'abilities',
    label: t('taskList.taskList.algorithmConfig'),
    component: 'CheckboxGroup',
    helpMessage: [t('taskList.taskList.algorithmConfig')],
    // required: true,
    // componentProps: {
    //   options: options,
    // },
    slot: 'abilitie',
    colProps: {
      span: 24,
    },
  },
];

export const addTaskSchema: FormSchema[] = [
  {
    field: 'taskId',
    label: t('taskList.taskList.id'),
    component: 'Input',
    required: true,
  },
  {
    field: 'urlList',
    label: t('taskList.taskList.videoSource'),
    component: 'Select',
    slot: 'urlList',
    required: true,
  },
  {
    field: 'abilities',
    label: t('taskList.taskList.algorithmConfig'),
    component: 'CheckboxGroup',
    helpMessage: [t('taskList.taskList.algorithmConfig')],
    // required: true,
    // componentProps: {
    //   options: options,
    // },
    slot: 'abilitie',
    colProps: {
      span: 24,
    },
  },
];

export const algoConfigFormSchema: FormSchema[] = [
  {
    label: t('dataSource.mediaServers.serverIp'),
    field: 'ip',
    component: 'Input',
    // colProps: { span: 8 },
    required: true,
  },
  {
    label: t('dataSource.mediaServers.serverPort'),
    field: 'port',
    component: 'Input',
    required: true,
  },
];

export const upUrlConfigFormSchema: FormSchema[] = [
  {
    label: t('taskList.taskList.upUrl'),
    field: 'ip',
    component: 'Input',
    required: true,
    rules: [
      {
        pattern: /^((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}$/g,
        message: 'IP地址不正确',
        trigger: 'blur',
      },
    ],
  },
];
