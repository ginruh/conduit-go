package comment

import "github.com/iyorozuya/real-world-app/internal/sqlc"

type CommentService interface {
	// List()
	// Create()
	// Delete()
}

type CommentServiceImpl struct {
	q *sqlc.Queries
}

func NewCommentService(q *sqlc.Queries) CommentServiceImpl {
	return CommentServiceImpl{q}
}
