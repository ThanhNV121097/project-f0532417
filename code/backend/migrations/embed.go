package migrations

import "embed"

// Files contains SQL migration files bundled into API binary.
//
//go:embed *.sql
var Files embed.FS
