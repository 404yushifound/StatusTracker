package service

import (
	"status_tracker/model"
	"status_tracker/provider"
	"status_tracker/util"
	"time"
	"strconv"
)

type TaskService struct {
	tasks        []model.Task
	emailService *provider.MockEmailProvider
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:        []model.Task{},
		emailService: provider.NewMockEmailProvider(),
	}
}

func (s *TaskService) AddTask(email string) model.Task {
	task := model.Task{
		ID:     strconv.FormatInt(time.Now().UnixNano(), 10),
		Email:  email,
		Status: model.PENDING,
	}
	s.tasks = append(s.tasks, task)
	util.Info("Task added for: " + email)
	return task
}

func (s *TaskService) ProcessTasks() {
	queue := s.tasks

	for i, task := range queue {
		success := s.emailService.Send(task.Email)

		if success {
			s.tasks[i].Status = model.SUCCESS
			util.Success("Email sent to: " + task.Email)
			continue
		}

		// Retry once if it failed
		util.Info("Retrying email to: " + task.Email)
		time.Sleep(2 * time.Second)

		retrySuccess := s.emailService.Send(task.Email)
		if retrySuccess {
			s.tasks[i].Status = model.SUCCESS
			util.Success("Retry succeeded for: " + task.Email)
		} else {
			s.tasks[i].Status = model.FAILED
			util.Error("Retry failed for: " + task.Email)
		}
	}
}

func (s *TaskService) GetTasks() []model.Task {
	return s.tasks
}


