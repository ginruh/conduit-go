package tag

import (
	"github.com/iyorozuya/real-world-app/internal/queries"
)

type TagService interface {
	// List()
}

type TagServiceImpl struct {
	q *queries.Queries
}

func NewTagService(q *queries.Queries) TagServiceImpl {
	return TagServiceImpl{q}
}
