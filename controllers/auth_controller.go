package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/blen/task_manager_api/db"
	"github.com/blen/task_manager_api/model"
	"github.com/blen/task_manager_api/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST /register -- Register a new user
func RegisterUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	// Give the first use the Admin role

	count, err := db.UserCollection.CountDocuments(context.Background(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error"})
		return
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	// Insert the user into the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.UserCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	userID := result.InsertedID.(primitive.ObjectID)

	token, err := utils.GenerateJWT(userID.Hex(), user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"token":   token,
	})
}

// POST /login -- Login a user
func LoginUser(c *gin.Context) {
	var input model.User
	var user model.User

	c.ShouldBindJSON(&input)

	// Find the user by email
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// err := db.UserCollection.FindOne(ctx, model.User{Email: input.Email}).Decode(&user)
	err := db.UserCollection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	// Check the password
	if !utils.CheckPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password is not correct"})
		return
	}

	token, err := utils.GenerateJWT(user.ID.Hex(), user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}

// Promote User to Admin
func PromoteUser(c *gin.Context) {
	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	_, err = db.UserCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"role": "admin"}}, // set role to "admin"
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Promotion failed"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})

}
