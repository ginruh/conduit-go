package seeders

import (
	"github.com/iyorozuya/real-world-app/internal/sqlc"
)

type Seed struct {
	q *sqlc.Queries
}

func New(q *sqlc.Queries) Seed {
	return Seed{q}
}

func (s Seed) SeedAll() {
	s.seedUsers()
}
