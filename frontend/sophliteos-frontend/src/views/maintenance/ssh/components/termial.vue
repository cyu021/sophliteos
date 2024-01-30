<template>
  <div>
    <div class="terminal-wrap">
      <div id="terminal" ref="terminal"></div>
    </div>
    <a-dropdown placement="top" v-if="ip">
      <template #overlay>
        <a-menu @click="handleMenuClick">
          <template v-if="CommandList">
            <a-menu-item v-for="command in CommandList" :key="command">
              {{ command.name + `   ` + command.command }}
            </a-menu-item>
          </template>
          <div v-else class="quickCmd-palcehoder">您还未创建快速命令</div>
        </a-menu>
      </template>
      <a-button class="quickDropdownName" postIcon="mingcute:down-fill">
        快速命令，请选择
      </a-button>
    </a-dropdown>
  </div>
</template>
<script lang="ts" setup>
  import { defineProps, nextTick, onMounted, ref, watch, onUnmounted } from 'vue';
  import { Menu, Dropdown } from 'ant-design-vue';
  // import { useI18n } from '/@/hooks/web/useI18n';
  import { historyTermial } from '/@/store/modules/terimal';
  import { Terminal } from 'xterm';
  import { FitAddon } from 'xterm-addon-fit';
  import 'xterm/css/xterm.css';
  import { getCommandList } from '/@/api/maintenance/index';
  // import { useUserStoreWithOut } from '/@/store/modules/user';
  // import { useMessage } from '/@/hooks/web/useMessage';
  // const { t } = useI18n();
  // const { createMessage } = useMessage();
  // const userStore = useUserStoreWithOut();
  // const token = userStore.getToken;
  const CommandList = ref<any[]>([]);
  const props = defineProps({
    ip: {
      type: String,
      default: '',
    },
    core: {
      type: Number,
      default: 0,
    },
    containerId: {
      type: String,
      default: '',
    },
    isCloseDrawer: {
      type: Boolean,
      default: false,
    },
  });
  // 用于判断是否是新建了term，但用的是stroe里的socket
  const isNewTem = ref(false);
  const AMenu = Menu;
  const AMenuItem = Menu.Item;
  const ADropdown = Dropdown;
  const storeHstory = historyTermial();
  const terminal = ref();
  const terminalCols = ref<Number>(1580);
  const terminalRows = ref<Number>(80);
  const pageHeight = window.innerHeight;
  if (pageHeight < 720) {
    // termial输入框高度设置
    terminalRows.value = 30;
  } else if (pageHeight < 920) {
    terminalRows.value = Math.floor(Math.floor(pageHeight - 720) / 16) + 30;
  } else if (pageHeight < 1050) {
    terminalRows.value = Math.floor(Math.floor(pageHeight - 920) / 18) + 42;
  } else if (pageHeight < 1200) {
    terminalRows.value = Math.floor(Math.floor(pageHeight - 1050) / 14) + 49;
  } else {
    terminalRows.value = Math.floor(Math.floor(pageHeight - 1200) / 16) + 59;
  }

  const term = new Terminal({
    cursorStyle: 'block', //光标样式
    cursorBlink: true, //光标闪烁
    rows: terminalRows.value as number, // 必须和后端一致
    cols: terminalCols.value as number,
    theme: {
      foreground: '#ECECEC', //字体
      background: '#000000', //背景色
      cursor: 'help', //设置光标
    },
  });
  const fitAddon = new FitAddon();
  const socket = ref();
  let socketUri = `ws://${window.location.host}/api/terminals`;

  // 监控 ip 的变化
  watch(
    () => props.ip,
    (ip) => {
      handleSocketCleanup();
      const hisSocket = storeHstory.getSocket(ip);
      // 0--全部ip标签关掉是显示的内容
      if (ip === 'null') {
        term.write('请选择主机，建立连接');
      }
      // 1--把之前的socket给页面
      else if (hisSocket) {
        // initSocket();
        socket.value = hisSocket;
        // 如果有以往的message的历史数据，就给term
        storeHstory.includesIp(ip) && term.write(storeHstory.getContent(ip));
        // 如果新建了term,要更新socket.onmessage
        if (isNewTem.value) socket.value.onmessage = onMsg;
        // 一般不会走到else，若initTerm没有在watch之前执行，走else，等一个事件循环看看是否需要更新socket.onmessage
        else {
          nextTick(() => {
            isNewTem.value && (socket.value.onmessage = onMsg);
          });
        }
        resize();
        term.focus();
      }
      // 2--没有存入历史的连接，新建，加入store
      else {
        initSocket();
        storeHstory.addSocketMap(ip, socket.value);
      }
    },
  );

  onMounted(async () => {
    initTerm();
    shouldExecuteContainerIdLogic.value = true;
    window.addEventListener('resize', resize);
  });
  onUnmounted(() => {
    window.removeEventListener('resize', resize);
  });

  term.onData((data) => {
    // 在这里处理用户输入，你可以发送给SSH服务器等
    socket.value && socket.value.send(sendProcess(data));
  });

  function handleMenuClick({ key }) {
    socket.value && socket.value.send(sendProcess(key.command));
    term.focus();
  }

  async function initTerm() {
    const result = await getCommandList();
    CommandList.value = result;
    if (terminal.value) {
      term.loadAddon(fitAddon);
      term.open(terminal.value);
      fitAddon.fit();
      terminalCols.value = term.cols;
      isNewTem.value = true;
    }
    term.focus();
    terminal.value = term;
  }

  async function initSocket() {
    let socketApi = `${socketUri}/ssh?cols=${terminalCols.value}&rows=${terminalRows.value}&ip=${props.ip}`;
    if (props.core) {
      socketApi = `${socketUri}/ssh?cols=${terminalCols.value}&rows=${terminalRows.value}&ip=${
        props.ip.split('-')[1]
      }&core=1`;
    }
    if (props.containerId) {
      socketApi = `${socketUri}/docker/ssh?cols=${terminalCols.value}&rows=${terminalRows.value}&containerid=${props.containerId}&user=&command=/bin/bash`;
    }
    await (socket.value = new WebSocket(socketApi));
    let isFirstMessage = true; // 是否已经收到首个消息
    // 监听 WebSocket 收到消息事件
    socket.value.onmessage = (event) => {
      term.write(decodeBase64Data(event.data) || '');
      if (props.ip) {
        if (!isFirstMessage || (isFirstMessage && storeHstory.getContent(props.ip).length <= 0)) {
          storeHstory.pushContent(props.ip, decodeBase64Data(event.data));
        }
        isFirstMessage = false;
      }
    };
    isNewTem.value = false;

    // 监听 WebSocket 错误事件
    socket.value.addEventListener('error', (event) => {
      term.write('连接建立失败,错误信息', event);
    });
    // 监听 关闭事件
    socket.value.addEventListener('close', () => {
      console.log('监测连接关闭');
    });
    term.focus();
  }

  function onMsg(event, isFirstMessage = false) {
    term.write(decodeBase64Data(event.data) || '');
    if (props.ip) {
      if (!isFirstMessage || (isFirstMessage && storeHstory.getContent(props.ip).length <= 0)) {
        storeHstory.pushContent(props.ip, decodeBase64Data(event.data));
      }
      isFirstMessage = false;
    }
  }

  // 发数据处理
  function sendProcess(message: string) {
    const sendMessage = btoa(unescape(encodeURIComponent(message)));
    return JSON.stringify({ type: 'cmd', data: sendMessage });
  }
  // 收数据处理
  function decodeBase64Data(jsonString) {
    try {
      const parsedData = JSON.parse(jsonString);
      const base64Data = parsedData?.data;
      const decodedData = decodeURIComponent(escape(atob(base64Data)));
      return decodedData;
    } catch (error) {
      return null;
    }
  }

  function resize() {
    fitAddon.fit();
    const params = {
      type: 'resize',
      cols: term.cols,
      rows: term.rows,
    };
    socket.value && socket.value.send(JSON.stringify(params));
  }

  // 清socket
  function handleSocketCleanup() {
    shouldExecuteContainerIdLogic.value = false;
    // 终端页面，不关闭socket，只有关闭标签或打开的标签大于5时，才关闭对应的socket；容器页面需要关闭
    if (
      props.containerId &&
      socket.value &&
      (socket.value.readyState === WebSocket.OPEN ||
        socket.value.readyState === 1 ||
        socket.value.readyState === WebSocket.CONNECTING ||
        socket.value.readyState === 0)
    ) {
      socket.value.close();
    }
    term.clear();
    term.write('\x1b[2J\x1b[H');
    term.blur();
  }

  // 合并容器下的terminal.vue的代码
  const shouldExecuteContainerIdLogic = ref(true); //控制两个watch 初始化互斥执行
  // 监控 containerId 的变化
  watch(
    () => props.containerId,
    () => {
      if (shouldExecuteContainerIdLogic.value) {
        handleSocketCleanup();
        setTimeout(() => {
          socket.value = null;
          initSocket();
          term.focus();
        }, 300);
      }
    },
  );

  // 监听 isCloseDrawer，若关闭弹窗，断开连接，若再次打开相同containerId ssh，重新建立连接
  watch(
    () => props.isCloseDrawer,
    () => {
      if (props.isCloseDrawer) {
        handleSocketCleanup();
      } else if (!shouldExecuteContainerIdLogic.value) {
        setTimeout(() => {
          socket.value = null;
          initSocket();
          term.focus();
        }, 300);
      }
    },
  );
</script>
<style scoped lang="less">
  .terminal-wrap {
    margin-top: 1px;
  }

  #terminal {
    overflow-y: auto;
  }

  .quickDropdownName {
    min-width: 200px;
  }

  .quickCmd-palcehoder {
    min-width: 200px;
    min-height: 45px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: gray;
  }
</style>
