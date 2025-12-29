package main

import (
	"spec-builder/backend/internal/db"
	"spec-builder/backend/internal/excel"
	"spec-builder/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	// Импорт данных при первом запуске (если таблица пуста)
	var count int64
	db.DB.Model(&db.DB).Table("cabinets").Count(&count)
	if count == 0 {
		excel.ImportFromXLSX("../init-data/Приложение №1 к приказу 03.23.xlsx")
	}

	r := gin.Default()
	r.Use(cors())

	api := r.Group("/api")
	{
		api.GET("/cabinets", handlers.GetCabinets)
		api.POST("/spec", handlers.GenerateSpec)
		api.POST("/export/xlsx", handlers.ExportSpecExcel)
	}

	r.Run(":8080")
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
