package db

import "github.com/quan12xz/simple_bank/util"

func RandomAmountTransfer() int64 {
	return util.RandomInt(1, 1000)
}
