package db

import (
	"fmt"
	"log"
)

type Tag struct {
	ID uint `json:"id"`

	UserID uint   `json:"user_id"`
	Name   string `json:"tag_name"`
	Color  string `json:"color"`
}

type Tags []Tag

func (u Tag) GetAll(userID string) (Tags, error) {

	rows, err := DB.Query(
		"SELECT id, tag_name, color FROM tags WHERE user_id = $1",
		userID,
	)

	if err != nil {
		return nil, err
	}

	tags := make(Tags, 0)
	var tag Tag

	for rows.Next() {
		err = rows.Scan(
			&tag.ID,
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
	fmt.Println("TAG", tag)
	fmt.Println("DB ", DB)
	_, err := DB.Query(
		"INSERT INTO tags (user_id, tag_name, color) VALUES ($1, $2, $3)",
		&tag.UserID, &tag.Name, &tag.Color,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (tag Tag) Update(userID string) error {
	_, err := DB.Query(
		"UPDATE tags SET (tag_name = $2, color = $3) WHERE id = $1, user_id = $2",
		&tag.ID,
		&tag.Name,
		&tag.Color,
		&userID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (tag Tag) Remove(userID string) error {
	_, err := DB.Query(
		"DELETE FROM tags WHERE id = $1, user_id = $2",
		tag.ID, userID,
	)
	if err != nil {
		return err
	}

	return nil
}
