package main

import (
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/cache"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/dao"
	"github.com/spike-engine/spike-web3-server/global"
	"github.com/spike-engine/spike-web3-server/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description
// @securityDefinitions.api_key api_key
// @in header
// @name Authorization
// @BasePath /
func main() {
	logger.SetLogLevel("*", "INFO")
	global.Viper = config.InitViper()
	global.GormClient = initialize.GormMysql()
	global.DbAccessor = dao.NewGormAccessor(global.GormClient)
	global.RedisClient = cache.ConnectRedis()
	//chain.NewBscListener()
	initialize.RunServer()
}
