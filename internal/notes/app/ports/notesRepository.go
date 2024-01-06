package ports

import (
	"context"
	"multitenancy/internal/notes/domain"
)

type NoteRepository interface {
	Save(context.Context, *string, *domain.Note) error
	List(context.Context, *string) ([]*domain.Note, error)
}
