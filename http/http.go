package http

import (
    "net/http"
    "log"

    "encoding/json"
    "github.com/rs/cors"
    "github.com/kperilla/habitapi/habitapi"
)

type Handler struct {
    UserService habitapi.UserService
    HabitGroupService habitapi.HabitGroupService
    HabitService habitapi.HabitService
    DeedService habitapi.DeedService
    RewardService habitapi.RewardService
}

func NewHandler(
    userService habitapi.UserService,
    habitGroupService habitapi.HabitGroupService,
    habitService habitapi.HabitService,
    deedService habitapi.DeedService,
    rewardService habitapi.RewardService,
) *Handler {
    return &Handler{
        UserService: userService,
        HabitGroupService: habitGroupService,
        HabitService: habitService,
        DeedService: deedService,
        RewardService: rewardService,
    }
}

type APIServer struct {
    addr string
}

func NewAPIServer(addr string) *APIServer {
    return &APIServer{addr: addr}
}

func (s *APIServer) Run(handler *Handler) error {
    router := http.NewServeMux()

    fs := http.FileServer(http.Dir("./views/static"))
    router.Handle("/static/", http.StripPrefix("/static/", fs))

    router.HandleFunc("GET /api/v1/users/{id}", handler.HandleGetUser)
    router.HandleFunc("POST /api/v1/users/", handler.HandleCreateUser)
    router.HandleFunc("GET /api/v1/users/", handler.HandleGetUsers)
    router.HandleFunc("PUT /api/v1/users/{id}", handler.HandleUpdateUser)
    router.HandleFunc("DELETE /api/v1/users/{id}", handler.HandleDeleteUser)

    router.HandleFunc("GET /api/v1/habit_groups/{id}", handler.HandleGetHabitGroup)
    router.HandleFunc("POST /api/v1/habit_groups/", handler.HandleCreateHabitGroup)
    router.HandleFunc("GET /api/v1/habit_groups/", handler.HandleGetHabitGroups)
    router.HandleFunc("PUT /api/v1/habit_groups/{id}", handler.HandleUpdateHabitGroup)
    router.HandleFunc("DELETE /api/v1/habit_groups/{id}", handler.HandleDeleteHabitGroup)
    router.HandleFunc("GET /habit_groups/", handler.HandleGetHabitGroupsView)

    router.HandleFunc("GET /api/v1/habits/{id}", handler.HandleGetHabit)
    router.HandleFunc("POST /api/v1/habits/", handler.HandleCreateHabit)
    router.HandleFunc("GET /api/v1/habits/", handler.HandleGetHabits)
    router.HandleFunc("PUT /api/v1/habits/{id}", handler.HandleUpdateHabit)
    router.HandleFunc("DELETE /api/v1/habits/{id}", handler.HandleDeleteHabit)
    router.HandleFunc("GET /habits/", handler.HandleGetHabitsView)

    router.HandleFunc("GET /api/v1/deeds/{id}", handler.HandleGetDeed)
    router.HandleFunc("POST /api/v1/deeds/", handler.HandleCreateDeed)
    router.HandleFunc("GET /api/v1/deeds/", handler.HandleGetDeeds)
    router.HandleFunc("PUT /api/v1/deeds/{id}", handler.HandleUpdateDeed)
    router.HandleFunc("DELETE /api/v1/deeds/{id}", handler.HandleDeleteDeed)
    router.HandleFunc("GET /deeds/", handler.HandleGetDeedsView)

    router.HandleFunc("GET /api/v1/rewards/{id}", handler.HandleGetReward)
    router.HandleFunc("POST /api/v1/rewards/", handler.HandleCreateReward)
    router.HandleFunc("GET /api/v1/rewards/", handler.HandleGetRewards)
    router.HandleFunc("PUT /api/v1/rewards/{id}", handler.HandleUpdateReward)
    router.HandleFunc("DELETE /api/v1/rewards/{id}", handler.HandleDeleteReward)
    router.HandleFunc("GET /rewards/", handler.HandleGetRewardsView)

    router.HandleFunc("GET /api/v1/", handler.HandleHealthcheck)
    router.HandleFunc("/", handler.HandleIndexView)

    // router.Handle("/static/", fs)
    // v1 := http.NewServeMux()
    // v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))
    ensureCorsHandler := cors.Default().Handler(router)

    server := http.Server{
        Addr: s.addr,
        Handler: ensureCorsHandler,
    }
    log.Println("Listening on ", s.addr)
    return server.ListenAndServe()
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
    if status != 200 {
        w.WriteHeader(status)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(v)
}
