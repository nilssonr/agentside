package postgres

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func mustCreateTime(t time.Time) pgtype.Timestamptz {
	ts := pgtype.Timestamptz{}
	ts.Scan(t)

	return ts
}

func mustCreateString(s string) pgtype.Text {
	t := pgtype.Text{}
	if len(s) > 0 {
		t.Scan(s)
	}

	return t
}
