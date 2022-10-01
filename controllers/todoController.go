package todoController

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"restAPI/models"
)

func FindAll(c *gin.Context) {
	var todos []models.Todo

	models.DB.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
}
func FindById(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	if err := models.DB.First(&todo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}
func Create(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "data added succesfully"})
}
func Update(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&todo).Where("id = ?", id).Updates(&todo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "update fail!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data updated successfully"})
}
func Delete(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if models.DB.Delete(&todo, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "delete fail!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data deleted successfully"})
}
