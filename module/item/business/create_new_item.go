package todobiz

import (
	"context"
	"errors"
	todomodel "go-with-clean-architecture/module/item/model"
)

type CreateTodoItemStorage interface {
	CreateItem(ctx context.Context, data *todomodel.ToDoItem) error
}

type createBiz struct {
	store CreateTodoItemStorage
}

func NewCreateToDoItemBiz(store CreateTodoItemStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewItem(ctx context.Context, data *todomodel.ToDoItem) error {
	if data.Title == "" {
		return errors.New("title can not be blank")
	}

	// do not allow "finished" status when creating a new task
	data.Status = "Doing" // set to default

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
