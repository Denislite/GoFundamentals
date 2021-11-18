package main

import (
	"fmt"
	"math/rand"
)

const secondsPerDay = 86400

func main() {

	distance := 62100000
	trip_type := ""
	company := ""
	fmt.Printf("%-14v %2v %-11v %v\n", "Spaceline", "Days", "Trip Type", "Price")
	fmt.Println("=====================================")

	for count := 1; count < 10; count++ {

		switch rand.Intn(3) {
		case 0:
			company = "Virgin Galactic"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Space Advetures"
		}

		speed := rand.Intn(15) + 16
		days := distance / speed / secondsPerDay
		price := speed + 20

		switch rand.Intn(2) {
		case 0:
			trip_type = "One-way"
		case 1:
			trip_type = "Round-trip"
			price *= 2
		}

		fmt.Printf("%-15v %3v %-11v $%4v \n", company, days, trip_type, price)
	}
}
