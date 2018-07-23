package db

import "time"

type Note struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`

	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
type Notes []Note

func (note Note) GetAll(id string) (Notes, error) {
	rows, err := DB.Query(
		"SELECT id, user_id, title, content, created_at FROM notes WHERE user_id = $1",
		&id,
	)

	if err != nil {
		return nil, err
	}

	notes := make(Notes, 0)

	for rows.Next() {
		err := rows.Scan(
			&note.ID,
			&note.UserID,
			&note.Title,
			&note.Content,
			&note.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (note Note) Create(newNote Note) (bool, error) {
	_, err := DB.Query(
		"INSERT INTO notes(user_id, title, content, created_at) VALUES ($1, $2, $3, $4)",
		&newNote.UserID,
		&newNote.Title,
		&newNote.Content,
		&newNote.CreatedAt,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (note Note) RemoveNote(id string) error {
	_, err := DB.Query(
		"DELETE FROM notes WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
