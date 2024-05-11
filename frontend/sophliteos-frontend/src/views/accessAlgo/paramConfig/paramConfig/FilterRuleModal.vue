<template>
    <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="Submit" width="70%" :height="1000">
      <div>{{ taskName + '-' + algorithmName }}</div>
      <Divider />

      <div style="padding: 14px; display: flex; flex-direction: row; align-items: center;">
        <Input v-model:value="cacheExtend.FilterName" placeholder="规则名称" />
      </div>

      <div style="padding: 14px; padding-bottom: 30px;">
        <Collapse v-model:activeKey="attributeActiveKey">

          <CollapsePanel key="1" header="属性规则">
            <ul>
              <li style="padding: 14px;" v-for="(item, index) in cacheExtend.Filter" :key="index">
                <div style="padding: 14px; border: 1px black dashed; display: flex; align-items: center; flex-direction: row;">
                  <div style="width: 150px; text-align: left;"> {{ nameOfKvpLabel(algorithmName, item.Label) + '(' + item.Label + ')' }} </div>
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
                            <Button style="width: 100px; " @click="opClick(item.Label, jtem.K)">
                              {{ jtem.Op }}
                              <DownOutlined style="opacity: 0.5;"/>
                            </Button>
                            <template #overlay>
                              <Menu>
                                <MenuItem v-for="(value, index) in ops" :key="index" @click="updateExtendFilterOp(item.Label, jtem.K, value)">
                                  {{ value }}
                                </MenuItem>
                              </Menu>
                            </template>
                          </Dropdown>
                        </div>
                        <div style="padding: 4px;">
                          <Input v-if="typeof jtem.V === 'number' || typeof jtem.V === 'string' && !isNaN(Number(jtem.V))" style="width: 250px;" v-model:value="jtem.V" />
                          <Dropdown v-else :trigger="['click']">
                            <Button style="width: 250px; " @click="vClick(item.Label, jtem.K)">
                              {{ jtem.V }}
                              <DownOutlined style="opacity: 0.5;"/>
                            </Button>
                            <template #overlay>
                              <Menu>
                                <MenuItem v-for="(value, index) in kVs" :key="index" @click="updateExtendFilter(item.Label, jtem.K, value)">
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
                        <Button type="primary" style="width: 150px; " @click="andClick(item.Label)">
                          AND
                        </Button>
                        <template #overlay>
                          <Menu>
                            <MenuItem v-for="value in kvps" :key="value" @click="ruleAdd(item.Annotator, item.Label, value)">
                              {{ value }}
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
              <Dropdown :trigger="['click']">
                <Button type="primary" style="width: 150px;" >
                  OR
                </Button>
                <template #overlay>
                  <Menu>
                    <MenuItem v-for="label in labels" :key="label.key" @click="attrAdd(label.key)">
                      {{ label.value }}
                    </MenuItem>
                  </Menu>
                </template>
              </Dropdown>
            </div>
          </CollapsePanel>

          <CollapsePanel key="2" header="时间规则">
            <ul>
              <li style="padding: 14px;" v-for="(item, index) in cacheExtend.FilterPeriod" :key="index">
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
    import { onMounted, ref, watch, computed, watchEffect, } from 'vue';
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
      if (cacheExtend.value.Filter) {
        cacheExtend.value.Filter.forEach((item) => {
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
      if (cacheExtend.value.Filter) {
        cacheExtend.value.Filter.forEach((item) => {
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
      if (cacheExtend.value.FilterPeriod) {
        cacheExtend.value.FilterPeriod.splice(index, 1);
      }
    }

    function periodUpdateWeekday(index, value) {
      if (cacheExtend.value.FilterPeriod) {
        cacheExtend.value.FilterPeriod[index].Weekday = value.join(',');
      }
    }

    function periodUpdateTimeStart(index, value) {
      if (cacheExtend.value.FilterPeriod) {
        cacheExtend.value.FilterPeriod[index].TimeStart = value.format('HHmm');
      }
    }

    function periodUpdateTimeEnd(index, value) {
      if (cacheExtend.value.FilterPeriod) {
        cacheExtend.value.FilterPeriod[index].TimeEnd = value.format('HHmm');
      }
    }

    function periodUpdateDateStart(index, value) {
      if (cacheExtend.value.FilterPeriod) {
        cacheExtend.value.FilterPeriod[index].DateStart = value.format('YYYYMMDD');
      }
    }

    function periodUpdateDateEnd(index, value) {
      if (cacheExtend.value.FilterPeriod) {
        cacheExtend.value.FilterPeriod[index].DateEnd = value.format('YYYYMMDD');
      }
    }

    function periodAdd() {
      if (!cacheExtend.value.FilterPeriod) {
        cacheExtend.value.FilterPeriod = [];
      }

      cacheExtend.value.FilterPeriod.push({
        DateStart: '20240101',
        DateEnd: '20240131',
        TimeStart: '0800',
        TimeEnd: '2000',
        Weekday: '0,1,2,3,4,5,6',
      });
    }

    function ruleAdd(annotator, label, key) {
      if (cacheExtend.value.Filter) {
        cacheExtend.value.Filter.forEach((item, index) => {
          if (label === item.Label) {
            const values = getOpsOfKvp(annotator, label, key);
            const ops = values.length <=0 ? numberOps : selectOps;

            console.log('rule add', ops, values);

            if (!item.Rule) { item.Rule = []; }

            item.Rule.push({
              Grp: index,
              Op: ops[0],
              K: key,
              V: values.length > 0 ? values[0] : 0,
            })
          }
        });
      }
    }

    function ruleDelete(label, index) {
      if (cacheExtend.value.Filter) {
        cacheExtend.value.Filter.forEach((item) => {
          if (label === item.Label) {
            item.Rule.splice(index, 1);
          }
        });
      }
    }

    function attrDelete(label) {
      if (cacheExtend.value.Filter) {
        let index = cacheExtend.value.Filter.findIndex(obj => obj.Label === label);

        if (index !== -1) {
          cacheExtend.value.Filter.splice(index, 1);
        }
      }
    }

    const currentLabelId = ref(''); // 当前label id 数字，比如1420
    const currentKvp = ref(''); // 当前 kvp key，比如 automobile_angle

    function opClick(labelId, kvp) {
      currentLabelId.value = labelId;
      currentKvp.value = kvp;
    }

    function andClick(labelId) {
      currentLabelId.value = labelId;
    }

    function vClick(labelId, kvp) {
      currentLabelId.value = labelId;
      currentKvp.value = kvp;
    }

    const kVs = computed(() => {
      return getOpsOfKvp(props.algorithmName, currentLabelId.value, currentKvp.value);
    })
    
    const labels = computed(() => {
      if (props.algorithmName && attrList.value && attrList.value[props.algorithmName]) {
        const data = attrList.value[props.algorithmName];
        return Object.keys(data).map((key) => {
          return {
            key: key,
            value: data[key].label,
          }
        })
      }

      return []
    })

    const kvps = computed(() => {
      console.log('kvp computed', attrList.value, currentLabelId.value)
      if (attrList.value) {
        if (attrList.value[props.algorithmName] && attrList.value[props.algorithmName][currentLabelId.value]) {
          const kvp = attrList.value[props.algorithmName] && attrList.value[props.algorithmName][currentLabelId.value].kvp;
          console.log('kvp computed kvp', Object.keys(kvp))

          return Object.keys(kvp);
        }
      }

      return [];
    })

    const ops = computed(() => {
      if (attrList.value) {
        if (attrList.value[props.algorithmName] && attrList.value[props.algorithmName][currentLabelId.value]) {
          const kvp = attrList.value[props.algorithmName] && attrList.value[props.algorithmName][currentLabelId.value].kvp;
          if (kvp && kvp[currentKvp.value]) {
            const k = kvp[currentKvp.value];
            if (k && k.length <= 0) {
              return numberOps;
            }
          }
        }
      }

      return selectOps;
    })

    function attrAdd(labelId) {
      currentLabelId.value = labelId;
      
      if (!cacheExtend.value.Filter) {
        cacheExtend.value.Filter = [];
      }

      cacheExtend.value.Filter.push({
        Annotator: props.algorithmName,
        Label: labelId,
        Rule: [],
      });
    }

    function nameOfKvpLabel(annotator, label) {
      if (attrList.value && attrList.value[annotator] && attrList.value[annotator][label] && attrList.value[annotator][label].label) {
        return attrList.value[annotator][label].label
      } else {
        return ''
      }
    }

    async function Submit() {
      const newValue = JSON.parse(JSON.stringify(cacheExtend.value))

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
        emit('success', newValue);
        closeModal();
      } finally {
        setModalProps({ confirmLoading: false });
      }
    }

    onMounted(() => {
      apis.attrList().then((res) => {
        attrList.value = res;
      });

    })

    watch(() => props.extend, async (newValue, oldValue) => {
      console.log('watch props.extend', props.extend)
      const value = await newValue;
      cacheExtend.value = value;
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

  </script>
  