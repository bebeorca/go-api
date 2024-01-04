package usercontroller

import (
	"example/go_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var daftaruser []models.User

	if err := models.DB.Preload("Pesanans").Find(&daftaruser).Error; err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": daftaruser})
}

func Create(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})

}
