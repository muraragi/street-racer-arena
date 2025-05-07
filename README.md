# W I P

## Local Development

To start the project locally, follow these steps:

1.  **Start PostgreSQL Database:**
    Navigate to the `backend` directory and run Docker Compose:

    ```bash
    cd backend
    docker compose up -d
    ```

2.  **Install Backend Dependencies:**
    Still in the `backend` directory, install the necessary dependencies:

    ```bash
    go mod tidy
    ```

3.  **Run Backend Server:**
    Use `air` to run the backend server with live reloading:

    ```bash
    air
    ```

4.  **Install Frontend Dependencies:**
    Navigate to the `frontend` directory and install the required packages:

    ```bash
    cd ../frontend
    npm install
    ```

5.  **Run Frontend Development Server:**
    Start the frontend development server:
    ```bash
    npm run dev
    ```

Now you should be able to access the frontend at `http://localhost:3000` and the backend API will be running at `http://localhost:8080`.

## CI/CD & Production Deployment

Briefly, on every push to the `main` branch:

- A GitHub Actions workflow builds and pushes Docker images for the backend and frontend (only if their code changed).
- The workflow then connects via SSH to the production server, copies over the updated `docker-compose.yml` and `Caddyfile`, pulls the new images, and restarts services with Docker Compose.
- Unused Docker images on the server are cleaned up automatically.

## Running Locally (Simplified)

A convenience script `run_local.sh` is provided to start all necessary services (backend, frontend, and Docker containers) with a single command.

1.  **Make the script executable** (only needed once):

    ```bash
    chmod +x run_local.sh
    ```

2.  **Run the script**:
    ```bash
    ./run_local.sh
    ```

This script uses `concurrently` (npm i -g concurrently) to run backend and frontend processes in parallel and display their logs in the same terminal window. If `concurrently` is not found, it will offer to run them sequentially.
