package user

import (
	"errors"

	"github.com/asdine/storm"
	// "github.com/mongodb/mongo-go-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	// "github.com/asdine/storm"
	// "github.com/mongodb/mongo-go-driver/bson/bsontype"
)

// user hold data fora single user
type User struct {
	Id   bsontype.Type `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role"`
}

const (
	dbPath = "userd.db"
)

var (
	ErrRecordInvalid = errors.New("record is invalid")
)

// All retives all the users
func All() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func One(id bsontype.Type) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func Delete(id bsontype.Type) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return err
	}
	return db.DeleteStruct(u)
}

// saves update or create a given record in a database
func (u *User) Save() error {
	if err := u.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(u)
}

// vaidate makes sure the the record contains alid data
func (u *User) validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
