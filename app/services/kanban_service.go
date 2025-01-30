package services

import (
	"app/db"
	"app/models"
	"log"
)

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := db.DB.Find(&tasks).Error; err != nil {
		log.Println("Error fetching tasks: ", err)
		return nil, err
	}
	return tasks, nil
}

func GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := db.DB.First(&task, id).Error; err != nil {
		log.Println("Error fetching task: ", err)
		return nil, err
	}
	return &task, nil
}

func CreateTask(title, status, priority string) (models.Task, error) {
	task := models.Task{Title: title, Status: status, Priority: priority}
	if err := db.DB.Create(&task).Error; err != nil {
		log.Println("Error creating task: ", err)
		return task, err
	}
	return task, nil
}

func UpdateTask(id uint, title, status, priority string) (*models.Task, error) {
	var task models.Task
	if err := db.DB.First(&task, id).Error; err != nil {
		log.Println("Error fetching task for update: ", err)
		return nil, err
	}
	task.Title = title
	task.Status = status
	task.Priority = priority
	if err := db.DB.Save(&task).Error; err != nil {
		log.Println("Error updating task: ", err)
		return nil, err
	}
	return &task, nil
}

func DeleteTask(id uint) error {
	var task models.Task
	if err := db.DB.First(&task, id).Error; err != nil {
		log.Println("Error fetching task for deletion: ", err)
		return err
	}
	if err := db.DB.Delete(&task).Error; err != nil {
		log.Println("Error deleting task: ", err)
		return err
	}
	return nil
}
