package db

//easyjson:json -snake_case
type Tag struct {
	ID uint `json:"id"`

	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
}

type Tags []Tag

func (u Tag) GetAll(userID string) (Tags, error) {

	rows, err := DB.Query(
		"SELECT id, name, color FROM tags WHERE user_id = $1",
		userID,
	)

	if err != nil {
		return nil, err
	}

	var tags Tags
	var tag Tag

	for rows.Next() {
		err = rows.Scan(
			&tag.ID,
			&tag.UserID,
			&tag.Name,
			&tag.Color,
		)

		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)

	}

	defer rows.Close()

	return tags, nil
}

func (tag Tag) Create() error {
	_, err := DB.Query(
		"INSERT INTO tags (user_id, name, color) VALUES ($1, $2, $3)",
		tag.UserID, tag.Name, tag.Color,
	)
	if err != nil {
		return err
	}
	return nil
}

func (tag Tag) Update() error {
	_, err := DB.Query(
		"UPDATE tags SET (name = $2, color = $3) WHERE id = $1",
		tag.ID,
		tag.Name,
		tag.Color,
	)

	if err != nil {
		return err
	}

	return nil
}

func (tag Tag) Remove(id string) error {
	_, err := DB.Query(
		"DELETE FROM tags WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
