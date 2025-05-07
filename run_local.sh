set -e

log() {
  echo "[RUNNER] $1"
}

log "Starting Docker Compose services in apps/backend..."
(cd apps/backend && docker-compose up -d)

if ! command -v concurrently &> /dev/null
then
    log "concurrently could not be found. Please install it globally to run backend and frontend simultaneously."
    log "You can typically install it with: npm install -g concurrently"
    log "Attempting to run backend and frontend sequentially. Press Ctrl+C to stop the backend and start the frontend."
    log "Starting backend with air in apps/backend..."
    (cd apps/backend/cmd/server && air)
    log "Starting frontend with npm run dev in apps/frontend..."
    (cd apps/frontend && npm run dev)
else
    log "Starting backend with air and frontend with npm run dev using concurrently..."
    concurrently --kill-others-on-fail \
      "cd apps/backend && air" \
      "cd apps/frontend && npm run dev"
fi

log "All processes started. Monitoring logs..."
wait 