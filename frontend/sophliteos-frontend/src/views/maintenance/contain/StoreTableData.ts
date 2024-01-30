import { BasicColumn } from '/@/components/Table/src/types/table';
import { useI18n } from '/@/hooks/web/useI18n';
const { t } = useI18n();

export function getStoreColumns(): BasicColumn[] {
  const arr = [
    {
      title: t('maintenance.containers.storeName'),
      dataIndex: 'addr',
      align: 'center',
    },
  ];
  return arr;
}
