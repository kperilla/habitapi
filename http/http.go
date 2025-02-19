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
    router.HandleFunc("GET /users/{id}", handler.HandleGetUser)
    router.HandleFunc("POST /users/", handler.HandleCreateUser)
    router.HandleFunc("GET /users/", handler.HandleGetUsers)
    router.HandleFunc("PUT /users/{id}", handler.HandleUpdateUser)
    router.HandleFunc("DELETE /users/{id}", handler.HandleDeleteUser)

    router.HandleFunc("GET /habit_groups/{id}", handler.HandleGetHabitGroup)
    router.HandleFunc("POST /habit_groups/", handler.HandleCreateHabitGroup)
    router.HandleFunc("GET /habit_groups/", handler.HandleGetHabitGroups)
    router.HandleFunc("GET /view/habit_groups/", handler.HandleGetHabitGroupsView)
    router.HandleFunc("PUT /habit_groups/{id}", handler.HandleUpdateHabitGroup)
    router.HandleFunc("DELETE /habit_groups/{id}", handler.HandleDeleteHabitGroup)


    router.HandleFunc("GET /habits/{id}", handler.HandleGetHabit)
    router.HandleFunc("POST /habits/", handler.HandleCreateHabit)
    router.HandleFunc("GET /habits/", handler.HandleGetHabits)
    router.HandleFunc("PUT /habits/{id}", handler.HandleUpdateHabit)
    router.HandleFunc("DELETE /habits/{id}", handler.HandleDeleteHabit)

    router.HandleFunc("GET /deeds/{id}", handler.HandleGetDeed)
    router.HandleFunc("POST /deeds/", handler.HandleCreateDeed)
    router.HandleFunc("GET /deeds/", handler.HandleGetDeeds)
    router.HandleFunc("PUT /deeds/{id}", handler.HandleUpdateDeed)
    router.HandleFunc("DELETE /deeds/{id}", handler.HandleDeleteDeed)

    router.HandleFunc("GET /rewards/{id}", handler.HandleGetReward)
    router.HandleFunc("POST /rewards/", handler.HandleCreateReward)
    router.HandleFunc("GET /rewards/", handler.HandleGetRewards)
    router.HandleFunc("PUT /rewards/{id}", handler.HandleUpdateReward)
    router.HandleFunc("DELETE /rewards/{id}", handler.HandleDeleteReward)

    router.HandleFunc("GET /", handler.HandleHealthcheck)

    v1 := http.NewServeMux()
    v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))
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
