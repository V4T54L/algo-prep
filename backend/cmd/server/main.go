package main

import (
	"app/internal/application/usecase"
	"app/internal/domain/service"
	"app/internal/infrastructure/handlers"
	memrepo "app/internal/infrastructure/repository"
	"app/internal/infrastructure/router"
	"app/internal/infrastructure/template"
	"log"
	"net/http"
)

func main() {
	// Repos
	// userRepo := memrepo.NewInMemoryUserRepo()
	postRepo := memrepo.NewInMemoryPostRepo()
	problemRepo := memrepo.NewInMemoryProblemRepo()
	contestRepo := memrepo.NewInMemoryContestRepo()

	// Services
	// userService := service.NewUserDomainService(userRepo)
	postService := service.NewPostDomainService(postRepo)
	problemService := service.NewProblemDomainService(problemRepo)
	contestService := service.NewContestService(contestRepo)

	// UseCases
	postUC := usecase.NewPostUseCase(postService)
	problemUC := usecase.NewProblemUseCase(problemService)
	contestUC := usecase.NewContestUseCase(contestService)

	// cfg, err := config.Load()
	// if err != nil {
	// 	log.Fatal("Error loading the config : ", err)
	// }
	// Renderer
	renderer := template.NewTemplRenderer()

	// Handlers
	postHandler := handlers.NewPostHandler(*postUC, *renderer)
	appHandler := handlers.NewAppHandler(*postUC, *renderer)
	problemHandler := handlers.NewProblemHandler(*problemUC, *renderer)
	contestHandler := handlers.NewContestHandler(*contestUC, *renderer)

	// Router
	r := router.NewRouter(postHandler, appHandler, problemHandler, contestHandler)

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", r)
}
