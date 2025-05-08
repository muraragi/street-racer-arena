# 🚗 Street Racing Arena

[![CI/CD](https://github.com/muraragi/street-racing-arena/actions/workflows/deploy.yml/badge.svg)](https://github.com/muraragi/street-racing-arena/actions/workflows/deploy.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A modern multiplayer street racing game platform. Race, compete, and climb the leaderboards! Built with a robust Go backend and a sleek Nuxt 3 frontend.

---

## 🛠️ Project Stack

- **Backend:**  
  [![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=rounded)](https://golang.org/)
  [![Gin](https://img.shields.io/badge/Gin-00B386?logo=go&logoColor=white&style=rounded)](https://gin-gonic.com/)
  [![GORM](https://img.shields.io/badge/GORM-FF7043?logo=go&logoColor=white&style=rounded)](https://gorm.io/)
  [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white&style=rounded)](https://www.postgresql.org/)

- **Frontend:**  
  [[![Nuxt](https://img.shields.io/badge/Nuxt-002E3B?logo=nuxt&logoColor=#00DC82)](https://nuxt.com/)
  [![Vue 3](https://img.shields.io/badge/Vue_3-4FC08D?logo=vue.js&logoColor=white&style=rounded)](https://vuejs.org/)

- **Web Server/Proxy:**  
  [![Caddy](https://img.shields.io/badge/Caddy-00C7B7?logo=caddy&logoColor=white&style=rounded)](https://caddyserver.com/)

- **Containerization:**  
  [![Docker Compose](https://img.shields.io/badge/Docker_Compose-2496ED?logo=docker&logoColor=white&style=rounded)](https://docs.docker.com/compose/)

- **CI/CD:**  
  [![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?logo=githubactions&logoColor=white&style=rounded)](https://github.com/features/actions)

---

## 🚀 Getting Started (Local Development)

Follow these steps to run the project locally:

1. **Start PostgreSQL Database:**
   ```bash
   cd apps/backend
   docker compose up -d
   ```
2. **Install Backend Dependencies:**
   ```bash
   go mod tidy
   ```
3. **Run Backend Server (with live reload):**
   ```bash
   air
   ```
4. **Install Frontend Dependencies:**
   ```bash
   cd ../../frontend
   npm install
   ```
5. **Run Frontend Development Server:**
   ```bash
   npm run dev
   ```

- Frontend: [http://localhost:3000](http://localhost:3000)
- Backend API: [http://localhost:8080](http://localhost:8080)

---

## 🐳 One-Command Local Run

A convenience script `run_local.sh` is provided to start all services (backend, frontend, and Docker containers) with a single command.

1. **Make the script executable** (only needed once):
   ```bash
   chmod +x run_local.sh
   ```
2. **Run the script:**
   ```bash
   ./run_local.sh
   ```

> This script uses [`concurrently`](https://www.npmjs.com/package/concurrently) to run backend and frontend processes in parallel. If not installed, it will offer to run them sequentially.

---

## 🚢 CI/CD & Production Deployment

- On every push to the `main` branch:
  - GitHub Actions builds and pushes Docker images for backend and frontend (only if their code changed).
  - The workflow connects via SSH to the production server, updates `docker-compose.yml` and `Caddyfile`, pulls new images, and restarts services.
  - Unused Docker images are cleaned up automatically.

---

## 📄 License

This project is licensed under the MIT License.
