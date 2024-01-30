import { BasicColumn } from '/@/components/Table/src/types/table';
import { useI18n } from '/@/hooks/web/useI18n';
const { t } = useI18n();

export function getCommandColumns(): BasicColumn[] {
  const arr = [
    {
      title: t('maintenance.ssh.commandName'),
      dataIndex: 'name',
      align: 'center',
      width: 160,
    },
    {
      title: t('maintenance.ssh.command'),
      dataIndex: 'command',
      align: 'center',
      width: 160,
    },
  ];
  return arr;
}
