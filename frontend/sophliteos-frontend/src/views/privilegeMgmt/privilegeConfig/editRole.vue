<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="submit">
    <BasicForm @register="registerForm" style="padding-top: 20px" />
  </BasicModal>
</template>

<script setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { rolePrivFormSchema, isAdd } from './tableData';
  import { useI18n } from '/@/hooks/web/useI18n';
  // import { ModDevice, AddDevice } from '/@/api/dataSource/index';
  import { PostRoleUpsertApi, GetRolePrivTplApi } from '/@/api/sys/user';
  import { ref } from 'vue';

  const { t } = useI18n();
  const emit = defineEmits(['success', 'register', 'error']);

  const title = ref();
  const [registerForm, { resetFields, validate, setFieldsValue }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 21 },
    schemas: rolePrivFormSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 20,
    },
    submitButtonOptions: {
      text: t('dataSource.mediaServers.save'),
    },
  });
  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    resetFields();
    setModalProps({ confirmLoading: false });
    // console.info("data = " + JSON.stringify(data))
    if (data.record == 'add') {
      title.value = t('dataSource.privilegeCfg.addRole');
      isAdd(false);
    } else {
      title.value = t('dataSource.privilegeCfg.editRole');
      isAdd(true);
      setFieldsValue({
        ...data.record,
      });
    }
  });
  async function submit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });

      // console.info("values1 = " + JSON.stringify(values))

      var privCfg;
      GetRolePrivTplApi({}).then((res) => {
        // console.info("GetRolePrivTplApi = " + JSON.stringify(res.data));
        privCfg = res.data.priv_tpl;
        for(const k in privCfg) {
          if(values[k] == true) { continue; }
          privCfg[k]['hide'] = true
        }
      }).then(() => {
        PostRoleUpsertApi({"role_name": values['role_name'], "sitemap_priv": privCfg}).then((resp) => {
          if(resp["code"] != 0) {
            emit('error', resp["msg"])
          } else {
            emit('success');
            closeModal();
          }
        })
      });
      

      // emit('success');
      // closeModal();
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
<style lang="less"></style>
