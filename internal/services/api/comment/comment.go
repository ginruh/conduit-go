package comment

import (
	"github.com/iyorozuya/real-world-app/internal/queries"
)

type CommentService interface {
	// List()
	// Create()
	// Delete()
}

type CommentServiceImpl struct {
	q *queries.Queries
}

func NewCommentService(q *queries.Queries) CommentServiceImpl {
	return CommentServiceImpl{q}
}
