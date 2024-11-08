package habitapi

type User struct {
    Name string
}

type UserService interface {
    User(string) (*User, error)
    Users() ([]*User, error)
    CreateUser(string) (*User, string, error)
    DeleteUser(string) error
}

type Habit struct {
    Name string
}

type HabitService interface {
    Habit(string) (*Habit, error)
    // Habits() ([]*Habit, error)
    CreateHabit(string) (*Habit, string, error)
    DeleteHabit(string) error
}

type HabitGroup struct {
    Name string
}

type HabitGroupService interface {
    HabitGroup(string) (*HabitGroup, error)
    // HabitGroups() ([]*HabitGroup, error)
    CreateHabitGroup(string) (*HabitGroup, string, error)
    DeleteHabitGroup(string) error
}
