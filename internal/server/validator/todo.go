package validator

import "github.com/pepeunlimited/billing/pkg/subscriptionrpc"

type TodoServerValidator struct {}


func NewTodoServerValidator() TodoServerValidator {
	return TodoServerValidator{}
}

func (TodoServerValidator) CreateTodo(params *subscriptionrpc.CreateTodoParams) error {
	return nil
}

func (TodoServerValidator) GetTodo(params *subscriptionrpc.GetTodoParams) error {
	return nil
}

func (TodoServerValidator) UpdateTodo(params *subscriptionrpc.UpdateTodoParams) error {
	return nil
}

func (TodoServerValidator) DeleteTodo(params *subscriptionrpc.DeleteTodoParams) error {
	return nil
}