package main

import (
	"fmt"
	"github.com/Hailemari/task_manager/router"
)

func main() {
	fmt.Println("Task Manager API")
	r := router.SetupRouter()
	r.Run(":8000") 
}
