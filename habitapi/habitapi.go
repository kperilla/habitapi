package habitapi

import "time"

type User struct {
    Name string
    PointTotal int
}

type UserService interface {
    User(string) (*User, error)
    Users() ([]*User, error)
    CreateUser(string) (*User, string, error)
    DeleteUser(string) error
}

type Habit struct {
    Name string
    Description string
    // Hex representation of ObjectId in order to keep imports clean
    UserId string
    HabitGroupId string
}

type HabitService interface {
    Habit(string) (*Habit, error)
    // Habits() ([]*Habit, error)
    CreateHabit(string) (*Habit, string, error)
    DeleteHabit(string) error
}

type HabitGroup struct {
    Name string
    Description string
    UserId string
}

type HabitGroupService interface {
    HabitGroup(string) (*HabitGroup, error)
    // HabitGroups() ([]*HabitGroup, error)
    CreateHabitGroup(string) (*HabitGroup, string, error)
    DeleteHabitGroup(string) error
}

type Deed struct {
    Timestamp time.Time
    Name string
    Description string
    UserId string
    HabitId string
}

type Reward struct {
    Name string
    Description string
    PointCost int
    UserId string
}
