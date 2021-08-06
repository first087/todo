package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var index int
var tasks map[int]*Task = make(map[int]*Task)

type Task struct {
	gorm.Model
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
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	db.AutoMigrate(&Task{})

	var task NewTaskTodo
	if err := c.Bind(&task); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// New(task.Task)
	db.Create(task)
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
