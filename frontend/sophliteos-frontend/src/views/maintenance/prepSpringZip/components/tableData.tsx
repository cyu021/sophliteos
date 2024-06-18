import { BasicColumn } from '/@/components/Table/src/types/table';
import { useI18n } from '/@/hooks/web/useI18n';
const { t } = useI18n();

export function getBasicColumns(): BasicColumn[] {
  const arr = [
    {
      title: t('maintenance.prepSpringZip.customAlgoName'),
      dataIndex: 'name',
      align: 'center',
      ellipsis: true,
    },
    {
      title: t('maintenance.prepSpringZip.createDate'),
      dataIndex: 'createTime',
      align: 'left',
      ellipsis: true,
    },
    {
      title: t('maintenance.prepSpringZip.action'),
      dataIndex: 'action',
      align: 'left',
      ellipsis: true,
    },
  ];
  return arr;
}
