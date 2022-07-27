package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/iyorozuya/real-world-app/internal/controllers/api"
	"github.com/iyorozuya/real-world-app/internal/middlewares"
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/seeders"
	"github.com/iyorozuya/real-world-app/internal/services/api/article"
	"github.com/iyorozuya/real-world-app/internal/services/api/auth"
	"github.com/iyorozuya/real-world-app/internal/services/api/comment"
	"github.com/iyorozuya/real-world-app/internal/services/api/tag"
	"github.com/iyorozuya/real-world-app/internal/services/api/user"
	"github.com/jmoiron/sqlx"
	"os"
)

func Bootstrap(r chi.Router, database *sqlx.DB) {
	q := queries.New(database)
	validate := validator.New()

	authController := api.NewAuthController(auth.NewAuthService(q), validate)
	userController := api.NewUserController(user.NewUserService(q), validate)
	articleController := api.NewArticleController(article.NewArticleService(q), validate)
	commentController := api.NewCommentController(comment.NewCommentService(q), validate)
	tagController := api.NewTagController(tag.NewTagService(q), validate)

	// No authentication required
	r.Group(func(r chi.Router) {
		// User login and registration
		r.Post("/users/login", authController.Login)
		r.Post("/users", authController.Register)

		// tags
		r.Get("/tags", tagController.List)
	})

	// User authenticated actions
	r.Group(func(r chi.Router) {
		r.Use(middlewares.JWT)

		// auth
		r.Get("/user", userController.GetCurrentUser)
		r.Put("/user", userController.UpdateUser)

		// profiles
		r.Post("/profiles/{username}/follow", userController.FollowUser)
		r.Delete("/profiles/{username}/follow", userController.UnfollowUser)

		// articles
		r.Get("/articles/feed", articleController.Feed)
		r.Post("/articles", articleController.Create)
		r.Put("/articles/{slug}", articleController.Update)
		r.Delete("/articles/{slug}", articleController.Delete)
		r.Post("/articles/{slug}/favorite", articleController.Favorite)
		r.Delete("/articles/{slug}/favorite", articleController.Unfavorite)

		// comments
		r.Post("/articles/{slug}/comments", commentController.Create)
		r.Delete("/articles/{slug}/comments/{id}", commentController.Delete)
	})

	// Optional authenticated routes
	r.Group(func(r chi.Router) {
		r.Use(middlewares.OptionalJWT)

		// profiles
		r.Get("/profiles/{username}", userController.GetProfile)

		// articles
		r.Get("/articles", articleController.List)
		r.Get("/articles/{slug}", articleController.Get)

		// comments
		r.Get("/articles/{slug}/comments", commentController.List)
	})

	// Seed db if RUN_SEEDERS is yes
	if os.Getenv("RUN_SEEDERS") == "yes" {
		seed := seeders.New(q)
		seed.SeedAll()
	}
}
