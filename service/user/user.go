package user

import (
	"log"
	"gwprj/crv"
	"gwprj/common"
)

var userInfoFields=[]map[string]interface{}{
	{"field":"id"},
	{"field":"version"},
	{
		"field":"roles",
		"fieldType":"many2many",
		"relatedModelID":"core_role",
		"fields":[]map[string]interface{}{
			{"field":"id"},
			{"field":"version"},
		},
	},
}

var roleFields=[]map[string]interface{}{
	{"field":"id"},
}

type UserBusi struct{
	UserRepository UserRepository
	CRVClient *crv.CRVClient
}

func (busi *UserBusi)UpdateUserRoles(userID string){
	//获取用户角色信息
	roles:=busi.UserRepository.GetUserRoles(userID)
	log.Println("rbiRoles:",roles)
	//获取CRV角色信息
	crvRoles:=busi.GetCRVRoles()
	log.Println("crvRoles:",crvRoles)
	//比较角色信息
	userRoles:=busi.GetUserRoles(roles,crvRoles)
	log.Println("userRoles:",userRoles)
	//获取CRV用户信息
	if len(userRoles)>0 {
		userInfo:=busi.GetCRVUserInfo(userID)
		if userInfo!=nil {
			log.Println("userInfo:",userInfo)
			busi.UpdateCRVUser(userInfo,userRoles)
		} else {
			busi.CreateCRVUser(userID,userRoles)
		}
	}
}

func (busi *UserBusi)GetCRVUserInfo(userID string)(map[string]interface{}){
	//查询数据
	commonRep:=crv.CommonReq{
		ModelID:"core_user",
		Fields:&userInfoFields,
		Filter:&map[string]interface{}{
			"id":userID,
		},
	}

	req,commonErr:=busi.CRVClient.Query(&commonRep,"")
	if commonErr!=common.ResultSuccess {
		return nil
	}

	if req.Error == true {
		log.Println("GetProjectData error:",req.ErrorCode,req.Message)
		return nil
	}

	if req.Result["list"]!=nil && len(req.Result["list"].([]interface{}))>0 {
		return req.Result["list"].([]interface{})[0].(map[string]interface{})
	}
	return nil
}

func (busi *UserBusi)UpdateCRVUser(userInfo map[string]interface{},userRoles []string){
	crvRoles:=userInfo["roles"].(map[string]interface{})["list"].([]interface{})

	roleList:=[]map[string]interface{}{}
	//删除没有的角色
	for _,crvRole:=range crvRoles {
		crvRoleID:=crvRole.(map[string]interface{})["id"].(string)
		hasRole:=false
		for _,userRole:=range userRoles {
			if crvRoleID==userRole {
				hasRole=true
				break
			}
		}
		if hasRole==false {
			roleList=append(roleList,map[string]interface{}{
				"id":crvRoleID,
				"version":crvRole.(map[string]interface{})["version"],
				"_save_type":"delete",
			})
		}
	}

	//添加新的角色
	for _,userRole:=range userRoles {
		hasRole:=false
		for _,crvRole:=range crvRoles {
			if userRole==crvRole.(map[string]interface{})["id"].(string) {
				hasRole=true
				break
			}
		}
		if hasRole==false {
			roleList=append(roleList,map[string]interface{}{
				"id":userRole,
				"_save_type":"create",
			})
		}
	}

	if len(roleList)==0 {
		return
	}

	commonRep:=crv.CommonReq{
		ModelID:"core_user",
		List:&[]map[string]interface{}{
			{
				"id":userInfo["id"],
				"_save_type":"update",
				"version":userInfo["version"],
				"roles":map[string]interface{}{
					"list":roleList,
					"fieldType":"many2many",
					"modelID":"core_role",
				},
			},
		},
	}

	busi.CRVClient.Save(&commonRep,"")
}

func (busi *UserBusi)CreateCRVUser(userID string,userRoles []string){
	//查询数据
  roleList:=[]map[string]interface{}{}
	for _,role:=range userRoles {
		roleList=append(roleList,map[string]interface{}{
			"id":role,
			"_save_type":"create",
		})
	}

	commonRep:=crv.CommonReq{
		ModelID:"core_user",
		List:&[]map[string]interface{}{
			{
				"id":userID,
				"password":"a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
				"_save_type":"create",
				"roles":map[string]interface{}{
					"list":roleList,
					"fieldType":"many2many",
					"modelID":"core_role",
				},
			},
		},
	}
	busi.CRVClient.Save(&commonRep,"")
}

func (busi *UserBusi)GetUserRoles(rbiRoles []UserRole,crvRoles []string)([]string){
	userRoles:=[]string{}
	for _,rbiRole:=range rbiRoles{
		for _,crvRole:=range crvRoles {
			roleStr:=rbiRole.ModuleType+"$"+rbiRole.ResourceType+"$"+crvRole
			log.Println("roleStr:",roleStr,"rbiRole.ResourceId:",rbiRole.ResourceId)
			if roleStr==rbiRole.ResourceId {
				userRoles=append(userRoles,crvRole)
			}
		}
	}
	return userRoles
}

func (busi *UserBusi)GetCRVRoles()([]string){
	//查询数据
	commonRep:=crv.CommonReq{
		ModelID:"core_role",
		Fields:&roleFields,
	}

	req,commonErr:=busi.CRVClient.Query(&commonRep,"")
	if commonErr!=common.ResultSuccess {
		return nil
	}

	if req.Error == true {
		log.Println("GetProjectData error:",req.ErrorCode,req.Message)
		return nil
	}

	roles:=[]string{}
	if len(req.Result["list"].([]interface{}))>0 {
		for _,role:=range req.Result["list"].([]interface{}){
			roleMap:=role.(map[string]interface{})
			roles=append(roles,roleMap["id"].(string))
		}
	}
	
	return roles
}