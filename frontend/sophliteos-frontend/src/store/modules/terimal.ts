import { defineStore } from 'pinia';

export const historyTermial = defineStore('historyTermial', {
  state: () => ({
    store: {
      linkIpList: new Array(0),
      socketMap: new Map(),
      maxLinkIp: 5,
    },
  }),
  actions: {
    addIp(ip) {
      this.store[ip] = '';
    },
    removeIp(ip) {
      delete this.store[ip];
    },
    pushContent(ip, content) {
      this.store[ip] += content;
    },
    includesIp(ip) {
      return this.store.hasOwnProperty(ip);
    },
    getContent(ip) {
      return this.store[ip];
    },
    // 添加一个socket,超出最大限制时关闭并删除第一个socket
    addSocketMap(ip, socket) {
      if (this.store.linkIpList.length >= this.store.maxLinkIp) {
        const delIp = this.store.linkIpList[0];
        this.deleteSocket(delIp);
      }
      this.store.socketMap.set(ip, socket);
      this.store.linkIpList.push(ip);
      this.addIp(ip);
    },
    // 页面关闭ip连接时，关闭并删除socket
    deleteSocket(ip) {
      const socket = this.store.socketMap.get(ip);
      socket && socket.close();
      socket && this.store.socketMap.delete(ip);
      this.store.linkIpList = this.store.linkIpList.filter((v) => v !== ip);
      this.includesIp(ip) && this.removeIp(ip);
    },
    // 返回socket
    getSocket(ip) {
      return this.store.socketMap.get(ip);
    },

    // 返回ipList
    getIpList() {
      return this.store.linkIpList;
    },
  },
});
