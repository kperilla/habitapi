package habitapi

type DTO[T any] interface {
    ToModel() T
}

type CreateUserDTO struct {
    Name string
}

func (dto *CreateUserDTO) ToModel() User {
    return User{Name: dto.Name, PointTotal: 0}
}

type UpdateUserDTO struct {
    Name string
}

func (dto *UpdateUserDTO) ToModel() User {
    return User{Name: dto.Name}
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

type UpdateDeedDTO struct {
    Name string
    Description string
    UserId string
    HabitId string
    // Maybe want timestamp for testing??
    // Timestamp time.Time
}

type CreateRewardDTO struct {
    Name string
    Description string
    PointCost int
    UserId string
}

type UpdateRewardDTO struct {
    Name string
    Description string
    PointCost int
    UserId string
}
