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

var countries = []Country{
	{Name: "USA"},
	{Name: "Canada"},
	{Name: "UK"},
}

func ListUsers() []User {
	return users
}

func GetUser(username string) (*User, error) {
	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}

	}
}

func ListCountries() []Country {
	return countries
}
