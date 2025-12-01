package main

import (
	"fmt"
	"sync"
)

/*
    🍔 Food Delivery App — Daily Revenue Problem Statement
    You are building a mini revenue tracker for a food-delivery company (like Swiggy/Zomato).

	👨‍🍳 There are 3 delivery boys:

	Delivery Boy	Money Collected Per Hour
	Arun			₹200
	Vivek			₹150
	Raju			₹120

	🕗 They each work for 8 hours a day.

*/

var wg sync.WaitGroup

type Delivery struct {
	Name                  string
	MoneyCollectedPerHour int
}

func main() {
	var totalRevenue int
	var revenue sync.Mutex
	fmt.Printf("Initial Total Revenue: ₹%d.00\n", totalRevenue)
	deliveries := []Delivery{
		{Name: "Arun", MoneyCollectedPerHour: 200},
		{Name: "Vivek", MoneyCollectedPerHour: 150},
		{Name: "Raju", MoneyCollectedPerHour: 120},
	}
	wg.Add(len(deliveries))
	for _, delivery := range deliveries {
		go func(delivery Delivery) {
			defer wg.Done()

			for hour := range 8 {
				revenue.Lock()
				temp := totalRevenue
				temp += delivery.MoneyCollectedPerHour
				totalRevenue = temp
				revenue.Unlock()
				fmt.Printf("In Hour %d, %s collected ₹%d.00\n", hour, delivery.Name, delivery.MoneyCollectedPerHour)
			}

		}(delivery)

	}
	wg.Wait()
	fmt.Printf("Final Total Revenue after 8 hours: ₹%d.00\n", totalRevenue)
}
