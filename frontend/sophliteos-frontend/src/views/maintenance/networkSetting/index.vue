<template>
  <a-tabs v-model:activeKey="activeKey" class="!m-4 !p-4 bg-white" animated>
    <a-tab-pane key="wan" :tab="t('maintenance.newworkSettings.wan')">
      <a-skeleton :loading="pageLoading" active>
        <a-form
          :model="wan"
          v-bind="formItemLayout"
          size="large"
          class="w-1/2 !mx-auto"
          :rules="wanRules"
          @finish="submitForm"
          v-show="!pageLoading"
        >
          <a-form-item
            v-for="item of formItemList"
            :key="item.field"
            :label="item.label"
            :name="item.field"
            :labelCol="{ style: 'width: 150px' }"
          >
            <a-select
              v-if="item.type === 'select'"
              ref="select"
              v-model:value="wan[item.field]"
              :options="item.options"
              @change="item.onChange"
            />
            <a-input
              v-if="item.type === 'input'"
              v-model:value="wan[item.field]"
              :placeholder="item.placeholder"
              :disabled="wan.ipType === 2"
            />
          </a-form-item>
          <a-form-item class="!pl-1/6">
            <a-button type="primary" html-type="submit" :loading="loading">{{
              t('sys.btn.confirm')
            }}</a-button>
          </a-form-item>
        </a-form>
      </a-skeleton>
    </a-tab-pane>
    <!-- <a-tab-pane key="lan1" :tab="t('maintenance.newworkSettings.lan')">
      <a-skeleton :loading="pageLoading" active>
        <a-form
          :model="lan1"
          v-bind="formItemLayout"
          size="large"
          class="w-1/2 !mx-auto"
          :rules="lanRules"
          @finish="submitForm"
        >
          <a-form-item
            v-for="item of formItemListLan"
            :key="item.field"
            :label="item.label"
            :name="item.field"
          >
            <a-select
              v-if="item.type === 'select'"
              ref="select"
              v-model:value="lan1[item.field]"
              :options="options"
              @change="handleChange"
              :disabled="item.field === 'ipType'"
            />
            <a-input
              v-if="item.type === 'input'"
              v-model:value="lan1[item.field]"
              :placeholder="item.placeholder"
            />
          </a-form-item>
          <a-form-item class="!pl-1/6">
            <a-button type="primary" html-type="submit" :loading="loading">{{
              t('sys.btn.confirm')
            }}</a-button>
          </a-form-item>
        </a-form>
      </a-skeleton>
    </a-tab-pane> -->
    <!-- <a-tab-pane key="core" :tab="t('maintenance.newworkSettings.core')" disabled>
      {{ t('maintenance.newworkSettings.core') }}
    </a-tab-pane> -->
    <!-- <a-tab-pane key="algoConfig" :tab="t('taskList.taskList.algoConfigTitle')">
      <a-skeleton :loading="pageLoading" active>
        <a-form
          :model="algoConfig"
          v-bind="formItemLayout"
          size="large"
          class="w-1/2 !mx-auto"
          :rules="algoConfigRules"
          @finish="submitFormAlgoConfig"
          v-show="!pageLoading"
        >
          <a-form-item
            v-for="item of formItemListAlgoConfig"
            :key="item.field"
            :label="item.label"
            :name="item.field"
          >
            <a-input
              v-if="item.type === 'input'"
              v-model:value="algoConfig[item.field]"
              :placeholder="item.placeholder"
            />
          </a-form-item>
          <a-form-item class="!pl-1/6">
            <a-button type="primary" html-type="submit" :loading="loading">{{
              t('sys.btn.confirm')
            }}</a-button>
          </a-form-item>
        </a-form>
      </a-skeleton>
    </a-tab-pane> -->
    <a-tab-pane key="upUrlConfig" :tab="t('taskList.taskList.upUrlConfig')">
      <a-skeleton :loading="pageLoading" active>
        <a-form
          :model="upUrlConfig"
          v-bind="formItemLayout"
          size="large"
          class="w-1/2 !mx-auto"
          :rules="upUrlConfigRules"
          @finish="submitFormUpUrlConfig"
          v-show="!pageLoading"
        >
          <a-form-item
            v-for="item of formItemListUpUrlConfig"
            :key="item.field"
            :label="item.label"
            :name="item.field"
            :label-col="{ style: {width: '200px'} }"
          >
            <a-input
              v-if="item.type === 'input'"
              v-model:value="upUrlConfig[item.field]"
              :placeholder="item.placeholder"
              :suffix="item.suffix"
            />
            <a-select
              v-if="item.type === 'select'"
              ref="select"
              v-model:value="upUrlConfig[item.field]"
              :options="item.options"
              @change="item.onChange"
            >
            </a-select>
            <template v-if="item.type === 'divider'">
              <label>{{item.field}}</label>
              <hr style="height: 3px; width: 100%; background-color: #bbb" />
            </template>
          </a-form-item>
          <a-form-item class="!pl-1/6">
            <a-button type="primary" html-type="submit" :loading="loading">{{
              t('sys.btn.confirm')
            }}</a-button>
          </a-form-item>
        </a-form>
      </a-skeleton>
    </a-tab-pane>
    <a-tab-pane key="rotateConfig" :tab="t('maintenance.rotateConfig.rotateConfig')">
      <a-skeleton :loading="pageLoading" active>
        <a-form
          :model="rotateCfgConfig"
          v-bind="formItemLayout"
          size="large"
          class="w-1/2 !mx-auto"
          :rules="rotateCfgConfigRules"
          @finish="submitFormRotateCfgConfig"
          v-show="!pageLoading"
        >
          <a-form-item
            v-for="item of formItemListRotateCfgConfig"
            :key="item.field"
            :label="item.label"
            :name="item.field"
            :labelCol="{ style: 'width: 150px' }"
          >
            <a-input
              v-if="item.type === 'input'"
              v-model:value="rotateCfgConfig[item.field]"
              :placeholder="item.placeholder"
              :suffix="item.suffix"
            />
          </a-form-item>
          <a-form-item class="!pl-1/6">
            <a-button type="primary" html-type="submit" :loading="loading">{{
              t('sys.btn.confirm')
            }}</a-button>
          </a-form-item>
        </a-form>
      </a-skeleton>
    </a-tab-pane>
    <a-tab-pane key="resourceAlarm" :tab="t('maintenance.threshold.title')">
      <a-skeleton :loading="pageLoading" active>
        <a-form
          :model="formState"
          v-bind="formItemLayout"
          class="flex flex-wrap justify-around"
          size="large"
        >
          <a-form-item
            v-for="item of formItemListAlarm"
            :key="item.field"
            :label="item.label"
            class="w-2/5"
            :labelCol="{ style: 'width: 150px' }"
          >
            <a-input-number
              v-model:value="formState[item.field]"
              :placeholder="placeholder"
              :max="100"
            >
              <template #addonAfter>
                <span v-if="item.unit" :style="{ color }">{{ item.unit }}</span>
                <percentage-outlined v-else :style="{ color }" />
              </template>
            </a-input-number>
          </a-form-item>
          <a-form-item class="w-2/5" />
          <a-form-item :wrapper-col="{ offset: 8, span: 16 }" class="w-2/5 !mr-1/2">
            <a-button type="primary" @click="submitFormAlarm" :loading="loading">{{
              t('sys.btn.confirm')
            }}</a-button>
          </a-form-item>
        </a-form>
      </a-skeleton>
    </a-tab-pane>
    <a-tab-pane key="timedateConfig" :tab="t('maintenance.timedateConfig.timedateConfig')">
      <a-skeleton :loading="pageLoading" active>
        <a-form
          :model="formTimedate"
          v-bind="formItemLayout"
          size="large"
          class="w-1/2 !mx-auto"
          @finish="submitFormTimedate"
          v-show="!pageLoading"
        >
          <a-form-item
            v-for="item of formTimedateConfig"
            :key="item.field"
            :label="item.label"
            :name="item.field"
            :labelCol="{ style: 'width: 150px' }"
          >
            <a-input
              v-if="item.type === 'input'"
              v-model:value="formTimedate[item.field]"
              :placeholder="item.placeholder"
            />
            <a-input
              v-if="item.type === 'text'"
              v-model:value="formTimedate[item.field]"
              :placeholder="item.placeholder"
              :disabled=true
            />
            <a-select 
              v-if="item.type === 'select'"
              mode="combobox" 
              show-search
              v-model:value="formTimedate['timezone']"
              @change="handleTZChange">
              <a-select-option 
                v-for="i in formTimedate['timezones']" 
                :key="i">
                {{i}}
              </a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item class="!pl-1/6">
            <a-button type="primary" html-type="submit" :loading="loading">{{
              t('sys.btn.confirm')
            }}</a-button>
          </a-form-item>
        </a-form>
      </a-skeleton>
    </a-tab-pane>
  </a-tabs>
  <BasicModal
    v-bind="$attrs"
    @register="register"
    :title="t('maintenance.restart.taskRestart')"
    :canFullscreen="false"
    @ok="handleRestartTasksOk"
    :closable="false"
    :okText="t('common.okText')"
    :ok-button-props="restartLoading"
    :cancel-button-props="restartCancel"
  >
    <div style="padding: 14px; white-space: pre-wrap">
      {{ restartTips }}
    </div>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { reactive, ref, onMounted, computed, watch, h } from 'vue';
  import type { UnwrapRef } from 'vue';
  import { ipGet, ipSet, getAlarm, setAlarm, getTimedate, setTimedate } from '/@/api/maintenance/index';
  import { Tabs, Modal } from 'ant-design-vue';
  import { ExclamationCircleOutlined, CopyOutlined } from '@ant-design/icons-vue';

  import { useI18n } from '/@/hooks/web/useI18n';
  import { copyTextToClipboard } from '/@/hooks/web/useCopyToClipboard';
  import { useMessage } from '/@/hooks/web/useMessage';

  import { IpSetParams, AlarmParams, TimedateParams } from '/@/api/maintenance/model/index';
  // import { number } from '@intlify/core-base';
  import { IpCheck, subnetMaskCheck, gatewayCheck, dnsCheck } from '/@/utils/validateFuncs';
  import { useDeviceInfo } from '/@/store/modules/overview';

  import { BasicModal, useModal } from '/@/components/Modal';

  const [register, { openModal, closeModal }] = useModal();

  const deviceStore = useDeviceInfo();
  const { createMessage } = useMessage();

  const { t } = useI18n();
  const ATabs = Tabs;
  const ATabPane = Tabs.TabPane;

  const activeKey = ref('wan');
  const wan: UnwrapRef<IpSetParams> = reactive({
    device: '',
    ipType: 1,
    ip: '',
    subnetMask: '',
    gateway: '',
    dns: '',
  });

  watch(
    () => wan.device,
    (value) => {
      const currentNetCard: any = ipData.wan.find((item: any) => item.netCardName === value);
      if (currentNetCard) {
        wan.ipType = currentNetCard?.dynamic + 1;
        wan.ip = currentNetCard?.ip || '';
        wan.subnetMask = currentNetCard?.netMask || '';
        wan.gateway = currentNetCard?.gateway || '';
        wan.dns = currentNetCard?.dns || '';
      }
    },
  );
  const wanRules = computed(() => {
    return wan.ipType === 1
      ? {
          ip: [
            {
              required: true,
              validator: IpCheck,
              trigger: 'blur',
            },
          ],
          subnetMask: [
            {
              required: true,
              validator: subnetMaskCheck,
              trigger: 'blur',
            },
          ],
          gateway: [
            {
              required: true,
              validator: gatewayCheck,
              trigger: 'blur',
            },
          ],
          dns: [
            {
              required: true,
              validator: dnsCheck,
              trigger: 'blur',
            },
          ],
        }
      : null;
  });

  const netMap = {
    wan,
    // lan1,
  };

  const formItemList = [
    {
      label: t('maintenance.newworkSettings.netCard'),
      field: 'device',
      placeholder: t('sys.form.placeholder'),
      type: 'select',
      options: [],
      onChange() {},
    },
    {
      label: t('maintenance.newworkSettings.ipType'),
      field: 'ipType',
      placeholder: t('sys.form.placeholder'),
      type: 'select',
      options: [
        {
          value: 1,
          label: t('maintenance.newworkSettings.staticIP'),
        },
        {
          value: 2,
          label: t('maintenance.newworkSettings.dynmicIP'),
        },
      ],
      onChange() {},
    },
    {
      label: t('maintenance.newworkSettings.ip'),
      field: 'ip',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
    {
      label: t('maintenance.newworkSettings.subnetMask'),
      field: 'subnetMask',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
    {
      label: t('maintenance.newworkSettings.gateway'),
      field: 'gateway',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
    {
      label: t('maintenance.newworkSettings.dns'),
      field: 'dns',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
  ];

  const formItemLayout = {
    labelCol: { span: 4 },
    wrapperCol: { span: 20 },
  };

  // 查询到的IP数据存储
  const ipData = reactive({
    wan: [],
  });
  const pageLoading = ref(true);
  const init = async () => {
    const result = await ipGet();
    pageLoading.value = false;
    if (result) {
      ipData.wan = result;
      if (!deviceStore.isSingleBoard) {
        ipData.wan = result.filter((item) => item.netCardName.startsWith('enp'));
      }
      setInitValue();
    }
  };
  const setInitValue = () => {
    formItemList[0].options = ipData.wan.map((item: any) => ({
      value: item.netCardName,
      label: item.netCardName,
    }));
    wan.device = formItemList[0].options[0]?.value as any;
  };
  // 复制地址
  const copyAddress = () => {
    const isSuccess = copyTextToClipboard(`${wan.ip}:8080`);
    isSuccess && createMessage.success(t('sys.copySuccess'));
  };
  const loading = ref(false);
  const submitForm = () => {
    const content =
      wan.ipType === 1
        ? h('div', { style: { 'text-align': 'center', color: '#000', margin: '30px 0 0 -30px' } }, [
            h('p', {}, t('maintenance.newworkSettings.ipRightOrNotNetword')),
            h('p', {}, t('maintenance.newworkSettings.restartAfterSubmit')),
            h('p', {}, [
              h('span', {}, `${t('maintenance.newworkSettings.useAddress') + wan.ip}:8080 `),
              h(CopyOutlined, {
                title: t('maintenance.newworkSettings.copyAddress'),
                style: { color: '#0960BC', cursor: 'pointer' },
                onClick: copyAddress,
              }),
              h('span', {}, t('maintenance.newworkSettings.reOpenPage')),
            ]),
          ])
        : h(
            'div',
            { style: { 'text-align': 'center', color: '#000', margin: '30px 0 0 -30px' } },
            t('maintenance.newworkSettings.dynIpDialogContent'),
          );
    Modal.confirm({
      title: t('sys.tips'),
      icon: h(ExclamationCircleOutlined),
      width: 600,
      content,
      onOk() {
        try {
          loading.value = true;
          const params = {
            ...netMap[activeKey.value],
          };
          ipSet(params)
            .then((res) => {
              createMessage.success(res.msg);
            })
            .catch(() => {
              createMessage.error('不支持修改ip');
            });

          // console.log(params);
        } catch (error) {
        } finally {
          loading.value = false;
        }
      },
      onCancel() {},
    });
  };

  onMounted(() => {
    if (!deviceStore.deviceType) {
      deviceStore.getDeviceInfo().then(() => {
        init();
      });
    } else {
      init();
    }
    // initAlgoConfig();
    initUpUrlConfig();
    initRotateCfgConfig();
    
    initResourceAlarmgConfig();
    initTimedateConfig();
  });

  ////////////////////////////////////////////////

  import { AlgoConfigSetParams } from '/@/api/maintenance/model/index';
  import { portCheck } from '/@/utils/validateFuncs';
  import { getAlgorithm, addAlgorithm, getTaskList, StartTask, StopTask } from '/@/api/task/index';

  const algoConfig: UnwrapRef<AlgoConfigSetParams> = reactive({
    ip: '',
    port: '',
  });

  const algoConfigRules = computed(() => {
    return {
      ip: [
        {
          required: true,
          validator: IpCheck,
          trigger: 'blur',
        },
      ],
      port: [
        {
          required: true,
          validator: portCheck,
          trigger: 'blur',
        },
      ],
    };
  });

  const algoConfigMap = {
    algoConfig,
  };

  const formItemListAlgoConfig = [
    {
      label: t('maintenance.newworkSettings.ip'),
      field: 'ip',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
    {
      label: t('maintenance.newworkSettings.serverPort'),
      field: 'port',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
  ];

  const ipDataAlgoConfig = reactive({
    algoConfig: [],
  });
  const initAlgoConfig = async () => {
    const result = await getAlgorithm();
    pageLoading.value = false;
    console.info('initAlgoConfig result=' + JSON.stringify(result));
    if (result) {
      ipDataAlgoConfig.algoConfig = result;
      algoConfig.ip = result.ip;
      algoConfig.port = result.port.toString();
    }
  };

  const submitFormAlgoConfig = () => {
    loading.value = true;
    var params = {};
    params['ip'] = algoConfigMap['algoConfig'].ip;
    params['port'] = Number(algoConfigMap['algoConfig'].port);
    console.info('addAlgorithm params = ' + JSON.stringify(params));
    const result = addAlgorithm(params);
    console.info('addAlgorithm result = ' + JSON.stringify(result));
    if (JSON.stringify(result) === '{}') {
      createMessage.success('update success');
    } else {
      createMessage.error('invalid operation');
    }
    loading.value = false;
    // const content =
    //   wan.ipType === 1
    //     ? h('div', { style: { 'text-align': 'center', color: '#000', margin: '30px 0 0 -30px' } }, [
    //         h('p', {}, t('maintenance.newworkSettings.ipRightOrNotNetword')),
    //         h('p', {}, t('maintenance.newworkSettings.restartAfterSubmit')),
    //         h('p', {}, [
    //           h('span', {}, `${t('maintenance.newworkSettings.useAddress') + wan.ip}:8080 `),
    //           h(CopyOutlined, {
    //             title: t('maintenance.newworkSettings.copyAddress'),
    //             style: { color: '#0960BC', cursor: 'pointer' },
    //             onClick: copyAddress,
    //           }),
    //           h('span', {}, t('maintenance.newworkSettings.reOpenPage')),
    //         ]),
    //       ])
    //     : h(
    //         'div',
    //         { style: { 'text-align': 'center', color: '#000', margin: '30px 0 0 -30px' } },
    //         t('maintenance.newworkSettings.dynIpDialogContent'),
    //       );
    // Modal.confirm({
    //   title: t('sys.tips'),
    //   icon: h(ExclamationCircleOutlined),
    //   width: 600,
    //   content,
    //   onOk() {
    //     try {
    //       loading.value = true;
    //       const params = {
    //         ...netMap[activeKey.value],
    //       };
    //       ipSet(params)
    //         .then((res) => {
    //           createMessage.success(res.msg);
    //         })
    //         .catch(() => {
    //           createMessage.error('不支持修改ip');
    //         });

    //       // console.log(params);
    //     } catch (error) {
    //     } finally {
    //       loading.value = false;
    //     }
    //   },
    //   onCancel() {},
    // });
  };

  // onMounted(() => {
  //   initAlgoConfig();
  // });

  /////////////////////////////////////////////////////

  import { UpUrlConfigSetParams } from '/@/api/maintenance/model/index';
  import { getUpUrl, addUpUrl } from '/@/api/task/index';

  import { upUrlParams } from '/@/api/task/model/index';

  import { protocolCheck } from '/@/utils/validateFuncs';

  const upUrlConfig: UnwrapRef<UpUrlConfigSetParams> = reactive({
    ip: '',
    port: '',
    protocol: '',
    endpoint: '',
    ip2: '',
    port2: '',
    protocol2: '',
    endpoint2: '',
    retryRows: 1000,
    enableRetry: false,
  });

  const upUrlConfigRules = computed(() => {
    return {
      protocol: [
        {
          required: true,
          trigger: 'blur',
        },
      ],
      ip: [
        {
          required: true,
          validator: IpCheck,
          trigger: 'blur',
        },
      ],
      port: [
        {
          required: true,
          validator: portCheck,
          trigger: 'blur',
        },
      ],
      endpoint: [
        {
          required: true,
          trigger: 'blur',
        },
      ],
      protocol2: [
        {
          required: false,
          trigger: 'blur',
        },
      ],
      ip2: [
        {
          required: false,
          validator: IpCheck,
          trigger: 'blur',
        },
      ],
      port2: [
        {
          required: false,
          validator: portCheck,
          trigger: 'blur',
        },
      ],
      endpoint2: [
        {
          required: false,
          trigger: 'blur',
        },
      ],
      retryRows: [
        {
          required: true,
          validator: rotateRetryRowsCheck,
          trigger: 'blur',
        },
      ],
    };
  });

  const upUrlConfigMap = {
    upUrlConfig,
  };

  const formItemListUpUrlConfig = [
    {
      field: t('maintenance.newworkSettings.pushEndpointMain'),
      type: 'divider',
    },
    {
      label: t('maintenance.newworkSettings.protocol'),
      field: 'protocol',
      placeholder: t('sys.form.placeholder'),
      type: 'select',
      options: [
        {
          value: "http",
          label: "http",
        },
        {
          value: "https",
          label: "https",
        },
      ],
      onChange() {},
    },
    {
      label: t('maintenance.newworkSettings.ip'),
      field: 'ip',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
    {
      label: t('maintenance.newworkSettings.serverPort'),
      field: 'port',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
    {
      label: t('maintenance.newworkSettings.serverEndpoint'),
      field: 'endpoint',
      placeholder: '/algorithm/upload',
      type: 'input',
    },
    {
      field: t('maintenance.newworkSettings.pushEndpointBackup'),
      type: 'divider',
    },
    {
      label: t('maintenance.newworkSettings.protocol'),
      field: 'protocol2',
      placeholder: t('sys.form.placeholder'),
      type: 'select',
      options: [
        {
          value: "http",
          label: "http",
        },
        {
          value: "https",
          label: "https",
        },
      ],
      onChange() {},
    },
    {
      label: t('maintenance.newworkSettings.ip'),
      field: 'ip2',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
    {
      label: t('maintenance.newworkSettings.serverPort'),
      field: 'port2',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
    {
      label: t('maintenance.newworkSettings.serverEndpoint'),
      field: 'endpoint2',
      placeholder: '/algorithm/upload',
      type: 'input',
    },
    {
      field: t('maintenance.newworkSettings.pushRetry'),
      type: 'divider',
    },
    {
      label: t('maintenance.newworkSettings.enableRetry'),
      field: 'enableRetry',
      placeholder: t('sys.form.placeholder'),
      type: 'select',
      options: [
        {
          value: false,
          label: "Disable",
        },
        {
          value: true,
          label: "Enable",
        },
      ],
      onChange() {},
    },
    {
      label: t('maintenance.newworkSettings.retryRows'),
      field: 'retryRows',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
      suffix: 'rows',
    },
  ];

  const ipDataUpUrlConfig = reactive({
    upUrlConfig: [],
  });
  const initUpUrlConfig = async () => {
    let result = await getUpUrl();
    result = result ? result : '{"uploadhost": "http://127.0.0.1:8081", "enableRetry": false, "retryRows": 1000}';
    pageLoading.value = false;
    let resultJson = JSON.parse(result)
    console.info('initUpUrlConfig result=' + JSON.stringify(resultJson));
    if (resultJson["Uploadhost"]) {
      let uploadhosts = resultJson["Uploadhost"].split(";")
      for(let i = 0; i < uploadhosts.length; i++) {
        if(i == 0) {
          let resultEles = uploadhosts[i].split('://');
          let resultEles2 = resultEles[1].split(':');
          ipDataUpUrlConfig.upUrlConfig['protocol'] = resultEles[0];
          ipDataUpUrlConfig.upUrlConfig['ip'] = resultEles2[0];
          let resultEles3 = resultEles2[1].split('/');
          ipDataUpUrlConfig.upUrlConfig['port'] = resultEles3[0];
          ipDataUpUrlConfig.upUrlConfig['endpoint'] = '/algorithm/upload';
          if (resultEles3.length > 2) {
            ipDataUpUrlConfig.upUrlConfig['endpoint'] = resultEles3
              .slice(1, resultEles3.length)
              .join('/');
          }
          if (!ipDataUpUrlConfig.upUrlConfig['endpoint'].startsWith('/')) {
            ipDataUpUrlConfig.upUrlConfig['endpoint'] = '/' + ipDataUpUrlConfig.upUrlConfig['endpoint'];
          }
          upUrlConfig.ip = ipDataUpUrlConfig.upUrlConfig['ip'];
          upUrlConfig.port = ipDataUpUrlConfig.upUrlConfig['port'];
          upUrlConfig.protocol = ipDataUpUrlConfig.upUrlConfig['protocol'];
          upUrlConfig.endpoint = ipDataUpUrlConfig.upUrlConfig['endpoint'];
        } else if(i == 1) {
          let resultEles = uploadhosts[i].split('://');
          let resultEles2 = resultEles[1].split(':');
          ipDataUpUrlConfig.upUrlConfig['protocol2'] = resultEles[0];
          ipDataUpUrlConfig.upUrlConfig['ip2'] = resultEles2[0];
          let resultEles3 = resultEles2[1].split('/');
          ipDataUpUrlConfig.upUrlConfig['port2'] = resultEles3[0];
          ipDataUpUrlConfig.upUrlConfig['endpoint2'] = '/algorithm/upload';
          if (resultEles3.length > 2) {
            ipDataUpUrlConfig.upUrlConfig['endpoint2'] = resultEles3
              .slice(1, resultEles3.length)
              .join('/');
          }
          if (!ipDataUpUrlConfig.upUrlConfig['endpoint2'].startsWith('/')) {
            ipDataUpUrlConfig.upUrlConfig['endpoint2'] = '/' + ipDataUpUrlConfig.upUrlConfig['endpoint2'];
          }
          upUrlConfig.ip2 = ipDataUpUrlConfig.upUrlConfig['ip2'];
          upUrlConfig.port2 = ipDataUpUrlConfig.upUrlConfig['port2'];
          upUrlConfig.protocol2 = ipDataUpUrlConfig.upUrlConfig['protocol2'];
          upUrlConfig.endpoint2 = ipDataUpUrlConfig.upUrlConfig['endpoint2'];
        }
      }
    }
    ipDataUpUrlConfig.upUrlConfig["enableRetry"] = resultJson["EnableRetry"]
    upUrlConfig.enableRetry = ipDataUpUrlConfig.upUrlConfig["enableRetry"]
    ipDataUpUrlConfig.upUrlConfig["retryRows"] = resultJson["RetryRows"]
    if(ipDataUpUrlConfig.upUrlConfig["retryRows"] == 0) {
      ipDataUpUrlConfig.upUrlConfig["retryRows"] = 1000
    }
    upUrlConfig.retryRows = ipDataUpUrlConfig.upUrlConfig["retryRows"]
  };

  const submitFormUpUrlConfig = async () => {
    loading.value = true;
    var params: upUrlParams = { protocol: '', ip: '', port: '', endpoint: '', enableRetry: false, retryRows: 1000, protocol2: '', ip2: '', port2: '', endpoint2: '' };
    params.protocol = upUrlConfigMap['upUrlConfig'].protocol;
    params.ip = upUrlConfigMap['upUrlConfig'].ip;
    params.port = upUrlConfigMap['upUrlConfig'].port;
    params.endpoint = upUrlConfigMap['upUrlConfig'].endpoint;
    params.enableRetry = upUrlConfigMap['upUrlConfig'].enableRetry;
    params.retryRows = Number(upUrlConfigMap['upUrlConfig'].retryRows);
    params.protocol2 = upUrlConfigMap['upUrlConfig'].protocol2;
    params.ip2 = upUrlConfigMap['upUrlConfig'].ip2;
    params.port2 = upUrlConfigMap['upUrlConfig'].port2;
    params.endpoint2 = upUrlConfigMap['upUrlConfig'].endpoint2;
    console.info('addUpUrl params = ' + JSON.stringify(params));
    const result = addUpUrl(params);
    console.info('addUpUrl result = ' + JSON.stringify(result));
    if (JSON.stringify(result) === '{}') {
      createMessage.success('update success');
    } else {
      createMessage.error('invalid operation');
    }
    loading.value = false;

    const r = await existRunningTasks();
    if (r.length > 0) {
      openModal();
      restartTips.value =
        t('maintenance.restart.taskRestartTipOK') +
        '\n' +
        t('maintenance.restart.taskRestartTipNO');
    }
  };

  const restartTips = ref(
    t('maintenance.restart.taskRestartTipOK') + '\n' + t('maintenance.restart.taskRestartTipNO'),
  );
  const restartLoading = ref({ loading: false });
  const restartCancel = ref({ disabled: false });
  const handleRestartTasksOk = async () => {
    restartLoading.value = { loading: true };
    restartCancel.value = { disabled: true };

    const list = await existRunningTasks();
    restartTips.value = '';
    for (const item of list) {
      await taskRestart(item);
    }

    restartLoading.value = { loading: false };
    restartCancel.value = { disabled: false };
    closeModal();
  };

  const existRunningTasks = async () => {
    const res = await getTaskList();
    const list = res.items.filter((item) => item.status == 1);
    return list;
  };

  const taskRestart = async (task) => {
    restartTips.value =
      restartTips.value + task.deviceName + ' ' + t('maintenance.restart.stopping');
    await StopTask({ taskId: task.taskId, deviceName: task.deviceName }).then();
    restartTips.value = restartTips.value + ' ' + t('maintenance.restart.starting');
    await StartTask({ taskId: task.taskId }).then();
    restartTips.value = restartTips.value + ' ' + t('maintenance.restart.startDone');
    restartTips.value = restartTips.value + '\n';
  };

  /////////////////////////////////////////////////////

  import { RotateCfgSetParams } from '/@/api/maintenance/model/index';
  import { getRotateCfg, modRotateCfg } from '/@/api/task/index';

  import { rotateCfgParams } from '/@/api/task/model/index';

  import { rotateRetentionCheck, rotateRetryRowsCheck } from '/@/utils/validateFuncs';
  import { Model } from 'echarts';

  const rotateCfgConfig: UnwrapRef<RotateCfgSetParams> = reactive({
    record: 72,
    serviceLog: 3,
  });

  const rotateCfgConfigRules = computed(() => {
    return {
      record: [
        {
          required: true,
          validator: rotateRetentionCheck,
          trigger: 'blur',
        },
      ],
      serviceLog: [
        {
          required: true,
          validator: rotateRetentionCheck,
          trigger: 'blur',
        },
      ],
    };
  });

  const rotateCfgConfigMap = {
    rotateCfgConfig,
  };

  const formItemListRotateCfgConfig = [
    {
      label: t('maintenance.rotateConfig.record'),
      field: 'record',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
      suffix: 'hours',
    },
    {
      label: t('maintenance.rotateConfig.serviceLog'),
      field: 'serviceLog',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
      suffix: 'days',
    },
  ];

  const ipDataRotateCfgConfig = reactive({
    rotateCfgConfig: [],
  });
  const initRotateCfgConfig = async () => {
    const result = await getRotateCfg();
    pageLoading.value = false;
    console.info('initRotateCfgConfig result=' + JSON.stringify(result));
    if (result) {
      ipDataRotateCfgConfig.rotateCfgConfig['record'] = result['data']['record'];
      ipDataRotateCfgConfig.rotateCfgConfig['serviceLog'] = result['data']['serviceLog'];
      // ipDataRotateCfgConfig.rotateCfgConfig['retryRows'] = result['data']['retryRows'];
      // ipDataRotateCfgConfig.rotateCfgConfig['enableRetry'] = result['data']['enableRetry'];
      rotateCfgConfig.record = ipDataRotateCfgConfig.rotateCfgConfig['record'];
      rotateCfgConfig.serviceLog = ipDataRotateCfgConfig.rotateCfgConfig['serviceLog'];
      // rotateCfgConfig.retryRows = ipDataRotateCfgConfig.rotateCfgConfig['retryRows'];
      // rotateCfgConfig.enableRetry = ipDataRotateCfgConfig.rotateCfgConfig['enableRetry'];
    }
  };

  const submitFormRotateCfgConfig = () => {
    loading.value = true;
    var params: rotateCfgParams = { record: 72, serviceLog: 3, retryRows: 1000, enableRetry: false };
    params.record = parseInt(rotateCfgConfigMap['rotateCfgConfig'].record, 10);
    params.serviceLog = parseInt(rotateCfgConfigMap['rotateCfgConfig'].serviceLog, 10);
    params.retryRows = parseInt(rotateCfgConfigMap['rotateCfgConfig'].retryRows, 10);
    params.enableRetry = rotateCfgConfigMap['rotateCfgConfig'].enableRetry
    // console.info('rotateCfgConfigMap = '+JSON.stringify(rotateCfgConfigMap));
    console.info('modRotateCfg params = ' + JSON.stringify(params));
    const result = modRotateCfg(params);
    console.info('modRotateCfg result = ' + JSON.stringify(result));
    if (JSON.stringify(result) === '{}') {
      createMessage.success('update success');
    } else {
      createMessage.error('invalid operation');
    }
    loading.value = false;
  };
  
  const initResourceAlarmgConfig = async () => {
    const resultAlarm = await getAlarm();
    console.info("getAlarm => "+JSON.stringify(resultAlarm))
    if (resultAlarm) {
      formState.fanSpeed = 9999;
      formState.boardTemperature = resultAlarm.boardTemperature;
      formState.coreTemperature = resultAlarm.coreTemperature;
      formState.cpuRate = Math.round(resultAlarm.cpuRate * 100);
      formState.totalMemoryScale = Math.round(resultAlarm.totalMemoryScale * 100);
      // formState.systemScale = resultAlarm.systemScale * 100;
      // formState.videoScale = resultAlarm.videoScale * 100;
      formState.tpuScale = Math.round(resultAlarm.tpuScale * 100);
      // formState.externalHardDiskRate = resultAlarm.externalHardDiskRate * 100;
      formState.diskRate = Math.round(resultAlarm.diskRate * 100);
      formState.tpuRate = Math.round(resultAlarm.tpuRate * 100);
    }
  };

  const placeholder = t('sys.form.phNumber');
  const color = '#0960BD';
  const formItemListAlarm = [
    // {
    //   label: t('maintenance.threshold.fanSpeed'),
    //   field: 'fanSpeed',
    //   unit: 'r/min',
    //   placeholder: t('sys.form.placeholder'),
    //   max: 1000000,
    // },
    {
      label: t('maintenance.threshold.boardTemperature'),
      field: 'boardTemperature',
      unit: '°C',
    },
    {
      label: t('maintenance.threshold.coreTemperature'),
      field: 'coreTemperature',
      unit: '°C',
    },
    {
      label: t('maintenance.threshold.cpuRate'),
      field: 'cpuRate',
      unit: '%',
    },
    {
      label: t('maintenance.threshold.totalMemoryScale'),
      field: 'totalMemoryScale',
      unit: '%',
    },
    // {
    //   label: t('maintenance.threshold.memoryScale'),
    //   field: 'systemScale',
    // },
    // {
    //   label: t('maintenance.threshold.videoScale'),
    //   field: 'videoScale',
    // },
    {
      label: t('maintenance.threshold.tpuScale'),
      field: 'tpuScale',
      unit: '%',
    },
    // {
    //   label: t('maintenance.threshold.externalHardDiskRate'),
    //   field: 'externalHardDiskRate',
    // },
    {
      label: t('maintenance.threshold.diskRate'),
      field: 'diskRate',
      unit: '%',
    },
    {
      label: t('overview.tpuUtililizationRate'),
      field: 'tpuRate',
      unit: '%',
    },
  ];
  const formState: UnwrapRef<AlarmParams> = reactive({
    fanSpeed: 0,
    boardTemperature: 0,
    coreTemperature: 0,
    cpuRate: 90,
    totalMemoryScale: 90,
    // systemScale: 90,
    // videoScale: 90,
    tpuScale: 90,
    // externalHardDiskRate: 90,
    diskRate: 90,
    tpuRate: 90,
  });

  const formTimedateConfig = [
    {
      label: t('maintenance.timedateConfig.localtime'),
      field: 'localtime',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
    {
      label: t('maintenance.timedateConfig.timezone'),
      field: 'timezone',
      placeholder: t('sys.form.placeholder'),
      type: 'select',
    },
    {
      label: t('maintenance.timedateConfig.ntpservers'),
      field: 'ntpservers',
      placeholder: t('sys.form.placeholder'),
      type: 'input',
    },
    {
      label: t('maintenance.timedateConfig.synced'),
      field: 'synced',
      placeholder: t('sys.form.placeholder'),
      type: 'text',
    },
    {
      label: t('maintenance.timedateConfig.ntpstat'),
      field: 'ntpstat',
      placeholder: t('sys.form.placeholder'),
      type: 'text',
    },
  ];

  const initTimedateConfig = async () => {
    const resultTimedate = await getTimedate();
    console.info("getTimedate => "+JSON.stringify(resultTimedate))
    if(resultTimedate) {
      formTimedate.localtime = resultTimedate.data.localtime
      formTimedate.localtimePast =formTimedate.localtime
      formTimedate.timezone = resultTimedate.data.timezone
      formTimedate.ntpservers = resultTimedate.data.ntpservers
      formTimedate.timezones = resultTimedate.data.timezones
      formTimedate.synced = resultTimedate.data.synced
      formTimedate.ntpstat = resultTimedate.data.ntpstat
    }
  };

  const formTimedate: UnwrapRef<TimedateParams> = reactive({
    localtimePast: '',
    localtime: '',
    timezone: '',
    ntpservers: '',
    timezones: [],
    synced: '',
    ntpstat: '',
  });

  const submitFormAlarm = async () => {
    try {
      loading.value = true;
      const params = {
        ...formState,
      };
      // 不需要转变的字段
      const staticFields = ['fanSpeed', 'boardTemperature', 'coreTemperature'];
      Object.keys(params).forEach((key) => {
        if (!staticFields.includes(key)) {
          params[key] = params[key] / 100;
        }
      });
      const isParams = Object.values(params).every((value) => value > 0);
      if (!isParams) {
        createMessage.error(`Config value must be greater than 0！`, 1);
        init();
      } else {
        await setAlarm(params)
          .then(() => {
            createMessage.success('Done');
          })
          .catch(() => {
            createMessage.error('Failed');
          });
      }
    } catch (error) {
      // createErrorModal({
      //   title: t('sys.api.errorTip'),
      //   content: (error as unknown as Error).message || t('sys.api.networkExceptionMsg'),
      //   getContainer: () => document.body,
      // });
    } finally {
      loading.value = false;
    }
  };

  const handleTZChange = async (value) => {
    console.info(`Selected: ${formTimedate['timezone']}`);
    formTimedate['timezone'] = value;
  };

  const submitFormTimedate = async () => {
    try {
      loading.value = true;
      let formTimedateReq = {...formTimedate}
      formTimedateReq.timezones = [];
      if(formTimedate.localtimePast === formTimedate.localtime) {
        formTimedateReq.localtime = '';
      }
      formTimedate.localtimePast = formTimedate.localtime;
      const params = {
        ...formTimedateReq,
      };
      await setTimedate(params)
        .then(() => {
          createMessage.success('Done');
        })
        .catch(() => {
          createMessage.error('Failed');
        });
    } catch (error) {
    } finally {
      loading.value = false;
    }
  };
</script>
