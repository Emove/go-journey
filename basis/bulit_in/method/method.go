package method

import "fmt"

type User struct {
	Name string
	Age  int8
}

type Manager struct {
	User
}

func (user User) Values() {
	fmt.Printf("values %p, %#v\n", &user, user)
}

func (user *User) Pointer() {
	fmt.Printf("pointer %p, %#v\n", user, user)
}
