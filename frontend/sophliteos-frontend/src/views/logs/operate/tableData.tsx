import { BasicColumn } from '/@/components/Table/src/types/table';
import { useI18n } from '/@/hooks/web/useI18n';

const { t } = useI18n();
export function getBasicColumns(): BasicColumn[] {
  return [
    {
      title: t('logs.operate.id'),
      dataIndex: 'recordId',
      width: 200,
      align: 'center',
    },
    {
      title: t('logs.operate.type1'),
      dataIndex: 'operationType',
      width: 200,
      align: 'left',
    },
    {
      title: t('logs.operate.funcName'),
      dataIndex: 'operationFunc',
      width: 200,
      align: 'left',
    },
    {
      title: t('logs.operate.people'),
      dataIndex: 'userName',
      width: 200,
      align: 'left',
    },
    {
      title: t('logs.operate.ip'),
      dataIndex: 'operationIp',
      width: 200,
      align: 'left',
    },
    {
      title: t('logs.operate.time'),
      width: 200,
      dataIndex: 'dataTime',
      align: 'left',
    },
    {
      title: t('logs.operate.content'),
      dataIndex: 'operationContent',
      align: 'left',
      ellipsis: true,
    },
  ];
}

export const getFormConfig = () => {
  return {
    labelWidth: 120,
    schemas: [
      {
        field: 'operationFunc',
        label: t('logs.operate.filter.name'),
        component: 'Select',
        slot: 'operationFunc',
        colProps: {
          xl: 12,
          xxl: 8,
        },
      },
      {
        field: 'username',
        label: t('logs.operate.filter.people'),
        component: 'Select',
        slot: 'username',
        colProps: {
          xl: 12,
          xxl: 8,
        },
      },
      {
        field: 'operationIp',
        label: t('logs.operate.filter.ip'),
        component: 'Select',
        slot: 'operationIp',
        colProps: {
          xl: 12,
          xxl: 8,
        },
      },
      {
        field: 'startTime',
        label: t('sys.table.startTime'),
        component: 'DatePicker',
        componentProps: {
          'show-time': true,
        },
        colProps: {
          xl: 12,
          xxl: 8,
        },
      },
      {
        field: 'endTime',
        label: t('sys.table.endTime'),
        component: 'DatePicker',
        componentProps: {
          'show-time': true,
        },
        colProps: {
          xl: 12,
          xxl: 8,
        },
      },
    ],
  };
};
