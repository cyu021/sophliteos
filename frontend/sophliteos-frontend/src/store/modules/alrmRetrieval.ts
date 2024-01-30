import { defineStore } from 'pinia';

export const alarmInfo = defineStore('info', {
  state: () => ({
    iamgeInfo: {
      type: Number(),
      time: Number(),
      itemsInBox: [{ confidence: '' }],
      deviceName: '',
      box: [
        {
          width: 0,
          height: 0,
          x: 0,
          y: 0,
        },
      ],
      bigImage: '',
      smallImage: '',
      taskId: '',
      srcId: '',
      frameIndex: '',
      Extend: {},
    },
  }),
  actions: {
    setInfo(value) {
      this.iamgeInfo = value;
    },
  },
  getters: {
    // 获取数据的 getter
    getInfo(): any {
      return this.iamgeInfo;
    },
  },
});
