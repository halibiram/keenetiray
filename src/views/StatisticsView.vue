<script setup lang="ts">
import { BarChart2, Globe, MapPin, ShieldCheck, Wifi } from 'lucide-vue-next';
import { Line } from 'vue-chartjs';
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  LineElement,
  CategoryScale,
  LinearScale,
  PointElement,
  Filler,
} from 'chart.js';
import { ref } from 'vue';

ChartJS.register(Title, Tooltip, Legend, LineElement, CategoryScale, LinearScale, PointElement, Filler);

const chartData = ref({
  labels: Array.from({ length: 10 }, (_, i) => `-${i * 10}s`),
  datasets: [
    {
      label: 'Download',
      borderColor: '#3182ce',
      backgroundColor: 'rgba(49, 130, 206, 0.2)',
      data: Array.from({ length: 10 }, () => Math.random() * 5),
      fill: true,
      tension: 0.4,
    },
    {
      label: 'Upload',
      borderColor: '#38a169',
      backgroundColor: 'rgba(56, 161, 105, 0.2)',
      data: Array.from({ length: 10 }, () => Math.random() * 2),
      fill: true,
      tension: 0.4,
    },
  ],
});

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  scales: {
    y: {
      beginAtZero: true,
      ticks: {
        callback: (value) => `${value.toFixed(1)} Mbps`,
      },
    },
  },
  plugins: {
    legend: {
      position: 'top',
    },
    title: {
      display: true,
      text: 'Real-time Bandwidth Usage',
    },
  },
};

const networkInfo = {
  publicIp: '123.45.67.89',
  location: 'Istanbul, Turkey',
  dnsLeak: 'Not Detected',
  latency: '120ms',
};
</script>

<template>
  <div>
    <h2 class="text-2xl font-bold mb-6 flex items-center"><BarChart2 class="mr-3" :size="24"/>İstatistikler</h2>

    <!-- Real-time Chart -->
    <div class="p-4 border rounded-lg mb-8">
      <h3 class="text-lg font-semibold mb-4 text-primary">Bandwidth Usage</h3>
      <div class="h-64">
        <Line :data="chartData" :options="chartOptions" />
      </div>
    </div>

    <!-- Network Info Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-4 text-center">
      <div class="bg-white/20 p-4 rounded-lg shadow-md">
        <h4 class="font-semibold text-secondary/80 flex items-center justify-center"><Globe class="mr-2" :size="16"/>Public IP</h4>
        <p class="text-xl font-bold">{{ networkInfo.publicIp }}</p>
      </div>
      <div class="bg-white/20 p-4 rounded-lg shadow-md">
        <h4 class="font-semibold text-secondary/80 flex items-center justify-center"><MapPin class="mr-2" :size="16"/>Konum</h4>
        <p class="text-xl font-bold">{{ networkInfo.location }}</p>
      </div>
      <div class="bg-white/20 p-4 rounded-lg shadow-md">
        <h4 class="font-semibold text-secondary/80 flex items-center justify-center"><ShieldCheck class="mr-2" :size="16"/>DNS Leak Test</h4>
        <p class="text-xl font-bold text-success">{{ networkInfo.dnsLeak }}</p>
      </div>
      <div class="bg-white/20 p-4 rounded-lg shadow-md">
        <h4 class="font-semibold text-secondary/80 flex items-center justify-center"><Wifi class="mr-2" :size="16"/>Gecikme</h4>
        <p class="text-xl font-bold">{{ networkInfo.latency }}</p>
      </div>
    </div>
  </div>
</template>
