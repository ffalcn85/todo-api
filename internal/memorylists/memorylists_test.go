package memorylists

import (
	"reflect"
	"testing"

	"github.com/ffalcn85/todo-api/internal/list"
)

var (
	singleTaskList = &todoList{
		TodoList: list.TodoList{
			ID:          "testID1",
			Name:        "Single Task List",
			Description: "A list with a single task object.",
			Tasks: []list.Task{
				list.Task{
					ID:        "1",
					Name:      "First Task",
					Completed: false,
				},
			},
		},
		taskMap: map[string]int{"1": 1},
	}
	twoTaskList = &todoList{
		TodoList: list.TodoList{
			ID:          "testID2",
			Name:        "Two Task List",
			Description: "A list with a two task objects.",
			Tasks: []list.Task{
				list.Task{
					ID:        "1",
					Name:      "First Task",
					Completed: true,
				},
				list.Task{
					ID:        "2",
					Name:      "Second Task",
					Completed: false,
				},
			},
		},
		taskMap: map[string]int{"1": 1, "2": 2},
	}
	testLists = &todoLists{
		lists: []*todoList{
			singleTaskList,
			twoTaskList,
		},
		listMap: map[string]*todoList{
			singleTaskList.ID: singleTaskList,
			twoTaskList.ID:    twoTaskList,
		},
	}
)

func Test_todoLists_GetLists(t *testing.T) {
	type fields struct {
		lists   []*todoList
		listMap map[string]*todoList
	}
	type args struct {
		filter string
		skip   int
		max    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []list.TodoList
		wantErr bool
	}{
		{
			name: "empty lists object",
			fields: fields{
				lists:   []*todoList{},
				listMap: map[string]*todoList{},
			},
			args:    args{filter: "", skip: 0, max: 50},
			want:    []list.TodoList{},
			wantErr: false,
		},
		{
			name: "single Task List Test",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{filter: "single", skip: 0, max: 50},
			want:    []list.TodoList{singleTaskList.TodoList},
			wantErr: false,
		},
		{
			name: "max 0 test",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{filter: "", skip: 0, max: 1},
			want:    []list.TodoList{singleTaskList.TodoList},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &todoLists{
				lists:   tt.fields.lists,
				listMap: tt.fields.listMap,
			}
			got, err := tl.GetLists(tt.args.filter, tt.args.skip, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoLists.GetLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoLists.GetLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoLists_GetList(t *testing.T) {
	type fields struct {
		lists   []*todoList
		listMap map[string]*todoList
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    list.TodoList
		wantErr bool
	}{
		{
			name: "invalid ID test",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{id: "invalid ID"},
			want:    list.TodoList{},
			wantErr: true,
		},
		{
			name: "valid ID test",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{id: singleTaskList.ID},
			want:    singleTaskList.TodoList,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &todoLists{
				lists:   tt.fields.lists,
				listMap: tt.fields.listMap,
			}
			got, err := tl.GetList(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoLists.GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoLists.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoLists_AddList(t *testing.T) {
	type fields struct {
		lists   []*todoList
		listMap map[string]*todoList
	}
	type args struct {
		list list.TodoList
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Adding Empty Task List",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "Adding Second Empty Task List",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "Adding List with duplicate ID",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{list: list.TodoList{ID: singleTaskList.ID}},
			wantErr: true,
		},
		{
			name: "Adding List with objects",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{list: list.TodoList{ID: "generic Test ID", Name: "testName", Description: "new", Tasks: singleTaskList.TodoList.Tasks}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &todoLists{
				lists:   tt.fields.lists,
				listMap: tt.fields.listMap,
			}
			if err := tl.AddList(tt.args.list); (err != nil) != tt.wantErr {
				t.Errorf("todoLists.AddList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_todoLists_AddTask(t *testing.T) {
	type fields struct {
		lists   []*todoList
		listMap map[string]*todoList
	}
	type args struct {
		listID string
		task   list.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Adding Task to List",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{listID: singleTaskList.ID, task: list.Task{}},
			wantErr: false,
		},
		{
			name: "Adding Second Empty Task to List",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{listID: singleTaskList.ID, task: list.Task{}},
			wantErr: false,
		},
		{
			name: "Attempting to add task with duplicate ID",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{listID: singleTaskList.ID, task: list.Task{ID: "1"}},
			wantErr: true,
		},
		{
			name: "Attempting to add task to invalid list",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{listID: "non-existent list", task: list.Task{}},
			wantErr: true,
		},
		{
			name: "Adding pre filled task to list",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{listID: singleTaskList.ID, task: list.Task{ID: "894752983475", Name: "testTask", Completed: true}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &todoLists{
				lists:   tt.fields.lists,
				listMap: tt.fields.listMap,
			}
			if err := tl.AddTask(tt.args.listID, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("todoLists.AddTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_todoLists_MarkTaskComplete(t *testing.T) {
	type fields struct {
		lists   []*todoList
		listMap map[string]*todoList
	}
	type args struct {
		listID string
		taskID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Marking Task as Completed",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{listID: singleTaskList.ID, taskID: "1"},
			wantErr: false,
		},
		{
			name: "Attempting to Mark non-existent task as Completed",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{listID: singleTaskList.ID, taskID: "non-existent Task"},
			wantErr: true,
		},
		{
			name: "Attempting to mark task completed in invalid list",
			fields: fields{
				lists:   testLists.lists,
				listMap: testLists.listMap,
			},
			args:    args{listID: "non-existent list", taskID: "1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &todoLists{
				lists:   tt.fields.lists,
				listMap: tt.fields.listMap,
			}
			if err := tl.MarkTaskComplete(tt.args.listID, tt.args.taskID); (err != nil) != tt.wantErr {
				t.Errorf("todoLists.MarkTaskComplete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name string
		want list.Lists
	}{
		{name: "Construct Test", want: &todoLists{lists: []*todoList{}, listMap: map[string]*todoList{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
