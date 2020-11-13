package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func main()	{
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")   //更改一下监听端口
	if port != ""{
		panic (r.Run(":" + port))
	}
	panic(r.Run()) //listen and serve on 0.0.0.0:8080
}

func InitConfig(){
	workDir,_ := os.Getwd()
	viper.SetConfigName("application")      //获取文件名
	viper.SetConfigType("yml")              //获取文件格式
	viper.AddConfigPath(workDir + "/config")   //获取文件路径
	err := viper.ReadInConfig()
	if err != nil{
		panic(err)
	}
}