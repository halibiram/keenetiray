<script setup lang="ts">
import { ref, computed } from 'vue';
import { FileText, Search, Download, ChevronsDown } from 'lucide-vue-next';

const logs = ref([
  { level: 'INFO', message: 'Proxy service starting...', timestamp: '2023-10-27 10:00:00' },
  { level: 'INFO', message: 'Connecting to server vmess://mock.server.com...', timestamp: '2023-10-27 10:00:01' },
  { level: 'SUCCESS', message: 'Connection established. Latency: 120ms', timestamp: '2023-10-27 10:00:02' },
  { level: 'WARN', message: 'High memory usage detected: 85%', timestamp: '2023-10-27 10:05:00' },
  { level: 'ERROR', message: 'Failed to authenticate with server. Check your credentials.', timestamp: '2023-10-27 10:10:00' },
  { level: 'INFO', message: 'Retrying connection...', timestamp: '2023-10-27 10:10:05' },
]);

const logLevelFilter = ref('ALL');
const searchTerm = ref('');
const autoScroll = ref(true);

const filteredLogs = computed(() => {
  return logs.value.filter(log => {
    const levelMatch = logLevelFilter.value === 'ALL' || log.level === logLevelFilter.value;
    const searchTermMatch = log.message.toLowerCase().includes(searchTerm.value.toLowerCase());
    return levelMatch && searchTermMatch;
  });
});

const getLogLevelClass = (level: string) => {
  switch (level) {
    case 'ERROR': return 'text-danger';
    case 'WARN': return 'text-warning';
    case 'SUCCESS': return 'text-success';
    default: return 'text-blue-400';
  }
};
</script>

<template>
  <div>
    <h2 class="text-2xl font-bold mb-6 flex items-center"><FileText class="mr-3" :size="24"/>Loglar</h2>

    <!-- Log Controls -->
    <div class="flex flex-wrap items-center gap-4 mb-4 p-4 bg-white/10 rounded-lg">
      <div class="flex-grow">
        <label for="search" class="sr-only">Search Logs</label>
        <div class="relative">
          <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <Search class="h-5 w-5 text-gray-400" />
          </div>
          <input type="text" v-model="searchTerm" id="search" class="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-white text-gray-900 placeholder-gray-500 focus:outline-none focus:placeholder-gray-400 focus:ring-1 focus:ring-primary focus:border-primary sm:text-sm" placeholder="Search logs...">
        </div>
      </div>
      <div>
        <label for="logLevel" class="sr-only">Log Level</label>
        <select v-model="logLevelFilter" id="logLevel" class="rounded-md border-gray-300 shadow-sm focus:border-primary focus:ring-primary">
          <option>ALL</option>
          <option>INFO</option>
          <option>WARN</option>
          <option>ERROR</option>
          <option>SUCCESS</option>
        </select>
      </div>
      <div class="flex items-center">
        <input type="checkbox" v-model="autoScroll" id="autoScroll" class="h-4 w-4 text-primary border-gray-300 rounded focus:ring-primary">
        <label for="autoScroll" class="ml-2 text-sm text-secondary">Auto-scroll</label>
      </div>
      <button class="bg-secondary text-white font-semibold py-2 px-4 rounded-lg shadow-md hover:bg-secondary/80 flex items-center">
        <Download class="mr-2" :size="16"/> Download Logs
      </button>
    </div>

    <!-- Log Viewer -->
    <div class="bg-k-dark-gray text-white font-mono text-sm rounded-lg p-4 h-96 overflow-y-auto">
      <div v-for="(log, index) in filteredLogs" :key="index" class="flex">
        <span class="text-gray-500 mr-4">{{ log.timestamp }}</span>
        <span :class="['font-bold w-20', getLogLevelClass(log.level)]">[{{ log.level }}]</span>
        <span>{{ log.message }}</span>
      </div>
    </div>
  </div>
</template>
