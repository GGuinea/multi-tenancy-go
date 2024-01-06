package domain

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Id        string
	Title     string
	Body      string
	CreatedAt time.Time
	Tag       string
}

type NewNoteRequestDto struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NoteResponseDto struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

func NewNoteFromRequestDto(requestDto *NewNoteRequestDto) *Note {
	return &Note{
		Id:        uuid.New().String(),
		Title:     requestDto.Title,
		Body:      requestDto.Body,
		CreatedAt: time.Now(),
	}
}

func CreateNoteResponseDto(note *Note) *NoteResponseDto {
	return &NoteResponseDto{
		Id:        note.Id,
		Title:     note.Title,
		Body:      note.Body,
		CreatedAt: note.CreatedAt.Format(time.RFC3339),
	}
}
