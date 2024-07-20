package storage

import (
	"time"

	"github.com/arthben/basic_grpc/grpc_server/pkg/contract"
	"github.com/google/uuid"
)

type MockDB struct {
	data contract.TodoItems
}

func NewMockDB() *MockDB {
	return &MockDB{}
}

func (m *MockDB) List() (*contract.TodoItems, error) {
	return &m.data, nil
}

func (m *MockDB) AddNew(p *contract.AddTodoReq) (*contract.TodoItem, error) {
	item := &contract.TodoItem{
		Id:          uuid.NewString(),
		Todo:        p.GetTodo(),
		Details:     p.GetDetails(),
		CreatedDate: time.Now().Local().String(),
	}
	m.data.Todolist = append(m.data.Todolist, item)
	return item, nil
}
