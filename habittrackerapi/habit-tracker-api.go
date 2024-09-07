package habittrackerapi

type User struct {
    Id   string
    Name string
}

type UserService interface {
    User(string) (*User, error)
    Users() ([]*User, error)
    CreateUser(User) (*User, error)
    DeleteUser(string) error
}
