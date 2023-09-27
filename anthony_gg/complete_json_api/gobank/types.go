package main

import "math/rand"

type Account struct {
	Id int
	FirstName string
	LastName string
	Number int64
	Balance int64
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		Id: rand.Intn(99999), 
		FirstName: firstName, 
		LastName: lastName,
		Number: int64(rand.Intn(9999999)),
	}
}