package task

import (
	"fiber-root/app/api"
	"fiber-root/app/model"
	"fiber-root/app/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

// FindAll 查询
// @Tags task
// @Summary 查询
// @Description 查询
// @Security jwt
// @Accept json
// @Produce json
// @Success 200 {object} api.ResponseHTTP{data=[]model.Task}
// @Router /api/task/list [get]
func FindAll(c *fiber.Ctx) error {
	todoLists, err := service.TaskService.TodoLists()
	if err != nil {
		return api.Response(c, err, nil)
	}
	return api.Response(c, err, todoLists)
}

// Save 保存
// @Summary 保存
// @Description 保存
// @Tags task
// @Security jwt
// @Accept json
// @Produce json
// @Param task body  model.Task true "save task"
// @Success 200 {object} api.ResponseHTTP{data=string}
// @Router /api/task/save [post]
func Save(c *fiber.Ctx) error {
	m := new(model.Task)
	if err := c.BodyParser(m); err != nil {
		return api.Response(c, err, nil)
	}
	id, err := service.TaskService.Save(m)
	if err != nil {
		log.Println("Error while save task:", err.Error())
	}
	return api.Response(c, err, id)
}

// ChangeStatus 修改
// @Summary 修改
// @Description 修改
// @Tags task
// @Security jwt
// @Accept json
// @Produce json
// @Param id path string true "task ID"
// @Param task body  model.Task true "task"
// @Success 200 {object} api.ResponseHTTP{data=string}
// @Router /api/task/{id} [put]
func ChangeStatus(c *fiber.Ctx) error {
	m := new(model.Task)
	taskId := c.Params("id")

	if err := c.BodyParser(m); err != nil {
		return api.Response(c, err, nil)
	}
	id, err := service.TaskService.ChangeStatus(taskId, m)
	if err != nil {
		return api.Response(c, err, nil)
	}
	return api.Response(c, err, id)
}

// Remove 删除
// @Summary 删除
// @Description 删除
// @Tags task
// @Security jwt
// @Accept json
// @Produce json
// @Param id path string true "task ID"
// @Success 200 {object} api.ResponseHTTP{data=string}
// @Router /api/task/{id} [delete]
func Remove(c *fiber.Ctx) error {
	taskId := c.Params("id")
	id, err := service.TaskService.Remove(taskId)
	if err != nil {
		return api.Response(c, err, nil)
	}
	return api.Response(c, err, id)
}
