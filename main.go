package main

import (
	"github.com/Prateek462003/trello-backend/controllers"
	"github.com/Prateek462003/trello-backend/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
)

// var db *sql.DB

// type Task struct {
// 	ID          int            `json:"id"`
// 	Title       string         `json:"title"`
// 	Description string         `json:"description"`
// 	Image       sql.NullString `json:"image"`
// 	ActivityId  int            `json:"activity_id"`
// }

// func (t Task) MarshalJSON() ([]byte, error) {
// 	type Alias Task
// 	return json.Marshal(&struct {
// 		Image *string `json:"image,omitempty"`
// 		Alias
// 	}{
// 		Image: func() *string {
// 			if t.Image.Valid {
// 				return &t.Image.String
// 			}
// 			return nil
// 		}(),
// 		Alias: (Alias)(t),
// 	})
// }

// type Activity struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// // All the Task Controllers and Activity Controllers are inside the ./controllers folder due to hosting problem all reside in a single file
// func getTasks(c *gin.Context) {
// 	c.Header("Content-Type", "application/json")
// 	rows, err := db.Query("SELECT id, title, description, image, activity_id FROM tasks")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var tasks []Task
// 	for rows.Next() {
// 		var task Task
// 		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Image, &task.ActivityId); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		tasks = append(tasks, task)
// 	}
// 	c.JSON(http.StatusOK, tasks)
// }

// func (t *Task) UnmarshalJSON(data []byte) error {
// 	type Alias Task
// 	aux := &struct {
// 		Image *string `json:"image,omitempty"`
// 		*Alias
// 	}{
// 		Alias: (*Alias)(t),
// 	}

// 	if err := json.Unmarshal(data, &aux); err != nil {
// 		return err
// 	}

// 	if aux.Image != nil {
// 		t.Image = sql.NullString{String: *aux.Image, Valid: true}
// 	} else {
// 		t.Image = sql.NullString{String: "", Valid: false}
// 	}

// 	return nil
// }

// func createTask(c *gin.Context) {
// 	var task Task
// 	if err := c.BindJSON(&task); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	_, err := db.Exec("INSERT INTO tasks (title, description, image, activity_id) VALUES ($1, $2, $3, $4)",
// 		task.Title, task.Description, task.Image, task.ActivityId)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, task)
// }

// func deleteTask(c *gin.Context) {
// 	id := c.Param("id")
// 	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
// }

// func getActivities(c *gin.Context) {
// 	c.Header("Content-Type", "application/json")
// 	rows, err := db.Query("SELECT id, name FROM activities")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var activities []Activity
// 	for rows.Next() {
// 		var activity Activity
// 		if err := rows.Scan(&activity.ID, &activity.Name); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		activities = append(activities, activity)
// 	}
// 	c.JSON(http.StatusOK, activities)
// }

// func createActivity(c *gin.Context) {
// 	var activity Activity
// 	if err := c.BindJSON(&activity); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	_, err := db.Exec("INSERT INTO activities (name) VALUES ($1)", activity.Name)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, activity)
// }

// func GetTasksByActivity(c *gin.Context) {
// 	activityID := c.Param("activity_id")

// 	rows, err := db.Query("SELECT id, title, description, image, activity_id FROM tasks WHERE activity_id = $1", activityID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer rows.Close()

// 	var tasks []Task
// 	for rows.Next() {
// 		var task Task
// 		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Image, &task.ActivityId); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		tasks = append(tasks, task)
// 	}

// 	if err = rows.Err(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, tasks)
// }

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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	router.GET("/activities", controllers.GetActivities)
	router.POST("/activities", controllers.CreateActivity)
	router.GET("/activities/:activity_id", controllers.GetTasksByActivity)

	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run("localhost:" + port)
}
