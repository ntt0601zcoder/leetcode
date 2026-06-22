// Package carfleet solves LeetCode 853. Car Fleet.
// https://leetcode.com/problems/car-fleet/
package carfleet

import (
	"slices"
)

type Car struct {
	position  int
	speed     int
	timeTaken float64
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))

	for i := 0; i < len(position); i++ {
		cars = append(cars, Car{
			position:  position[i],
			speed:     speed[i],
			timeTaken: float64(target-position[i]) / float64(speed[i]),
		})
	}

	slices.SortFunc(cars, func(c1, c2 Car) int {
		return c2.position - c1.position
	})

	fleets := 0
	var lead float64

	for _, car := range cars {
		if car.timeTaken > lead {
			fleets++
			lead = car.timeTaken
		}
	}

	return fleets
}
