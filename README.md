# SDUI Example (Go backend + Vite React frontend)

This repository contains a demonstration of Server-Driven UI (SDUI):

- A Go backend using Gin that exposes navigation and screen endpoints.
- A React frontend (Vite) that fetches navigation and screen definitions and renders them.

## Run backend (PowerShell)

From the repo root run:

```powershell
cd ./cmd/server
go run .
```

The server listens on port 8080.

To run the server in the background in PowerShell (useful for local testing):

```powershell
# starts the server and detaches; logs are written to server.log
Start-Process -NoNewWindow -FilePath pwsh -ArgumentList "-NoProfile -Command cd ./cmd/server; go run . *> server.log 2>&1"
```

## Run frontend (PowerShell)

```powershell
cd ./frontend
npm install
npm run dev
```

Open http://localhost:5173 in your browser. The frontend expects the backend at http://localhost:8080.

Notes:
- Ensure Go 1.21+ and Node 18+ (or newer) are installed.
- The repository contains a Go backend (Gin) and a minimal Vite + React frontend. The frontend files are scaffolded but you must run `npm install` to fetch dependencies before using `npm run dev`.
- I built the Go server locally during setup to verify compilation; the frontend was not launched here (npm install not run). If you want, I can start the backend and run the frontend dev server here as well.


## Quick Start

To run:

- **Backend:**  
    ```powershell
    cd .\cmd\server
    go run .
    ```

- **Frontend:**  
    ```powershell
    cd .\frontend
    npm install
    npm run dev
    ```
    Then open [http://localhost:5173](http://localhost:5173) in your browser.