# MyLottery


### 功能实现

   外部分为用户和管理员两部分。用户模块实现了，登录，注册，邮箱验证，重置密码，个人资料修改，跨域Token，抽奖，对奖品的操作。管理员模块实现了登录，token验证，对奖品，抽奖活动，用户，用户抽到的奖品四个模块的增删改查。
		内部实现了跨域，token鉴权，mysql，redis处理数据，emai的控制，grpc协议编程，熔断处理，文件配置及实时监听，日志持久化，实时，循环分割及软连接。swagger，postman接口文档。功能逻辑的详细注释说明。错误处理，逻辑封装，Makefile命令，doker部署等功能

### 技术应用

前端：vue3框架，element-ui、vuesax界面组件库，tinymce富文本组件库等
后端：mysql、redis数据库。gin、gorm、go-micro框架。prorobuf生成文件，rabbitmq消息队列，consul服务发现，hystrix熔断处理，viper配置管理，logurs、rotatelogs日志管理，swaggo接口文档，cors跨域，jwt跨域鉴权，md5加密，email邮箱配置等

### 安装教程


#### 直接安装访问
需要配置go环境，并安装mysql：8.0.28，创建数据库mylottery
安装redis:5.0.14 

需要下载consul，rabbitmq并开启这两项服务
需要安装go-micro工具

修改config.yaml文件
执行命令go mod tidy下载go依赖包
go run main.go执行


### 使用说明

需要在windows下或者docker下安装mysql,redis
如果主机名不是localhost，端口号想修改，需要修改config/config.yaml

swagger导入依赖后，swag init初始化 进入路由界面即可http://localhost:9002/swagger/index.html
授权token格式 Bearer+空格+登录返回reponse数据中token中的字符串

### 参与贡献

Fork 本仓库
新建 Feat_xxx 分支
提交代码
新建 Pull Request

### 项目实例

前台：
![前台](https://github.com/shisanxiaobai/MyLottery/blob/main/image/Snipaste_2022-11-11_21-42-11.png)

后台：
![后台](https://github.com/shisanxiaobai/MyLottery/blob/main/image/Snipaste_2022-11-11_21-42-39.png)
