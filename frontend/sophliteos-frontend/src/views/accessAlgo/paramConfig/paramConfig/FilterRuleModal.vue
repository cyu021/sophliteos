<template>
    <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="Submit" width="70%" :height="1000">
      <div>{{ taskName + '-' + algorithmName }}</div>
      <Divider />

      <div style="padding: 14px; display: flex; flex-direction: row; align-items: center;">
        <Input v-model:value="extend.FilterName" placeholder="规则名称" />
      </div>

      <div style="padding: 14px; padding-bottom: 30px;">
        <Collapse v-model:activeKey="attributeActiveKey">

          <CollapsePanel key="1" header="Attribute conditions">
            <ul>
              <li style="padding: 14px;" v-for="(item, index) in extend.Filter" :key="index">
                <div style="padding: 14px; border: 1px black dashed; display: flex; align-items: center; flex-direction: row;">
                  <div style="width: 150px; font-size: large; text-align: left;"> {{ item.Label + (item.LabelName ? ' - ' + item.LabelName : '') }} </div>
                  <div style="display: flex; flex-direction: column; flex-grow: 1; padding-left: 48px;">
                    <ul>
                      <li 
                        style="align-items: center; height: 48px; justify-content: center; display: flex; margin: 8px; border: 1px gray dashed; flex-grow: 1; flex-direction: row;" 
                        v-for="(jtem, jndex) in item.Rule" :key="jndex">
                        <div style="width: 250px; text-align: center; ">
                          {{ jtem.K }}
                        </div>
                        <div style="padding: 4px;">
                          <Dropdown :trigger="['click']">
                            <Button style="width: 100px; ">
                              {{ jtem.Op }}
                              <DownOutlined style="opacity: 0.5;"/>
                            </Button>
                            <template #overlay>
                              <Menu>
                                <MenuItem v-for="(value, index) in jtem.ops" :key="index" @click="updateExtendFilterOp(item.Label, jtem.K, value)">
                                  {{ value }}
                                </MenuItem>
                              </Menu>
                            </template>
                          </Dropdown>
                        </div>
                        <div style="padding: 4px;">
                          <Input v-if="typeof jtem.V === 'number' || typeof jtem.V === 'string' && !isNaN(Number(jtem.V))" style="width: 250px;" v-model:value="jtem.V" />
                          <Dropdown v-else :trigger="['click']">
                            <Button style="width: 250px; ">
                              {{ jtem.V }}
                              <DownOutlined style="opacity: 0.5;"/>
                            </Button>
                            <template #overlay>
                              <Menu>
                                <MenuItem v-for="(value, index) in jtem.Vs" :key="index" @click="updateExtendFilter(item.Label, jtem.K, value)">
                                  {{ value }}
                                </MenuItem>
                              </Menu>
                            </template>
                          </Dropdown>
                        </div>
                        <div>
                          <Button danger type="text" @click="ruleDelete(item.Label, jndex)">
                            <DeleteOutlined/>
                          </Button>
                        </div>
                      </li>
                    </ul>
                    <div style="display: flex; align-items: center; justify-content: center; padding-top: 24px;">
                      <Dropdown :trigger="['click']">
                        <Button type="primary" style="width: 150px; ">
                          AND
                        </Button>
                        <template #overlay>
                          <Menu>
                            <MenuItem v-for="(value, key) in item.kvps" :key="key" @click="ruleAdd(item.Annotator, item.Label, key)">
                              {{ key }}
                            </MenuItem>
                          </Menu>
                        </template>
                      </Dropdown>
                    </div>
                  </div>
                  <div style="align-items: center; justify-content: center; ">
                    <Button danger type="text" @click="attrDelete(item.Label)">
                      <DeleteOutlined/>
                    </Button>
                  </div>
                </div>
              </li>
            </ul>
            <div style="padding: 16px;">
              <!-- <Button type="primary" @click="attrAdd">OR</Button> -->
              <Dropdown :trigger="['click']">
                <Button type="primary" style="width: 150px; ">
                  OR
                </Button>
                <template #overlay>
                  <Menu>
                    <MenuItem v-for="(value, key) in extend.FilterOps" :key="key" @click="attrAdd(key)">
                      {{ value.label }}
                    </MenuItem>
                  </Menu>
                </template>
              </Dropdown>
            </div>
          </CollapsePanel>

          <CollapsePanel key="2" header="Time conditions">
            <ul>
              <li style="padding: 14px;" v-for="(item, index) in extend.FilterPeriod" :key="index">
                <div style="padding: 14px; border: 1px black dashed; display: flex; flex-direction: row; align-items: center;"> 
                  <div style="flex-grow: 1;">
                    <div style="display: flex; flex-direction: row;">
                      <div style="display: flex; flex-direction: row; align-items: center; padding: 8px;">
                        <div style="width: 120px;">开始时间</div>
                        <TimePicker :value="dayjs(item.TimeStart || '0000', 'HHmm')" @change="periodUpdateTimeStart(index, $event)"/>
                      </div>
                      <div style="display: flex; flex-direction: row; align-items: center; padding: 8px;">
                        <div style="width: 120px;">结束时间</div>
                        <TimePicker :value="dayjs(item.TimeEnd || '0000', 'HHmm')" @change="periodUpdateTimeEnd(index, $event)"/>
                      </div>
                    </div>

                    <div style="display: flex; flex-direction: row;">
                      <div style="display: flex; flex-direction: row; align-items: center; padding: 8px;">
                        <div style="width: 120px;">开始日期</div>
                        <DatePicker :value="dayjs(item.DateStart || dayjs(), 'YYYYMMDD')" @change="periodUpdateDateStart(index, $event)"/>
                      </div>
                      <div style="display: flex; flex-direction: row; align-items: center; padding: 8px;">
                        <div style="width: 120px;">结束日期</div>
                        <DatePicker :value="dayjs(item.DateEnd || dayjs(), 'YYYYMMDD')" @change="periodUpdateDateEnd(index, $event)"/>
                      </div>
                    </div>

                    <div style="padding: 14px;">
                      <CheckboxGroup :value="item.Weekday.split(',')" :options="weekdays" @change="periodUpdateWeekday(index, $event)"/>
                    </div>
                  </div>
                  <div style="align-items: center; justify-content: center; ">
                    <Button danger type="text" @click="periodDelete(index)">
                      <DeleteOutlined/>
                    </Button>
                  </div>
                </div>
              </li>
            </ul>
            <div style="padding: 16px;">
              <Button type="primary" @click="periodAdd">OR</Button>
            </div>
          </CollapsePanel>
        </Collapse>
      </div>
    </BasicModal>
  </template>
  
  <script lang="ts" setup>
    import { Divider, Input, Collapse, CollapsePanel, Dropdown, Button, Menu, MenuItem, Space, CheckboxGroup, Checkbox, TimePicker, DatePicker, } from 'ant-design-vue';
    import { BasicModal, useModalInner } from '/@/components/Modal';
    import { useI18n } from '/@/hooks/web/useI18n';
    import { onMounted, ref, } from 'vue';
    import apis from './api';
    import { DeleteOutlined, DownOutlined, } from '@ant-design/icons-vue';
    import dayjs from 'dayjs';

    const { t } = useI18n();
    const title = t('paramConfig.param.filterRule');
    const emit = defineEmits(['success', 'register']);

    const attributeActiveKey = ref(['1', '2'])

    const attrList = ref()

    const numberOps = ['=', '>', '<', '≥', '≤'];
    const selectOps = ['='];

    const weekdays = [
      { label: '日', value: '0' },
      { label: '一', value: '1' },
      { label: '二', value: '2' },
      { label: '三', value: '3' },
      { label: '四', value: '4' },
      { label: '五', value: '5' },
      { label: '六', value: '6' }]

    const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
      setModalProps({ confirmLoading: false });

    });

    function getKvps(annotator, label) {
      if (attrList.value && attrList.value[annotator] && attrList.value[annotator][label] && attrList.value[annotator][label].kvp) {
        const kvp = attrList.value[annotator][label].kvp;
        return kvp;
      }

      return null;
    }

    function getOpsOfKvp(annotator, label, key) {
      const kvp = getKvps(annotator, label)
      if (kvp) {
        if (kvp[key]) {
          return kvp[key]
        }
      }

      return [];
    }

    function updateExtendFilter(label, key, value) {
      if (extend.value.Filter) {
        extend.value.Filter.forEach((item) => {
          if (label === item.Label) {
            item.Rule.forEach((rule) => {
              if (rule.K === key) {
                rule.V = value;
              }
            });
          }
        });
      }
    }

    function updateExtendFilterOp(label, key, op) {
      if (extend.value.Filter) {
        extend.value.Filter.forEach((item) => {
          if (label === item.Label) {
            item.Rule.forEach((rule) => {
              if (rule.K === key) {
                rule.Op = op;
              }
            });
          }
        });
      }
    }

    function periodDelete(index) {
      if (extend.value.FilterPeriod) {
        extend.value.FilterPeriod.splice(index, 1);
      }
    }

    function periodUpdateWeekday(index, value) {
      if (extend.value.FilterPeriod) {
        extend.value.FilterPeriod[index].Weekday = value.join(',');
      }
    }

    function periodUpdateTimeStart(index, value) {
      if (extend.value.FilterPeriod) {
        extend.value.FilterPeriod[index].TimeStart = value.format('HHmm');
      }
    }

    function periodUpdateTimeEnd(index, value) {
      if (extend.value.FilterPeriod) {
        extend.value.FilterPeriod[index].TimeEnd = value.format('HHmm');
      }
    }

    function periodUpdateDateStart(index, value) {
      if (extend.value.FilterPeriod) {
        extend.value.FilterPeriod[index].DateStart = value.format('YYYYMMDD');
      }
    }

    function periodUpdateDateEnd(index, value) {
      if (extend.value.FilterPeriod) {
        extend.value.FilterPeriod[index].DateEnd = value.format('YYYYMMDD');
      }
    }

    function periodAdd() {
      if (extend.value.FilterPeriod) {
        extend.value.FilterPeriod.push({
          DateStart: '',
          DateEnd: '',
          TimeStart: '',
          TimeEnd: '',
          Weekday: ''
        });
      }
    }

    function ruleAdd(annotator, label, key) {
      if (extend.value.Filter) {
        extend.value.Filter.forEach((item, index) => {
          if (label === item.Label) {
            const values = getOpsOfKvp(annotator, label, key);
            const ops = values.length <=0 ? numberOps : selectOps;

            console.log('rule add', ops, values);

            item.Rule.push({
              Grp: index,
              Op: ops[0],
              ops: ops,
              Vs: values,
              K: key,
              V: values.length > 0 ? values[0] : 0,
            })
          }
        });
      }
    }

    function ruleDelete(label, index) {
      if (extend.value.Filter) {
        extend.value.Filter.forEach((item) => {
          if (label === item.Label) {
            item.Rule.splice(index, 1);
          }
        });
      }
    }

    function attrDelete(label) {
      if (extend.value.Filter) {
        let index = extend.value.Filter.findIndex(obj => obj.Label === label);

        if (index !== -1) {
          extend.value.Filter.splice(index, 1);
        }
      }
    }

    function attrAdd(label) {
      if (extend.value.Filter) {
        const kvps = getKvps(props.algorithmName, label);
        const key = Object.keys(kvps)[0]
        const values = getOpsOfKvp(props.algorithmName, label, key)
        const ops = values.length <=0 ? numberOps : selectOps;
        const name = nameOfKvpLabel(props.algorithmName, label);

        const ruleCount = extend.value.Filter.length || 0;

        extend.value.Filter.push({
          Annotator: props.algorithmName,
          Label: label,
          LabelName: name,
          Rtsp: '',
          kvps: kvps,
          Rule: [{
            Grp: ruleCount + 1,
            Op: ops[0],
            ops: ops,
            Vs: values,
            K: key,
            V: values.length > 0 ? values[0] : 0,
          }],
        });
      }
    }

    function getFilterOps() {
      const annotator = props.algorithmName;
      const ops = attrList.value[annotator];
      return ops;
    }

    function updateExtend() {
      if (extend.value.Filter && attrList.value) {
        extend.value.Filter.forEach((item) => {
          const annotator = item.Annotator;
          const label = item.Label;

          const name = nameOfKvpLabel(annotator, label);
          item.LabelName = name;
          item.kvps = getKvps(annotator, label);

          item.Rule.forEach((rule) => {
            const values = getOpsOfKvp(annotator, label, rule.K);
            rule.ops = values.length <=0 ? numberOps : selectOps;
            rule.Vs = values;
          });
        });
      }

      if (extend.value && attrList.value) {
        extend.value.FilterOps = getFilterOps();
      }
    }

    function nameOfKvpLabel(annotator, label) {
      if (attrList.value && attrList.value[annotator] && attrList.value[annotator][label] && attrList.value[annotator][label].label) {
        return attrList.value[annotator][label].label
      } else {
        return ''
      }
    }

    async function Submit() {
      const newValue = JSON.parse(JSON.stringify(extend.value))

      delete newValue.FilterOps;
      newValue.Filter.forEach(filter => {
        delete filter.LabelName;
        delete filter.kvps;

        filter.Rule.forEach(rule => {
          delete rule.Vs;
          delete rule.ops;
        });
      });

      console.log('submit', newValue)

      try {
        // const values = await validate();
        // setModalProps({ confirmLoading: true });
        // // TODO custom api
        // const minBox = { width: 0, height: 0 };
        // minBox.width = Number(values.width);
        // minBox.height = Number(values.height);
        // values.ability = revOption[values.ability];
        // // console.log(values);
        // await algoConfigMod({
        //   ability: values.ability,
        //   minBox: minBox,
        //   interval: Number(values.interval),
        //   threshold: Number(values.threshold),
        // });
        emit('success', newValue);
        closeModal();
      } finally {
        setModalProps({ confirmLoading: false });
      }
    }

    const extend = ref({
      "FilterName": "五一期间晚上出现的老年男子与小萝莉",
      "FilterPeriod":
      [
        {
          "DateStart": "20240501",
          "DateEnd": "20240505",
          "TimeStart": "1800",
          "TimeEnd": "0800",
          "Weekday": "0,1,2,3,4,5,6"
        }
      ],
      "Filter":
      [
        {
          "Annotator": "full_structure",
          "Label": "1420",
          "Rtsp": "rtsp://192.168.0.158:8554/sample",
          "Rule":
          [
            {
                "Grp": 1,
                "Op": "=",
                "K": "special_vehicle",
                "V": "common_vehicle"
            },
            {
                "Grp": 1,
                "Op": "=",
                "K": "vehicle_color",
                "V": "yellow"
            }
          ]
        },
      ]
    })

    onMounted(async ()=> {
      apis.attrList().then((res) => {
        attrList.value = res;
        updateExtend()
      });
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
    });

  </script>
  