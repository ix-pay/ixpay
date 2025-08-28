# ixpay

让你爱上支付


-----------------------------------------------------------


## 1.初始化Go模块‌

在项目根目录执行（替换your-module-name为实际模块名）：

```
go mod init github.com/ix-pay/ixpay
```
## 2.设置国内代理加速下载‌

配置GOPROXY解决网络问题
```
go env -w GOPROXY=https://goproxy.cn,direct
```

## 3.‌安装依赖包‌

执行以下命令安装缺失的GoPay组件
```
go get github.com/go-pay/gopay/alipay

go get github.com/go-pay/gopay/wechat/v3
```

## 4.‌‌验证依赖‌

检查go.mod文件是否自动生成并包含以下内容：
```
require (
   github.com/go-pay/gopay v1.5.9 // 版本可能不同
)
```

查包版本：
```
go list -m -versions github.com/go-pay/gopay
```
验证依赖图：
```
go mod graph | grep gopay
```

## 5.‌清理并重新运行‌

执行以下命令确保依赖完整：
```
go clean -modcache

go mod tidy

go run main.go
```
## 6.安装支持swagger
```
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

### 核心代码实现
// main.go
```
// @title API文档
// @version 1.0
// @description ixpay项目Swagger集成
// @host localhost:8989
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
   r := gin.Default()
   // 开发环境启用文档
   if gin.Mode() != gin.ReleaseMode {
       url := ginSwagger.URL("/swagger/doc.json")

       r.GET("/swagger/\*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
   }

   // 注册路由
   routes.SetupRouter(r)
   r.Run()
}
```

### 接口注释规范
```
// GetProfile
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 基础功能
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param userId path string true "用户ID"
// @Success 200 {object} utils.RespData{data=models.ProfileUser} "成功响应"
// @Failure 400 {object} utils.RespData{data=string} "失败消息"
// @Router /auth/profile/{userId} [get]
```

### 请求token
```
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTY2NTMxMTUsInVzZXJfaWQiOjYxNjI2OTY0MzA2MDg3NTI2NH0.RlVOKvJaGRtabc5FzDVCFnODKXGbvXNmARB9zSyClXA
```

### 文档生成命令
安装swag
```
// go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/swaggo/swag/cmd/swag@v1.8.12
```
```
swag init -g main.go --output docs --parseDependency --parseInternal
```

<http://127.0.0.1:8989/swagger/index.html>

## 热加载
```
go install github.com/air-verse/air@latest
```
// .vscode/launch.json
```
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Debug ixpay",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}/main.go",
      "args": [],
      "cwd": "${workspaceFolder}",
      "env": {},
      "showLog": true,
      "debugAdapter": "dlv-dap",
    }
  ]
}
```