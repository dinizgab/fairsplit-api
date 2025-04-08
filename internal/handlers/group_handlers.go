package handlers

import (
	"database/sql"
	"net/http"

	_ "embed"

	"github.com/dinizgab/split-api/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

//go:embed sql/get_group_by_id.sql
var getGroupByIdQuery string

// TODO - Add the created group to the user who created it
func CreateNewGroup(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		var group entity.Group
		if err := c.BindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.QueryRow(
			c,
		   "INSERT INTO groups (name, value, due_day) VALUES ($1, $2, $3) RETURNING id",
			group.Name, group.Value, group.DueDay,
		).Scan(&group.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "group created successfully",
			"id":      group.ID,
		})
	}
}

func GetGroupByID(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var group entity.Group
		group.Users = []entity.User{}
		rows, err := db.Query(
			c,
			getGroupByIdQuery,
		 	id,
		)
	 	if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var (
				userId sql.NullString
				username sql.NullString
			)
			err := rows.Scan(
				&group.ID,
				&group.Name,
				&group.Value,
				&group.DueDay,
				&userId,
				&username,
			)
			if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }

			if userId.Valid && username.Valid {
				user := entity.User{
					ID:       userId.String,
					Username: username.String,
				}
				group.Users = append(group.Users, user)
			}
		}

		if rows.Err() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": rows.Err()})
            return
		}

		c.JSON(http.StatusOK, group)
	}
}

func AddUserToGroup(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var user entity.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if _, err := db.Exec(
			c,
			"INSERT INTO user_groups (group_id, user_id) VALUES ($1, $2)",
			id, user.ID,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "user added to group successfully",
		})
	}
}
