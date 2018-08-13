package db

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDBTags(t *testing.T) {
	db := Connect("test_notes")

	tag := Tag{
		Name:   "test",
		Color:  "default",
		UserID: 1,
	}

	err := tag.Create()
	if err != nil {
		t.Error("error in tag.Create()", err)
	}

	tag2 := Tag{}

	err = db.QueryRow(
		"SELECT tag_name, color, user_id FROM tags WHERE tag_name=$1",
		&tag.Name,
	).Scan(
		&tag2.Name,
		&tag2.Color,
		&tag2.UserID,
	)

	if !reflect.DeepEqual(tag, tag2) {
		t.Error("expected\n" + fmt.Sprint(tag) + "\ngot\n" + fmt.Sprint(tag2))
	}

}
