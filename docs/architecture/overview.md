# Architecture overview

## Stack

- Frontend: Next.js 15 App Router, TypeScript, Tailwind v3, ESLint.
- Backend: Go 1.22+ HTTP server using `net/http` and pgx PostgreSQL driver.
- Database: PostgreSQL 16. Docker Compose runs local stack.

## Layout

```text
code/backend/cmd/api/     API entry point
code/backend/migrations/  ordered SQL migration pairs
code/frontend/app/        App Router shell and frozen shared tokens
code/frontend/components/ story components
code/frontend/lib/mock/   UI-stage fixtures, removed by API stage
```

`app/page.tsx` is composition root only. Story UI adds one default component import and one element. Interactive components start with literal first line `"use client"`.

## Data flow

Browser calls API through `NEXT_PUBLIC_API_URL`. API validates input, reads/writes PostgreSQL, and returns JSON. Backend applies migrations before accepting traffic. `/healthz` returns 200 only after migration and database ping succeed.

## Decisions

| Decision | Rejected | Tradeoff |
|---|---|---|
| One global greeting row | Multi-row/history | Meets scope; no visitor ownership or history. |
| pgx + stdlib HTTP | ORM/framework | Small dependency and explicit SQL; later complex data may justify query layer. |
| SQL migration pairs tracked in `schema_migrations` | Manual schema setup | Startup owns empty runtime database; migrations run during boot. |
| Last completed save wins | Version conflicts | Matches SRS; concurrent edits may overwrite earlier saves. |

## Conventions

- SQL uses lowercase snake_case. Go exports only cross-package symbols.
- API paths use `/v1/...`, never `/api/...`; proxy owns `/api` prefix.
- JSON errors share `{ "error": { "code": "...", "message": "..." } }`.
- CSS modules use tokens from `app/globals.css`; no visual literals or `var()` fallbacks.
- No secrets committed. Every service tracks `.env.example`.

## Environment

| Service | Keys |
|---|---|
| Backend | `DATABASE_URL`, `PORT`, optional `APP_PORT` fallback |
| Frontend | `NEXT_PUBLIC_API_URL`, `API_ORIGIN` |
| Compose | `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`, optional port/memory overrides |

## Run

1. Copy root `.env.example` to `.env` and change local password if wanted.
2. Run `docker compose --profile local up --build` from repository root.
3. Frontend serves `http://localhost:3000`; backend health endpoint serves `http://localhost:8080/healthz`.

## Compatibility and rollout

Migrations are filename-ordered, recorded after success, and rerunnable. Migration runner executes transactionally except files containing `CREATE INDEX CONCURRENTLY`, which PostgreSQL requires outside transactions. First deploy creates seed greeting. No external services or credentials.
