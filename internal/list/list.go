// Package list defines the Lists interface and TodoList and Task list structures
// required for implementing a TodoList API.
package list

// Lists provides an interface for specifying the methods necessary for managing a set
// of to-do lists.
type Lists interface {
	GetLists(filter string, skip, max int) ([]TodoList, error)
	GetList(id string) (TodoList, error)
	AddList(list TodoList) error
	AddTask(listID string, task Task) error
	MarkTaskComplete(listID, taskID string) error
}

// TodoList defines the properties required for a list of tasks.
type TodoList struct {
	ID          string
	Name        string
	Description string
	Tasks       []Task
}

// Task defines the properties required for a TodoList task object.
type Task struct {
	ID        string
	Name      string
	Completed bool
}
