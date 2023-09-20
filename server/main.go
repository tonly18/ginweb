package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"server/config"
	"server/global"
	"server/router"
)

func main() {
	//Running Environment
	fmt.Printf("[Server Running Environment: %v]\n", global.SERVER_RUN_ENV)
	fmt.Printf("[Gin Mode: %v]\n", global.GIN_MODE)

	//init router
	r := router.InitRouter()

	//listen: 0.0.0.0:8080
	if err := endless.ListenAndServe(config.Config.Http.Host+":"+config.Config.Http.Port, r); err != nil {
		fmt.Println(err)
	}

	//over
	fmt.Println("Game Server is ready to stop!")
}
