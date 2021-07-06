package main

import (
	"fmt"

	accounts "github.com/Heongilee/learngo/banking"
)

func main() {
	account := accounts.NewAccount("heongilee")
	account = account.Deposit(100)
	fmt.Println(account.GetBalance())
}
