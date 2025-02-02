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
    Name string `json:"name" bson:"name,omitempty"`
    PointTotal int `json:"point_total" bson:"point_total,omitempty"`
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
    ID bson.ObjectID `json:"_id" bson:"_id"`
    Name string `json:"name" bson:"name,omitempty"`
    Description string `json:"description" bson:"description,omitempty"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id,omitempty"`
    HabitGroupId bson.ObjectID `json:"habit_group_id" bson:"habit_group_id,omitempty"`
}

type HabitService interface {
    GetById(string) (*Habit, error)
    List() ([]*Habit, error)
    Create(CreateHabitDTO) (*Habit, error)
    Update(string, UpdateHabitDTO) (*Habit, error)
    Delete(string) error
}

type HabitGroup struct {
    ID bson.ObjectID `json:"_id" bson:"_id"`
    Name string `json:"name" bson:"name,omitempty"`
    Description string `json:"description" bson:"description,omitempty"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id,omitempty"`
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
    ID bson.ObjectID `json:"_id" bson:"_id"`
    Timestamp time.Time `json:"timestamp" bson:"timestamp,omitempty"`
    Name string `json:"name" bson:"name,omitempty"`
    Description string `json:"description" bson:"description,omitempty"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id,omitempty"`
    HabitId bson.ObjectID `json:"habit_id" bson:"habit_id,omitempty"`
}

type DeedService interface {
    GetById(string) (*Deed, error)
    List() ([]*Deed, error)
    Create(CreateDeedDTO) (*Deed, error)
    Update(string, UpdateDeedDTO) (*Deed, error)
    Delete(string) error
}

type Reward struct {
    ID bson.ObjectID `json:"_id" bson:"_id"`
    Name string `json:"name" bson:"name,omitempty"`
    Description string `json:"description" bson:"description,omitempty"`
    PointCost int `json:"point_cost" bson:"point_cost,omitempty"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id,omitempty"`
    IsEarned bool `json:"is_earned" bson:"is_earned,omitempty"`
    EarnedTimestamp time.Time `json:"earned_timestamp" bson:"earned_timestamp,omitempty"`
}

type RewardService interface {
    GetById(string) (*Reward, error)
    List() ([]*Reward, error)
    Create(CreateRewardDTO) (*Reward, error)
    Update(string, UpdateRewardDTO) (*Reward, error)
    Delete(string) error
}
