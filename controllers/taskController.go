package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/blen/task_manager_api/db"
	"github.com/blen/task_manager_api/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GET /tasks -- Get all tasks
func GetTasks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error fetching tasks"})
		return
	}
	defer cursor.Close(ctx)

	var tasks []model.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error decoding tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

// GET /tasks/:id -- Get task by ID
func GetTaskById(c *gin.Context) {
	idParam := c.Param("id")
	taskId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task model.Task
	err = db.TaskCollection.FindOne(ctx, bson.M{"_id": taskId}).Decode(&task)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

// POST /tasks -- Create a new task
func CreateTask(c *gin.Context) {
	var newTask model.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}

	newTask.ID = primitive.NewObjectID() // generate new ObjectID

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.TaskCollection.InsertOne(ctx, newTask)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating task"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newTask)
}

// PUT /tasks/:id -- Update an existing task
func UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	taskId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task ID"})
		return
	}

	var updateTask model.Task
	if err := c.BindJSON(&updateTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}

	updateFields := bson.M{
		"title":       updateTask.Title,
		"description": updateTask.Description,
		"due_date":    updateTask.DueDate,
		"status":      updateTask.Status,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := db.TaskCollection.UpdateOne(ctx, bson.M{"_id": taskId}, bson.M{"$set": updateFields})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error updating task"})
		return
	}

	if result.MatchedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	// Return the updated task
	updateTask.ID = taskId
	c.IndentedJSON(http.StatusOK, updateTask)
}

// DELETE /tasks/:id -- Delete a task
func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	taskId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := db.TaskCollection.DeleteOne(ctx, bson.M{"_id": taskId})
	if err != nil || result.DeletedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
