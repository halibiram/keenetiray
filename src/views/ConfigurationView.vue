<script setup lang="ts">
import { useMainStore } from '@/stores/main';
import { storeToRefs } from 'pinia';
import { Save, TestTubeDiagonal, Cog } from 'lucide-vue-next';

const mainStore = useMainStore();
const { config } = storeToRefs(mainStore);

import { onMounted } from 'vue';

const { config, loading } = storeToRefs(mainStore);
const { saveConfig, fetchConfig } = mainStore;

onMounted(() => {
  fetchConfig();
});

async function saveConfiguration() {
  await saveConfig(config.value);
  alert('Configuration saved!');
}

function testConnection() {
  // In a real app, this would trigger a connection test
  alert('Testing connection...');
}
</script>

<template>
  <div>
    <h2 class="text-2xl font-bold mb-6 flex items-center"><Cog class="mr-3" :size="24"/>Konfigürasyon</h2>

    <form @submit.prevent="saveConfiguration" class="space-y-8">
      <!-- Server Konfigürasyonu -->
      <div class="p-4 border rounded-lg">
        <h3 class="text-lg font-semibold mb-4 text-primary">Server Konfigürasyonu</h3>
        <div class="space-y-4">
          <div>
            <label for="serverUrl" class="block text-sm font-medium text-secondary">Server URL (VMess/VLESS/SS/Trojan)</label>
            <input type="text" id="serverUrl" v-model="config.serverUrl" class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-primary focus:border-primary sm:text-sm">
          </div>
          <div>
            <label for="localPort" class="block text-sm font-medium text-secondary">Local SOCKS5 Port</label>
            <input type="number" id="localPort" v-model="config.localPort" class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-primary focus:border-primary sm:text-sm">
          </div>
        </div>
      </div>

      <!-- Proxy Modu -->
      <div class="p-4 border rounded-lg">
        <h3 class="text-lg font-semibold mb-4 text-primary">Proxy Modu</h3>
        <div class="flex space-x-4">
          <label v-for="mode in ['manual', 'transparent', 'pac']" :key="mode" class="flex items-center">
            <input type="radio" :value="mode" v-model="config.proxyMode" name="proxyMode" class="focus:ring-primary h-4 w-4 text-primary border-gray-300">
            <span class="ml-2 text-sm text-secondary">{{ mode.charAt(0).toUpperCase() + mode.slice(1) }}</span>
          </label>
        </div>
      </div>

      <!-- DNS Ayarları -->
      <div class="p-4 border rounded-lg">
        <h3 class="text-lg font-semibold mb-4 text-primary">DNS Ayarları</h3>
        <div>
          <label for="dnsServers" class="block text-sm font-medium text-secondary">DNS Sunucuları (virgülle ayırın)</label>
          <input type="text" id="dnsServers" v-model="config.dnsServers" class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-primary focus:border-primary sm:text-sm">
        </div>
      </div>

      <!-- Routing Rules -->
      <div class="p-4 border rounded-lg">
        <h3 class="text-lg font-semibold mb-4 text-primary">Routing Kuralları</h3>
        <div class="space-y-2">
          <label class="flex items-center">
            <input type="checkbox" v-model="config.rules.bypassChina" class="h-4 w-4 text-primary border-gray-300 rounded focus:ring-primary">
            <span class="ml-2 text-sm text-secondary">Bypass China IPs</span>
          </label>
          <label class="flex items-center">
            <input type="checkbox" v-model="config.rules.blockAds" class="h-4 w-4 text-primary border-gray-300 rounded focus:ring-primary">
            <span class="ml-2 text-sm text-secondary">Block Ads</span>
          </label>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="flex justify-end space-x-4 pt-4">
        <button type="button" @click="testConnection" class="bg-secondary text-white font-semibold py-2 px-4 rounded-lg shadow-md hover:bg-secondary/80 transition-colors flex items-center">
          <TestTubeDiagonal class="mr-2" :size="16"/> Test Et
        </button>
        <button type="submit" class="bg-success text-white font-semibold py-2 px-4 rounded-lg shadow-md hover:bg-success/80 transition-colors flex items-center">
          <Save class="mr-2" :size="16"/> Kaydet
        </button>
      </div>
    </form>
  </div>
</template>
