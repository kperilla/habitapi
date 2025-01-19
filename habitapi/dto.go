package habitapi

type CreateUserDTO struct {
    Name string
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
