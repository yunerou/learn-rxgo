package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/reactivex/rxgo/v2"
)

type DB struct {
	id int
}

func (d *DB) GetID() int {
	return d.id
}

func connectDB() (*DB, error) {
	num := rand.Intn(10)
	fmt.Printf("DB number %d \n", num)
	if num%5 == 0 {
		return &DB{
			id: num,
		}, nil
	}
	return nil, fmt.Errorf("connect error")
}

func test2_1() {
	observable := rxgo.Defer([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		db, err := connectDB()
		if err == nil {
			next <- rxgo.Of(db)
		} else {
			next <- rxgo.Error(err)
		}
	}}).
		Retry(10, func(err error) bool {
			return err != nil
		}).
		ForEach(
			func(item interface{}) {
				fmt.Println("next")
				fmt.Printf("DB successful: %d \n", item.(*DB).GetID())
			}, func(err error) {
				fmt.Println("error")
				fmt.Printf("error: %v \n", err)
			}, func() {
				fmt.Println("done")
			})
	<-observable
}

func test2_2() {
	observable := rxgo.Defer([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		db, err := connectDB()
		if err == nil {
			next <- rxgo.Of(db)
		} else {
			next <- rxgo.Error(err)
		}
	}}).
		Retry(10, func(err error) bool {
			return err != nil
		}).
		Take(1)

	item := <-observable.Observe()
	if item.E != nil {
		fmt.Println("error")
		panic("error")
	}
	db := item.V.(*DB)
	fmt.Println(db.GetID())
}
