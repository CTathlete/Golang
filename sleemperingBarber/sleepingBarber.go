package main

import(
	"fmt"
	"sync/atomic"
	"errors"
	"sync"
	"math/rand"
	"time"
)

func main() {
	const maxChairs = 3

	shop := ShopIsOpened(maxChairs)
	defer shop.close()
// Close the shop in 50 milliseconds.
	t = time.NewTimer(50 * time.Millisecond)
	<-t.c
}