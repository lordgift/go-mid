package main

import (
	"fmt"
	"go-mid/mq"
	"time"
)

func main() {

	//go loop("first")
	//go loop("second")

	higherOrder(func(isOk bool) {
		if isOk {
			fmt.Print("99999")
		}
	})

	//s := services.CreateService()
	//r := services.CreateRouter(s)
	//r.Run()

	mq.Consume()
	mq.Publish()
}

func higherOrder(completion func(bool)) {
	completion(true)
}

func loop(my string) {
	for i := 1; i < 20; i++ {
		fmt.Printf("%s:%d\n", my, i)
		time.Sleep(1 * time.Second)
	}

}
