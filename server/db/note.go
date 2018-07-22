package db

type Note struct {
	ID     uint `json:"id"`
	UserID uint `json:""`

	Title   string `json:""`
	Content string `json:""`
}
type Notes []Note

func (note Note) GetAll(id string) (Notes, error) {
	rows, err := DB.Query(
		"SELECT id, user_id, title, content FROM notes WHERE user_id = $1",
		&id,
	)

	if err != nil {
		return nil, err
	}

	var notes Notes

	for rows.Next() {
		err := rows.Scan(
			&note.ID,
			&note.UserID,
			&note.Title,
			&note.Content,
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
		"INSERT INTO notes(title, content) VALUES ($1, $2)",
		&newNote.Title,
		&newNote.Content,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}
