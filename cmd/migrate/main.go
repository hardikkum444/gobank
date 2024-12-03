package main

import(
    "github.com/hardikkum444/go-backend-api/cmd/api"
    "log"
)

func main() {
    server := api.NewAPIServer(":8000", nil)
    if err := server.Run(); err != nil {
        log.Fatal(err)
    }
}
