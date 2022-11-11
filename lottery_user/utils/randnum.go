package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandNum(len int) string {
	rand.Seed(time.Now().UnixNano())
	res := ""
	for i := 0; i < len; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}
