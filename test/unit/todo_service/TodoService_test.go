package todoService_test

import (
	"server/mocks"
	"server/model"
	todoSvc "server/todo_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoneTaskRepo(t *testing.T) {
	// Arrange
	var (
		task               = &model.Task{Id: 1, Title: "sweep", Status: "doing"}
		mockTodoRepo       = mocks.ITodoRepo{}
		tds                = todoSvc.NewTodoService(&mockTodoRepo)
		expectedErr  error = nil
	)
	//Giả lập hàm Save() của TodoRepo là trả về nil
	mockTodoRepo.On("Save", task).Return(nil)

	// Act
	actualErr := tds.DoneTask(task)

	// Assert
	assert.Equal(t, actualErr, expectedErr)
	// Check hàm Save được gọi
	mockTodoRepo.AssertNumberOfCalls(t, "Save", 0)
}
