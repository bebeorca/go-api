package main

import (
	"example/go_api/controllers/pesanancontroller"
	"example/go_api/controllers/usercontroller"
	"example/go_api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	//user_controller
	r.GET("/api/daftaruser", usercontroller.Index)
	r.POST("/api/register", usercontroller.Create)

	//pesanan_controller
	r.GET("/api/daftarpesanan", pesanancontroller.Index)
	r.GET("/api/pesanan/:id", pesanancontroller.Show)
	r.POST("/api/pesanan", pesanancontroller.Create)
	r.PUT("/api/pesanan/:id", pesanancontroller.Update)
	r.DELETE("/api/pesanan/", pesanancontroller.Delete)

	r.Run()
}
