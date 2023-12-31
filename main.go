package main

import (
	"os"

	"GIN_GORM/common"
	"GIN_GORM/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	//gin.SetMode(gin.ReleaseMode)
	//db_student := common.GetDB()
	//defer db_student.Close()
	common.GetDB_User()
	// common.GetDB_Teacher()

	r := gin.Default()
	r = router.CollectRouter(r)
	r.Run(":9999")
}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
