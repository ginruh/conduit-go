package seeders

import (
	"github.com/iyorozuya/real-world-app/internal/queries"
)

type Seed struct {
	q *queries.Queries
}

func New(q *queries.Queries) Seed {
	return Seed{q}
}

func (s Seed) SeedAll() {
	s.seedUsers()
}
