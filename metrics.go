package aerospike

import "time"

var flow func(string, string, time.Time) = nil

func InitializeMetrics(flowFunc func(string, string, time.Time)) {
	flow = flowFunc
}

func FlowMetrics(isRead bool, step string, startedAt time.Time) {
	if flow != nil {
		if isRead {
			flow("aerospike-read", step, startedAt)
		} else {
			flow("aerospike-write", step, startedAt)
		}
	}
}
