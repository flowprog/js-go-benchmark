package main

import (
	"io"
	"log"
	"os"

	TaskController "stress/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server..")

	//db.Init()

	// 设置 Gin 为 Release 模式
	gin.SetMode(gin.ReleaseMode)
	// 打开 /dev/null 文件
	f, _ := os.Open(os.DevNull)

	// 将 Gin 的日志输出重定向到 /dev/null
	gin.DefaultWriter = io.Writer(f)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("/", TaskController.GetTasks)
			tasks.POST("/", TaskController.CreateTask)
			tasks.PUT("/:id", TaskController.UpdateTask)
			tasks.DELETE("/:id", TaskController.DeleteTask)
		}
	}

	r.Run()
}
