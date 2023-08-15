package project

import (
	"regexp"
	"strings"
	"log"
	"encoding/json"
	"fmt"
)

type QueryResult struct {
	ModelID string `json:"modelID"`
	ViewID *string `json:"viewID,omitempty"`
	Value *string `json:"value,omitempty"`
	Total int `json:"total"`
	Summaries *map[string]interface{} `json:"summaries,omitempty"`
	List []map[string]interface{} `json:"list"`
}

var re=regexp.MustCompile(`%{([A-Z|a-z|_|0-9|.]*)}`)

func ContentFilter(content string,data map[string]interface{})(string){
	replaceItems:=re.FindAllStringSubmatch(content,-1)
	
	if replaceItems!=nil {
		log.Println(content)
		for _,replaceItem:=range replaceItems {
			log.Printf("ContentFilter replaceItem:%s,%s \n",replaceItem[0],replaceItem[1])
			repalceStr:=getfilterDataString(replaceItem[1],data)
			content=strings.Replace(content,replaceItem[0],repalceStr,-1)
		}
		log.Println(content)
	}
	return content
}

func getfilterDataString(path string,data map[string]interface{})(string){
	/*
		data中按模型保存的查询结果数据，数据结构如下
		{
			modelID:{
				ModelID
				Value
				Total
				List [
					{
						fieldName:value
						fieldName:{ //对于关联字段，器值得结构和第一层的结构一致，允许多层级关联嵌套
							modelID
							value
							total
							list:[...] 
						},
						...
					},
					...
				]
			}
		}
		基于以上数据结构，在查询条件中引用某个字段值的情况和使用方式如下：
		1、第一层结构通过modelID区分取值于哪个model，后续层级都是通过关联字段引用的，使用关联字段名称来表示，
		   不同层级件使用点号间隔，距离如下：
		   core_user.id：表示获取core_user表中的id字段的值；
		   core_user.roles.id：表示获取core_user表的reles关联字段表中的id字段的值；
		2、通常对于每个层级都存在多条记录的情况，将会自动获取所有层级记录中的所有值，并进行去重处理，去重后的值生成
		   如下字符串形式：role1","role2","role3，因此在配置查询条件时，应该使用类似
		   {Op.in:["%{core_user.roles.id}"]}这样的形式，程序会将变量%{core_user.roles.id}替换为role1","role2","role3
		   替换后的字符将改为：{Op.in:["role1","role2","role3"]}
	*/
	//首先对path按照点好拆分
	values:=[]string{}
	pathNodes:=strings.Split(path, ".")
	getPathData(pathNodes,0,data,&values)
	//将value转为豆号分割的字符串
	if len(values)>0 {
		valueStr:=strings.Join(values, "\",\"")
		log.Println(valueStr)
		return valueStr
	}
	return ""  //如果没有匹配上就用空字符串替换
}

func getPathData(path []string,level int,data map[string]interface{},values *[]string){
	pathNode:=path[level]

	dataStr, _ := json.Marshal(data)
	log.Printf("getPathData pathNode:%s,level:%d,data:%s",pathNode,level,string(dataStr))

	dataNode,ok:=data[pathNode]
	if !ok {
		log.Println("getPathData no pathNode ",pathNode)
		return
	}

	//如果当前层级为最后一层
	if len(path)==(level+1) {
		switch dataNode.(type) {
			case string:
				sVal, _ := dataNode.(string)   
				*values=append(*values,sVal) 
			case int64:
				iVal,_:=dataNode.(int64)
				sVal:=fmt.Sprintf("%d",iVal)
				*values=append(*values,sVal) 
			default:
				log.Printf("getPathData not supported value type %T!\n", dataNode)
		}
	} else {
		//如果不是最后一级，则数据中应该存在list属性
		log.Printf("dataNode type is %T",dataNode)
		result,ok:=dataNode.(map[string]interface{})
		if !ok {
			log.Println("getPathData dataNode is not a map[string]interface{} ")
			return
		}

		list:=result["list"].([]interface{})

		for _,row:=range list {
			rowMap,ok:=row.(map[string]interface{})
			if ok {
				log.Println("getPathData dataNode list member is not a map ")
				getPathData(path,level+1,rowMap,values)
			}
		}
		return
	}
}