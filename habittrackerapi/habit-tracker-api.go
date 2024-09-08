package habittrackerapi

type User struct {
    Name string
}

type UserService interface {
    User(string) (*User, error)
    // Users() ([]*User, error)
    CreateUser(string) (*User, string, error)
    DeleteUser(string) error
}
