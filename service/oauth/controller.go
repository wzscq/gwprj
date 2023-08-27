package oauth

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"fmt"
)

type getLoginPageReq struct {
	AppID     string   `json:"appID"`
}

type getLoginPageRsp struct {
	Url     string   `json:"url"`
}

type oauthLoginReq struct {
	UserID     string  `json:"userID"`
  Password  string   `json:"password"`
	AppID     string   `json:"appID"`
	RedirectUri string   `json:"redirectUri"`
	ClientID string `json:"clientID"`
}

type oauthBackReq struct {
  OAuthCode  string   `json:"oauthCode"`
	AppID     string   `json:"appID"`
}

type accessTokenReq struct {
	ClientID string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
	Code string `form:"code"`
}

type userInfoReq struct {
	Authorization string `json:"Authorization"`
}

type OauthToken struct {
	AccessToken string `json:"access_token"`
}

type OAuthController struct {
	OAuthCache *OAuthCache
	BackUrl string
}

const (
	API_OAUTH2_AUTHORIZE="oauth2_authorize"
	API_OAUTH2_ACCESSTOKEN="oauth2_accessToken"
	API_OAUTH2_USERINFO="oauth2_userInfo"
)

func (controller *OAuthController)login(c *gin.Context) {
	log.Println("start OAuthController login")
	userID:=c.Query("userId")
	//redirectUri:=c.Query("redirectUri")
	token:=getOAuthToken()
  log.Println("userID:",userID,"token:",token)
	//将token存入缓存
	controller.OAuthCache.SetCache(userID,token)

	c.Header("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	c.Header("Pragma", "no-cache") // HTTP 1.0.
	c.Header("Expires", "0") // Proxies.

	//重定向web到给定的回调地址
	url:=fmt.Sprintf(controller.BackUrl,url.QueryEscape(token))
	//url:="http://localhost:3000/mainframe/#/OAuthBack/gwprj?code="+
	c.Redirect(http.StatusMovedPermanently, url)
	log.Println("end OAuthController login Redirect to",url)
}

func (controller *OAuthController)accessToken(c *gin.Context){
	//这里暂时不对client密码做校验
	log.Println("start OAuthController accessToken")
	var req accessTokenReq
	if err := c.ShouldBind(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Println("end OAuthController accessToken with error")
		return 
    }	
	log.Println(req)
	//这里不做处理，直接将code作为token返回
	token:=&OauthToken{
		AccessToken:req.Code,
	}
	//rsp:=common.CreateResponse(nil,token)
	c.IndentedJSON(http.StatusOK, token)
	log.Println("end OAuthController accessToken")
}

func (controller *OAuthController)userInfo(c *gin.Context){
	log.Println("start OAuthController userInfo")
	var req userInfoReq
	if err := c.ShouldBindHeader(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Println("end OAuthController userInfo with error")
		return 
    }	
	log.Println(req)
	if len(req.Authorization)<7 {
		log.Println("OAuthController userInfo request token is too short")
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Println("end OAuthController userInfo with error")
		return 
	}
	token:=req.Authorization[6:]
	log.Println("token is:",token)
	
	userID,err:=controller.OAuthCache.GetUserID(token)
	if err!=nil {
		log.Println("OAuthController userInfo GetUserID error",err)
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Println("end OAuthController userInfo with error")
		return 
	}

	result:=map[string]string{
		"id":userID,
	}
	//rsp:=common.CreateResponse(nil,token)
	c.IndentedJSON(http.StatusOK, result)
	log.Println("end OAuthController userInfo",result)
}

func (controller *OAuthController) Bind(router *gin.Engine) {
	log.Println("Bind OAuthController")
	router.GET("/oauth/login", controller.login)
	router.POST("/oauth/accessToken", controller.accessToken)
	router.GET("/oauth/userInfo", controller.userInfo)
}