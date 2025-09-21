import { defineStore } from 'pinia';
import * as api from '@/api/mock';

export const useMainStore = defineStore('main', {
  state: () => ({
    proxyRunning: false,
    connected: false,
    uptime: '00:00:00',
    traffic: {
      upload: '0 GB',
      download: '0 GB',
    },
    ping: 0,
    config: {
      serverUrl: '',
      localPort: 1080,
      proxyMode: 'pac',
      dnsServers: '',
      rules: {
        bypassChina: false,
        blockAds: false,
      },
      logLevel: 'info',
      bufferSize: 0,
    },
    loading: false,
  }),
  actions: {
    async fetchStatus() {
      this.loading = true;
      const status = await api.getStatus();
      this.proxyRunning = status.running;
      this.connected = status.connected;
      this.uptime = status.uptime;
      this.traffic = {
        upload: (status.traffic.upload / 1024).toFixed(2) + ' GB',
        download: (status.traffic.download / 1024).toFixed(2) + ' GB',
      };
      this.ping = status.ping;
      this.loading = false;
    },
    async startProxy() {
      this.loading = true;
      await api.startProxy();
      await this.fetchStatus(); // Refresh status after action
      this.loading = false;
    },
    async stopProxy() {
      this.loading = true;
      await api.stopProxy();
      await this.fetchStatus(); // Refresh status after action
      this.loading = false;
    },
    async fetchConfig() {
      this.loading = true;
      const fetchedConfig = await api.getConfig();
      this.config = fetchedConfig;
      this.loading = false;
    },
    async saveConfig(newConfig) {
      this.loading = true;
      const updatedConfig = await api.updateConfig(newConfig);
      this.config = updatedConfig.config;
      this.loading = false;
    },
  },
});
