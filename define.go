package tasks

import "runtime"

var (
	MaxWorker = runtime.NumCPU()
	MaxQueue  = 512
)

type Job interface {
	Exec() error
}

var JobQueue chan Job
