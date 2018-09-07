package models

// Task type to hold the information about an
type Task struct {
	ID         string
	Title      string
	IsComplete bool
	Created    int64
	Updated    int64
}

// Tasks type, which is a slice for holding multiple Task objects.
type Tasks []*Task
