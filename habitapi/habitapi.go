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

type Service[T any, CDTO any, UDTO any] interface {
    GetById(id string) (*T, error)
    List() ([]*T, error)
    Create(dto CDTO) (*T, string, error)
    Update(id string, dto UDTO) (*T, error)
    Delete(id string) error
}

type UserService interface {
    GetById(string) (*User, error)
    List() ([]*User, error)
    Create(CreateUserDTO) (*User, error)
    Update(string, UpdateUserDTO) (*User, error)
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
    GetById(string) (*Habit, error)
    List() ([]*Habit, error)
    Create(CreateHabitDTO) (*Habit, error)
    Update(string, UpdateHabitDTO) (*Habit, error)
    Delete(string) error
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
    Create(CreateHabitGroupDTO) (*HabitGroup, string, error)
    Update(string, UpdateHabitGroupDTO) (*HabitGroup, string, error)
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

type DeedService interface {
    GetById(string) (*Deed, error)
    List() ([]*Deed, error)
    Create(CreateDeedDTO) (*Deed, string, error)
    Update(string, UpdateDeedDTO) (*Deed, string, error)
    Delete(string) error
}

type Reward struct {
    ID string
    Name string
    Description string
    PointCost int
    UserId string
}

type RewardService interface {
    GetById(string) (*Reward, error)
    List() ([]*Reward, error)
    Create(CreateRewardDTO) (*Reward, string, error)
    Update(string, UpdateRewardDTO) (*Reward, string, error)
    Delete(string) error
}
