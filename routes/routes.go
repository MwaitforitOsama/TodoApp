package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mwaitforitosama/TodoApp/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Route("/v1/todo", loadTodoRoutes)
	return router
}

func loadTodoRoutes(router chi.Router) {
	todoHandler := &handler.Todo{}
	router.Post("/", todoHandler.Create)
	router.Get("/", todoHandler.List)
	router.Get("/{id}", todoHandler.GetByID)
	router.Put("/{id}", todoHandler.UpdateByID)
	router.Delete("/{id}", todoHandler.DeleteByID)

}
