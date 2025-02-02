package habitapi

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DTO[T any] interface {
    ToModel() T
}

type CreateUserDTO struct {
    Name string `json:"name" validate:"required"`
}

func (dto *CreateUserDTO) ToModel() User {
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

func (dto *UpdateUserDTO) ToModel() User {
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

func (dto *CreateHabitGroupDTO) ToModel() HabitGroup {
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

func (dto *UpdateHabitGroupDTO) ToModel() HabitGroup {
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

func (dto *CreateHabitDTO) ToModel() Habit {
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

func (dto *UpdateHabitDTO) ToModel() Habit {
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

func (dto *CreateDeedDTO) ToModel() Deed {
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
    // Maybe want timestamp for testing??
    // Timestamp time.Time
}

func (dto *UpdateDeedDTO) ToModel() Deed {
    return Deed{
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
        HabitId: dto.HabitId,
    }
}

type CreateRewardDTO struct {
    Name string `json:"name" bson:"name" validate:"required"`
    Description string `json:"description" bson:"description"`
    PointCost int `json:"point_cost" bson:"point_cost" validate:"required"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id" validate:"required"`
}

func (dto *CreateRewardDTO) ToModel() Reward {
    return Reward{
        ID: bson.NewObjectID(),
        Name: dto.Name,
        Description: dto.Description,
        PointCost: dto.PointCost,
        UserId: dto.UserId,
    }
}

type UpdateRewardDTO struct {
    Name string `json:"name" bson:"name,omitempty"`
    Description string `json:"description" bson:"description,omitempty"`
    PointCost int `json:"point_cost" bson:"point_cost,omitempty"`
    UserId bson.ObjectID `json:"user_id" bson:"user_id,omitempty"`
}

func (dto *UpdateRewardDTO) ToModel() Reward {
    return Reward{
        Name: dto.Name,
        Description: dto.Description,
        PointCost: dto.PointCost,
        UserId: dto.UserId,
    }
}
