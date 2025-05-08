# 🚗 Street Racing Arena

[![CI/CD](https://github.com/your-org/street-racing-arena/actions/workflows/ci.yml/badge.svg)](https://github.com/your-org/street-racing-arena/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A modern multiplayer street racing game platform. Race, compete, and climb the leaderboards! Built with a robust Go backend and a sleek Nuxt 3 frontend.

---

## 🛠️ Project Stack

- **Backend:** [Go](https://golang.org/) · [Gin](https://gin-gonic.com/) · [GORM](https://gorm.io/) · [PostgreSQL](https://www.postgresql.org/)
- **Frontend:** [Nuxt 3](https://nuxt.com/) · [Vue 3](https://vuejs.org/)
- **Web Server/Proxy:** [Caddy](https://caddyserver.com/)
- **Containerization:** [Docker Compose](https://docs.docker.com/compose/)
- **CI/CD:** GitHub Actions

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
