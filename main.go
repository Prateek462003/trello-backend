package main

import (
	"database/sql"

	"github.com/Prateek462003/trello-backend/controllers"
	"github.com/Prateek462003/trello-backend/database"
	"github.com/gin-gonic/gin"

	// "github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Activity struct {
	ID                  int    `json:"id"`
	TaskID              int    `json:"task_id"`
	ActivityName        string `json:"activity_name"`
	ActivityDescription string `json:"activity_description"`
}

func getTasks(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	rows, err := db.Query("SELECT id, title, description FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO tasks (title, description) VALUES ($1, $2)", task.Title, task.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db, err = sql.Open("postgres", os.Getenv("URI"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	database.Init()
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	port := os.Getenv("PORT")
	if port == " " {
		port = "8080"
	}
	router.Run("localhost:" + port)
}
