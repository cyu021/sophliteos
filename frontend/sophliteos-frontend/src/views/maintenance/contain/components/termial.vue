<template>
  <div>
    <div class="terminal-wrap">
      <div id="terminal" ref="terminal"></div>
    </div>
  </div>
</template>
<script lang="ts" setup>
  import { defineProps, onMounted, ref, watch } from 'vue';
  import { Terminal } from 'xterm';
  import { FitAddon } from 'xterm-addon-fit';
  import 'xterm/css/xterm.css';
  import { useUserStoreWithOut } from '/@/store/modules/user';

  const props = defineProps(['containerId', 'isCloseDrawer']);
  const userStore = useUserStoreWithOut();
  const token = userStore.getToken;
  const terminal = ref<Terminal | null>(null);
  const terminalCols = ref<Number>(1580);
  const terminalRows = ref<Number>(80);
  const pageHeight = window.innerHeight;
  if (pageHeight < 720) {
    terminalRows.value = 30;
  } else {
    terminalRows.value = parseInt(parseInt(pageHeight - 720) / 15) + 30; // termial输入框高度设置
  }

  const term = new Terminal({
    rendererType: 'canvas',
    windowsMode: true, // 根据窗口换行
    cursorStyle: 'underline', //光标样式
    cursorBlink: true, //光标闪烁
    rows: terminalRows.value, // 必须和后端一致
    cols: terminalCols.value,
    theme: {
      foreground: '#ECECEC', //字体
      background: '#000000', //背景色
      cursor: 'help', //设置光标
      lineHeight: 20,
    },
    terminalType: 'xterm-256color',
    encoding: 'UTF-8',
  });
  const fitAddon = new FitAddon();
  const socket = ref(null);
  let socketUri = `ws://${window.location.host}/api/terminals/docker/ssh?`;
  let shouldExecuteContainerIdLogic = true; //控制两个watch 初始化互斥执行

  // 监控 containerId 的变化
  watch(
    () => props.containerId,
    () => {
      if (shouldExecuteContainerIdLogic) {
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
      } else if (!shouldExecuteContainerIdLogic) {
        setTimeout(() => {
          socket.value = null;
          initSocket();
          term.focus();
        }, 300);
      }
    },
  );

  onMounted(async () => {
    shouldExecuteContainerIdLogic = true;
    await initTerm();
  });

  function initTerm() {
    if (terminal.value) {
      setTimeout(() => {
        term.loadAddon(fitAddon);
        term.open(terminal.value);
        fitAddon.fit();
      }, 100);
    }
    term.focus();
    terminal.value = term;
  }
  term.onData((data) => {
    // 在这里处理用户输入，你可以发送给SSH服务器等
    socket.value.send(sendProcess(data));
  });
  async function initSocket() {
    const socketApi =
      socketUri +
      `cols=${fitAddon.proposeDimensions()?.cols}&rows=${terminalRows.value}&containerid=${
        props.containerId
      }&user=&command=/bin/bash`;
    await (socket.value = new WebSocket(socketApi, [], {
      headers: {
        Authorization: token,
      },
    }));
    // 监听 WebSocket 收到消息事件
    socket.value.addEventListener('message', (event) => {
      term.write(decodeBase64Data(event.data));
    });

    // 监听 WebSocket 错误事件
    socket.value.addEventListener('error', (event) => {
      term.write('连接建立失败,错误信息', event);
    });
    // 监听 关闭事件
    socket.value.addEventListener('close', () => {
      console.log('连接已断开');
    });
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
  // 清socket
  function handleSocketCleanup() {
    shouldExecuteContainerIdLogic = false;
    if (
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
</script>
<style scoped lang="less">
  .terminal-wrap {
    margin-top: 1px;
    width: 100%;
  }

  #terminal {
    height: 70vh;
    width: 100%;
    overflow: hidden;
  }
</style>
