package project

import (
	"testing"
	"log"
	"gwprj/crv"
)

func TestCreatePDF(t *testing.T){
	log.Println("TestCreatePDF ...")
	
	crvClinet:=&crv.CRVClient{
		Server:"http://localhost:8200",
		Token:"5ff61df532082114923014268c145e1f",
		AppID:"gwprj",
	}

	ids:=[]string{"p003"}

	list:=GetProjectData(&ids,crvClinet,"03da9897e272a2b396678ae080347151")

	//log.Println(list)

	tmp,err:=LoadTempleteFromJsonFile("../templetes/closingreport.json")
	if err!=nil {
		log.Println(err)
		return
	}

	for _,data:=range list {
		//对数据做个处理
		mapData:=GetReportData(data.(map[string]interface{}))
		pdf:=CreatePdf(tmp,mapData)
		err=pdf.WritePdf("test.pdf")
		if err!=nil {
			log.Println(err)
			return
		}
	}
}