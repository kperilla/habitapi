package habitapi

import (
	"time"
)

type User struct {
    ID string
    Name string
    PointTotal int
}

type UserService interface {
    User(string) (*User, error)
    Users() ([]*User, error)
    CreateUser(CreateUserDTO) (*User, string, error)
    DeleteUser(string) error
}

type Habit struct {
    ID string
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
    ID string
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
    ID string
    Timestamp time.Time
    Name string
    Description string
    UserId string
    HabitId string
}

type Reward struct {
    ID string
    Name string
    Description string
    PointCost int
    UserId string
}
