<script setup lang="ts">
import { onMounted } from 'vue';
import { RouterView } from 'vue-router';
import Tabs from '@/components/Tabs.vue';
import { useMainStore } from './stores/main';
import { storeToRefs } from 'pinia';

const mainStore = useMainStore();
const { loading } = storeToRefs(mainStore);

onMounted(() => {
  mainStore.fetchStatus();
  mainStore.fetchConfig();
});
</script>

<template>
  <div class="relative min-h-screen text-white">
    <!-- Loading Overlay -->
    <div v-if="loading" class="absolute inset-0 bg-black/50 z-50 flex items-center justify-center">
      <div class="w-16 h-16 border-4 border-white border-t-transparent rounded-full animate-spin"></div>
    </div>

    <header class="p-4 sm:p-6 md:p-8">
      <div class="container mx-auto max-w-5xl">
        <h1 class="text-4xl sm:text-5xl font-bold">
          🚀 K2Ray
        </h1>
        <p class="text-lg sm:text-xl text-white/80">
          Keenetic Router V2Ray Client
        </p>
      </div>
    </header>

    <main class="p-4 sm:p-6 md:p-8 -mt-16">
      <div class="container mx-auto max-w-5xl">
        <div class="bg-card/95 backdrop-blur-lg rounded-xl shadow-lg p-6 text-primary border border-white/30">
          <Tabs />
          <RouterView />
        </div>
      </div>
    </main>
  </div>
</template>
