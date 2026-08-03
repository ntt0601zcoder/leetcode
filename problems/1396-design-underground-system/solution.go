// Package designundergroundsystem solves LeetCode 1396. Design Underground System.
// https://leetcode.com/problems/design-underground-system/
package designundergroundsystem

type CheckInData struct {
	from    string
	startAt int
}

type RouteData struct {
	totalTime int
	count     int
}

type UndergroundSystem struct {
	checkInInfo map[int]CheckInData
	routeStats  map[string]RouteData
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		checkInInfo: make(map[int]CheckInData, 8),
		routeStats:  make(map[string]RouteData, 8),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.checkInInfo[id] = CheckInData{from: stationName, startAt: t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	checkInData := this.checkInInfo[id]

	route, existed := this.routeStats[this.getRouteID(checkInData.from, stationName)]
	if !existed {
		route = RouteData{}
	}
	route.count++
	route.totalTime += t - checkInData.startAt

	this.routeStats[this.getRouteID(checkInData.from, stationName)] = route
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	route := this.routeStats[this.getRouteID(startStation, endStation)]

	return float64(route.totalTime) / float64(route.count)
}

func (this *UndergroundSystem) getRouteID(from, to string) string {
	return from + "_" + to
}
