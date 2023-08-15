package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"gwprj/project"
	"gwprj/common"
	"gwprj/crv"
	"time"
	"log"
	"os"
)

func main() {
	//设置log打印文件名和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	//初始化时区
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone

	confFile:="conf/conf.json"
	if len(os.Args)>1 {
		confFile=os.Args[1]
		log.Println(confFile)
	}
	
	conf:=common.InitConfig(confFile)

	//crvClinet 用于到crvframeserver的请求
	crvClinet:=&crv.CRVClient{
		Server:conf.CRV.Server,
		Token:conf.CRV.Token,
		AppID:conf.CRV.AppID,
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:true,
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	projectController:=&project.ProjectController{
		CRVClient:crvClinet,
	}

	projectController.Bind(router)

	router.Run("0.0.0.0:8300")
}