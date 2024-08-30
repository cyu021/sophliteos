<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="
      modelSelectDisabled
        ? t('maintenance.ruleTemplate.edit')
        : t('maintenance.ruleTemplate.create')
    "
    @ok="Submit"
    width="70%"
    :height="1000"
  >
    <div style="display: flex; flex-direction: column; align-content: space-between; padding: 14px">
      <div style="padding-bottom: 24px">
        {{ t('maintenance.ruleTemplate.modelName') + ':' }}
        <Select
          :value="modelDisplayName"
          :options="
            aiModelNameOptions.map((i) => {
              return { value: i, label: i };
            })
          "
          style="padding-left: 4px; width: 300px"
          :disabled="modelSelectDisabled"
          @select="aiModelSelectChange"
        />
      </div>
      <div v-if="ruleTemplateSelectShow"
        >{{ t('maintenance.ruleTemplate.copy') }}
        <Select
          v-model:value="currentFilterTemplateName"
          :options="
            (filterTemplate[modelDisplayName] || []).map((i) => {
              return { value: i.FilterName, label: i.FilterName };
            })
          "
          style="width: 300px"
          @select="updateFilterTemplate"
        />
      </div>
    </div>
    <Divider />

    <div style="padding: 14px; display: flex; flex-direction: row; align-items: center">
      <Input
        :disabled="modelSelectDisabled"
        v-model:value="cacheExtend.FilterName"
        :placeholder="t('maintenance.ruleTemplate.templateName')"
      />
    </div>

    <div style="padding: 14px; padding-bottom: 30px">
      <Collapse v-model:activeKey="attributeActiveKey">
        <CollapsePanel key="1" :header="t('paramConfig.filter.attributeRule')">
          <ul>
            <li v-for="(item, index) in cacheExtend.Filter" :key="index">
              <div
                style="
                  padding: 14px;
                  border: 1px black dashed;
                  display: flex;
                  align-items: center;
                  flex-direction: row;
                "
              >
                <div style="width: 150px; text-align: left">
                  {{ nameOfKvpLabel(algorithmName, item.Label) + '(' + item.Label + ')' }}
                </div>
                <div
                  style="display: flex; flex-direction: column; flex-grow: 1; padding-left: 48px"
                >
                  <ul>
                    <li v-for="(jtem, jndex) in item.Rule" :key="jndex">
                      <div
                        style="
                          align-items: center;
                          height: 48px;
                          justify-content: center;
                          display: flex;
                          border: 1px gray dashed;
                          flex-grow: 1;
                          flex-direction: row;
                        "
                      >
                        <div style="width: 250px; text-align: center">
                          {{ jtem.K }}
                        </div>
                        <div style="padding: 4px">
                          <Dropdown :trigger="['click']">
                            <Button style="width: 100px" @click="opClick(item.Label, jtem.K)">
                              {{ jtem.Op }}
                              <DownOutlined style="opacity: 0.5" />
                            </Button>
                            <template #overlay>
                              <Menu>
                                <MenuItem
                                  v-for="value in ops"
                                  :key="value"
                                  @click="updateExtendFilterOpByIndex(index, jndex, value)"
                                >
                                  {{ value }}
                                </MenuItem>
                              </Menu>
                            </template>
                          </Dropdown>
                        </div>
                        <div style="padding: 4px">
                          <Input
                            v-if="
                              typeof jtem.V === 'number' ||
                              (typeof jtem.V === 'string' && !isNaN(Number(jtem.V)))
                            "
                            style="width: 250px"
                            v-model:value="jtem.V"
                          />
                          <Dropdown v-else :trigger="['click']">
                            <Button style="width: 250px" @click="vClick(item.Label, jtem.K)">
                              {{ jtem.V }}
                              <DownOutlined style="opacity: 0.5" />
                            </Button>
                            <template #overlay>
                              <Menu>
                                <MenuItem
                                  v-for="value in kVs"
                                  :key="value"
                                  @click="updateExtendFilterByIndex(index, jndex, value)"
                                >
                                  {{ value }}
                                </MenuItem>
                              </Menu>
                            </template>
                          </Dropdown>
                        </div>
                        <div>
                          <Button danger type="text" @click="ruleDeleteByIndex(index, jndex)">
                            <DeleteOutlined />
                          </Button>
                        </div>
                      </div>
                      <div
                        v-if="jndex < item.Rule.length - 1"
                        style="
                          height: 18px;
                          display: flex;
                          justify-content: center;
                          align-items: center;
                        "
                      >
                        <div
                          style="
                            border-radius: 6px;
                            height: 28px;
                            width: 58px;
                            font-size: 16px;
                            background-color: #a0a0a0;
                            text-align: center;
                            display: flex;
                            justify-content: center;
                            align-items: center;
                            color: white;
                          "
                        >
                          <div>AND</div>
                        </div>
                      </div>
                    </li>
                  </ul>
                  <div
                    style="
                      display: flex;
                      align-items: center;
                      justify-content: center;
                      padding-top: 24px;
                    "
                  >
                    <Dropdown :trigger="['click']">
                      <Button type="primary" style="width: 150px" @click="andClick(item.Label)">
                        AND
                      </Button>
                      <template #overlay>
                        <Menu>
                          <MenuItem
                            v-for="value in kvps"
                            :key="value"
                            @click="ruleAdd(item.Annotator, index, value)"
                          >
                            {{ value }}
                          </MenuItem>
                        </Menu>
                      </template>
                    </Dropdown>
                  </div>
                </div>
                <div style="align-items: center; justify-content: center">
                  <Button danger type="text" @click="attrDeleteByIndex(index)">
                    <DeleteOutlined />
                  </Button>
                </div>
              </div>
              <div
                v-if="index < cacheExtend.Filter.length - 1"
                style="
                  height: 18px;
                  display: flex;
                  justify-content: flex-start;
                  align-items: center;
                "
              >
                <div
                  style="
                    margin-left: 24px;
                    border-radius: 6px;
                    height: 38px;
                    width: 38px;
                    font-size: 20px;
                    background-color: #a0a0a0;
                    text-align: center;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    color: white;
                  "
                >
                  <div>OR</div>
                </div>
              </div>
            </li>
          </ul>
          <div style="padding: 16px">
            <Dropdown :trigger="['click']">
              <Button type="primary" style="width: 150px"> OR </Button>
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

        <CollapsePanel key="2" :header="t('paramConfig.filter.timeRules')">
          <ul>
            <li v-for="(item, index) in cacheExtend.FilterPeriod" :key="index">
              <div
                style="
                  padding: 14px;
                  border: 1px black dashed;
                  display: flex;
                  flex-direction: row;
                  align-items: center;
                "
              >
                <div style="flex-grow: 1">
                  <div style="display: flex; flex-direction: row">
                    <div
                      style="display: flex; flex-direction: row; align-items: center; padding: 8px"
                    >
                      <div style="width: 120px">{{ t('paramConfig.filter.startTime') }}</div>
                      <TimePicker
                        format="HH:mm"
                        :value="dayjs(item.TimeStart || '0000', 'HHmm')"
                        @change="periodUpdateTimeStart(index, $event)"
                      />
                    </div>
                    <div
                      style="display: flex; flex-direction: row; align-items: center; padding: 8px"
                    >
                      <div style="width: 120px">{{ t('paramConfig.filter.endTime') }}</div>
                      <TimePicker
                        format="HH:mm"
                        :value="dayjs(item.TimeEnd || '2359', 'HHmm')"
                        @change="periodUpdateTimeEnd(index, $event)"
                      />
                    </div>
                  </div>

                  <div style="display: flex; flex-direction: row">
                    <div
                      style="display: flex; flex-direction: row; align-items: center; padding: 8px"
                    >
                      <div style="width: 120px">{{ t('paramConfig.filter.startDate') }}</div>
                      <DatePicker
                        :value="dayjs(item.DateStart || dayjs().month(0).date(1), 'YYYYMMDD')"
                        @change="periodUpdateDateStart(index, $event)"
                      />
                    </div>
                    <div
                      style="display: flex; flex-direction: row; align-items: center; padding: 8px"
                    >
                      <div style="width: 120px">{{ t('paramConfig.filter.endDate') }}</div>
                      <DatePicker
                        :value="
                          dayjs(
                            item.DateEnd ||
                              dayjs()
                                .year(dayjs().year() + 10)
                                .month(11)
                                .date(31),
                            'YYYYMMDD',
                          )
                        "
                        @change="periodUpdateDateEnd(index, $event)"
                      />
                    </div>
                  </div>

                  <div style="padding: 14px">
                    <CheckboxGroup
                      :value="item.Weekday.split(',')"
                      :options="weekdays"
                      @change="periodUpdateWeekday(index, $event)"
                    />
                  </div>
                </div>
                <div style="align-items: center; justify-content: center">
                  <Button danger type="text" @click="periodDelete(index)">
                    <DeleteOutlined />
                  </Button>
                </div>
              </div>
              <div
                v-if="index < cacheExtend.FilterPeriod.length - 1"
                style="
                  height: 18px;
                  display: flex;
                  justify-content: flex-start;
                  align-items: center;
                "
              >
                <div
                  style="
                    margin-left: 24px;
                    border-radius: 6px;
                    height: 38px;
                    width: 38px;
                    font-size: 20px;
                    background-color: #a0a0a0;
                    text-align: center;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    color: white;
                  "
                >
                  <div>OR</div>
                </div>
              </div>
            </li>
          </ul>
          <div style="padding: 16px">
            <Button type="primary" @click="periodAdd">OR</Button>
          </div>
        </CollapsePanel>
      </Collapse>
    </div>
  </BasicModal>
</template>

<script lang="ts" setup>
  import {
    Select,
    SelectOption,
    Divider,
    Input,
    Collapse,
    CollapsePanel,
    Dropdown,
    Button,
    Menu,
    MenuItem,
    CheckboxGroup,
    TimePicker,
    DatePicker,
  } from 'ant-design-vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { onMounted, ref, watch, computed } from 'vue';
  import apis from './api';
  import { DeleteOutlined, DownOutlined } from '@ant-design/icons-vue';
  import dayjs from 'dayjs';

  const aiModelName = ref('');
  const aiModelNameOptions = ref([]);

  const { t } = useI18n();
  const emit = defineEmits(['success', 'register']);

  const attributeActiveKey = ref(['1', '2']);

  const attrList = ref();

  const filterTemplate = ref();
  const currentFilterTemplateName = ref('');

  const numberOps = ['=', '!=', '>', '<'];
  const selectOps = ['=', '!='];

  const weekdays = [
    { label: t('paramConfig.filter.sunday'), value: '0' },
    { label: t('paramConfig.filter.monday'), value: '1' },
    { label: t('paramConfig.filter.tuesday'), value: '2' },
    { label: t('paramConfig.filter.wednesday'), value: '3' },
    { label: t('paramConfig.filter.thursday'), value: '4' },
    { label: t('paramConfig.filter.friday'), value: '5' },
    { label: t('paramConfig.filter.saturday'), value: '6' },
  ];

  const [registerModal, { setModalProps, closeModal, getVisible }] = useModalInner(async (data) => {
    setModalProps({ confirmLoading: false });
  });

  function getKvps(annotator, label) {
    if (
      attrList.value &&
      attrList.value[annotator] &&
      attrList.value[annotator][label] &&
      attrList.value[annotator][label].kvp
    ) {
      const kvp = attrList.value[annotator][label].kvp;
      return kvp;
    }

    return null;
  }

  function getOpsOfKvp(annotator, label, key) {
    const kvp = getKvps(annotator, label);
    if (kvp) {
      if (kvp[key]) {
        return kvp[key];
      }
    }

    return [];
  }

  function updateExtendFilterByIndex(labelIndex, keyIndex, value) {
    if (cacheExtend.value.Filter) {
      if (cacheExtend.value.Filter.length > labelIndex) {
        if (
          cacheExtend.value.Filter[labelIndex] &&
          cacheExtend.value.Filter[labelIndex].Rule.length > keyIndex
        ) {
          cacheExtend.value.Filter[labelIndex].Rule[keyIndex].V = value;
        }
      }
    }
  }

  function updateExtendFilterOpByIndex(labelIndex, keyIndex, op) {
    if (cacheExtend.value.Filter) {
      if (cacheExtend.value.Filter.length > labelIndex) {
        if (
          cacheExtend.value.Filter[labelIndex] &&
          cacheExtend.value.Filter[labelIndex].Rule.length > keyIndex
        ) {
          cacheExtend.value.Filter[labelIndex].Rule[keyIndex].Op = op;
        }
      }
    }
  }

  const updateFilterTemplate = (e) => {
    currentFilterTemplateName.value = e;
    const template = filterTemplate.value[modelDisplayName.value].filter(
      (i) => i.FilterName === e,
    )[0];
    const newExtend = unformat(template || {});
    newExtend.FilterName = '';
    cacheExtend.value = newExtend;
  };

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

    var currYear = new Date().getFullYear();
    cacheExtend.value.FilterPeriod.push({
      DateStart: currYear + '0101',
      DateEnd: currYear + 10 + '1231',
      TimeStart: '0000',
      TimeEnd: '2359',
      Weekday: '0,1,2,3,4,5,6',
    });
  }

  function ruleAdd(annotator, labelIndex, key) {
    if (cacheExtend.value.Filter) {
      if (cacheExtend.value.Filter.length > labelIndex) {
        const item = cacheExtend.value.Filter[labelIndex];
        const values = getOpsOfKvp(annotator, item.Label, key);
        const ops = values.length <= 0 ? numberOps : selectOps;

        if (!item.Rule) {
          item.Rule = [];
        }

        item.Rule.push({
          Op: ops[0],
          K: key,
          V: values.length > 0 ? values[0] : 0,
        });
      }
    }
  }

  function ruleDeleteByIndex(labelIndex, index) {
    if (cacheExtend.value.Filter) {
      if (cacheExtend.value.Filter.length > labelIndex) {
        cacheExtend.value.Filter[labelIndex].Rule.splice(index, 1);
      }
    }
  }

  function attrDeleteByIndex(labelIndex) {
    if (cacheExtend.value.Filter) {
      if (cacheExtend.value.Filter.length > labelIndex) {
        cacheExtend.value.Filter.splice(labelIndex, 1);
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
    return getOpsOfKvp(modelDisplayName.value, currentLabelId.value, currentKvp.value);
  });

  const labels = computed(() => {
    console.log('computed labels: ', modelDisplayName.value, attrList.value);
    if (modelDisplayName.value && attrList.value && attrList.value[modelDisplayName.value]) {
      const data = attrList.value[modelDisplayName.value];
      return Object.keys(data).map((key) => {
        return {
          key: key,
          value: data[key].label,
        };
      });
    }

    return [];
  });

  const kvps = computed(() => {
    if (attrList.value) {
      if (
        attrList.value[modelDisplayName.value] &&
        attrList.value[modelDisplayName.value][currentLabelId.value]
      ) {
        const kvp =
          attrList.value[modelDisplayName.value] &&
          attrList.value[modelDisplayName.value][currentLabelId.value].kvp;
        return Object.keys(kvp);
      }
    }

    return [];
  });

  const ops = computed(() => {
    if (attrList.value) {
      if (
        attrList.value[modelDisplayName.value] &&
        attrList.value[modelDisplayName.value][currentLabelId.value]
      ) {
        const kvp =
          attrList.value[modelDisplayName.value] &&
          attrList.value[modelDisplayName.value][currentLabelId.value].kvp;
        if (kvp && kvp[currentKvp.value]) {
          const k = kvp[currentKvp.value];
          if (k && k.length <= 0) {
            return numberOps;
          }
        }
      }
    }

    return selectOps;
  });

  const modelDisplayName = computed(() => {
    if (props.algorithmName && props.algorithmName.trim().length > 0) {
      return props.algorithmName;
    }

    return aiModelName.value;
  });

  const modelSelectDisabled = computed(() => {
    if (props.algorithmName && props.algorithmName.trim().length > 0) {
      return true;
    }

    return false;
  });

  const ruleTemplateSelectShow = computed(() => {
    if (props.algorithmName && props.algorithmName.trim().length > 0) {
      return false;
    }

    if (aiModelName.value && aiModelName.value.trim().length > 0) {
      return true;
    }

    return false;
  });

  function attrAdd(labelId) {
    currentLabelId.value = labelId;

    if (!cacheExtend.value.Filter) {
      cacheExtend.value.Filter = [
        {
          Annotator: props.algorithmName,
          Label: labelId,
          Rule: [],
        },
      ];
      return;
    }

    cacheExtend.value.Filter.push({
      Annotator: props.algorithmName,
      Label: labelId,
      Rule: [],
    });
  }

  function nameOfKvpLabel(annotator, label) {
    if (
      attrList.value &&
      attrList.value[annotator] &&
      attrList.value[annotator][label] &&
      attrList.value[annotator][label].label
    ) {
      return attrList.value[annotator][label].label;
    } else {
      return '';
    }
  }

  function format() {
    const cacheE = JSON.parse(JSON.stringify(cacheExtend.value));
    const filter = cacheE.Filter;
    let newFilter = {};
    if (filter && filter.length > 0) {
      filter.forEach((element) => {
        const label = element.Label;
        const annotator = element.Annotator;
        const rule = element.Rule;

        if (!newFilter[label]) {
          newFilter[label] = {
            Annotator: annotator,
            Label: label,
            Rule: [],
          };
        }

        const existFilter = newFilter[label];
        let grp = -1;
        existFilter.Rule.forEach((element) => {
          if (element.Grp > grp) {
            grp = element.Grp;
          }
        });

        grp = grp + 1;

        rule.forEach((element) => {
          element.Grp = grp;
          existFilter.Rule.push(element);
        });
      });
    }

    cacheE.Filter = Object.values(newFilter);

    return cacheE;
  }

  function unformat(originExtend) {
    const origin = JSON.parse(JSON.stringify(originExtend));
    const filter = origin.Filter || [];
    let cache = {};
    filter.forEach((element) => {
      const label = element.Label;
      const annotator = element.Annotator;
      const rule = element.Rule;

      rule.forEach((r) => {
        const key = label + r.Grp;

        if (!cache[key]) {
          cache[key] = {
            Annotator: annotator,
            Label: label,
            Rule: [],
          };
        }

        cache[key].Rule.push(r);
      });
    });

    origin.Filter = Object.values(cache);
    return origin;
  }

  async function Submit() {
    const newValue = format();
    await apis.updateRuleTemplate(modelDisplayName.value, newValue);
    closeModal();
    emit('success');
  }

  watch(getVisible, (newValue, oldValue) => {
    if (!newValue) {
      return;
    }
    aiModelName.value = '';
    currentFilterTemplateName.value = '';

    apis.ruleTemplateList().then((res) => {
      filterTemplate.value = res;
    });

    apis.modelList().then((res) => {
      aiModelNameOptions.value = res.map((item) => item.annotator_name);
    });
  });

  const aiModelSelectChange = (e) => {
    aiModelName.value = e;
    currentFilterTemplateName.value = '';
    cacheExtend.value = {};
  };

  onMounted(() => {
    apis.attrList().then((res) => {
      attrList.value = res;
    });
  });

  watch(
    () => props.extend,
    async (newValue, oldValue) => {
      const value = await newValue;
      cacheExtend.value = unformat(value || {});
    },
  );

  const props = defineProps({
    algorithmName: {
      type: String,
      default: '',
    },
    extend: {
      type: Object,
      default: {},
    },
  });

  const cacheExtend = ref({});
</script>
