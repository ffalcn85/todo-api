package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ffalcn85/todo-api/internal/list"
)

type httpTransport struct {
	Lists list.Lists
}

func (h *httpTransport) getListsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter := vars["searchString"]

	max, err := strconv.Atoi(vars["limit"])
	if err != nil {
		max = 50
	}

	skip, err := strconv.Atoi(vars["skip"])
	if err != nil {
		skip = 0
	}

	todoLists, err := h.Lists.GetLists(filter, skip, max)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(todoLists)
}

func (h *httpTransport) getListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	todoList, err := h.Lists.GetList(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(todoList)
}

func (h *httpTransport) addListHandler(w http.ResponseWriter, r *http.Request) {
	todoList := list.TodoList{}
	err := json.NewDecoder(r.Body).Decode(&todoList)
	if err != nil && err.Error() != "EOF" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("invalid input, object invalid"))
		return
	}

	err = h.Lists.AddList(todoList)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("an existing item already exists"))
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *httpTransport) addTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var task list.Task
	listID := vars["id"]
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil && err.Error() != "EOF" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("invalid input, object invalid"))
		return
	}

	err = h.Lists.AddTask(listID, task)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("an existing item already exists"))
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *httpTransport) markTaskCompleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["id"]
	taskID := vars["taskId"]

	err := h.Lists.MarkTaskComplete(listID, taskID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("an existing item already exists"))
	}

	w.WriteHeader(http.StatusCreated)
}
