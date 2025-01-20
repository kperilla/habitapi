package habitapi

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
    return User{Name: dto.Name, PointTotal: 0}
}

type CreateHabitGroupDTO struct {
    Name string
    Description string
    UserId string
}

type CreateHabitDTO struct {
    Name string
    Description string
    UserId string
    HabitGroupId string
}

type CreateDeedDTO struct {
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
