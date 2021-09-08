package user

import (
	"errors"

	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
)

// User struct for user
type User struct {
	Id   bson.ObjectId `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role"`
}

const (
	dbPath = "user.db"
)

// errors
var (
	ErrRecordInvalid = errors.New("record invalid")
)

// All rerieve all users from db
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

// One retrieve one user from db
func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	u := new(User)
	err = db.One("Id", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Delete delete user from db
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	u := new(User)
	err = db.One("Id", id, u)
	if err != nil {
		return err
	}
	return db.DeleteStruct(u)
}

// Save save user to db
func (u *User) Save() error {
	if err := u.Validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(u)
}

// Validate validate user
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
