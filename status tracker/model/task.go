package model

type TaskStatus string

const (
	PENDING TaskStatus = "PENDING"
	SUCCESS TaskStatus = "SUCCESS"
	FAILED  TaskStatus = "FAILED"
)

type Task struct {
	ID     string
	Email  string
	Status TaskStatus
}
