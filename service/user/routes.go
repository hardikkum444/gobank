package user

import(
    "github.com/gorilla/mux"
    "net/http"
)

type Handler struct {

}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/login", h.handleLogin).Methods("POST")
}

func(h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){

}
