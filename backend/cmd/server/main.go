// Package main is the composition root for the server.
// It wires all Clean Architecture layers together via constructor injection
package main

import (
	"log"
	"time"

	"example.com/interview-question-009/config"
	"example.com/interview-question-009/config/database"
	"example.com/interview-question-009/internal/handler"
	"example.com/interview-question-009/internal/middleware"
	"example.com/interview-question-009/internal/repository"
	"example.com/interview-question-009/internal/usecase"
	"example.com/interview-question-009/pkg"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	fiberlogger "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func main() {
	cfg := config.Load(".env.local")
	pkg.InitLogger(cfg.AppEnv)
	db := database.Init()

	// Infrastructure layer — GORM repository implementations
	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Application layer — use cases receive repository interfaces
	postUC := usecase.NewPostUseCase(postRepo)
	commentUC := usecase.NewCommentUseCase(commentRepo)
	authUC := usecase.NewAuthUseCase(userRepo, cfg.JWTSecret)

	// Interface adapter layer — handlers receive use-case interfaces
	postHandler := handler.NewPostHandler(postUC)
	commentHandler := handler.NewCommentHandler(commentUC)
	authHandler := handler.NewAuthHandler(authUC)

	// ─── Fiber App ────────────────────────────────────────────────────────────
	app := fiber.New(fiber.Config{
		ProxyHeader: "X-Real-IP",
		BodyLimit: 10 * 1024 * 1024,
	})

	// ─── Middleware ───────────────────────────────────────────────────────────
	app.Use(recover.New())
	app.Use(requestid.New())

	loggerCfg := fiberlogger.Config{
		Format: "[${time}] ${status} - ${method} ${path} - ${latency}\n",
	}
	if cfg.IsProduction() {
		loggerCfg.Format = `{"time":"${time}","status":${status},"latency":"${latency}","method":"${method}","path":"${url}","ip":"${ip}","rid":"${locals:requestid}"}` + "\n"
	}
	app.Use(fiberlogger.New(loggerCfg))
	app.Use(helmet.New())
	app.Use(limiter.New(limiter.Config{
		Max: 60,
		Expiration: 1 * time.Minute,
		LimitReached: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests. Please try again later.",
			})
		},
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSOrigins,
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	api := app.Group("/api")

	// Public routes
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)

	// Post routes — reads are public, writes require JWT
	jwtGuard := middleware.JWTProtected(authUC)
	posts := api.Group("/posts")
	posts.Get("/:postID", postHandler.GetPost)
	posts.Get("/:postID/comments", commentHandler.GetComments)
	posts.Post("/:postID/comments", jwtGuard, commentHandler.CreateComment)
	posts.Delete("/:postID/comments/:commentID", jwtGuard, commentHandler.DeleteComment)

	// ─── Listen ───────────────────────────────────────────────────────────────
	addr := ":" + cfg.AppPort
	log.Printf("Server running on %s (env=%s)", addr, cfg.AppEnv)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
