import { BasicColumn, FormSchema } from '/@/components/Table';
import { useI18n } from '/@/hooks/web/useI18n';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';
import { ref } from 'vue';

const { t } = useI18n();

export function getBasicColumns(): BasicColumn[] {
  return [
    {
      title: t('dataSource.acctMgmt.acctType'),
      dataIndex: 'label',
      width: 140,
      align: 'center',
      ellipsis: false,
    },
    {
      title: t('dataSource.acctMgmt.acctName'),
      dataIndex: 'user_id',
      width: 140,
      align: 'center',
      ellipsis: false,
    },
    {
      title: t('dataSource.acctMgmt.acctRole'),
      dataIndex: 'role',
      width: 120,
      align: 'center',
      ellipsis: true,
    },
    {
      title: t('dataSource.acctMgmt.acctStatus'),
      dataIndex: 'status',
      width: 80,
      align: 'left',
      customRender: ({ record }) => {
        const status = record.status;
        const enable = status === '1';
        const color = enable ? 'green' : 'red';
        const text = enable ? t('common.ON') : t('common.OFF');
        return h(Tag, { color: color }, () => text);
      },
    },
  ];
}

const disable = ref();
export function isAdd(param) {
  disable.value = param;
}

export const acctSchema: FormSchema[] = [
  {
    field: 'label',
    component: 'Select',
    label: t('dataSource.acctMgmt.acctType'),
    componentProps: {
      options: [
        {
          label: 'Web',
          value: 'Web',
        },
      ],
    },
    required: true,
  },
  {
    field: 'user_id',
    component: 'Input',
    label: t('dataSource.acctMgmt.acctName'),
    componentProps: {
      disabled: disable,
    },
    required: true,
  },
  {
    field: 'role',
    component: 'CheckboxGroup',
    label: t('dataSource.acctMgmt.acctRole'),
    componentProps: {
      options: [
        {
          label: 'Admin',
          value: 'Admin',
        },
      ],
    },
    required: true,
  },
];
