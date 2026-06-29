// Package timebasedkeyvaluestore solves LeetCode 981. Time Based Key Value Store.
// https://leetcode.com/problems/time-based-key-value-store/
package timebasedkeyvaluestore

type entry struct {
	timestamp int
	value     string
}

type TimeMap struct {
	store map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]entry)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], entry{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	entries, ok := this.store[key]

	if !ok {
		return ""
	}

	earliest, latest, result := 0, len(entries)-1, ""

	for earliest <= latest {
		mid := earliest + (latest-earliest)/2
		entry := entries[mid]

		if entry.timestamp <= timestamp {
			result = entry.value
			earliest = mid + 1
		} else {
			latest = mid - 1
		}
	}

	return result
}
