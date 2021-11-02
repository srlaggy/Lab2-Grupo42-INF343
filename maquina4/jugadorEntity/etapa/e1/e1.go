package e1

import (
    "math/rand"
	"time"
	"fmt"
)

const (
	min = 1
	max = 10
)

func RandomNumber() int64{
	rand.Seed(time.Now().UnixNano())
	aux := rand.Intn(max - min) + min
	return int64(aux)
}

// muerto 0 - vivo 1
func LiveOrDead(num int64){
	if num==0{
		fmt.Printf("Tu estas muerto, BANG!!")
	} else {
		fmt.Printf("Te salvaste esta vez")
	}
}