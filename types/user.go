// Package types provides shared data types used across the application.
// It offers a centralized location for defining and working with types
// that are utilized in multiple areas of the application, ensuring
// consistency and ease of maintenance.
package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const USER string = "user"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	LastName string             `json:"lastName" bson:"lastName"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

func ValidateUser(u *User) bool { return true }

func (u *User) NewUser(name, lastName, email, password string) {
	u.Name = name
	u.LastName = lastName
	u.Email = email
	u.Password = password
}

func (u *User) String() string {
	return fmt.Sprintf("%s %s", u.Name, u.LastName)
}
