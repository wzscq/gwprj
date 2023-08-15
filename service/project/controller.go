package project

import (
	"gwprj/common"
	"gwprj/crv"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

type ProjectController struct {
	CRVClient *crv.CRVClient
}

func (controller *ProjectController)downloadReport(c *gin.Context){
	log.Println("ProjectController start downloadReport")
	
	var header crv.CommonHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		log.Println(err)
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		log.Println("end downloadReport with error")
		return
	}	
	
	var rep crv.CommonReq
	if err := c.BindJSON(&rep); err != nil {
		log.Println(err)
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		return
  	}	

	if rep.SelectedRowKeys==nil || len(*rep.SelectedRowKeys)==0 {
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		log.Println("downloadReport error：request list is empty")
		return
	}

	//获取数据
	res:=GetProjectData(rep.SelectedRowKeys,controller.CRVClient,header.Token)
	log.Println(res)

	if res!=nil && len(res)>0 {
		var fileName string
		if len(res)==1 {
			fileName=GetReportFileName(res[0].(map[string]interface{}))+".pdf"
		} else {
			fileName=GetBatchID()+".zip"
		}

		fileName=url.QueryEscape(fileName)
		
		c.Header("Content-Type", "application/octet-stream")
    	c.Header("Content-Disposition", "attachment; filename="+fileName)
    	c.Header("Content-Transfer-Encoding", "binary")
	
		//生成报表
		CreateReports("closingreport",res,c.Writer)
	}
	
	log.Println("ProjectController end downloadReport")
}

//Bind bind the controller function to url
func (controller *ProjectController) Bind(router *gin.Engine) {
	log.Println("Bind ProjectController")
	router.POST("project/downloadReport", controller.downloadReport)
}