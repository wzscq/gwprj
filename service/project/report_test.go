package project

import (
	"testing"
	"log"
)

func TestCreatePDF(t *testing.T){
	log.Println("TestCreatePDF ...")
	tmp,err:=LoadTempleteFromJsonFile("../templetes/closingreport20231015.json")
	if err!=nil {
		log.Println(err)
		return
	}

	data:=map[string]interface{}{
		"approver0":map[string]interface{}{
			"list":[]interface{}{
				map[string]interface{}{
					"approver":map[string]interface{}{
						"list":[]interface{}{
							map[string]interface{}{
								"id":"architect1",
								"user_name_zh":"事务所1",
							},
						},
						"modelID":"core_user",
						"total":1,
						"value":"architect1",
					},
					"create_time":"2023-10-15 09:23:08",
					"id":"3",
					"project_id":"test20231015001",
					"role":"事务所审批",
					"update_time":"2023-10-15 09:23:30",
				},
			},
			"modelID":"gw_project_approver",
			"total":1,
		},
		"approver1":map[string]interface{}{
			"list":[]interface{}{
				map[string]interface{}{
					"approver":map[string]interface{}{
						"list":[]interface{}{
							map[string]interface{}{
								"id":"approver1",
								"user_name_zh":"审批人1",
							},
						},
						"modelID":"core_user",
						"total":1,
						"value":"approver1",
					},
					"create_time":"2023-10-15 09:23:30",
					"id":"4",
					"project_id":"test20231015001",
					"role":"财务审批",
					"update_time":"2023-10-15 09:23:50",
				},
			},
			"modelID":"gw_project_approver",
			"total":1,
		},
		"approver2":map[string]interface{}{
			"list":[]interface{}{
				map[string]interface{}{
					"approver":map[string]interface{}{
						"list":[]interface{}{
							map[string]interface{}{
								"id":"approver2",
								"user_name_zh":"审批人2",
							},
						},
						"modelID":"core_user",
						"total":1,
						"value":"approver2",
					},
					"create_time":"2023-10-15 09:23:30",
					"id":"5",
					"project_id":"test20231015001",
					"role":"财务审批",
					"update_time":"2023-10-15 09:28:16",
				},
			},
			"modelID":"gw_project_approver",
			"total":1,
		},
		"end_date":"2023-10-03 00:00:00",
		"id":"test20231015001",
		"pm":"test1",
		"project_name":"test20231015001",
		"start_date":"2023-10-01 00:00:00",
		"statement":"项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项 目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项 目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项 目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项 目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况项目概况",
	}

	pdf:=CreatePdf(tmp,data)

	err=pdf.WritePdf("test.pdf")
	if err!=nil {
		log.Println(err)
		return
	}
}