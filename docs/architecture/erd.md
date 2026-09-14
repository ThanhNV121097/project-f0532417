# ERD

## `greetings`

One global persisted greeting. Exactly one row, identified by fixed primary key `id = 1`.

| Column | Type | Constraints | Purpose |
|---|---|---|---|
| `id` | `smallint` | primary key, `CHECK (id = 1)` | Singleton identity |
| `text` | `text` | not null, `CHECK (char_length(text) BETWEEN 1 AND 200)` | Trimmed greeting |
| `updated_at` | `timestamptz` | not null, default `now()` | Latest save time |

Seed migration inserts `(1, 'Hello, World!')` with conflict ignored. API recreates this row if absent, per SRS.

## Relationships

No relationships. This module stores one global value; users, sessions, history, and external entities are out of scope.
