package serviceimpl

import (
	"errors"
	"testing"

	"github.com/phucnh/go-app-sample/errs"
	repoMocks "github.com/phucnh/go-app-sample/repository/mocks"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/phucnh/go-app-sample/testdata"
)

var ErrUnexpectedTask = errors.New("Unexpected error")

type TaskServiceTestSuite struct {
	suite.Suite
	userRepository *repoMocks.UserRepository
	taskRepository *repoMocks.TaskRepository
}

func (s *TaskServiceTestSuite) SetupTest() {
	userRepositoryMock := new(repoMocks.UserRepository)
	taskRepositoryMock := new(repoMocks.TaskRepository)
	s.userRepository = userRepositoryMock
	s.taskRepository = taskRepositoryMock
}

func (s *TaskServiceTestSuite) TestCreateTask_ShouldReturnSuccess() {
	task := testdata.Task1
	user := testdata.User1
	s.userRepository.
		On("GetUserByID", task.Assignee.String).
		Return(user, nil)

	s.taskRepository.
		On("CreateTask", task).
		Return(task, nil)

	taskService := NewTaskService(s.userRepository, s.taskRepository)
	createdTask, err := taskService.CreateTask(task)
	s.NoError(err)
	s.Equal(task, createdTask)
}

func (s *TaskServiceTestSuite) TestCreateTask_ShouldReturnServerErrorWhenGetUser() {
	task := testdata.Task1
	s.userRepository.
		On("GetUserByID", task.Assignee.String).
		Return(nil, ErrUnexpectedTask)

	taskService := NewTaskService(s.userRepository, s.taskRepository)
	createdTask, err := taskService.CreateTask(task)
	s.Error(err)
	s.Nil(createdTask)
}

func (s *TaskServiceTestSuite) TestCreateTask_ShouldReturnErrUserNotFound() {
	task := testdata.Task1
	s.userRepository.
		On("GetUserByID", task.Assignee.String).
		Return(nil, nil)

	taskService := NewTaskService(s.userRepository, s.taskRepository)
	createdTask, err := taskService.CreateTask(task)
	s.Equal(errs.ErrUserNotFound, err)
	s.Nil(createdTask)
}

func (s *TaskServiceTestSuite) TestCreateTask_ShouldReturnServerErrorWhenCreateTask() {
	task := testdata.Task1
	user := testdata.User1
	s.userRepository.
		On("GetUserByID", task.Assignee.String).
		Return(user, nil)

	s.taskRepository.
		On("CreateTask", task).
		Return(nil, gorm.ErrPrimaryKeyRequired)

	taskService := NewTaskService(s.userRepository, s.taskRepository)
	createdTask, err := taskService.CreateTask(task)
	s.Error(err)
	s.Nil(createdTask)
}

func TestTaskService(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TaskServiceTestSuite))
}
