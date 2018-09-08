package models

// Task type to hold the information about an
type Task struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"complete"`
	Created    int64
	Updated    int64
}

// Tasks type, which is a slice for holding multiple Task objects.
type Tasks []*Task
