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
	{"field":"pm"},
	{
		"field":"approver0",
		"fieldType":"one2many",
		"relatedModelID":"gw_project_approver",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"desc"},
		},
		"filter":map[string]interface{}{
			"role":"事务所审批",
		},
		"pagination":map[string]interface{}{
			"current":1,
			"pageSize":1,
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
	{
		"field":"approver1",
		"fieldType":"one2many",
		"relatedModelID":"gw_project_approver",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"asc"},
		},
		"filter":map[string]interface{}{
			"role":"财务审批",
		},
		"pagination":map[string]interface{}{
			"current":1,
			"pageSize":1,
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
	{
		"field":"approver2",
		"fieldType":"one2many",
		"relatedModelID":"gw_project_approver",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"asc"},
		},
		"filter":map[string]interface{}{
			"role":"财务审批",
		},
		"pagination":map[string]interface{}{
			"current":2,
			"pageSize":1,
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
	{
		"field":"approver3",
		"fieldType":"one2many",
		"relatedModelID":"gw_project_approver",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"asc"},
		},
		"filter":map[string]interface{}{
			"role":"财务审批",
		},
		"pagination":map[string]interface{}{
			"current":3,
			"pageSize":1,
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
	{
		"field":"approver4",
		"fieldType":"one2many",
		"relatedModelID":"gw_project_approver",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"asc"},
		},
		"filter":map[string]interface{}{
			"role":"财务审批",
		},
		"pagination":map[string]interface{}{
			"current":4,
			"pageSize":1,
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
	{
		"field":"approver5",
		"fieldType":"one2many",
		"relatedModelID":"gw_project_approver",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"asc"},
		},
		"filter":map[string]interface{}{
			"role":"财务审批",
		},
		"pagination":map[string]interface{}{
			"current":5,
			"pageSize":1,
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
	
	if len(*ids)==0 {
		return nil
	}

	res:=[]interface{}{}
	for _,id:=range *ids {
		//查询数据
		commonRep:=crv.CommonReq{
			ModelID:"gw_project",
			Filter:&map[string]interface{}{
				"id":id,
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

		if len(req.Result["list"].([]interface{}))>0 {
			res=append(res,req.Result["list"].([]interface{})[0])
		}
	}
	return res
}

func GetReportFileName(project map[string]interface{})(string){
	projectID:=project["id"].(string)
	projectName:=project["project_name"].(string)
	return projectID+"_"+projectName
}

func GetReportData(data map[string]interface{})(map[string]interface{}){
	return data
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
		//对数据做个处理
		mapData:=GetReportData(data.(map[string]interface{}))
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

func CreateReport(tmpName string,data map[string]interface{},outPath string){
	//加载模板
	tmp,err:=LoadTempleteFromJsonFile("./templetes/"+tmpName+".json")
	if err!=nil {
		log.Println(err)
		return
	}

	pdf:=CreatePdf(tmp,data)
	if pdf!=nil {
		//保存文件
		pdf.WritePdf(outPath+"/report.pdf")
	}
}
