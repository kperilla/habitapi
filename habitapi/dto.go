package habitapi

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// type DTO[T any] interface {
//     ToModel() T
// }
type DTO interface {
    ToModel() interface{}
}

type CreateUserDTO struct {
    Name string `json:"name" validate:"required"`
}

func (dto *CreateUserDTO) ToModel() interface{} {
    return User{
        ID: bson.NewObjectID(),
        Name: dto.Name,
        PointTotal: 0,
    }
}

func (dto *CreateUserDTO) Validate() error {
    return nil
}

// TODO: I may want a user-accessible and a svc-accessible version
type UpdateUserDTO struct {
    Name string `json:"name" bson:"name,omitempty"`
    PointTotal int `json:"point_total" bson:"point_total,omitempty"`
}

func (dto *UpdateUserDTO) ToModel() interface{} {
    return User{
        Name: dto.Name,
        PointTotal: dto.PointTotal,
    }
}

func (dto *UpdateUserDTO) Validate() error {
    return nil
}

type CreateHabitGroupDTO struct {
    Name string `json:"name" validate:"required"`
    Description string `json:"description"`
    UserId bson.ObjectID `json:"user_id" validate:"required"`
}

func (dto *CreateHabitGroupDTO) ToModel() interface{} {
    return HabitGroup{
        ID: bson.NewObjectID(),
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
    }
}

type UpdateHabitGroupDTO struct {
    Name string `json:"name" bson:"name,omitempty"`
    Description string `json:"description" bson:"description,omitempty"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id,omitempty"`
}

func (dto *UpdateHabitGroupDTO) ToModel() interface{} {
    return HabitGroup{
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
    }
}

type CreateHabitDTO struct {
    Name string `json:"name" bson:"name" validate:"required"`
    Description string `json:"description" bson:"description"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id" validate:"required"`
    HabitGroupId bson.ObjectID `json:"habit_group_id" bson:"habit_group_id" validate:"required"`
}

func (dto *CreateHabitDTO) ToModel() interface{} {
    return Habit{
        ID: bson.NewObjectID(),
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
        HabitGroupId: dto.HabitGroupId,
    }
}

type UpdateHabitDTO struct {
    Name string `json:"name" bson:"name,omitempty"`
    Description string `json:"description" bson:"description,omitempty"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id,omitempty"`
    HabitGroupId bson.ObjectID `json:"habit_group_id" bson:"habit_group_id,omitempty"`
}

func (dto *UpdateHabitDTO) ToModel() interface{} {
    return Habit{
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
    }
}

type CreateDeedDTO struct {
    Name string `json:"name" bson:"name" validate:"required"`
    Description string `json:"description" bson:"description"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id" validate:"required"`
    HabitId bson.ObjectID `json:"habit_id" bson:"habit_id" validate:"required"`
    Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

func (dto *CreateDeedDTO) ToModel() interface{} {
    fmt.Println()
    fmt.Println("Timestamp")
    fmt.Println(dto.Timestamp)
    timestamp := dto.Timestamp
    if timestamp.IsZero() {
        timestamp = time.Now()
    }

    return Deed{
        ID: bson.NewObjectID(),
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
        HabitId: dto.HabitId,
        Timestamp: timestamp,
    }
}

type UpdateDeedDTO struct {
    Name string `json:"name" bson:"name,omitempty"`
    Description string `json:"description" bson:"description,omitempty"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id,omitempty"`
    HabitId bson.ObjectID `json:"habit_id" bson:"habit_id,omitempty"`
    Timestamp time.Time `json:"timestamp" bson:"timestamp,omitempty"`
}

func (dto *UpdateDeedDTO) ToModel() interface{} {
    return Deed{
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
        HabitId: dto.HabitId,
        Timestamp: dto.Timestamp,
    }
}

type CreateRewardDTO struct {
    Name string `json:"name" bson:"name" validate:"required"`
    Description string `json:"description" bson:"description"`
    PointCost int `json:"point_cost" bson:"point_cost" validate:"required"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id" validate:"required"`
    IsEarned bool `json:"is_earned" bson:"is_earned"`
    EarnedTimestamp time.Time `json:"earned_timestamp" bson:"earned_timestamp"`
}

func (dto *CreateRewardDTO) ToModel() interface{} {
    timestamp := dto.EarnedTimestamp
    if !dto.IsEarned {
        timestamp = time.Time{}
    }
    return Reward{
        ID: bson.NewObjectID(),
        Name: dto.Name,
        Description: dto.Description,
        PointCost: dto.PointCost,
        UserId: dto.UserId,
        IsEarned: dto.IsEarned,
        EarnedTimestamp: timestamp,
    }
}

type UpdateRewardDTO struct {
    Name string `json:"name" bson:"name,omitempty"`
    Description string `json:"description" bson:"description,omitempty"`
    PointCost int `json:"point_cost" bson:"point_cost,omitempty"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id,omitempty"`
    IsEarned bool `json:"is_earned" bson:"is_earned,omitempty"`
    EarnedTimestamp time.Time `json:"earned_timestamp" bson:"earned_timestamp,omitempty"`
}

func (dto *UpdateRewardDTO) ToModel() interface{} {
    return Reward{
        Name: dto.Name,
        Description: dto.Description,
        PointCost: dto.PointCost,
        UserId: dto.UserId,
        IsEarned: dto.IsEarned,
        EarnedTimestamp: dto.EarnedTimestamp,
    }
}
