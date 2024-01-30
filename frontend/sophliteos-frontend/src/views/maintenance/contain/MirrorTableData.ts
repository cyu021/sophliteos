import { BasicColumn } from '/@/components/Table/src/types/table';
import { useI18n } from '/@/hooks/web/useI18n';
const { t } = useI18n();
const formatSize = (size) => {
  return `${size.toFixed(2)} M`; // 保留两位小数并加上单位M
};
const formatTimestamp = (timestamp) => {
  const date = new Date(timestamp * 1000);
  return date.toLocaleString();
};
const formatTag = (parm, data) => {
  return `${data.Repository}:${data.Tag}`;
};
export function getMirrorColumns(): BasicColumn[] {
  const arr = [
    {
      title: t('maintenance.containers.ID'),
      dataIndex: 'ID',
      align: 'center',
    },
    {
      title: t('maintenance.containers.tag'),
      dataIndex: ['Repository', 'Tag'],
      align: 'center',
      format: formatTag,
      ellipsis: false,
    },
    {
      title: t('maintenance.containers.creat'),
      dataIndex: 'Created',
      align: 'center',
      format: formatTimestamp,
    },
    {
      title: t('maintenance.containers.mirrorSize'),
      dataIndex: 'Size',
      align: 'center',
      format: formatSize,
    },
  ];
  return arr;
}

export const mirrorFromDockerHub = {
  label: t('maintenance.containers.dockerHubLabel'),
  value: 'docker.io',
};

export const getPullMirrorColumns = (): BasicColumn[] => [
  {
    title: t('maintenance.containers.mirrorName'),
    dataIndex: 'name',
    align: 'center',
    width: 100,
    ellipsis: false,
  },
  {
    title: t('maintenance.containers.pullCMD'),
    dataIndex: 'cmd',
    align: 'center',
    width: 150,
    ellipsis: false,
  },
  {
    title: t('maintenance.containers.pullStatus'),
    dataIndex: 'status',
    align: 'center',
    width: 80,
  },
];

export const pullMirrorStatusMap = new Map([
  [0, { label: t('maintenance.pullStatus.succsess'), color: 'green' }],
  [1, { label: t('maintenance.pullStatus.pulling'), color: 'orange' }],
  [2, { label: t('maintenance.pullStatus.failed'), color: 'brown' }],
]);
