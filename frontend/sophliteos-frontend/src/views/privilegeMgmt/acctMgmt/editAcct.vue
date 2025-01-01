<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="submit">
    <BasicForm @register="registerForm" style="padding-top: 20px" />
  </BasicModal>
</template>

<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { acctSchema, isAdd } from './tableData';
  import { useI18n } from '/@/hooks/web/useI18n';
  // import { ModDevice, AddDevice } from '/@/api/dataSource/index';
  import { PostAcctUpsertApi } from '/@/api/sys/user';
  import { ref } from 'vue';
  import md5 from 'crypto-js/md5';

  const { t } = useI18n();
  const emit = defineEmits(['success', 'register', 'error']);

  const title = ref();
  const [registerForm, { resetFields, validate, setFieldsValue }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 21 },
    schemas: acctSchema,
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
    // console.info("editAcct data = " + JSON.stringify(data))
    if (data.record == 'add') {
      title.value = t('dataSource.acctMgmt.addAcct');
      isAdd(false);
    } else {
      title.value = t('dataSource.acctMgmt.editAcct');
      isAdd(true);
      setFieldsValue({
        ...data.record,
      });
    }
    // console.info(acctSchema);
  });
  async function submit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });

      // console.info("values1 = " + JSON.stringify(values))
      // if (title.value == t('dataSource.acctMgmt.editAcct')) {
      //   // await ModDevice({
      //   //   name: values.name,
      //   //   protocol: values.protocol,
      //   //   ptzType: values.ptzType,
      //   //   url: values.url,
      //   //   deviceId: deviceId.value,
      //   // });
      // } else {
        // await AddDevice(values);
        values['status'] = '1'
        if(values['pwd'] == null || values['pwd'] == "") {
          values['pwd'] = md5(md5(values['user_id']).toString()).toString()
        }
        // console.info("values2 = " + JSON.stringify(values))
        // await PostAcctUpsertApi(values)
        PostAcctUpsertApi(values).then((resp) => {
          console.info("PostAcctUpsertApi = " + JSON.stringify(resp))
          if(resp["code"] != 0) {
            emit('error', resp["msg"])
          } else {
            emit('success');
            closeModal();
          }
        })
      // }

      // emit('success');
      // closeModal();
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
<style lang="less"></style>
