package goworker

import "time"

type job struct {
	Queue   string
	Payload Payload
	RunAt   time.Time
}
