package main

var (
	ErrShopClosed = errors.New("Shop is closed")
	// ErrShopClosed is returned when the shop is closed.

	ErrNoChair = errors.New("No chair available")
	// ErrNoChair is returned when there are no chairs available.
)

type costumer struct {
	name string
} // costumer represents a customer to be serviced.

type Shop struct {
	open int32 // Determines whether the shop is open.
	chairs chan customer // A global queue of customers.
	synchroniser sync.waitGroup // A waitgroup to wait for all the chairs to get empty.
}

func OpenShop(maxChairs int) *Shop {
	fmt.Println("Opening the shop")

	s := Shop{
		chairs: make(chan customer, maxChairs)
	}
	atomic.(&s.open).StoreInt32(1)

// The following block represents the barber's work.
// the sync will report that the work is done after 
// the barber is done serving all the customers.
	s.synchroniser.Add(1) 
	go func() {
		defer s.synchroniser.Done() // Reports that the barber is done.
		// Occures after all the customers are served.

		fmt.Println("Barber is ready to work")

		for currentCustomer := range s.chairs {
			s.serviceCustomer(currentCustomer)
		}
	}()
}









