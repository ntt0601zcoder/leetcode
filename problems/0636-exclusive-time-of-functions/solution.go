// Package exclusivetimeoffunctions solves LeetCode 636. Exclusive Time Of Functions.
// https://leetcode.com/problems/exclusive-time-of-functions/
package exclusivetimeoffunctions

import (
	"strconv"
	"strings"
)

type ExecInfo struct {
	taskID    int
	action    string
	timestamp int
	idle      int
}

func exclusiveTime(n int, logs []string) []int {
	result := make([]int, n)
	stack := make([]*ExecInfo, 0, len(logs))

	for _, log := range logs {
		curExec := parseExecInfo(log)

		if curExec.action == "start" {
			stack = append(stack, curExec)
			continue
		}

		preExec := stack[len(stack)-1]
		totalTime := curExec.timestamp - preExec.timestamp + 1
		result[curExec.taskID] += totalTime - preExec.idle
		stack = stack[:len(stack)-1]

		if len(stack) != 0 {
			stack[len(stack)-1].idle += totalTime
		}

	}

	return result
}

func parseExecInfo(log string) *ExecInfo {
	splittedLog := strings.Split(log, ":")

	info := &ExecInfo{}
	info.taskID, _ = strconv.Atoi(splittedLog[0])
	info.action = splittedLog[1]
	info.timestamp, _ = strconv.Atoi(splittedLog[2])

	return info
}
