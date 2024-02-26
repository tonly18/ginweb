package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"server/config"
	"server/global"
	"server/router"
	"server/tasks"
	"server/utils"
)

func main() {
	//Running Environment
	fmt.Printf("[Server] Running Environment: %v\n", global.SERVER_RUN_ENV)
	fmt.Printf("[Server] Running Gin Mode: %v\n", global.GIN_MODE)

	//listen signal
	sig := utils.NewSignal()
	sig.Waiter()

	//cron service: register tasks
	tasks.TasksRegister(sig)

	//Gin Service: Listen 0.0.0.0:8000
	router := router.InitRouter()
	if err := endless.ListenAndServe(fmt.Sprintf(`%v:%v`, config.Config.Http.Host, config.Config.Http.Port), router); err != nil {
		fmt.Printf("[Main Server] the main goroutine has exited, error:%v\n", err)
	}

	//over
	fmt.Println("[Main Server] Server is ready to stop!")
}
