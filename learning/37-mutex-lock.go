package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

type Account struct {
	balance int
	lock    sync.Mutex
}

func (acc *Account) GetBalance() int {
	acc.lock.Lock()
	defer acc.lock.Unlock()

	return acc.balance
}

func (acc *Account) Withdraw(val int) {
	acc.lock.Lock()
	defer acc.lock.Unlock()

	if val > acc.balance {
		println("Error: not enough balance in account")
	} else {
		acc.balance -= val
		printf("%d withdrawn successfully\n", val)
		printf("Updated balance: %d\n", acc.balance)
	}
}

func main() {
	var acc Account
	acc.balance = 100
	println("Initial Balance:", acc.GetBalance())
	println()

	random := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 12; i++ {
		amount := random.Intn(16) + 5
		go acc.Withdraw(amount)
	}

	time.Sleep(2 * time.Second)
}
