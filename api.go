package main

import(
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "log"
    "fmt"
)

type APIServer struct {
    listenAddr string
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
    w.WriteHeader(status)
    w.Header().Set("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
    Error string
}

type apifunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandlerFunc(f apifunc) (http.HandlerFunc) {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := f(w, r); err != nil {
            writeJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
        }
    }
}

func newAPIServer(listenAddr string) *APIServer {
    return &APIServer{
        listenAddr: listenAddr,
    }
}

func (s *APIServer) Run() {
    router := mux.NewRouter()

    router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount))

    log.Println("JSON api server running on port ", s.listenAddr)

    log.Panic(http.ListenAndServe(s.listenAddr, router))
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
    if r.Method == http.MethodGet {
        return s.handleGetAccount(w, r)
    }
    if r.Method == http.MethodPost {
        return s.handleCreateAccount(w, r)
    }
    if r.Method == http.MethodDelete {
        return s.handleDeleteAccount(w, r)
    }

    return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
    return nil
}
