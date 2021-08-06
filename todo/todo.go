package todo

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vmihailenco/msgpack/v5"
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

type Serializer interface {
	Decode(io.Reader, interface{}) error
	Encode(io.Writer, interface{}) error
}

type JSONSerializer struct{}

func (j JSONSerializer) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func (j JSONSerializer) Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func NewJSONSerializer() JSONSerializer {
	return JSONSerializer{}
}

type MessagePackSerializer struct{}

func (MessagePackSerializer) Decode(r io.Reader, v interface{}) error {
	return msgpack.NewDecoder(r).Decode(v)
}

func (MessagePackSerializer) Encode(w io.Writer, v interface{}) error {
	return msgpack.NewEncoder(w).Encode(v)
}

func NewMessagePackSerializer() MessagePackSerializer {
	return MessagePackSerializer{}
}

type App struct {
	serialize Serializer
}

func NewApp(serialize Serializer) *App {
	return &App{
		serialize: serialize,
	}
}

func (app *App) AddTask(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var task NewTaskTodo
	if err := app.serialize.Decode(r.Body, &task); err != nil {
		// if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
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
	if err := json.NewEncoder(rw).Encode(List()); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
