<script setup lang="ts">
import { Info, GitBranch, Users } from 'lucide-vue-next';
</script>

<template>
  <div>
    <h2 class="text-2xl font-bold mb-6 flex items-center"><Info class="mr-3" :size="24"/>Hakkında</h2>

    <div class="space-y-6 text-secondary">
      <div class="p-4 border rounded-lg">
        <h3 class="text-lg font-semibold mb-2 text-primary flex items-center"><GitBranch class="mr-2"/>Proje Hakkında</h3>
        <p>
          K2Ray, Keenetic router'lar için geliştirilmiş modern ve kullanıcı dostu bir V2Ray istemci arayüzüdür.
          Bu arayüz, teknik bilgisi olmayan kullanıcıların bile V2Ray proxy hizmetlerini kolayca yönetebilmesini hedefler.
        </p>
      </div>

      <div class="p-4 border rounded-lg">
        <h3 class="text-lg font-semibold mb-2 text-primary">Versiyon</h3>
        <p>
          Versiyon: 1.0.0 (Beta)
        </p>
      </div>

      <div class="p-4 border rounded-lg">
        <h3 class="text-lg font-semibold mb-2 text-primary flex items-center"><Users class="mr-2"/>Geliştiriciler</h3>
        <p>
          Bu proje, açık kaynak topluluğunun bir ürünüdür. Katkıda bulunan herkese teşekkürler.
        </p>
      </div>
    </div>
  </div>
</template>
