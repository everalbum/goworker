package goworker

import "time"

type Job struct {
	Queue   string
	Payload Payload
	RunAt   time.Time
}
