import { FormSchema } from '/@/components/Form';
import { useI18n } from '/@/hooks/web/useI18n';

const { t } = useI18n();
const colProps = {
  span: 6,
};

export const schemas: FormSchema[] = [
  {
    field: 'srcId',
    label: t('alarmRetrieval.alarm.channel'),
    component: 'Select',
    slot: 'srcId',
    colProps,
    defaultValue: '',
  },
  {
    field: 'taskId',
    label: t('alarmRetrieval.alarm.taskSelect'),
    component: 'Select',
    slot: 'taskId',
    colProps,
    defaultValue: '',
  },
  {
    field: 'types',
    label: t('alarmRetrieval.alarm.search'),
    component: 'Select',
    slot: 'type',
    defaultValue: [],
    colProps,
  },
  {
    field: 'beginTime',
    component: 'DatePicker',
    componentProps: {
      valueFormat: 'YYYY-MM-DD',
    },
    label: t('alarmRetrieval.alarm.startTime'),
    colProps,
  },
  {
    field: 'endTime',
    component: 'DatePicker',
    label: t('alarmRetrieval.alarm.EndTime'),
    componentProps: {
      valueFormat: 'YYYY-MM-DD',
    },
    colProps,
  },
];
