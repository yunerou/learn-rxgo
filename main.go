package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Customer struct {
	ID        int
	Age       int
	TaxNumber string
}

func producer(ch chan rxgo.Item) {
	for i := 0; i < 30; i++ {
		num := rand.Intn(20)
		ch <- rxgo.Of(&Customer{
			ID:        i,
			Age:       num,
			TaxNumber: "aa",
		})
	}
	close(ch)
}

func getTaxNumber(customer *Customer) (string, error) {
	time.Sleep(time.Millisecond * time.Duration(customer.Age*100))
	num := rand.Intn(1000)
	return fmt.Sprintf("%d-DD%d-%d", customer.ID, num, customer.Age), nil
}

func main() {
	// create
	ch := make(chan rxgo.Item)
	go producer(ch)

	observable := rxgo.FromChannel(ch)

	pipe := observable.
		// Filter(func(item interface{}) bool {
		// 	// Filter operation
		// 	customer := item.(*Customer)
		// 	return customer.Age > 2
		// }).
		Map(func(_ context.Context, item interface{}) (interface{}, error) {
			// Enrich operation
			customer := item.(*Customer)
			taxNumber, err := getTaxNumber(customer)
			if err != nil {
				return nil, err
			}
			customer.TaxNumber = taxNumber
			return customer, nil
		},
			// Create multiple instances of the map operator
			rxgo.WithPool(3),
			// Serialize the items emitted by their Customer.ID
			// rxgo.Serialize(func(item interface{}) int {
			// 	customer := item.(*Customer)
			// 	return customer.ID
			// }),
			rxgo.WithBufferedChannel(3)).
		ForEach(
			func(item interface{}) {
				fmt.Printf("next: %v\n", item)
			}, func(err error) {
				fmt.Printf("error: %v\n", err)
			}, func() {
				fmt.Println("done")
			})

	<-pipe
	// observe
	// for item := range pipe.Observe() {
	// 	if item.Error() {
	// 		fmt.Println("Customer Error :: ")
	// 	}
	// 	fmt.Println("Customer are :: ", item.V)
	// }
}
