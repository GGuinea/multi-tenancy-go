package adapters

import (
	"context"
	"multitenancy/internal/notes/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NotesRepository struct {
	DbPool *pgxpool.Pool
}

func NewNotesRepository(dbPool *pgxpool.Pool) *NotesRepository {
	return &NotesRepository{DbPool: dbPool}
}

func (r *NotesRepository) Save(ctx context.Context, tenant *string, note *domain.Note) error {
	tx, err := r.DbPool.Begin(ctx)

	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "SET search_path TO "+*tenant)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "INSERT INTO note VALUES ($1, $2, $3, $4)", note.Id, note.Title, note.Body, note.CreatedAt)

	if err != nil {
		return err
	}

	tx.Exec(ctx, "SET search_path TO public")
	err = tx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *NotesRepository) List(ctx context.Context, tenant *string) ([]*domain.Note, error) {
	tx, err := r.DbPool.Begin(ctx)

	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, "SET search_path TO "+*tenant)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, "SELECT id, title, body, created_at FROM note")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notes := make([]*domain.Note, 0)
	for rows.Next() {
		note := new(domain.Note)
		err := rows.Scan(&note.Id, &note.Title, &note.Body, &note.CreatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	tx.Exec(ctx, "SET search_path TO public")
	err = tx.Commit(ctx)

	if err != nil {
		return nil, err
	}

	return notes, nil
}
