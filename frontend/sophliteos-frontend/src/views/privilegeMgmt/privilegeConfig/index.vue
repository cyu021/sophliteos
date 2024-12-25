<template>
  <div style="margin: 20px; display: flex; flex-direction: column">
    <div style="font-size: 20px; margin-bottom: 20px"> {{ t('dashboard.privilegeConfig') }} </div>
    <Space>
      <Button
        type="primary"
        style="margin-left: 16px"
        @click="handleCreateRole('add')"
      >
        {{ t('dataSource.privilegeCfg.addRole') }}
      </Button>
    </Space>
  </div>
  <Divider />
  <Table :columns="columns" :data-source="dataSource">
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'action'">
        <Space>
          <Button v-if="userRole == 'Admin' && record['role_name'] != 'Admin'" type="primary" @click="handleUpdateRole(record)">{{ t('dataSource.privilegeCfg.editRole') }}</Button>
          <Button v-if="userRole == 'Admin' && record['role_name'] != 'Admin'" type="danger" @click="handleDeleteRole(record)">{{ t('algorithm.delete') }}</Button>
        </Space>
      </template>
    </template>
  </Table>
  <editRole @register="RoleModal" @success="RoleSuccess" />
</template>

<script setup>
  import { ref, onMounted } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Table, Divider, Button, Space } from 'ant-design-vue';
  import { GetRoleListApi, PostAcctGetApi, PostRoleDeleteApi, GetRolePrivTplApi, PostRoleGetApi } from '/@/api/sys/user';
  import { useUserStore } from '/@/store/modules/user';
  import { useModal } from '/@/components/Modal';
  import editRole from './editRole.vue';
  import { rolePrivFormSchema, isAdd } from './tableData';
  
  import { useMessage } from '/@/hooks/web/useMessage';
  const { createMessage } = useMessage();
  const { t } = useI18n();
  
  const userStore = useUserStore();

  var userRole = "";
  
  const columns = ref([
    {
      title: t('dataSource.privilegeCfg.roleName'),
      dataIndex: 'role_name',
      key: 'role_name',
      align: 'center',
    },
    {
      title: t('algorithm.action'),
      dataIndex: 'action',
      key: 'action',
    },
  ]);
  
  const refresh = () => {
    GetRoleListApi({}).then((res) => {
      dataSource.value = res.data.roleList;
      console.info("RoleList dataSource.value = " + JSON.stringify(dataSource.value));
    });
    PostAcctGetApi({"user_id": userStore.$state.userInfo.username}).then((res) => {
      userRole = res.data.acct.role;
    });    
    if(rolePrivFormSchema.length == 1) {
    GetRolePrivTplApi({}).then((res) => {
      // console.info("GetRolePrivTplApi = " + JSON.stringify(res.data));
      var privTpl = res.data.priv_tpl;
      var menuTree = {};
      for(const k in privTpl) {
        if(privTpl[k].parent == "") {
          if(!(k in menuTree)) {
            menuTree[k] = {"hide": false, "kids": {}};
          }
        } else {
          if(!(privTpl[k].parent in menuTree)) {
            menuTree[privTpl[k].parent] = {"hide": false, "kids": {}};
          }
          menuTree[privTpl[k].parent]['kids'][k] = {"hide": false}
        }
      }
      // console.info("menuTree = " + JSON.stringify(menuTree));
      for(const k in menuTree) {
        var opt = {
          field: k,
          component: 'Checkbox',
          label: t(k),
          required: false,
          defaultValue: (menuTree[k]['hide']==false)? true: false,
          labelWidth: 250,
        }
        rolePrivFormSchema.push(opt)
        for(const k2 in menuTree[k]['kids']) {
          var opt2 = {
            field: k2,
            component: 'Checkbox',
            label: t(k) + " > " + t(k2),
            required: false,
            defaultValue: (menuTree[k]['kids'][k2]['hide']==false && menuTree[k]['hide']==false)? true: false,
            labelWidth: 250,
          }
          rolePrivFormSchema.push(opt2)
        }
      }
    });
    }
  }

  var dataSource = ref([]);

  onMounted(() => {
    refresh()
  });

  
  const [RoleModal, { openModal: OpenRoleModal }] = useModal();

  function handleCreateRole(record) {
    OpenRoleModal(true, { record });
  }

  function handleUpdateRole(record) {
    PostRoleGetApi(record).then((res) => {
      for(const k in res.data.sitemap_priv) {
        record[k] = (res.data.sitemap_priv[k]['hide'] == false) ? true: false;
      }
      OpenRoleModal(true, { record });
    })
  }

  function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
  
  function RoleSuccess() {
    createMessage.success('Upsert role success');
    refresh();
  }
  
  function handleDeleteRole(record) {
    var result = PostRoleDeleteApi(record).then((result) => {
      console.info("PostRoleDeleteApi result = " + JSON.stringify(result))
      createMessage.success('Delete role success');
      refresh();
    })
    // sleep(1000);
  }

</script>
