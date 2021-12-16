package service

import (
	"fiber-root/app/model"
	"fiber-root/db"
	"time"
)

var TaskService = &taskService{}

type taskService struct{}

//查询所有
func (t *taskService) TodoLists() ([]model.Task, error) {
	var tasks []model.Task
	result := db.GetDB().Find(&tasks)
	if result.Error != nil {
		return tasks, result.Error
	}
	return tasks, nil
}

//新增
func (t *taskService) Save(task *model.Task) (uint, error) {
	task.CreateTime = time.Now()
	result := db.GetDB().Create(task)
	if result.Error != nil {
		return 0, result.Error
	}
	return task.ID, nil
}

//修改
func (t *taskService) ChangeStatus(id string, task *model.Task) (interface{}, error) {
	result := db.GetDB().Model(task).Where("id", id).Update("status", task.Status)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}

//删除
func (t *taskService) Remove(id string) (interface{}, error) {
	result := db.GetDB().Delete(&model.Task{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}
