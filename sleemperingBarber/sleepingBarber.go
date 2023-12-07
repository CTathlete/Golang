package main

import(
	"time"
)

func main() {
	const maxChairs = 100

	shop := ShopIsOpened(maxChairs)
	defer shop.Close()
// Close the shop in 50 seconds.
	t := time.NewTimer(50 * time.Second)
	<-t.C
}