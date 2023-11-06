package repository

import "server/model"

type ITodoRepo interface {
	Save(*model.Task) error
}

type todoRepoImpl struct{}

func NewTodoRepo() ITodoRepo {
	return &todoRepoImpl{}
}

func (*todoRepoImpl) Save(task *model.Task) error {
	return nil
}
