package todoService

import (
	"server/model"
	"server/repository"
)

type todoService struct {
	todoRepo repository.ITodoRepo
}

func NewTodoService(tdr repository.ITodoRepo) *todoService {
	return &todoService{
		todoRepo: tdr,
	}
}

// Cập nhật status thành done
func (tds *todoService) DoneTask(task *model.Task) error {
	task.Status = "done"

	err := tds.todoRepo.Save(task)

	return err
}
