// package main

// // SAGA example ^^^^ do not work

// import (
// 	"context"
// 	"fmt"

// 	"github.com/reactivex/rxgo/v2"
// )

// type Order struct {
// 	id      int
// 	product []int
// 	price   int
// 	status  string
// }
// type Payer struct {
// 	id     int
// 	name   string
// 	credit int
// }
// type Inventory struct {
// 	product []int
// }
// type CreateOrderEvent struct {
// 	id    int
// 	order *Order
// 	payer *Payer
// }

// func doPayment(e *CreateOrderEvent) (func(), error) {
// 	// make a payment
// 	paymentAction := func() {
// 		e.payer.credit -= e.order.price
// 		fmt.Printf("Payment successful: ID=%d $%d \n", e.order.id, e.order.price)
// 	}
// 	revertPaymentAction := func() {
// 		e.payer.credit += e.order.price
// 		fmt.Printf("Revert Payment successful: ID=%d $%d \n", e.order.id, e.order.price)
// 	}
// 	// Do
// 	if e.payer.credit >= e.order.price {
// 		paymentAction()
// 		return revertPaymentAction, nil
// 	}
// 	return func() {}, fmt.Errorf("not enough credit")
// }

// func doUpdateStatus(e *CreateOrderEvent) (func(), error) {
// 	oldStatus := e.order.status
// 	// Update status
// 	updateStatusAction := func() {
// 		e.order.status = "paid"
// 		fmt.Printf("Update status successful: ID=%d $%s \n", e.order.id, e.order.status)
// 	}
// 	revertAction := func() {
// 		e.order.status = oldStatus
// 		fmt.Printf("Revert update status successful: ID=%d $%s \n", e.order.id, e.order.status)
// 	}
// 	// Do
// 	updateStatusAction()
// 	return revertAction, nil
// }

// func doCheckProduct(e *CreateOrderEvent, inventory *Inventory) (func(), error) {
// 	// check product
// 	checkProductAction := func() bool {
// 		fmt.Printf("Check product successful: ID=%d \n", e.order.id)
// 		return isSubset(e.order.product, inventory.product)
// 	}
// 	// Do
// 	if checkProductAction() {
// 		return func() {}, nil
// 	}
// 	return func() {}, fmt.Errorf("not enough product")
// }

// func test3() {
// 	// Dump data
// 	inventory := &Inventory{
// 		product: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
// 	}
// 	order := Order{
// 		id:      1,
// 		product: []int{2, 3},
// 		price:   100,
// 		status:  "pending",
// 	}
// 	payer := Payer{
// 		id:     1,
// 		name:   "John",
// 		credit: 1000,
// 	}
// 	createOrderEvent := CreateOrderEvent{
// 		id:    1,
// 		order: &order,
// 		payer: &payer,
// 	}
// 	// Create pipeline
// 	ch := make(chan rxgo.Item, 1)
// 	ch <- rxgo.Of(createOrderEvent)
// 	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy)
// 	//
// 	payment := observable.Map(func(_ context.Context, e interface{}) (interface{}, error) {
// 		return doPayment(e.(*CreateOrderEvent))
// 	})
// 	//checkproduct
// 	//status
// }

// func isSubset(first, second []int) bool {
// 	set := make(map[int]int)
// 	for _, value := range second {
// 		set[value] += 1
// 	}

// 	for _, value := range first {
// 		if count, found := set[value]; !found {
// 			return false
// 		} else if count < 1 {
// 			return false
// 		} else {
// 			set[value] = count - 1
// 		}
// 	}
// 	return true
// }
