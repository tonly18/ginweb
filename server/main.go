package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/pprof"
	"server/config"
	"server/global"
	"server/router"
)

func main() {
	//Running Environment
	fmt.Printf("[Server] Running Environment: %v\n", global.SERVER_RUN_ENV)
	fmt.Printf("[Server] Running Gin Mode: %v\n", global.GIN_MODE)

	//init router
	router := router.InitRouter()

	//pprof
	pprof.Register(router)

	//listen: 0.0.0.0:8000
	if err := endless.ListenAndServe(fmt.Sprintf(`%v:%v`, config.Config.Http.Host, config.Config.Http.Port), router); err != nil {
		fmt.Println(err)
	}

	//over
	fmt.Println("[Server] server is ready to stop!")
}
