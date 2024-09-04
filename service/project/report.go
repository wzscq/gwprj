package project

import (
	"gwprj/common"
	"gwprj/crv"
	"log"
	"archive/zip"
	"io"
	"time"
	"strconv"
)

var ProjectFields=[]map[string]interface{}{
	{"field":"id"},
	{"field":"project_name"},
	{"field":"statement"},
	{"field":"plan_number"},
	{"field":"branch"},
	{"field":"department"},
	{
		"field":"create_user",
		"fieldType":"many2one",
		"relatedModelID":"core_user",
		"fields":[]map[string]interface{}{
			{"field":"id"},
			{"field":"user_name_zh"},
		},
	},
	{
		"field":"dh",
		"fieldType":"one2many",
		"relatedModelID":"gw_audit_rec",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"desc"},
		},
		"filter":map[string]interface{}{
			"role":"业务部门负责人",
			"audit_result":"1",
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
			{"field":"comment"},
			{"field":"project_id"},
			{"field":"update_time"},
			{"field":"create_time"},
		},
	},
	{
		"field":"ssds",
		"fieldType":"one2many",
		"relatedModelID":"gw_audit_rec",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"desc"},
		},
		"filter":map[string]interface{}{
			"role":"安监专责",
			"audit_result":"1",
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
			{"field":"comment"},
			{"field":"project_id"},
			{"field":"update_time"},
			{"field":"create_time"},
		},
	},
	{
		"field":"ssdh",
		"fieldType":"one2many",
		"relatedModelID":"gw_audit_rec",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"desc"},
		},
		"filter":map[string]interface{}{
			"role":"安监负责人",
			"audit_result":"1",
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
			{"field":"comment"},
			{"field":"project_id"},
			{"field":"update_time"},
			{"field":"create_time"},
		},
	},
	{
		"field":"bl",
		"fieldType":"one2many",
		"relatedModelID":"gw_audit_rec",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"desc"},
		},
		"filter":map[string]interface{}{
			"role":"业务分管领导",
			"audit_result":"1",
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
			{"field":"comment"},
			{"field":"project_id"},
			{"field":"update_time"},
			{"field":"create_time"},
		},
	},
	{
		"field":"ssbl",
		"fieldType":"one2many",
		"relatedModelID":"gw_audit_rec",
		"relatedField":"project_id",
		"sorter":[]map[string]interface{}{
			{"field":"update_time","order":"desc"},
		},
		"filter":map[string]interface{}{
			"role":"安监分管领导",
			"audit_result":"1",
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
			{"field":"comment"},
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
				"id":map[string]interface{}{
					"Op.eq":id,
				},
			},
			Fields:&ProjectFields,
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
	projectID:=project["project_id"].(string)
	projectName:=project["project_name"].(string)
	return projectID+"_"+projectName
}

func GetReportID(overTime,reportID string)(string){
	//字符串转换为date对象
	date,_:=time.Parse("2006-01-02 15:04:05",overTime)
	//日期格式化为字符串
	overTime=date.Format("20060102150405")
	//reportID转换为整数
	id,_:=strconv.Atoi(reportID)
	//取1000的余数
	id=id%1000
	//转为字符串
	reportID=strconv.Itoa(id)
	//不足3位的补0
	for len(reportID)<3 {
		reportID="0"+reportID
	}
	return overTime+reportID
}

func GetReportData(data map[string]interface{})(map[string]interface{}){
	//最后一个审批时间作为报告时间
	overTime:=getfilterDataString("ssbl.create_time",data)
	log.Println("GetReportData overTime:"+overTime)
	//生成报告编号，用最后一个审批时间+审批记录ID
	reportID:=getfilterDataString("ssbl.id",data)
	reportID=GetReportID(overTime,reportID)
	branch:=getfilterDataString("branch",data)
	if branch=="眉山公司本部" {
		branch=""
	}
	
	//转换日期格式
	newData:=map[string]interface{}{
		"project_id":data["id"],
		"project_name":data["project_name"],
		"statement":data["statement"],
		"plan_number":data["plan_number"],
		"branch":branch,
		"department":data["department"],
		"over_time":overTime[0:10],
		"report_id":reportID,
		"create_user":data["create_user"],
		"dh":data["dh"],
		"ssds":data["ssds"],
		"ssdh":data["ssdh"],
		"bl":data["bl"],
		"ssbl":data["ssbl"],
	}

	return newData
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
