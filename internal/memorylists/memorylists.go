package memorylists

import (
	"errors"
	"strings"

	"github.com/ffalcn85/todo-api/internal/list"
	"github.com/google/uuid"
)

type todoLists struct {
	lists   []*todoList
	listMap map[string]*todoList
}

type todoList struct {
	list.TodoList
	taskMap map[string]int
}

// Create acts as the constructor for the memory list's un-exported objects.
func Create() list.Lists {
	t := &todoLists{
		lists:   []*todoList{},
		listMap: map[string]*todoList{},
	}
	return t
}

func (t *todoLists) GetLists(filter string, skip, max int) ([]list.TodoList, error) {
	result := []list.TodoList{}
	for i, list := range t.lists[skip:] {
		if i >= max {
			break
		}
		if strings.Contains(list.Name, filter) ||
			strings.Contains(list.Description, filter) {
			result = append(result, list.TodoList)
		}
	}
	return result, nil
}

func (t *todoLists) GetList(id string) (list.TodoList, error) {
	result := t.listMap[id]
	if nil == result {
		return list.TodoList{}, errors.New("Invalid id supplied")
	}
	return result.TodoList, nil
}

func (t *todoLists) AddList(list list.TodoList) error {
	listToAdd := &todoList{
		TodoList: list,
		taskMap:  map[string]int{},
	}

	if listToAdd.ID == "" {
		listToAdd.ID = uuid.New().String()
	}

	if t.listMap[listToAdd.ID] != nil {
		return errors.New("an existing item already exists")
	}

	t.listMap[listToAdd.ID] = listToAdd
	t.lists = append(t.lists, listToAdd)
	return nil
}

func (t *todoLists) AddTask(listID string, task list.Task) error {
	todo := t.listMap[listID]

	if nil == todo {
		return errors.New("Invalid list ID")
	}

	if task.ID == "" {
		task.ID = uuid.New().String()
	}

	if 0 != todo.taskMap[task.ID] {
		return errors.New("Task already exists")
	}

	todo.Tasks = append(todo.Tasks, task)
	todo.taskMap[task.ID] = len(todo.Tasks)

	return nil
}

func (t *todoLists) MarkTaskComplete(listID, taskID string) error {
	todo := t.listMap[listID]

	if nil == todo {
		return errors.New("Invalid list ID")
	}

	if 0 == todo.taskMap[taskID] {
		return errors.New("Task not found")
	}

	todo.Tasks[todo.taskMap[taskID]-1].Completed = true
	return nil
}
