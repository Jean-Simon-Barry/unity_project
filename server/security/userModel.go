package security

import "github.com/globalsign/mgo/bson"

// Data model
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Mail     string        `json:"mail"`
	Name     string        `json:"name"`
	Password string        `json:"password"`
	Roles    []string      `json:"roles"`
	Surname  string        `json:"surname"`
	Username string        `json:"username"`
	Token    string        `json:"token"`
}
