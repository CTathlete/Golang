package main

import(
	"fmt"
	"sync/atomic"
	"errors"
	"sync"
	"math/rand"
	"time"
)

var (
	ErrShopClosed = errors.New("Shop is closed")
	// ErrShopClosed is returned when the shop is closed.

	ErrNoChair = errors.New("No chair available")
	// ErrNoChair is returned when there are no chairs available.
)

type customer struct {
	name string
} // costumer represents a customer to be serviced.

type Shop struct {
	open int32 // Determines whether the shop is open.
	chairs chan customer // A global queue of customers.
	synchroniser sync.WaitGroup 
	// A waitgroup to wait for all the chairs to get empty.
}

func ShopIsOpened(maxChairs int) *Shop {
	fmt.Println("Opening the shop")

	s := Shop{
		chairs: make(chan customer, maxChairs),
	}
	atomic.StoreInt32(&s.open, 1)

// The following block represents the barber's work.
// the sync will report that the work is done after 
// the barber is done serving all the customers.
	s.synchroniser.Add(1) 
	go func() {
		defer s.synchroniser.Done() // Reports that the barber is done.
		// Occures after all the customers are served.

		fmt.Println("Barber is ready to work")

		for currentCustomer := range s.chairs {
			s.serveCustomer(currentCustomer)
		}
	}()

	// The following block represents customers arriving at random times.
	// The arrival happens in a seperate goroutine (thread)
	// so customers arrive in parallel to the barber working.
	go func() {
		var id int64

		for {
			time.Sleep(time.Duration(rand.Intn(7))*time.Second)

			name := fmt.Sprintf("customer-%d", atomic.AddInt64(&id, 1))
			if err := s.newCustArrive(name); err!= nil {
				if err == ErrShopClosed {
					break
					// Customers stop arriving when the shop is closed.
				}
			}

		}
	}()

	return &s
}

func (s *Shop) Close() {
	// Closing prevents new customers from entering
	// the shop. After announcing that the shop is closed,
	// this method waits for all the chairs to get empty.
	// Only then, the shop is closed.
	fmt.Println("Closing the shop.")
	defer fmt.Println("Shop is closed.")

	atomic.StoreInt32(&s.open, 0)
	// Marking the shop as closed.

	close(s.chairs)	
	// marks that the channel of unserved
	// customers is closed.
	s.synchroniser.Wait()
	// We stay here until the counter of the 
	// waitgroup(synchroniser) is zero.
}

func (s *Shop) serveCustomer(cust customer) {
	fmt.Println("Customer", cust.name, "is served.")

	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	fmt.Println("Barber is done serving customer", cust.name)

	if len(s.chairs) == 0 && atomic.LoadInt32(&s.open) == 1 {
		fmt.Println("Shop is empty, barber is taking a nap.")
	}
}

func (s *Shop) newCustArrive(name string) error {
	if atomic.LoadInt32(&s.open) == 0 {
		fmt.Println("Customer", name, "leaves, since the shop is closed.")
		return ErrShopClosed
	}

	fmt.Println("Customer", name, "arrives.")

	select {
		case s.chairs <- customer{name:name}:
			fmt.Println("Customer" + name + "takes a seat and waits.")

		default:
			fmt.Println("Customer" + name + "leaves, since there are no chairs available.")
	}
	return nil
}







