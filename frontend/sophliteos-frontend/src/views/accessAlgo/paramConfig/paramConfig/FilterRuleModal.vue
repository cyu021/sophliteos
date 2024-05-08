<template>
    <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="Submit" width="60%" :height="1000">
      <div>#任务名称 + 算法包名称</div>
      <Divider />

      <div style="padding: 14px; display: flex; flex-direction: row; align-items: center;">
        <Input v-model:value="name" placeholder="规则名称" />
      </div>

      <div style="padding: 14px; padding-bottom: 30px;">
        <Collapse v-model:activeKey="attributeActiveKey">
          <CollapsePanel key="1" header="Attribute conditions">
            <ul>
              <li style="padding: 14px;" v-for="(item, index) in attrConditions" :key="index">
                <div style="padding: 14px; border: 1px black dashed; display: flex; align-items: flex-start; flex-direction: row;">
                  <div style="width: 150px; padding: 48px; align-items: center; justify-content: center; display: flex;">
                    <Dropdown>
                      <Button type="primary">
                        {{ item.name }}
                        <DownOutlined/>
                      </Button>
                      <template #overlay>
                        <Menu>
                          <MenuItem>
                            <div>face</div>
                          </MenuItem>
                          <MenuItem>
                            <div>full_structure</div>
                          </MenuItem>
                          <MenuItem>
                            <div>fire</div>
                          </MenuItem>
                        </Menu>
                      </template>
                    </Dropdown>
                  </div>
                  <div style="display: flex; flex-direction: column; flex-grow: 1; padding-left: 48px;">
                    <ul>
                      <li 
                        style="align-items: center; height: 48px; justify-content: center; display: flex; margin: 8px; border: 1px gray dashed; flex-grow: 1; flex-direction: row;" 
                        v-for="(jtem, jndex) in item.values" :key="jndex">
                        <div style="padding: 4px;">
                          <Dropdown>
                            <Button type="primary" style="width: 250px; ">
                              {{ jtem.key }}
                              <DownOutlined/>
                            </Button>
                            <template #overlay>
                              <Menu>
                                <MenuItem>
                                  <div>face</div>
                                </MenuItem>
                                <MenuItem>
                                  <div>full_structure</div>
                                </MenuItem>
                                <MenuItem>
                                  <div>fire</div>
                                </MenuItem>
                              </Menu>
                            </template>
                          </Dropdown>
                        </div>
                        <div style="padding: 4px;">
                          <Dropdown>
                            <Button type="primary" style="width: 150px; ">
                              {{ jtem.opt }}
                              <DownOutlined/>
                            </Button>
                            <template #overlay>
                              <Menu>
                                <MenuItem>
                                  <div>face</div>
                                </MenuItem>
                                <MenuItem>
                                  <div>full_structure</div>
                                </MenuItem>
                                <MenuItem>
                                  <div>fire</div>
                                </MenuItem>
                              </Menu>
                            </template>
                          </Dropdown>
                        </div>
                        <div style="padding: 4px;">
                          <Input v-if="typeof jtem.value === 'number'" style="width: 250px;" v-model:value="jtem.value" />
                          <Dropdown v-else>
                            <Button type="primary" style="width: 250px; ">
                              {{ jtem.value }}
                              <DownOutlined style="opacity: 0.5;"/>
                            </Button>
                            <template #overlay>
                              <Menu>
                                <MenuItem>
                                  <div>face</div>
                                </MenuItem>
                                <MenuItem>
                                  <div>full_structure</div>
                                </MenuItem>
                                <MenuItem>
                                  <div>fire</div>
                                </MenuItem>
                              </Menu>
                            </template>
                          </Dropdown>
                        </div>
                        <div>
                          <Button danger type="text" @click="attrConditionsDelete(item.name, jndex)">
                            <DeleteOutlined/>
                          </Button>
                        </div>
                      </li>
                    </ul>
                    <div style="display: flex; align-items: center; justify-content: center;">
                      <Button type="primary" @click="attrConditionsAdd(item.name)">AND</Button>
                    </div>
                  </div>
                </div>
              </li>
            </ul>
            <div style="padding: 16px;">
              <Button type="primary" @click="attrConditionsOr">OR</Button>
            </div>
          </CollapsePanel>
          <CollapsePanel key="2" header="Time conditions">
            <p>123</p>
          </CollapsePanel>
        </Collapse>
      </div>

    </BasicModal>
  </template>
  
  <script lang="ts" setup>
    import { Divider, Input, Collapse, CollapsePanel, Dropdown, Button, Menu, MenuItem, Space, } from 'ant-design-vue';
    import { BasicModal, useModalInner } from '/@/components/Modal';
    import { useI18n } from '/@/hooks/web/useI18n';
    import { onMounted, ref, computed, } from 'vue';
    import apis from './api';
    import { DeleteOutlined, DownOutlined } from '@ant-design/icons-vue';

    const { t } = useI18n();
    const title = t('paramConfig.param.filterRule');
    const emit = defineEmits(['success', 'register']);

    const name = ref<string>('')
    const attributeActiveKey = ref(['1', '2'])

    const attrList = ref()

    const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
      setModalProps({ confirmLoading: false });

    });

    function attrConditionsAdd(name) {
      attrConditions.value.forEach((item) => {
        if (item.name === name) {
          item.values.push({
            key: '',
            opt: '',
            value: 0
          })
        }
      })
    }

    function attrConditionsDelete(name, index) {
      attrConditions.value.forEach((item) => {
        if (item.name === name) {
          item.values.splice(index, 1)
        }
      })
    }

    function attrConditionsOr() {
      attrConditions.value.push({
        label: '',
        name: '',
        values: [{
          key: '',
          opt: '',
          value: 0
        }]
      })
    }

    async function Submit() {
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
        // emit('success');
        // closeModal();
      } finally {
        setModalProps({ confirmLoading: false });
      }
    }

    const attrConditions = ref([
      {
        label: '12345',
        name: 'fire',
        values: [
          {
            key: 'age',
            opt: '>',
            value: 12,
          },
          {
            key: 'gender',
            opt: '=',
            value: 'male',
          }
        ]
      },
      {
        label: '12346',
        name: 'full_structure',
        values: [
          {
            key: 'color',
            opt: '=',
            value: 'red',
          },
          {
            key: 'size',
            opt: '>',
            value: 150,
          }
        ]
      },
    ])

    onMounted(async ()=> {
      console.log('on mounted')
      apis.attrList().then((res) => {
        attrList.value = res;
      });
    })
  </script>
  