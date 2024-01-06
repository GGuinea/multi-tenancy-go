package app

import (
	"context"
	"multitenancy/internal/notes/app/ports"
	"multitenancy/internal/notes/domain"
)

type NoteService struct {
	noteRepository ports.NoteRepository
}

func NewNoteService(noteRepository ports.NoteRepository) *NoteService {
	return &NoteService{noteRepository: noteRepository}
}

func (s *NoteService) CreateNote(ctx context.Context, tenantName string, requestDto *domain.NewNoteRequestDto) (*string, error) {
	note := domain.NewNoteFromRequestDto(requestDto)

	err := s.noteRepository.Save(ctx, &tenantName, note)

	if err != nil {
		return nil, err
	}

	return &note.Id, nil
}

func (s *NoteService) ListNotes(ctx context.Context, tenantName string) ([]*domain.NoteResponseDto, error) {
	notes, err := s.noteRepository.List(ctx, &tenantName)

	if err != nil {
		return nil, err
	}

	var response []*domain.NoteResponseDto

	for _, note := range notes {
		response = append(response, domain.CreateNoteResponseDto(note))
	}

	return response, nil
}
