package main

import(
    "fmt"
    "net/http"
)

type APIserver struct {
    listenAddr string
}

func NewAPIServer(listenAddr string) *APIserver {
    return &APIserver {
        listenAddr: listenAddr,
    }
}

func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
}
