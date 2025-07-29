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

	// Services
	// userService := service.NewUserDomainService(userRepo)
	postService := service.NewPostDomainService(postRepo)

	// UseCases
	postUC := usecase.NewPostUseCase(postService)

	// cfg, err := config.Load()
	// if err != nil {
	// 	log.Fatal("Error loading the config : ", err)
	// }
	// Renderer
	renderer := template.NewTemplRenderer()

	// Handlers
	postHandler := handlers.NewPostHandler(postUC, *renderer)

	// Router
	r := router.NewRouter(postHandler)

	log.Println("Server running at :8080")
	http.ListenAndServe(":8080", r)
}
