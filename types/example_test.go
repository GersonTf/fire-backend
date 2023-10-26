package types

import "fmt"

func ExampleUser_NewUser() {
	u := &User{}
	u.NewUser("John", "Fooson", "john@example.com", "password123")

	fmt.Println(u)
	// Output: John Fooson
}
