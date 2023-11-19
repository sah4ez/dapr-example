package models

type User struct {
	ID   int
	Name string
}

func (User) Mock() User {
	return User{
		ID:   123,
		Name: "some user name mock",
	}
}
