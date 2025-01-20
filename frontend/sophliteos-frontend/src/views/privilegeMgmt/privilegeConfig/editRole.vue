<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="Submit" >
    <div style="padding: 14px; display: flex; flex-direction: row; align-items: center;">
      <Input v-if="cacheExtend.FilterNameExists == ''" v-model:value="cacheExtend.FilterName" :disabled=false :placeholder="t('paramConfig.filter.name')" />
      <Input v-else v-model:value="cacheExtend.FilterNameExists" :disabled=true />
    </div>

    <div style="padding: 14px; padding-bottom: 30px;">
      <Collapse v-model:defaultActiveKey="optionsActiveKey">
        <CollapsePanel key="dashboard.basicInfor" :header="t('dashboard.basicInfor')" >
          <CheckboxGroup :value="cacheExtend.sitemap_opts['dashboard.basicInfor'].split(',')" :options="options['dashboard.basicInfor']" @change="updateSitemapPriv('dashboard.basicInfor', $event)" />
        </CollapsePanel>
        <CollapsePanel key="dashboard.accessAlgo" :header="t('dashboard.accessAlgo')" >
          <CheckboxGroup :value="cacheExtend.sitemap_opts['dashboard.accessAlgo'].split(',')" :options="options['dashboard.accessAlgo']" @change="updateSitemapPriv('dashboard.accessAlgo', $event)" />
        </CollapsePanel>
        <CollapsePanel key="dashboard.eventManage" :header="t('dashboard.eventManage')" >
          <CheckboxGroup :value="cacheExtend.sitemap_opts['dashboard.eventManage'].split(',')" :options="options['dashboard.eventManage']" @change="updateSitemapPriv('dashboard.eventManage', $event)" />
        </CollapsePanel>
        <CollapsePanel key="dashboard.algoModel" :header="t('dashboard.algoModel')" >
          <CheckboxGroup :value="cacheExtend.sitemap_opts['dashboard.algoModel'].split(',')" :options="options['dashboard.algoModel']" @change="updateSitemapPriv('dashboard.algoModel', $event)" />
        </CollapsePanel>
        <CollapsePanel key="dashboard.maintenance" :header="t('dashboard.maintenance')" >
          <CheckboxGroup :value="cacheExtend.sitemap_opts['dashboard.maintenance'].split(',')" :options="options['dashboard.maintenance']" @change="updateSitemapPriv('dashboard.maintenance', $event)" />
        </CollapsePanel>
        <CollapsePanel key="dashboard.logs" :header="t('dashboard.logs')" >
          <CheckboxGroup :value="cacheExtend.sitemap_opts['dashboard.logs'].split(',')" :options="options['dashboard.logs']" @change="updateSitemapPriv('dashboard.logs', $event)" />
        </CollapsePanel>
        <CollapsePanel key="dashboard.privilegeMgmt" :header="t('dashboard.privilegeMgmt')" >
          <CheckboxGroup :value="cacheExtend.sitemap_opts['dashboard.privilegeMgmt'].split(',')" :options="options['dashboard.privilegeMgmt']" @change="updateSitemapPriv('dashboard.privilegeMgmt', $event)" />
        </CollapsePanel>
      </Collapse>
    </div>
  </BasicModal>
</template>

<script lang="ts" setup>
  import { Divider, Input, Collapse, CollapsePanel, Dropdown, Button, Menu, MenuItem, CheckboxGroup, TimePicker, DatePicker, } from 'ant-design-vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { onMounted, ref, watch, computed } from 'vue';
  import apis from './api';
  import { PostRoleGetApi, PostRoleUpsertApi } from '/@/api/sys/user';
  import { DeleteOutlined, DownOutlined } from '@ant-design/icons-vue';
  import dayjs from 'dayjs';

  const { t } = useI18n();
  const title = ref();
  const emit = defineEmits(['success', 'register', 'error']);

  var optionsActiveKey = ref([
      'dashboard.accessAlgo',
      'dashboard.privilegeMgmt',
      'dashboard.eventManage',
      'dashboard.algoModel',
      'dashboard.basicInfor',
      'dashboard.logs',
      'dashboard.maintenance'
    ]);

  const options = {
    "dashboard.accessAlgo": [
      {label: t('dashboard.AlgoParamConfig'), value: '0'},
      {label: t('dashboard.task'), value: '1'},
      {label: t('dashboard.videoManage'), value: '2'}
    ],
    "dashboard.privilegeMgmt": [
      {label: t('dashboard.acctManage'), value: '3'},
      {label: t('dashboard.privilegeConfig'), value: '4'}
    ],
    "dashboard.eventManage": [
      {label: t('dashboard.alarmRetrieval'), value: '5'},
      {label: t('dashboard.eventList'), value: '6'}
    ],
    "dashboard.algoModel": [
      {label: t('dashboard.algorithm'), value: '7'},
      {label: t('dashboard.prepSpringZip'), value: '8'}
    ],
    "dashboard.basicInfor": [
      {label: t('dashboard.ovewview'), value: '9'}
    ],
    "dashboard.logs": [
      {label: t('dashboard.logDownload'), value: '10'},
      {label: t('dashboard.operate'), value: '11'},
      {label: t('dashboard.warning'), value: '12'}
    ],
    "dashboard.maintenance": [
      {label: t('dashboard.restart'), value: '13'},
      {label: t('dashboard.ruleTemplate'), value: '14'},
      {label: t('dashboard.softUpdate'), value: '15'},
      {label: t('dashboard.sysSetting'), value: '16'},
      {label: t('dashboard.sysSoft'), value: '17'}
    ]
  };

  const optionsKeyMap = {
    'dashboard.AlgoParamConfig': '0',
    'dashboard.task': '1',
    'dashboard.videoManage': '2',
    'dashboard.acctManage': '3',
    'dashboard.privilegeConfig': '4',
    'dashboard.alarmRetrieval': '5',
    'dashboard.eventList': '6',
    'dashboard.algorithm': '7',
    'dashboard.prepSpringZip': '8',
    'dashboard.ovewview': '9',
    'dashboard.logDownload': '10',
    'dashboard.operate': '11',
    'dashboard.warning': '12',
    'dashboard.restart': '13',
    'dashboard.ruleTemplate': '14',
    'dashboard.softUpdate': '15',
    'dashboard.sysSetting': '16',
    'dashboard.sysSoft': '17'
  };

  const optionsValMap = {
    '0': 'dashboard.AlgoParamConfig',
    '1': 'dashboard.task',
    '2': 'dashboard.videoManage',
    '3': 'dashboard.acctManage',
    '4': 'dashboard.privilegeConfig',
    '5': 'dashboard.alarmRetrieval',
    '6': 'dashboard.eventList',
    '7': 'dashboard.algorithm',
    '8': 'dashboard.prepSpringZip',
    '9': 'dashboard.ovewview',
    '10': 'dashboard.logDownload',
    '11': 'dashboard.operate',
    '12': 'dashboard.warning',
    '13': 'dashboard.restart',
    '14': 'dashboard.ruleTemplate',
    '15': 'dashboard.softUpdate',
    '16': 'dashboard.sysSetting',
    '17': 'dashboard.sysSoft'
  };

  const reqVals = {};

  const [registerModal, { setModalProps, closeModal, getVisible }] = useModalInner(async (data) => {
    setModalProps({ confirmLoading: false });
    console.info("editRole data=" + JSON.stringify(data));
    
    cacheExtend.value.sitemap_opts = {
      "dashboard.accessAlgo": "",
      "dashboard.privilegeMgmt": "",
      "dashboard.eventManage": "",
      "dashboard.algoModel": "",
      "dashboard.basicInfor": "",
      "dashboard.logs": "",
      "dashboard.maintenance": ""
    };
    cacheExtend.value.FilterName = "";
    cacheExtend.value.FilterNameExists = "";
    
    title.value = t('dataSource.privilegeCfg.addRole');
    reqVals["role_name"] = "Admin";
    if(data["record"] != "add") {
      cacheExtend.value.FilterNameExists = data["record"]["role_name"]
      title.value = t('dataSource.privilegeCfg.editRole');
      reqVals["role_name"] = cacheExtend.value.FilterNameExists;
    }
    PostRoleGetApi(reqVals).then((resp) => {
      // console.info("PostRoleGetApi = " + JSON.stringify(resp))
      if(resp["code"] != 0) {
        emit('error', resp["msg"])
      } else {
        for(const k in resp.data.sitemap_priv) {
          var v = resp.data.sitemap_priv[k];
          if(v["parent"] != "" ) {
            if(v["hide"] == false) {
              cacheExtend.value.sitemap_opts[v["parent"]] += optionsKeyMap[k]+","
            }
          }
        }
      }
      // console.info("registerModal:editRole cacheExtend=" + JSON.stringify(cacheExtend));
    });
  });

  function updateSitemapPriv(k, value) {
    cacheExtend.value.sitemap_opts[k] = value.join(',');
  }

  async function Submit() {
    console.info("submit cacheExtend=" + JSON.stringify(cacheExtend));
    var roleName = cacheExtend.value.FilterNameExists;
    if(roleName == "") {
      roleName = cacheExtend.value.FilterName;
    }
    if(roleName == "") {
      emit('error', 'roleName must not be blank');
      return;
    }

    var privCfg = {};

    reqVals["role_name"] = "Admin";
    PostRoleGetApi(reqVals).then((resp) => {
      console.info("editRole:Submit:PostRoleGetApi = " + JSON.stringify(resp))
      if(resp["code"] != 0) {
        emit('error', resp["msg"])
        return;
      } else {
        // for(const k in resp.data.sitemap_priv) {
        //   var v = resp.data.sitemap_priv[k];
        //   if(v["parent"] != "" ) {
        //     if(v["hide"] == false) {
        //       cacheExtend.value.sitemap_opts[v["parent"]] += optionsKeyMap[k]+","
        //     }
        //   }
        // }
        privCfg = resp.data.sitemap_priv;
        console.info("privCfg =" + JSON.stringify(privCfg));
        for(const k in privCfg) {
          privCfg[k]['hide'] = true;
        }
        console.info("privCfg2 =" + JSON.stringify(privCfg));
        for(const k in cacheExtend.value.sitemap_opts) {
          console.info("checking " + k + " => " + cacheExtend.value.sitemap_opts[k])
          if(cacheExtend.value.sitemap_opts[k] != "") {
            console.info("show " + k);
            privCfg[k]['hide'] = false;
            let optVals = cacheExtend.value.sitemap_opts[k].split(',');
            for(const v in optVals) {
              let idx = optVals[v];
              let privName = optionsValMap[idx];
              if(privName == undefined) {
                continue;
              }
              console.info("show " + idx + "/" + privName);
              privCfg[privName]['hide'] = false;
            }
          }
        }
        console.info("registerModal:editRole:PostRoleUpsertApi cacheExtend=" + JSON.stringify(cacheExtend));
        console.info("registerModal:editRole:PostRoleUpsertApi privCfg=" + JSON.stringify(privCfg));
        PostRoleUpsertApi({"role_name": roleName, "sitemap_priv": privCfg}).then((resp) => {
          if(resp["code"] != 0) {
            emit('error', resp["msg"])
          } else {
            // emit('success');
            closeModal();
          }
          console.info("registerModal:editRole:PostRoleUpsertApi = " + JSON.stringify(resp));
        })
      }
    });

    try {
      emit('success');
    } finally {
      closeModal();
    }
  }

  onMounted(() => {
  })

  const props = defineProps({
    taskName: {
      type: String,
      default: null,
    },
    algorithmName: {
      type: String,
      default: null,
    },
    extend: {
      type: Object,
      default: {},
    }
  });

  const cacheExtend = ref({});
  cacheExtend.value.sitemap_opts = {
    "dashboard.accessAlgo": "",
    "dashboard.privilegeMgmt": "",
    "dashboard.eventManage": "",
    "dashboard.algoModel": "",
    "dashboard.basicInfor": "",
    "dashboard.logs": "",
    "dashboard.maintenance": ""
  };
  cacheExtend.value.FilterName = "";
  cacheExtend.value.FilterNameExists = "";

</script>
