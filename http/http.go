package http

import (
    "net/http"
    "log"

    "encoding/json"
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
    router.HandleFunc("GET /", handler.HandleHealthcheck)
    router.HandleFunc("GET /users/{id}", handler.HandleGetUser)
    router.HandleFunc("GET /users/", handler.HandleGetUsers)
    router.HandleFunc("POST /users/", handler.HandleCreateUser)
    // TODO: Delete Users
    // router.HandleFunc("DELETE /users/{id}", handler.HandleDeleteUser)
    router.HandleFunc("GET /habit_groups/{id}", handler.HandleGetHabitGroup)
    router.HandleFunc("GET /habit_groups/", handler.HandleGetHabitGroups)
    router.HandleFunc("POST /habit_groups/", handler.HandleCreateHabitGroup)

    router.HandleFunc("GET /habits/{id}", handler.HandleGetHabit)
    router.HandleFunc("GET /habits/", handler.HandleGetHabit)
    router.HandleFunc("POST /habits/", handler.HandleCreateHabit)

    router.HandleFunc("GET /deeds/{id}", handler.HandleGetDeed)
    router.HandleFunc("GET /deeds/", handler.HandleGetDeeds)
    router.HandleFunc("POST /deeds/", handler.HandleCreateDeed)

    router.HandleFunc("GET /rewards/{id}", handler.HandleGetReward)
    router.HandleFunc("GET /rewards/", handler.HandleGetRewards)
    router.HandleFunc("POST /rewards/", handler.HandleCreateReward)

    v1 := http.NewServeMux()
    v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

    server := http.Server{
        Addr: s.addr,
        Handler: router,
    }
    log.Println("Listening on ", s.addr)
    return server.ListenAndServe()
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
    w.WriteHeader(status)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(v)
}
