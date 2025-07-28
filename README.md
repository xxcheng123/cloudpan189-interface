# CloudPan189 Interface

一个用于天翼云盘189的Go语言接口库，提供了常用的API封装，支持文件操作、分享管理、用户信息获取等功能。

> **注意**: 本项目目前只实现了天翼云盘的部分常用接口，并非完整的API覆盖。具体实现的接口请参考 [Client接口定义](client/client.go)。

## 功能特性

- 🔐 **扫描登录**: 使用二维码登录，免输入账号密码
- 📁 **文件操作**: 获取文件信息、下载链接、文件夹信息等
- 🎬 **媒体播放**: 支持获取VLC视频播放URL
- 🔗 **分享管理**: 获取分享信息、列出分享目录等
- 👤 **用户管理**: 获取用户信息和权限
- 🛠 **调试模式**: 支持调试模式便于开发

## 接口覆盖说明

本项目目前实现的接口有限，主要包括：

- 文件和文件夹的基本信息获取
- 文件下载链接获取
- 分享相关操作
- 用户信息和权限查询
- 视频播放URL获取

如需查看完整的接口列表和方法签名，请参考 [client/client.go](client/client.go) 中的 `Client` 接口定义。

如果您需要其他未实现的接口，欢迎提交 Issue 或 Pull Request。

## 安装

```bash
go get github.com/xxcheng123/cloudpan189-interface
```

## 快速开始

### 1. 创建客户端

```go
package main

import (
    "context"
    "fmt"
    "github.com/xxcheng123/cloudpan189-interface/client"
)

func main() {
    // 创建客户端
    c := client.New()
    
    // 启用调试模式（可选）
    c = c.WithDebug(true)
    
    // 设置认证Token
    token := client.NewAuthToken("your_access_token", 1756359689000)
    c = c.WithToken(token)
}
```

### 2. 二维码登录

```go
import (
    "fmt"
    "time"
    "github.com/skip2/go-qrcode"
)

// 初始化登录，获取UUID
uuidInfo, err := client.LoginInit()
if err != nil {
    panic(err)
}

fmt.Printf("请扫描二维码登录，UUID: %s\n", uuidInfo.UUID)

// 生成二维码（二维码内容就是UUID）
qrCode, err := qrcode.New(uuidInfo.UUID, qrcode.Low)
if err != nil {
    panic(err)
}

// 打印二维码到控制台
fmt.Println(qrCode.ToString(false))

// 轮询登录状态
var success = false
for count := 0; count < 30; count++ { // 最多尝试30次
    tokenResp, err := client.LoginQuery(uuidInfo.UUID)
    if err != nil {
        fmt.Printf("等待扫码登录... (%d/30)\n", count+1)
        time.Sleep(5 * time.Second)
        continue
    }
    
    fmt.Printf("登录成功！AccessToken: %s\n", tokenResp.AccessToken)
    fmt.Printf("过期时间: %d\n", tokenResp.ExpiresIn)
    
    // 创建带Token的客户端
    token := client.NewAuthToken(tokenResp.AccessToken, tokenResp.ExpiresIn)
    c = c.WithToken(token)
    
    success = true
    break
}

if !success {
    panic("登录超时，请重试")
}
```

### 3. 文件操作

```go
ctx := context.Background()

// 获取文件信息
fileInfo, err := c.GetFileInfo(ctx, "file_id")
if err != nil {
    panic(err)
}
fmt.Printf("文件名: %s, 大小: %d\n", fileInfo.FileName, fileInfo.FileSize)

// 获取文件下载链接
downloadResp, err := c.GetFileDownload(ctx, "file_id")
if err != nil {
    panic(err)
}
fmt.Printf("下载链接: %s\n", downloadResp.FileDownloadUrl)

// 获取文件夹信息
folderInfo, err := c.GetFolderInfo(ctx, "folder_id")
if err != nil {
    panic(err)
}
fmt.Printf("文件夹: %s, 文件数量: %d\n", folderInfo.Name, folderInfo.FileCount)
```

### 4. 分享操作

```go
// 获取分享信息
shareInfo, err := c.GetShareInfo(ctx, "share_code")
if err != nil {
    panic(err)
}
fmt.Printf("分享文件: %s\n", shareInfo.FileName)

// 列出分享目录
shareDir, err := c.ListShareDir(ctx, shareInfo.ShareId, "folder_id")
if err != nil {
    panic(err)
}
fmt.Printf("目录文件数量: %d\n", len(shareDir.FileListAO.FileList))
```

### 5. 用户信息

```go
// 获取用户信息
userInfo, err := c.GetUserInfo(ctx)
if err != nil {
    panic(err)
}
fmt.Printf("用户: %s\n", userInfo.LoginName)

// 获取用户权限
privileges, err := c.GetUserPrivileges(ctx)
if err != nil {
    panic(err)
}
fmt.Printf("用户权限: %+v\n", privileges)
```

## 配置选项

### 调试模式

```go
// 启用调试模式
client := client.New().WithDebug(true)

// 或者
client := client.New().WithDebug() // 默认启用
```

### Token认证

```go
// 创建认证Token
token := client.NewAuthToken("access_token", expireTime)
client := client.New().WithToken(token)
```

### 请求选项

大部分API方法都支持可选参数，例如：

```go
// 获取文件下载链接时指定分享ID
downloadResp, err := c.GetFileDownload(ctx, "file_id", func(req *client.GetFileDownloadRequest) {
    req.ShareId = 12322116912755
    req.Short = true // 获取短地址
})

// 获取分享信息时指定访问码
shareInfo, err := c.GetShareInfo(ctx, "share_code", func(req *client.GetShareInfoRequest) {
    req.AccessCode = "access_code"
})
```

## 错误处理

库提供了统一的错误处理机制：

```go
resp, err := client.GetFileInfo(ctx, "file_id")
if err != nil {
    // 处理网络错误或其他系统错误
    fmt.Printf("请求失败: %v\n", err)
    return
}

// API返回的错误会在响应结构中
if resp.ResCode != 0 {
    fmt.Printf("API错误: %s\n", resp.ResMessage)
}
```

## 许可证

本项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 免责声明

本项目仅供学习和研究使用，请遵守天翼云盘的服务条款。使用本库产生的任何问题，作者不承担责任。