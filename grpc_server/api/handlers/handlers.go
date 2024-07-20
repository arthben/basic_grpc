package handlers

import (
	"context"
	"fmt"

	"github.com/arthben/basic_grpc/grpc_server/internal/storage"
	"github.com/arthben/basic_grpc/grpc_server/pkg/contract"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type APIHandlers struct {
	contract.UnimplementedThingsTodoServer
	database *storage.MockDB
}

func NewAPIHandlers() *APIHandlers {
	// this handlers will implement all method on the ThingsTodoServer interface
	// see file contract/todo_grpc.pb.go
	return &APIHandlers{
		database: storage.NewMockDB(),
	}
}

// implement AddTask function from ThingsTodoServer interface
func (a *APIHandlers) AddTask(
	ctx context.Context,
	p *contract.AddTodoReq,
) (*contract.TodoItem, error) {
	item, err := a.database.AddNew(p)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed Save New Task - %v", err.Error()))
	}
	return item, nil
}

// implement Tasklist function from ThingsTodoServer interface
func (a *APIHandlers) TaskList(ctx context.Context, _ *emptypb.Empty) (*contract.TodoItems, error) {
	items, err := a.database.List()
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, status.Error(codes.NotFound, "No Data Found")
	}
	return items, nil
}
