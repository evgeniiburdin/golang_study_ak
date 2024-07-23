package repository

type UserRepository interface {
	CreateUser(username, password string) error
	ReadUser(username string) (string, error)
	UpdateUser(username, newPassword string) error
	DeleteUser(username string) error
}

type UserRepositoryProxy interface {
	CreateUser(username, password string) error
	ReadUser(username string) (string, error)
	UpdateUser(username, newPassword string) error
	DeleteUser(username string) error
}

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Unset(key string) error
}
