package notes

import (
	"multitenancy/internal/notes/app"
	"multitenancy/internal/notes/domain"

	"github.com/gin-gonic/gin"
)

type NoteRouter struct {
	noteService *app.NoteService
}

func BuildRoutes(notePath *gin.RouterGroup, deps *DependencyTree) {
	noteRouter := &NoteRouter{noteService: deps.NoteService}

	notePath.POST("/", noteRouter.CreateNote)
	notePath.GET("/", noteRouter.ListNotes)
}

func (r *NoteRouter) CreateNote(c *gin.Context) {
	var requestDto domain.NewNoteRequestDto

	err := c.BindJSON(&requestDto)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tenantName := c.GetHeader("X-Tenant")

	id, err := r.noteService.CreateNote(c.Request.Context(), tenantName, &requestDto)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"id": id})
}

func (r *NoteRouter) ListNotes(c *gin.Context) {
	tenantName := c.GetHeader("X-Tenant")

	notes, err := r.noteService.ListNotes(c.Request.Context(), tenantName)

	if err != nil {
		c.JSON(500, gin.H{"error": "Tenant not found"})
		return
	}

	c.JSON(200, gin.H{"notes": notes})
}
