package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var index int
var tasks map[int]*Task = make(map[int]*Task)

type Task struct {
	Title string
	Done  bool
}

func List() map[int]*Task {
	return tasks
}

type NewTaskTodo struct {
	Task string `json:"task"`
}

func New(task string) {
	defer func() {
		index++
	}()

	tasks[index] = &Task{
		Title: task,
		Done:  false,
	}
}

func AddTask(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var task NewTaskTodo
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	New(task.Task)
}

func MarkDone(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	index, err := strconv.Atoi(vars["index"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks[index].Done = true
}

func GetTodo(rw http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(rw).Encode(tasks); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
