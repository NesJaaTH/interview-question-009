# Interview Question No.9 — IT 08-1 Comment Page

A single-page web app showing a post with a comment section.  
Users sign in and submit comments by pressing **Enter**.

**Stack:** Go + Fiber v3 + GORM + SQLite · Vue 3 + TypeScript + Tailwind v4 + Pinia · Vue CLI

---

## Prerequisites

| Tool | Version |
|------|---------|
| Go | 1.22+ |
| Node.js | 18+ |
| npm | 9+ |

> No GCC required — the project uses a pure-Go SQLite driver.

---

## Project structure

```
interview-question/
├── backend/
│   ├── cmd/server/        # entry point
│   ├── config/            # env config + database init
│   ├── internal/
│   │   ├── domain/        # entities & repository interfaces
│   │   ├── repository/    # GORM implementations
│   │   ├── usecase/       # business logic
│   │   ├── handler/       # HTTP handlers (Fiber)
│   │   └── middleware/    # JWT guard
│   └── .env.example
└── frontend/
    ├── public/            # HTML template
    ├── src/
    │   ├── api/           # axios clients
    │   ├── assets/        # global CSS
    │   ├── stores/        # Pinia auth store
    │   ├── views/         # LoginPage
    │   └── components/    # Post, CommentSection
    ├── vue.config.ts      # Vue CLI config + dev proxy
    └── .env.example
```

---

## Backend setup

### 1. Install dependencies

```bash
cd backend
go mod tidy
```

### 2. Create the env file

```bash
cp .env.example .env.local
```

Open `.env.local` and fill in the values:

```env
APP_PORT=7809
APP_ENV=development
CORS_ORIGINS=*,http://localhost:5173
JWT_SECRET=your-secret-key-here
```

> `JWT_SECRET` can be any long random string in development.

### 3. Run the server

```bash
go run ./cmd/server
```

The server starts on `http://localhost:7809`.  
On first run it creates `data.db` and seeds:
- 1 post by **Change can**
- 1 comment by **Blend 285**
- 1 user account: username `blend285` / password `blend285`

### 4. Run tests

```bash
go test ./...
```

---

## Frontend setup

### 1. Install dependencies

```bash
cd frontend
npm install
```

### 2. Create the env file

```bash
cp .env.example .env.local
```

Open `.env.local` and fill in the values:

```env
VUE_APP_API_BASE_URL=/api
VUE_APP_BACKEND_URL=http://localhost:7809
```

> `VUE_APP_BACKEND_URL` tells the Vue CLI dev-server proxy where to forward `/api` requests.  
> In a production build, set `VUE_APP_API_BASE_URL` to the full backend URL instead.

### 3. Start the dev server

```bash
npm run dev
```

Open `http://localhost:5173`.

### 4. Build for production

```bash
npm run build
# output → frontend/dist/
```

---

## API reference

All endpoints are prefixed with `/api`.

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| `POST` | `/auth/login` | — | Sign in, returns JWT token |
| `GET` | `/posts/:id` | — | Get a post |
| `GET` | `/posts/:id/comments` | — | List comments |
| `POST` | `/posts/:id/comments` | JWT | Create a comment |
| `DELETE` | `/posts/:id/comments/:commentId` | JWT | Delete a comment |

**Login request**
```json
{ "username": "blend285", "password": "blend285" }
```

**Login response**
```json
{
  "token": "<jwt>",
  "user": { "id": 1, "username": "blend285", "display_name": "Blend 285" }
}
```

Protected routes require the header:
```
Authorization: Bearer <token>
```

---

## Demo credentials

| Field | Value |
|-------|-------|
| Username | `blend285` |
| Password | `blend285` |
