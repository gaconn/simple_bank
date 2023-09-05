package db

import "github.com/quan12xz/simple_bank/util"

func RandomAmount() int64 {
	return util.RandomInt(0, 1000)
}
