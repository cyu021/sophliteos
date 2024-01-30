import { BasicColumn } from '/@/components/Table/src/types/table';
import { useI18n } from '/@/hooks/web/useI18n';
const { t } = useI18n();

export function getContainersColumns(): BasicColumn[] {
  const arr: BasicColumn[] = [
    {
      title: t('maintenance.containers.containerName'),
      dataIndex: 'name',
      align: 'center',
      width: 120,
      ellipsis: false,
    },
    {
      title: t('maintenance.containers.mirror'),
      dataIndex: 'image',
      align: 'center',
      width: 150,
      ellipsis: false,
    },
    {
      title: t('maintenance.containers.containerState'),
      dataIndex: 'running',
      align: 'center',
      width: 120,
    },
    {
      title: t('maintenance.containers.containerPort'),
      dataIndex: 'ports',
      align: 'center',
      width: 120,
    },
    {
      title: t('maintenance.containers.containerStarttime'),
      dataIndex: 'createdAt',
      align: 'center',
      width: 120,
    },
    {
      title: t('maintenance.containers.runningTime'),
      dataIndex: 'status',
      align: 'center',
      width: 120,
    },
  ];
  return arr;
}
