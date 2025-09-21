<script setup lang="ts">
import { ref } from 'vue';
import { useMainStore } from '@/stores/main';
import { Download, QrCode, Link } from 'lucide-vue-next';

const mainStore = useMainStore();
const activeTab = ref('url');
const manualUrl = ref('');
const parsedPreview = ref(null);

function parseUrl() {
  // Mock parsing logic
  if (manualUrl.value.startsWith('vmess://')) {
    parsedPreview.value = {
      protocol: 'vmess',
      address: 'mock.server.com',
      port: 443,
      id: 'mock-uuid',
      remarks: 'Mock Server'
    };
  } else {
    parsedPreview.value = { error: 'Invalid or unsupported URL format.' };
  }
}

function saveFromUrl() {
  if (parsedPreview.value && !parsedPreview.value.error) {
    mainStore.updateConfig({ serverUrl: manualUrl.value });
    alert('Configuration saved from URL!');
  } else {
    alert('Cannot save. Please parse a valid URL first.');
  }
}
</script>

<template>
  <div>
    <h2 class="text-2xl font-bold mb-6 flex items-center"><Download class="mr-3" :size="24"/>İçe Aktar</h2>

    <!-- Inner Tab Navigation -->
    <div class="flex border-b mb-6">
      <button @click="activeTab = 'url'" :class="['px-4 py-2', activeTab === 'url' ? 'border-b-2 border-primary text-primary' : 'text-secondary/80']">
        <Link class="inline-block mr-2" :size="16"/>Manuel URL
      </button>
      <button @click="activeTab = 'qr'" :class="['px-4 py-2', activeTab === 'qr' ? 'border-b-2 border-primary text-primary' : 'text-secondary/80']">
        <QrCode class="inline-block mr-2" :size="16"/>QR Kod
      </button>
      <button @click="activeTab = 'sub'" :class="['px-4 py-2', activeTab === 'sub' ? 'border-b-2 border-primary text-primary' : 'text-secondary/80']">
        <Link class="inline-block mr-2" :size="16"/>Abonelik
      </button>
    </div>

    <!-- Tab Content -->
    <div>
      <!-- Manual URL Import -->
      <div v-if="activeTab === 'url'" class="space-y-4">
        <div>
          <label for="manualUrl" class="block text-sm font-medium text-secondary">Yapılandırma URL'si</label>
          <textarea id="manualUrl" v-model="manualUrl" rows="4" class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-primary focus:border-primary sm:text-sm" placeholder="vmess://..."></textarea>
        </div>
        <button @click="parseUrl" class="bg-primary text-white font-semibold py-2 px-4 rounded-lg shadow-md hover:bg-primary/80">
          Ayrıştır ve Önizle
        </button>
        <div v-if="parsedPreview" class="p-4 bg-gray-100 rounded-lg">
          <h4 class="font-semibold mb-2">Önizleme</h4>
          <pre class="text-sm bg-k-dark-gray text-white p-2 rounded"><code>{{ JSON.stringify(parsedPreview, null, 2) }}</code></pre>
        </div>
        <button v-if="parsedPreview && !parsedPreview.error" @click="saveFromUrl" class="bg-success text-white font-semibold py-2 px-4 rounded-lg shadow-md hover:bg-success/80">
          Yapılandırmayı Kaydet
        </button>
      </div>

      <!-- QR Code Scanner -->
      <div v-if="activeTab === 'qr'">
        <div class="p-4 border rounded-lg text-center">
          <p class="text-secondary mb-4">QR kod tarayıcı özelliği yakında eklenecektir.</p>
          <div class="w-full h-48 bg-gray-200 rounded-md flex items-center justify-center">
            <p class="text-gray-500">Kamera Önizlemesi</p>
          </div>
        </div>
      </div>

      <!-- Subscription Import -->
      <div v-if="activeTab === 'sub'">
        <div class="p-4 border rounded-lg text-center">
          <p class="text-secondary">Abonelik ile içe aktarma özelliği yakında eklenecektir.</p>
        </div>
      </div>
    </div>
  </div>
</template>
