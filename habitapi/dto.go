package habitapi

import "go.mongodb.org/mongo-driver/v2/bson"

type DTO[T any] interface {
    ToModel() T
}

type CreateUserDTO struct {
    Name string `json:"name"`
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
    Name string
    Description string
    UserId string
}

func (dto *CreateHabitGroupDTO) ToModel() HabitGroup {
    return HabitGroup{
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
    }
}

type UpdateHabitGroupDTO struct {
    Name string
    Description string
    UserId string
}

func (dto *UpdateHabitGroupDTO) ToModel() HabitGroup {
    return HabitGroup{
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
    }
}

type CreateHabitDTO struct {
    Name string
    Description string
    UserId string
    HabitGroupId string
}

func (dto *CreateHabitDTO) ToModel() Habit {
    return Habit{
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
    }
}

type UpdateHabitDTO struct {
    Name string
    Description string
    UserId string
}

func (dto *UpdateHabitDTO) ToModel() Habit {
    return Habit{
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
    }
}

type CreateDeedDTO struct {
    Name string
    Description string
    UserId string
    HabitId string
    // Maybe want timestamp for testing??
    // Timestamp time.Time
}

func (dto *CreateDeedDTO) ToModel() Deed {
    return Deed{
        Name: dto.Name,
        Description: dto.Description,
        UserId: dto.UserId,
        HabitId: dto.HabitId,
    }
}

type UpdateDeedDTO struct {
    Name string
    Description string
    UserId string
    HabitId string
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
    Name string
    Description string
    PointCost int
    UserId string
}

func (dto *CreateRewardDTO) ToModel() Reward {
    return Reward{
        Name: dto.Name,
        Description: dto.Description,
        PointCost: dto.PointCost,
        UserId: dto.UserId,
    }
}

type UpdateRewardDTO struct {
    Name string
    Description string
    PointCost int
    UserId string
}

func (dto *UpdateRewardDTO) ToModel() Reward {
    return Reward{
        Name: dto.Name,
        Description: dto.Description,
        PointCost: dto.PointCost,
        UserId: dto.UserId,
    }
}
