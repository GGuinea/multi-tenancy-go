package notes

import (
	"multitenancy/internal/notes/app"
	"multitenancy/internal/notes/app/adapters"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteDepencies struct {
	DbPool *pgxpool.Pool
}

type DependencyTree struct {
	NoteService *app.NoteService
}

func NewNoteDependencies(deps *NoteDepencies) *DependencyTree {
	if deps == nil {
		panic("NoteDepencies is nil")
	}

	noteRepository := adapters.NewNotesRepository(deps.DbPool)
	noteService := app.NewNoteService(noteRepository)

	return &DependencyTree{NoteService: noteService}
}
