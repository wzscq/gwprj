package common

type CommonRsp struct {
	ErrorCode int `json:"errorCode"`
	Message string `json:"message"`
	Error bool `json:"error"`
	Result interface{} `json:"result"`
	Params map[string]interface{} `json:"params"`
}

type CommonError struct {
	ErrorCode int `json:"errorCode"`
	Params map[string]interface{} `json:"params"`
	Message string `json:"message"`
}

const (
	ResultSuccess = 10000000
	ResultReadTempleteFileError = 30100001
	ResultSaveFileError = 30100002
	ResultCreatePdfError = 30100003
	ResultSaveDataError = 30100004
	ResultQueryRequestError = 30100005
	ResultWrongRequest = 30100006
)

var errMsg = map[int]CommonRsp{
	ResultSuccess:CommonRsp{
		ErrorCode:ResultSuccess,
		Message:"操作成功",
		Error:false,
	},
	ResultReadTempleteFileError:CommonRsp{
		ErrorCode:ResultReadTempleteFileError,
		Message:"读取模板文件失败，请联系管理管理员处理",
		Error:true,
	},
	ResultSaveFileError:CommonRsp{
		ErrorCode:ResultSaveFileError,
		Message:"保存文件时发生错误，请联系管理管理员处理",
		Error:true,
	},
	ResultCreatePdfError:CommonRsp{
		ErrorCode:ResultCreatePdfError,
		Message:"生成PDF文件时发生错误，请联系管理管理员处理",
		Error:true,
	},
	ResultSaveDataError:CommonRsp{
		ErrorCode:ResultSaveDataError,
		Message:"保存数据到数据时发生错误，请与管理员联系处理",
		Error:true,
	},
	ResultQueryRequestError:CommonRsp{
		ErrorCode:ResultQueryRequestError,
		Message:"从数据库查询数据时发生错误，请与管理员联系处理",
		Error:true,
	},
	ResultWrongRequest:CommonRsp{
		ErrorCode:ResultWrongRequest,
		Message:"请求参数错误，请与管理员联系处理",
		Error:true,
	},
}

func CreateResponse(err *CommonError,result interface{})(*CommonRsp){
	if err==nil {
		commonRsp:=errMsg[ResultSuccess]
		commonRsp.Result=result
		return &commonRsp
	}

	commonRsp,ok:=errMsg[err.ErrorCode]
	if !ok {
		commonRsp=CommonRsp{
			ErrorCode:err.ErrorCode,
			Message:err.Message,
			Error:true,
		}
	}
	commonRsp.Result=result
	commonRsp.Params=err.Params
	return &commonRsp
}

func CreateError(errorCode int,params map[string]interface{})(*CommonError){
	return &CommonError{
		ErrorCode:errorCode,
		Params:params,
	}
}