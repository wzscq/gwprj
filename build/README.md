docker run --name gwprjservice -d -p8302:80 -v /root/gwprjservice/font:/services/gwprjservice/font -v /root/gwprjservice/templetes:/services/gwprjservice/templetes -v /root/gwprjservice/conf:/services/gwprjservice/conf  wangzhsh/gwprjservice:0.1.0


2023-06-22 第一次创建

2023-10-10 更新以下问题
1、数据库表gw_project.project_name字段长度从50改为100
2、nginx配置文件中对上传内容的限制从50m改为1024m
3、配置文件修改gw_project/forms/pmEdit、gw_project/forms/detail、  gw_project/forms/approvalDetail 
4、修改最外层operatios中/gwprj/function的权限配置，允许所有人访问{roles:"*"}
5、上传文件的方式改进
   替换crvframe镜像
   修改crvframe配置文件conf.json，增加upload相应配置项
   修改gw_project/forms/pmEdit、gw_project/forms/detail中对应文件上传控件
6、导出报告格式完善，部分字段值需要更长


2023-10-15
1、修改了项目结项报告的模板文件
