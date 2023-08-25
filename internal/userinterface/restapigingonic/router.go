package restapigingonic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(projects ProjectsServicePort) http.Handler {
	var handler = &handler{
		projects: projects,
	}

	var router = gin.Default()

	router.Handle("GET", "/info", handler.info)

	router.Handle("GET", "/getAll", handler.getAll)

	router.Handle("GET", "/getByID/:id", handler.getByID)

	router.Handle("POST", "/save", handler.save)
	return router
}
