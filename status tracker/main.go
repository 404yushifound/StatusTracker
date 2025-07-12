package main
import (
	"fmt"
	"status_tracker/service"
)
func main() {
	taskService := service.NewTaskService()

	taskService.AddTask("ayushi@example.com")
	taskService.AddTask("fail@example.com")
	taskService.AddTask("fail@example.com")
	taskService.AddTask("fail@example.com")
	taskService.AddTask("client@example.com") // will be blocked, then allowed later

	taskService.ProcessTasks()

	fmt.Println("\nFinal Task Status:")
	for _, task := range taskService.GetTasks() {
		fmt.Printf("- %s: %s\n", task.Email, task.Status)
	}
}

