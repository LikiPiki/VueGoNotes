package note

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
  "time"
  "github.com/asaskevich/govalidator"
  "fmt"
)

type Note struct {
  Collection *mgo.Collection `json:"-" bson:"-"`

  Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
  User bson.ObjectId `json:"user" bson:"user"`

  Title string `json:"title" bson:"title"`
  Content string `json:"content" bson:"content"`
  CreatedDate time.Time `json:"date" bson:"date"`
  Tags []string `json:"tags" bson:"tags"`
}

func (note Note) ValidateNote() (bool, error) {
  fl, err := govalidator.ValidateStruct(&note)
  return fl, err
}

func (note *Note) CreateNote() (bool, error) {
  err := note.Collection.Insert(Note{
    Title: note.Title,
    Content: note.Content,
    CreatedDate: time.Now(),
    User: note.User,
  })
  if err != nil {
    return false, err
  }
  return true, nil
}

func (note *Note) UpdateNote() (bool, error) {
  err := note.Collection.UpdateId(note.Id, note)
  if err != nil {
    return false, err
  }
  return true, nil
}

func (note *Note) FindNote() (*Note, error) {
  err := note.Collection.FindId(note.Id).One(&note)
  if err != nil {
    return note, err
  }
  return note, nil
}

func (note *Note) DeleteNote() (bool, error) {
  err := note.Collection.RemoveId(note.Id)
  if err != nil {
    return false, err
  }
  return true, nil
}

func (note *Note) GetAllById() ([]Note, error) {
  notes := make([]Note, 0)
  fmt.Println("note", note.User)
  err := note.Collection.Find(bson.M{"user": note.User}).Limit(100).Sort("-date").All(&notes)
  return notes, err
}


