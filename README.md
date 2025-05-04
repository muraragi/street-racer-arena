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
