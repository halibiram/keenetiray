<script setup>
import { ref, onMounted } from 'vue'

const servers = ref([])
const newServer = ref({ name: '', address: '', port: null })

async function fetchServers() {
  try {
    const response = await fetch('/api/servers')
    if (!response.ok) throw new Error('Network response was not ok')
    servers.value = await response.json()
  } catch (error) {
    console.error('Failed to fetch servers:', error)
    alert('Error fetching servers. Is the backend running?')
  }
}

async function addServer() {
  if (!newServer.value.name || !newServer.value.address || !newServer.value.port) {
    alert('Please fill out all fields.');
    return;
  }
  try {
    const response = await fetch('/api/servers', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        ...newServer.value,
        port: parseInt(newServer.value.port, 10) // Ensure port is an integer
      }),
    })
    if (!response.ok) throw new Error('Failed to add server')

    // As the mock backend always returns the same ID, we just refetch.
    // In a real app, the backend would return the newly created object.
    newServer.value = { name: '', address: '', port: null }
    fetchServers()

  } catch (error) {
    console.error('Failed to add server:', error)
    alert('Error adding server. Is the backend running?')
  }
}

onMounted(fetchServers)
</script>

<template>
  <div id="app">
    <header>
      <h1>V2Ray Keenetic Control Panel</h1>
    </header>

    <main>
      <section class="server-list">
        <h2>Server List</h2>
        <ul>
          <li v-for="server in servers" :key="server.id">
            <strong>{{ server.name }}</strong> ({{ server.address }}:{{ server.port }})
          </li>
        </ul>
        <p v-if="!servers || servers.length === 0">No servers configured.</p>
      </section>

      <hr>

      <section class="add-server">
        <h2>Add New Server</h2>
        <form @submit.prevent="addServer">
          <input type="text" v-model="newServer.name" placeholder="Server Name" required>
          <input type="text" v-model="newServer.address" placeholder="Address (e.g., domain.com)" required>
          <input type="number" v-model="newServer.port" placeholder="Port (e.g., 443)" required>
          <button type="submit">Add Server</button>
        </form>
      </section>
    </main>
  </div>
</template>

<style>
:root {
  font-family: Inter, system-ui, Avenir, Helvetica, Arial, sans-serif;
  line-height: 1.5;
  font-weight: 400;
  color-scheme: light dark;
  color: rgba(255, 255, 255, 0.87);
  background-color: #242424;
}

#app {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
}

hr {
  margin: 2rem 0;
  border-color: #444;
}

section {
  margin-bottom: 2rem;
}

form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  max-width: 400px;
}

input {
  padding: 0.8rem;
  border-radius: 4px;
  border: 1px solid #555;
}

button {
  padding: 0.8rem;
  border-radius: 4px;
  border: 1px solid transparent;
  cursor: pointer;
  background-color: #1a1a1a;
  transition: border-color 0.25s;
}
button:hover {
  border-color: #646cff;
}
</style>
