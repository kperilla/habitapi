package habitapi

import (
	"time"

    "go.mongodb.org/mongo-driver/v2/bson"
)

type ApiResource interface {
    SetID(id string)
}

type User struct {
    ID bson.ObjectID `json:"_id" bson:"_id"`
    Name string `json:"name"`
    PointTotal int `json:"point_total"`
}

// func (u *User) SetID(id string) {
//     u.ID = id
// }

type Service[T any, CDTO any, UDTO any] interface {
    GetById(id string) (*T, error)
    List() ([]*T, error)
    Create(dto CDTO) (*T, error)
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
    Create(CreateHabitGroupDTO) (*HabitGroup, error)
    Update(string, UpdateHabitGroupDTO) (*HabitGroup, error)
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
    Create(CreateDeedDTO) (*Deed, error)
    Update(string, UpdateDeedDTO) (*Deed, error)
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
    Create(CreateRewardDTO) (*Reward, error)
    Update(string, UpdateRewardDTO) (*Reward, error)
    Delete(string) error
}
