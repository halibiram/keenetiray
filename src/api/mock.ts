// --- Mock Database ---
let db = {
  status: {
    running: true,
    connected: true,
    uptime: '12:34:56',
    traffic: {
      upload: 1200, // in MB
      download: 3400, // in MB
    },
    ping: 120,
  },
  config: {
    serverUrl: 'vmess://initial-mock-server',
    localPort: 1080,
    proxyMode: 'pac',
    dnsServers: '8.8.8.8, 1.1.1.1',
    rules: {
      bypassChina: true,
      blockAds: true,
    },
    logLevel: 'warning',
    bufferSize: 2048,
  }
};

const SIMULATED_DELAY = 500; // ms

// --- API Functions ---

export const getStatus = () => {
  return new Promise((resolve) => {
    setTimeout(() => {
      db.status.uptime += 1; // Simulate uptime increasing
      resolve({ ...db.status });
    }, SIMULATED_DELAY);
  });
};

export const startProxy = () => {
  return new Promise((resolve) => {
    setTimeout(() => {
      db.status.running = true;
      db.status.connected = true;
      resolve({ success: true, status: db.status });
    }, SIMULATED_DELAY);
  });
};

export const stopProxy = () => {
  return new Promise((resolve) => {
    setTimeout(() => {
      db.status.running = false;
      db.status.connected = false;
      resolve({ success: true, status: db.status });
    }, SIMULATED_DELAY);
  });
};

export const getConfig = () => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve({ ...db.config });
    }, SIMULATED_DELAY);
  });
};

export const updateConfig = (newConfig) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      db.config = { ...db.config, ...newConfig };
      resolve({ success: true, config: db.config });
    }, SIMULATED_DELAY);
  });
};
