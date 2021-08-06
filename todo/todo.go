package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func AddTask(c *gin.Context) {
	var task NewTaskTodo
	if err := c.Bind(&task); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	New(task.Task)
}

func MaskDone(c *gin.Context) {
	index := c.Param("index")
	i, err := strconv.Atoi(index)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	tasks[i].Done = true
}

func GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, List())
}
