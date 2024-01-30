import { BasicColumn } from '/@/components/Table/src/types/table';
import { useI18n } from '/@/hooks/web/useI18n';
const { t } = useI18n();

export function getHostColumns(): BasicColumn[] {
  const arr = [
    {
      title: t('maintenance.ssh.hostName'),
      dataIndex: 'name',
      align: 'center',
      width: 160,
    },
    {
      title: t('maintenance.ssh.addr'),
      dataIndex: 'addr',
      align: 'center',
      width: 160,
    },
    {
      title: t('maintenance.ssh.user'),
      dataIndex: 'user',
      align: 'center',
      width: 160,
    },
    {
      title: t('maintenance.ssh.port'),
      dataIndex: 'port',
      align: 'center',
      width: 160,
    },
    {
      title: t('maintenance.ssh.description'),
      dataIndex: 'description',
      align: 'center',
      width: 160,
    },
  ];
  return arr;
}
