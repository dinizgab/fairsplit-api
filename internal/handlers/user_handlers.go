package handlers

import (
	"net/http"

	"github.com/dinizgab/split-api/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// TODO - Check if this is necessary
// maybe is better to just use SSO or OAuth
func RegisterUser(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := db.QueryRow(
			c,
			"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
			user.Username, user.Email, user.Password,
		).Scan(&user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "group created successfully",
			"id":      user.ID,
		})
	}
}
