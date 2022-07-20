package tag

import "github.com/iyorozuya/real-world-app/internal/sqlc"

type TagService interface {
	// List()
}

type TagServiceImpl struct {
	q *sqlc.Queries
}

func NewTagService(q *sqlc.Queries) TagServiceImpl {
	return TagServiceImpl{q}
}
