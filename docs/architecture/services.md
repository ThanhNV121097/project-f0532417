# Service contracts

Base paths below reach backend after proxy strips `/api`.

## Error envelope

All non-success responses use:

```json
{"error":{"code":"INVALID_GREETING","message":"Greeting must contain 1 to 200 characters."}}
```

`code` is stable uppercase snake case. `message` is safe plain text. No database detail is returned.

## Read greeting

`GET /v1/greeting`

Request body: none.

Success `200`:

```json
{"greeting":"Hello, World!"}
```

If singleton row is absent, API recreates it then returns `200`. Database failure returns `500` with `INTERNAL_ERROR`.

## Save greeting

`PUT /v1/greeting`

Request:

```json
{"greeting":"Saved value"}
```

Backend trims surrounding whitespace, then validates 1–200 characters. Last completed request wins.

Success `200`:

```json
{"greeting":"Saved value"}
```

Malformed JSON or missing/non-string `greeting`: `400`, `INVALID_REQUEST`. Empty, whitespace-only, or over-limit greeting: `422`, `INVALID_GREETING`. Database failure: `500`, `INTERNAL_ERROR`.

## Health

`GET /healthz`

Request body: none. Returns `200` with `{"status":"ok"}` only when migration succeeded and database `SELECT 1` succeeds. Otherwise returns `503` with `SERVICE_UNAVAILABLE` envelope.
