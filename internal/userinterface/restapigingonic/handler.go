package restapigingonic

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"local/internal/domain"
)

type handler struct {
	projects ProjectsServicePort
}

func (obj *handler) info(c *gin.Context) {
	c.Writer.Header().Add("Cache-Control", "no-cache")
	c.JSON(http.StatusOK, gin.H{"alive": true})
}

func (obj *handler) getAll(c *gin.Context) {
	projects, err := obj.projects.GetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func (obj *handler) getByID(c *gin.Context) {
	ID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	project, err := obj.projects.GetByID(c.Request.Context(), ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func (obj *handler) save(c *gin.Context) {
	var requestBody domain.Project

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := obj.projects.Save(c.Request.Context(), requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
