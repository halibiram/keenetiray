# V2Ray Keenetic Control Panel

This project is a V2Ray proxy application designed specifically for Keenetic routers. It provides a web-based management interface to control the V2Ray service on the router.

## Features (Planned)

*   **Web UI**: A responsive web interface to manage V2Ray server configurations.
*   **V2Ray Core Integration**: Manages the V2Ray-core process for traffic routing.
*   **System Integration**: Integrates with Keenetic RouterOS for startup and network management.
*   **Multiple Protocol Support**: Supports VMess, VLESS, Shadowsocks, etc.

## Project Structure

*   `/cmd/v2ray-keenetic`: The Go backend application (API server).
*   `/web`: The Vue.js frontend application (Web UI).
*   `/internal`: Internal Go packages for the backend.
*   `/scripts`: Build and deployment scripts.
*   `/configs`: Example configuration files.
*   `/docs`: Project documentation.

## Development

This project is in the early stages of development.

### Backend

The backend is written in Go using the Gin framework. To run the backend:

```bash
# From the project root
# Note: Building the binary is currently not supported in this environment due to size limits.
# Use 'go run' for development.
go run ./cmd/v2ray-keenetic
```

### Frontend

The frontend is a Vue.js 3 application built with Vite. To run the frontend development server:

```bash
# From the project root
cd web
npm install
npm run dev
```

The frontend will be available at `http://localhost:5173` and will proxy API requests to the backend running on `http://localhost:8080`.
