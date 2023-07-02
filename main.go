package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/semeru1-go/modules"
	"github.com/purawaktra/semeru1-go/utils"
)

func main() {
	utils.InitConfig()
	utils.InitLog()

	utils.Log("============= SEMERU1 RUNTIME STARTED =============")
	// create gin engine
	engine := gin.New()

	// create dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		utils.MysqlUser,
		utils.MysqlPassword,
		utils.MysqlHost,
		utils.MysqlPort,
		utils.MysqlScheme)

	gormInstance, err := utils.CreateGorm(dsn)
	if err != nil {
		utils.Error(err, "main", "")
	}

	// declare architecture
	repository := modules.CreateSemeru1Repo(gormInstance)
	usecase := modules.CreateSemeru1Usecase(repository)
	controller := modules.CreateSemeru1Controller(usecase)
	requestHandler := modules.CreateSemeru1RequestHandler(controller)
	router := modules.CreateSemeru1Router(engine, requestHandler)

	// init routing
	router.Init("/semeru1/api/v1")

	utils.Log("============= SEMERU1 ENGINE STARTING =============")
	utils.Log(fmt.Sprintf("App Address = %s:%s", utils.AppHost, utils.AppPort))

	// start http api engine
	err = engine.Run(fmt.Sprintf("%s:%s", utils.AppHost, utils.AppPort))
	if err != nil {
		utils.Fatal(err, "main", "")
		panic(err)
	}

	utils.Log("============= SEMERU1 ENGINE FAILED =============")
}
