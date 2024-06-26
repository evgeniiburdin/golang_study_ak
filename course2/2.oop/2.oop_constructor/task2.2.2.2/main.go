package main

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func WithUsername(Username string) UserOption {
	return func(u *User) {
		u.Username = Username
	}
}

func WithEmail(Email string) UserOption {
	return func(u *User) {
		u.Email = Email
	}
}

func WithRole(Role string) UserOption {
	return func(u *User) {
		u.Role = Role
	}
}

func NewUser(ID int, options ...UserOption) *User {
	user := &User{
		ID: ID,
	}
	for _, option := range options {
		option(user)
	}
	return user
}
