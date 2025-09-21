<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'

interface Server {
  id: string;
  name: string;
  protocol: 'vmess' | 'vless' | 'shadowsocks' | 'trojan';
  address: string;
  port: number;
  userID: string;
  alterID: number;
  security: string;
  network: 'tcp' | 'ws' | 'grpc';
  wsPath: string;
  grpcSvcName: string;
  tls: string;
  sni: string;
}

const servers = ref<Server[]>([])
const serviceStatus = ref('Not Running')

const newServer = ref<Omit<Server, 'id'>>({
  name: '',
  protocol: 'vmess',
  address: '',
  port: 443,
  userID: '',
  alterID: 0,
  security: 'auto',
  network: 'ws',
  wsPath: '/',
  grpcSvcName: '',
  tls: 'tls',
  sni: ''
})

async function apiCall(url: string, options?: RequestInit) {
  try {
    const response = await fetch(url, options)
    if (!response.ok) {
      const errorBody = await response.json()
      throw new Error(errorBody.error || 'Network response was not ok')
    }
    return response
  } catch (error) {
    console.error('API call failed:', error)
    if (error instanceof Error) {
      alert(`Error: ${error.message}`)
    } else {
      alert('An unknown error occurred')
    }
    throw error
  }
}

async function fetchServers() {
  const response = await apiCall('/api/servers')
  servers.value = await response.json()
}

async function fetchStatus() {
  const response = await apiCall('/api/service/status')
  const data = await response.json()
  serviceStatus.value = data.status
}

async function addServer() {
  const serverData = { ...newServer.value }
  await apiCall('/api/servers', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(serverData),
  })
  fetchServers()
}

async function deleteServer(id: string) {
  if (!confirm('Are you sure you want to delete this server?')) return
  await apiCall(`/api/servers/${id}`, { method: 'DELETE' })
  fetchServers()
}

async function startService(id: string) {
  await apiCall('/api/service/start', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ id }),
  })
  fetchStatus()
}

async function stopService() {
  await apiCall('/api/service/stop', { method: 'POST' })
  fetchStatus()
}

onMounted(() => {
  fetchServers()
  fetchStatus()
})
</script>

<template>
  <div id="app">
    <header>
      <h1>V2Ray Keenetic Control Panel</h1>
      <div class="status-bar">
        <span>Service Status: <strong>{{ serviceStatus }}</strong></span>
        <button @click="stopService" class="stop-btn">Stop Service</button>
      </div>
    </header>

    <main>
      <section class="server-list">
        <h2>Server List</h2>
        <div v-if="!servers || servers.length === 0" class="no-servers">No servers configured.</div>
        <ul v-else>
          <li v-for="server in servers" :key="server.id">
            <div class="server-info">
              <strong>{{ server.name }}</strong>
              <span>{{ server.protocol }} @ {{ server.address }}:{{ server.port }}</span>
            </div>
            <div class="server-actions">
              <button @click="startService(server.id)" class="start-btn">Start</button>
              <button @click="deleteServer(server.id)" class="delete-btn">Delete</button>
            </div>
          </li>
        </ul>
      </section>

      <hr>

      <section class="add-server">
        <h2>Add New Server</h2>
        <form @submit.prevent="addServer">
          <input type="text" v-model="newServer.name" placeholder="Name" required>
          <input type="text" v-model="newServer.address" placeholder="Address (e.g., domain.com)" required>
          <input type="number" v-model.number="newServer.port" placeholder="Port" required>
          <input type="text" v-model="newServer.userID" placeholder="User ID (UUID)" required>

          <select v-model="newServer.protocol">
            <option value="vmess">VMess</option>
            <option value="vless" disabled>VLESS (not yet supported)</option>
          </select>

          <select v-model="newServer.network">
            <option value="tcp">TCP</option>
            <option value="ws">WebSocket</option>
            <option value="grpc">gRPC</option>
          </select>

          <input v-if="newServer.network === 'ws'" type="text" v-model="newServer.wsPath" placeholder="WebSocket Path">
          <input v-if="newServer.network === 'grpc'" type="text" v-model="newServer.grpcSvcName" placeholder="gRPC Service Name">

          <input type="text" v-model="newServer.sni" placeholder="SNI (for TLS)">

          <button type="submit">Add Server</button>
        </form>
      </section>
    </main>
  </div>
</template>

<style>
/* Basic styling for layout and readability */
:root { font-family: Inter, system-ui, Avenir, Helvetica, Arial, sans-serif; color-scheme: light dark; color: rgba(255, 255, 255, 0.87); background-color: #242424; }
#app { max-width: 900px; margin: 0 auto; padding: 2rem; }
header { margin-bottom: 2rem; }
.status-bar { display: flex; justify-content: space-between; align-items: center; padding: 1rem; background-color: #2a2a2a; border-radius: 8px; }
hr { margin: 2rem 0; border-color: #444; }
section { margin-bottom: 2rem; }
form { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 1rem; }
input, select, button { padding: 0.8rem; border-radius: 4px; border: 1px solid #555; background-color: #1a1a1a; }
button { cursor: pointer; transition: border-color 0.25s; }
button:hover { border-color: #646cff; }
.server-list ul { list-style: none; padding: 0; }
.server-list li { display: flex; justify-content: space-between; align-items: center; padding: 1rem; background-color: #2a2a2a; border-radius: 8px; margin-bottom: 1rem; }
.server-info { display: flex; flex-direction: column; gap: 0.25rem; }
.server-actions { display: flex; gap: 0.5rem; }
.start-btn { background-color: #004d00; }
.stop-btn { background-color: #6b0000; }
.delete-btn { background-color: #4b0000; }
</style>
