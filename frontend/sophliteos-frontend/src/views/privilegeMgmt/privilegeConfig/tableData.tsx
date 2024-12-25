import { BasicColumn, FormSchema } from '/@/components/Table';
import { useI18n } from '/@/hooks/web/useI18n';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';
import { ref } from 'vue';
// import { options } from '/@/components/Data/algoData';
const { t } = useI18n();

const disable = ref();
export function isAdd(param) {
  disable.value = param;
}

export const rolePrivFormSchema: FormSchema[] = [
  {
    field: 'role_name',
    component: 'Input',
    label: t('dataSource.privilegeCfg.roleName'),
    componentProps: {
      disabled: disable,
    },
    required: true,
  },
];
