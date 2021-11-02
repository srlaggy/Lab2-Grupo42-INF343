package e1

import (
    "math/rand"
	"time"
)

const (
	min = 6
	max = 10
)

func RandomNumber() int64{
	rand.Seed(time.Now().UnixNano())
	aux := rand.Intn(max - min) + min
	return int64(aux)
}