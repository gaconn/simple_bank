package db

import (
	"math/rand"

	"github.com/quan12xz/simple_bank/util"
)

func RandomOwner() string {
	return util.RandomString(5)
}
func RandomBalance() int64 {
	return util.RandomInt(0, 1000)
}
func RandomCurrency() string {
	currency := []string{"USD", "EUR", "CAD"}
	n := len(currency)
	return currency[rand.Intn(n)]
}
