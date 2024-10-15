package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func newAccount(FirstName, LastName string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		FirstName: FirstName,
		LastName:  LastName,
		Number:    int64(rand.Intn(1000000)),
	}
}
