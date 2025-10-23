# SDUI Demo (Go Backend + React Frontend)

This repository demonstrates **Server-Driven UI (SDUI)**, a pattern where the server defines the UI structure and components, and the client renders them dynamically. This approach allows for flexible, backend-controlled user interfaces without requiring frontend deployments for UI changes.

## Quick Start

To quickly get the application running:

1. Run the backend: `make server`
2. In another terminal, run the frontend: `make frontend`
3. Open `http://localhost:5173` in your browser.

Ensure you have Go 1.21+ and Node.js 18+ installed.

## Architecture

- **Backend**: Go server using the Gin framework, serving JSON representations of UI components and navigation.
- **Frontend**: React application built with Vite, fetching UI definitions from the backend and rendering them using reusable component library.

The backend exposes RESTful endpoints that return arrays of `Component` objects, each with a `type` and `props`. The frontend maps these to React components.

## Features

The demo includes several screens/pages:

- **Home**: Welcome page with banner and navigation button.
- **Products**: Displays featured products in a grid layout.
- **Profile**: User profile overview with stats.
- **List & Detail**: A list of items with clickable details.
- **Detail**: Detailed view for individual items.
- **Forms**: Contact form with input fields and submit button.
- **Articles**: List of articles with title, content, author, and date.
- **Combined**: An overview page combining elements from all other screens (products, profile, list, forms, articles).

## Supported Components

The system supports the following UI component types:

- `heading`: Headings with configurable level (h1-h6).
- `paragraph`: Text paragraphs.
- `image`: Images with src and alt attributes.
- `button`: Clickable buttons with action IDs.
- `product_list`: Container for product cards.
- `product_card`: Individual product display with name, price, and image.
- `stat_item`: Key-value stat displays.
- `divider`: Visual separators.
- `list`: Ordered/unordered lists.
- `list_item`: Individual list items with click actions.
- `form_input`: Form inputs (text, email, textarea).
- `form_button`: Form submission buttons.
- `article`: Article displays with title, content, author, and date.

## API Endpoints

All endpoints are under `/api/ui`:

- `GET /api/ui/nav`: Returns navigation menu items.
- `GET /api/ui/home`: Home screen components.
- `GET /api/ui/products`: Products screen components.
- `GET /api/ui/profile`: Profile screen components.
- `GET /api/ui/list-detail`: List & Detail screen components.
- `GET /api/ui/detail/:id`: Detail screen for specific item.
- `GET /api/ui/forms`: Forms screen components.
- `GET /api/ui/articles`: Articles screen components.
- `GET /api/ui/combined`: Combined overview screen components.
- `GET /health`: Simple health check.

## Installation and Running

### Prerequisites

- Go 1.21 or later
- Node.js 18 or later

### Backend

From the repository root:

```powershell
cd ./cmd/server
go run .
```

The server starts on `http://localhost:8080`.

To run in background (PowerShell):

```powershell
Start-Process -NoNewWindow -FilePath pwsh -ArgumentList "-NoProfile -Command cd ./cmd/server; go run . *> server.log 2>&1"
```

### Frontend

```powershell
cd ./frontend
npm install
npm run dev
```

Open `http://localhost:5173` in your browser.

### Using Makefile

The repository includes a Makefile for convenience:

```powershell
make server    # Runs the Go backend
make frontend  # Installs deps and runs the React dev server
```

## Development

### Adding New Components

1. Define the component type in `internal/services/ui.go`.
2. Add the React component in `frontend/src/components/`.
3. Register it in `ComponentRenderer.jsx`.

### Extending Screens

Add new endpoints in `internal/api/routes.go` and implement the handler in `internal/services/ui.go`.

### Building for Production

Frontend:

```powershell
cd ./frontend
npm run build
```

Backend:

```powershell
go build ./cmd/server
```

## Project Structure

```
.
├── cmd/server/          # Go server entry point
├── internal/
│   ├── api/             # HTTP handlers and routes
│   ├── middleware/      # CORS middleware
│   └── services/        # UI service implementations
├── frontend/
│   ├── src/
│   │   ├── components/  # React UI components
│   │   ├── screens/     # Screen-specific components
│   │   └── ...
│   └── package.json
├── go.mod
├── Makefile
└── README.md
```

## Contributing

This is a demo project. To extend:

- Add new component types in the backend service.
- Implement corresponding React components.
- Update the navigation and routing as needed.

For questions or issues, please open a GitHub issue.