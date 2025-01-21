package habitapi

import (
	"time"
)

type ApiResource interface {
    SetID(id string)
}

type User struct {
    ID string
    Name string
    PointTotal int
}

// func (u *User) SetID(id string) {
//     u.ID = id
// }

// type Service[T any, DTO any] interface {
//     Create(dto DTO) (*T, string, error)
//     GetById(id string) (*T, error)
//     List() ([]*T, error)
//     Update(id string, dto DTO) (*T, error)
//     Delete(id string) error
// }

type UserService interface {
    GetById(string) (*User, error)
    List() ([]*User, error)
    Create(CreateUserDTO) (*User, error)
    Delete(string) error
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

// NOTE: I wasn't able to make this work from within the generic functions
// func (hg *HabitGroup) SetID(id string) {
//     hg.ID = id
// }

type HabitGroupService interface {
    GetById(string) (*HabitGroup, error)
    List() ([]*HabitGroup, error)
    Create(string) (*HabitGroup, string, error)
    Delete(string) error
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
