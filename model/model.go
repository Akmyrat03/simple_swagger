package model

type User struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	IsActive bool   `json:"isActive"`
}

var users = []User{
	{Username: "john", FullName: "John Doe", Email: "john@example.com", IsActive: true},
	{Username: "jane", FullName: "Jane Doe", Email: "jane@example.com", IsActive: false},
}

func CreateUser(user User) error {
	users = append(users, user)
	return nil
}

type Country struct {
	Name string `json:"name"`
}
