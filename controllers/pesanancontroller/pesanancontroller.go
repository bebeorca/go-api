package pesanancontroller

import (
	"encoding/json"
	"example/go_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var daftarpesanan []models.Pesanan

	models.DB.Find(&daftarpesanan)
	c.JSON(http.StatusOK, gin.H{"data": daftarpesanan})
}

func Show(c *gin.Context) {
	var pesanan models.Pesanan
	id := c.Param("id")

	if err := models.DB.First(&pesanan, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Pesanan": pesanan})

}

func Create(c *gin.Context) {
	var pesanan models.Pesanan

	if err := c.ShouldBindJSON(&pesanan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&pesanan)
	c.JSON(http.StatusOK, gin.H{"Pesanan": pesanan})

}

func Update(c *gin.Context) {
	var pesanan models.Pesanan
	id := c.Param("id")

	if err := c.ShouldBindJSON(&pesanan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&pesanan).Where("id = ?", id).Updates(&pesanan).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak Dapat Memperbarui Pesanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Diperbarui."})

}

func Delete(c *gin.Context) {
	var pesanan models.Pesanan

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&pesanan, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak Dapat Menghapus Pesanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus."})

}
