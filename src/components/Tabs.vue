<script setup lang="ts">
import { RouterLink } from 'vue-router';
import { Cog, Download, BarChart2, FileText, Info } from 'lucide-vue-next';

const tabs = [
  { name: 'Konfigürasyon', icon: Cog, path: '/configuration' },
  { name: 'İçe Aktar', icon: Download, path: '/import' },
  { name: 'İstatistikler', icon: BarChart2, path: '/stats' },
  { name: 'Loglar', icon: FileText, path: '/logs' },
  { name: 'Hakkında', icon: Info, path: '/about' },
];
</script>

<template>
  <div class="mb-6">
    <nav class="flex space-x-2 sm:space-x-4 border-b-2 border-white/20">
      <RouterLink
        v-for="tab in tabs"
        :key="tab.name"
        :to="tab.path"
        class="flex items-center px-3 py-3 text-sm font-medium text-secondary/80 hover:text-primary transition-colors"
        active-class="border-b-2 border-primary text-primary"
      >
        <component :is="tab.icon" class="mr-2" :size="18" />
        <span class="hidden sm:inline">{{ tab.name }}</span>
      </RouterLink>
    </nav>
  </div>
</template>

<style scoped>
/* Scoped styles for any specific adjustments if needed */
.router-link-exact-active {
  border-bottom-width: 2px;
  --tw-border-opacity: 1;
  border-color: var(--primary);
  color: var(--primary);
}
</style>
