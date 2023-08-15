package project

import (
	"gwprj/common"
	"gwprj/crv"
	"log"
	"archive/zip"
	"io"
)

var projectFields=[]map[string]interface{}{
	{"field":"id"},
	{"field":"project_name"},
	{"field":"start_date"},
	{"field":"end_date"},
	{"field":"statement"},
	{
		"field":"approvers",
		"fieldType":"one2many",
		"relatedModelID":"gw_project_approver",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"desc"},
		},
		"fields":[]map[string]interface{}{
			{"field":"id"},
			{
				"field":"approver",
				"fieldType":"many2one",
				"relatedModelID":"core_user",
				"fields":[]map[string]interface{}{
					{"field":"id"},
					{"field":"user_name_zh"},
				},
			},
			{"field":"role"},
			{"field":"project_id"},
			{"field":"update_time"},
			{"field":"create_time"},
		},
	},
}

func GetProjectData(ids *[]string,crvClinet *crv.CRVClient,token string)([]interface{}){
	//查询数据
	commonRep:=crv.CommonReq{
		ModelID:"gw_project",
		Filter:&map[string]interface{}{
			"id":map[string]interface{}{
				"Op.in":ids,
			},
		},
		Fields:&projectFields,
	}

	req,commonErr:=crvClinet.Query(&commonRep,token)
	if commonErr!=common.ResultSuccess {
		return nil
	}

	if req.Error == true {
		log.Println("GetProjectData error:",req.ErrorCode,req.Message)
		return nil
	}

	return req.Result["list"].([]interface{})
}

func GetReportFileName(project map[string]interface{})(string){
	projectID:=project["id"].(string)
	projectName:=project["project_name"].(string)
	return projectID+"_"+projectName
}

func CreateReports(tmpName string,list []interface{},w io.Writer)(int){
	//加载模板
	tmp,err:=LoadTempleteFromJsonFile("./templetes/"+tmpName+".json")
	if err!=nil {
		return common.ResultReadTempleteFileError
	}
	
	//如果下载的报告数量大于1，则打包成zip
	var zipWriter *zip.Writer
	if len(list)>1 {
		zipWriter= zip.NewWriter(w)
		defer zipWriter.Close()
	}

	for _,data:=range list {
		mapData:=data.(map[string]interface{})
		pdf:=CreatePdf(tmp,mapData)
		if pdf!=nil {
			if zipWriter!=nil {
				outFileName:=GetReportFileName(mapData)+".pdf"
				fileHeader:=&zip.FileHeader{
					Name:outFileName,
				}
				fileWriter,err:=zipWriter.CreateHeader(fileHeader)
				if err!=nil {
					log.Println(err)
				} 
				pdf.Write(fileWriter)
			} else {
			
				pdf.Write(w)
			}
		}
	}

	return common.ResultSuccess
}
