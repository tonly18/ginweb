package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"server/config"
	"server/global"
	"server/router"
	"server/utils"
)

func main() {
	//Running Environment
	fmt.Printf("[Server Running Environment: %v]\n", global.SERVER_RUN_ENV)
	fmt.Printf("[Gin Mode: %v]\n", global.GINMODE)

	//init router
	r := router.RouterInit()

	//listen: 0.0.0.0:8080
	err := endless.ListenAndServe(config.Config.Http.Host+":"+config.Config.Http.Port, r)
	if err != nil {
		fmt.Println(err)
	}

	//程序结束后，完成清理工作
	utils.Finish()

	//over
	fmt.Println("Game Server is ready to stop...")
}
