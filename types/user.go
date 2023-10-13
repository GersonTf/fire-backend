package types

import "go.mongodb.org/mongo-driver/bson/primitive"

const USER string = "user"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

func ValidateUser(u *User) bool { return true }

func (u *User) NewUser(name, email, password string) {
	u.Name = name
	u.Email = email
	u.Password = password
}
