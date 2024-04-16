<template>
  <div style="padding: 20px; display: flex; flex-direction: column">
    <div style="font-size: 20px; padding-bottom: 20px">选择视频流</div>
    <a-select
      v-model:value="value"
      show-search
      placeholder="Select a video stream"
      style="flex-grow: 1"
      :options="options"
      :filter-option="filterOption"
      @focus="handleFocus"
      @blur="handleBlur"
      @change="handleChange"
    />
  </div>
  <a-divider />

  <div style="display: flex; flex-grow: 1; height: 100vh">
    <div
      style="flex-grow: 1; display: flex; flex-direction: column; padding: 20px; flex: 1"
    >
      <div
        ref="video-container"
        style="margin-top: 20px; position: relative"
        :style="{ width: videoWidth + 'px', height: videoHeight + 'px' }"
      >
        <video
          id="video"
          ref="video"
          @loadedmetadata="adjustVideoSize"
          style="position: absolute; top: 0; left: 0"
        ></video>
        <div
          v-for="(view, index) in views"
          :key="index"
          class="view"
          style="position: absolute; top: 0; left: 0"
        >
          {{ view }}
        </div>
        <canvas
          ref="canvas"
          width="1920"
          height="1080"
          :style="{ width: videoWidth + 'px', height: videoHeight + 'px' }"
          style="position: absolute; top: 0; left: 0"
        ></canvas>
      </div>
    </div>
    <div
      style="flex-grow: 1; display: flex; flex: 1; flex-direction: column; padding: 20px"
    >
      <div style="background-color: white; overflow-y: auto; height: 100vh">
        <ul style="list-style: none; padding: 10px">
          <li
            v-for="(item, index) in itemList"
            :key="index"
            style="
              display: flex;
              flex-direction: row;
              margin-bottom: 10px;
              border: 1px solid #e8e8e8;
              border-radius: 10px;
              height: 140px;
              padding: 10px;
            "
          >
            <div
              style="
                display: flex;
                flex-direction: column;
                align-items: center;
                justify-content: center;
                padding: 10px;
              "
            >
              <img
                style="
                  background-color: gray;
                  object-fit: contain;
                  height: 120px;
                  width: 120px;
                  max-width: 100%;
                  max-height: 100%;
                "
                :src="`data:image/png;base64,${item.ImageBase64}`"
              />
            </div>
            <div
              style="
                flex-grow: 1;
                display: flex;
                flex-direction: column;
                max-width: 100%;
                align-items: flex-start;
                justify-content: center;
              "
            >
              <div>{{ "TaskID: " + item.TaskID }}</div>
              <div>{{ "SrcID: " + item.SrcID }}</div>
              <div>{{ "label: " + item.Extend.label }}</div>
              <div>{{ "Type: " + getTitleOfType(item.Type) }}</div>
            </div>
            <a-popover title="Detail" trigger="click" placement="topLeft">
              <template #content>
                <div style="white-space: pre-line">
                  {{ JSON.stringify(item.Extend, null, 2) }}
                </div>
              </template>
              <a-button type="primary" style="align-self: center">Detail</a-button>
            </a-popover>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, onMounted } from "vue";
import { useI18n } from "/@/hooks/web/useI18n";
import { Divider, Space, Select, Image, Popover } from "ant-design-vue";
import apis from "./apis.js";
import mpegts from "mpegts.js";
import ws from "./ws.js";

const ASelect = Select;
const ADivider = Divider;
const ASpace = Space;
const AImage = Image;
const APopover = Popover;

const views = ref([]);

const videoWidth = ref(800);
const videoHeight = ref(600);

const video = ref();
let player = null;

const itemList = ref([]);
const options = ref([]);

let deviceName = null;

const handleChange = (value) => {
  console.log(`selected ${value}`);
  itemList.value = [];

  deviceName = options.value.filter((item) => item.value == value)[0].label;

  apis.preview({ deviceId: value }).then((res) => {
    play(res);

    ws.connect().then(() => {
      return ws.listen((data) => {
        console.log("onMounted 收到消息:", data, deviceName);

        if (!deviceName) {
          return;
        }

        const filterResult = data.filter((item) => {
          console.log('filter:', item.SrcID, deviceName);
          return item.SrcID == deviceName;
        });

        console.log('filterResult', filterResult);
        if (filterResult.length == 0) {
          return;
        }

        let newList = []
        for (let i = 0; i < filterResult.length; i++) {
          const element = filterResult[i];

          if (element.AnalyzeEvents && element.AnalyzeEvents.length > 0) {
            // 有数据
            const list = element.AnalyzeEvents;
            list.forEach((item) => {
              item.SrcID = element.SrcID;
              item.TaskID = element.TaskID;
            })
            newList = newList.concat(list);
          }          
        }

        const list = newList.concat(itemList.value);
        console.log(list);
        itemList.value = list;

        // const box = data[0].AnalyzeEvents[0].Box;
        // drawRect(box.LeftTopX, box.LeftTopY, box.RightBtmX - box.LeftTopX, box.RightBtmY - box.LeftTopY)
      });
    });
  });
};
const handleBlur = () => {
  console.log("blur");
};
const handleFocus = () => {
  console.log("focus");
};
const filterOption = (input, option) => {
  return false;
  // return option.value.toLowerCase().indexOf(input.toLowerCase()) >= 0;
};
const value = ref(undefined);

const adjustVideoSize = () => {
  videoHeight.value =
    (video.value.videoHeight / video.value.videoWidth) * videoWidth.value;
};

const play = (url) => {
  console.log("play", url);
  
  if (player) {
    player.pause();
    player.unload();
    player = null;
  }

  if (mpegts.getFeatureList().mseLivePlayback) {
    console.log("mseLivePlayback");

    player = mpegts.createPlayer({
      type: "mse", // could also be mpegts, m2ts, flv
      isLive: true,
      url: url,
    });
    player.attachMediaElement(video.value);
    player.load();
    player.play();
  }
};

const canvas = ref();
const drawRect = (x, y, width, height) => {
  const ctx = canvas.value.getContext("2d");

  ctx.clearRect(0, 0, canvas.value.width, canvas.value.height);

  ctx.beginPath();
  ctx.strokeStyle = "red";
  ctx.lineWidth = 2;

  ctx.rect(x, y, width, height);
  ctx.stroke();
};

let types = [];

const getTitleOfType = (type) => {
  console.log(types, type);
  if (types.length == 0) return type;

  return types.filter((item) => item.type == type)[0].name;
};

onMounted(() => {
  console.log("on mounted");

  ws.create(`ws://${window.location.host}/ws`)

  apis.getAbilitesTypes().then((res) => {
    console.log("===>>>>", res);
    types = res;
  });

  apis.videosList().then((res) => {
    console.log('video list:', res)
    options.value = res.map((item) => ({
      value: item.deviceId,
      label: item.name,
    }));
  });
});
</script>
