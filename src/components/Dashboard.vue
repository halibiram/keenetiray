<script setup lang="ts">
import { useMainStore } from '@/stores/main';
import { storeToRefs } from 'pinia';
import { Power, PowerOff, Zap, Wifi, WifiOff, Timer, ArrowUp, ArrowDown, RefreshCw, TestTubeDiagonal } from 'lucide-vue-next';

const mainStore = useMainStore();
const { proxyRunning, connected, uptime, traffic, ping, loading } = storeToRefs(mainStore);

const { startProxy, stopProxy, fetchStatus } = mainStore;

function handleToggle() {
  if (proxyRunning.value) {
    stopProxy();
  } else {
    startProxy();
  }
}
</script>

<template>
  <div>
    <!-- Main Start/Stop Toggle -->
    <div class="text-center mb-8">
      <button
        @click="handleToggle"
        :disabled="loading"
        :class="[
          'w-32 h-32 sm:w-40 sm:h-40 rounded-full text-white font-bold text-2xl sm:text-3xl shadow-2xl transition-all duration-300 ease-in-out transform hover:scale-105 flex flex-col items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed',
          proxyRunning ? 'bg-success hover:bg-green-500' : 'bg-danger hover:bg-red-500',
        ]"
      >
        <component :is="proxyRunning ? Power : PowerOff" :size="48" class="mb-2"/>
        <span>{{ proxyRunning ? 'ON' : 'OFF' }}</span>
      </button>
    </div>

    <!-- Status Grid -->
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-4 text-center">
      <div class="bg-white/20 p-4 rounded-lg shadow-md">
        <h4 class="font-semibold text-secondary/80 flex items-center justify-center"><Zap class="mr-2" :size="16"/>Proxy Durumu</h4>
        <p class="text-2xl font-bold" :class="[proxyRunning ? 'text-success' : 'text-danger']">
          {{ proxyRunning ? 'Çalışıyor' : 'Durduruldu' }}
        </p>
      </div>
      <div class="bg-white/20 p-4 rounded-lg shadow-md">
        <h4 class="font-semibold text-secondary/80 flex items-center justify-center"><component :is="connected ? Wifi : WifiOff" class="mr-2" :size="16"/>Bağlantı Durumu</h4>
        <p class="text-2xl font-bold" :class="[connected ? 'text-success' : 'text-danger']">
          {{ connected ? 'Bağlı' : 'Kopuk' }}
        </p>
        <span v-if="connected" class="text-sm text-gray-600">Ping: {{ ping }}ms</span>
      </div>
      <div class="bg-white/20 p-4 rounded-lg shadow-md">
        <h4 class="font-semibold text-secondary/80 flex items-center justify-center"><Timer class="mr-2" :size="16"/>Çalışma Süresi</h4>
        <p class="text-2xl font-bold">{{ uptime }}</p>
      </div>
      <div class="bg-white/20 p-4 rounded-lg shadow-md">
        <h4 class="font-semibold text-secondary/80">Trafik İstatistiği</h4>
        <div class="flex justify-center items-center text-lg font-semibold">
          <ArrowUp class="text-blue-500" :size="20"/> {{ traffic.upload }}
          <ArrowDown class="text-purple-500 ml-4" :size="20"/> {{ traffic.download }}
        </div>
      </div>
    </div>

    <!-- Action Buttons -->
    <div class="flex flex-col sm:flex-row justify-center items-center space-y-4 sm:space-y-0 sm:space-x-4 mt-8">
      <button @click="fetchStatus" :disabled="loading" class="w-full sm:w-auto bg-primary text-white font-semibold py-2 px-4 rounded-lg shadow-md hover:bg-primary/80 transition-colors flex items-center justify-center disabled:opacity-50">
        <RefreshCw class="mr-2" :size="16"/> Yenile
      </button>
      <button :disabled="loading" class="w-full sm:w-auto bg-secondary text-white font-semibold py-2 px-4 rounded-lg shadow-md hover:bg-secondary/80 transition-colors flex items-center justify-center disabled:opacity-50">
        <TestTubeDiagonal class="mr-2" :size="16"/> Test Bağlantısı
      </button>
    </div>
  </div>
</template>
