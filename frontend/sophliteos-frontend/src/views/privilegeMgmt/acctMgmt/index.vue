<template>
  <div style="margin: 20px; display: flex; flex-direction: column">
    <div style="font-size: 20px; margin-bottom: 20px"> {{ t('dashboard.acctManage') }} </div>
    <Space>
      <Button
        type="primary"
        style="margin-left: 16px"
        @click="handleCreateAcct('add')"
      >
        {{ t('dataSource.acctMgmt.addAcct') }}
      </Button>
    </Space>
  </div>
  <Divider />
  <Table :columns="columns" :data-source="dataSource">
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'status'">
        <Space>
          <p v-if="record.status === '1'">{{ t('algorithm.state_on') }}</p>
          <p v-else>{{ t('algorithm.state_off') }}</p>
        </Space>
      </template>
      <template v-else-if="column.key === 'action'">
        <Space>
          <Button v-if="userStore.$state.userInfo.username != record.user_id && record.user_id != 'admin'" type="primary" @click="handleUpdateAcctRole(record)">{{ t('dataSource.acctMgmt.editAcct') }}</Button>
          <Button v-if="userStore.$state.userInfo.username != record.user_id && record.user_id != 'admin'" type="danger" @click="handleDeleteAcct(record)">{{ t('algorithm.delete') }}</Button>
        </Space>
      </template>
    </template>
  </Table>
  <editAcct @register="AcctModal" @success="AcctSuccess" />
</template>

<script setup>
  import { ref, onMounted } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Table, Divider, Button, Space } from 'ant-design-vue';
  import { GetAcctListApi, PostAcctDeleteApi, PostAcctUpsert } from '/@/api/sys/user';
  import { useUserStore } from '/@/store/modules/user';
  import { useModal } from '/@/components/Modal';
  import editAcct from './editAcct.vue';
  import { acctSchema, isAdd } from './tableData';
  
  import { useMessage } from '/@/hooks/web/useMessage';
  const { createMessage } = useMessage();
  const { t } = useI18n();
  
  const userStore = useUserStore();
  
  const columns = ref([
    {
      title: t('dataSource.acctMgmt.acctType'),
      dataIndex: 'label',
      key: 'label',
    },
    {
      title: t('dataSource.acctMgmt.acctName'),
      dataIndex: 'user_id',
      key: 'user_id',
    },
    {
      title: t('dataSource.acctMgmt.acctRole'),
      dataIndex: 'role',
      key: 'role',
    },
    {
      title: t('dataSource.acctMgmt.acctStatus'),
      dataIndex: 'status',
      key: 'status',
    },
    {
      title: t('algorithm.action'),
      dataIndex: 'action',
      key: 'action',
    },
  ]);
  
  const refresh = () => {
    GetAcctListApi({}).then((res) => {
      dataSource.value = res.data.acctList;
      console.info("AcctList dataSource.value = " + JSON.stringify(dataSource.value));
      // console.info("userInfo = " + JSON.stringify(userStore.$state.userInfo));
      var roleList = []
      for(let i=0; i < dataSource.value.length; i++) {
        roleList.push(dataSource.value[i].role)
      }
      // console.info("roleList = " + JSON.stringify(roleList));
      const acctRoleSet = new Set(roleList)
      // console.info("acctRoleSet = " + JSON.stringify(acctRoleSet));
      var acctRoleOptions = [];
      acctRoleSet.forEach(function(role) {
        var roleOption = {}
        roleOption['label'] = role
        roleOption['value'] = role
        acctRoleOptions.push(roleOption)
      })
      // console.info("acctRoleOptions = " + JSON.stringify(acctRoleOptions));
      // console.info("options1 = " + JSON.stringify(acctSchema[2].componentProps.options));
      acctSchema[2].componentProps.options = acctRoleOptions
      // console.info("options2 = " + JSON.stringify(acctSchema[2].componentProps.options));
    });
  }

  var dataSource = ref([]);

  onMounted(() => {
    refresh()
  });

  const [AcctModal, { openModal: OpenAcctModal }] = useModal();

  function handleCreateAcct(record) {
    OpenAcctModal(true, { record });
  }
  
  function AcctSuccess() {
    createMessage.success('Create account success');
    refresh();
  }

  function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
  
  function handleDeleteAcct(record) {
    var result = PostAcctDeleteApi(record)
    console.info("PostAcctDeleteApi result = " + JSON.stringify(result))
    sleep(1000);
    createMessage.success('Delete account success');
    refresh();
  }

  function handleUpdateAcctRole(record) {
    OpenAcctModal(true, {record});
  }

</script>
