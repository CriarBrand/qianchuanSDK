# qianchuanSDK

[![Go](https://img.shields.io/badge/Go-1.16%2B-00ADD8?logo=go)](https://go.dev/)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENSE)
[![GitHub Stars](https://img.shields.io/github/stars/CriarBrand/qianchuanSDK?style=flat)](https://github.com/CriarBrand/qianchuanSDK/stargazers)
[![GitHub Forks](https://img.shields.io/github/forks/CriarBrand/qianchuanSDK?style=flat)](https://github.com/CriarBrand/qianchuanSDK/network/members)
[![Maintenance](https://img.shields.io/badge/status-maintenance%20restarting-orange)](https://github.com/CriarBrand/qianchuanSDK)

巨量引擎千川开放平台的非官方 Go SDK。

Unofficial Go SDK for the Ocean Engine Qianchuan Open Platform.

> **项目状态：恢复维护中**
>
> 本项目最初完成于数年前，现有代码可能无法完全兼容千川开放平台的最新接口。
> 维护者已在 2026 年重新回到千川一线业务，并开始基于真实生产需求恢复和升级本项目。
>
> 在新版适配完成前，请在生产环境使用前自行核对官方接口文档。

## 项目简介

`qianchuanSDK` 封装了巨量引擎千川开放平台常用接口，帮助 Go 开发者减少在身份授权、请求构造、参数序列化、响应解析和错误处理上的重复工作。

项目目前包含或曾实现以下能力：

* OAuth 授权及 Access Token 管理
* 广告账户与授权账户查询
* 广告计划创建、查询及管理
* 广告创意查询及管理
* 图片、视频等素材查询
* 广告账户及计划报表查询
* 商品和抖音号授权信息查询
* 行业、兴趣、行为和关键词工具
* DMP 人群包相关接口

本项目不是巨量引擎或字节跳动官方维护的 SDK。

## 为什么重新维护这个项目

这个项目并不是为了展示而临时建立的仓库。

几年前，我在开发千川相关工具时，发现 Go 生态中几乎没有可以直接使用的千川 SDK，因此开始将实际开发中使用的接口整理、封装并开源，希望其他开发者不必再从签名、请求和响应结构开始重复造轮子。

后来由于工作方向变化，项目暂停了维护。

现在，我重新回到抖音直播和千川投放的一线业务，正在搭建一个面向真实生产环境的多广告账户素材数据分析系统，用于：

* 自动同步多个千川广告账户的数据
* 汇总短视频素材的消耗、成交和 ROI 表现
* 分析素材测试、放量、衰退及生命周期
* 识别潜力素材、爆款素材和低效素材
* 减少运营团队每天手工下载和整理报表的工作
* 为投放复盘和预算决策提供可追溯的数据依据

`qianchuanSDK` 将重新成为这套系统的底层基础设施。

这一次恢复维护，不只是为了让旧代码重新运行，也是希望将真实业务中遇到的问题沉淀为一个更可靠、更清晰、更容易参与的中文开源项目。

## 当前状态

现有版本主要基于早期千川开放平台接口开发。

目前正在进行以下工作：

* [ ] 核对现有接口与最新官方文档的差异
* [ ] 更新 OAuth 授权和 Token 管理逻辑
* [ ] 适配新版广告账户、计划、创意和素材接口
* [ ] 补充账户、计划、素材和直播相关报表接口
* [ ] 统一请求参数和响应错误处理
* [ ] 补充单元测试和 Mock 测试
* [ ] 增加 GitHub Actions 持续集成
* [ ] 提供可直接运行的示例项目
* [ ] 完善接口文档和迁移说明
* [ ] 建立版本兼容性说明

在这些工作完成前，建议把当前代码视为：

1. 千川 Go SDK 的历史实现；
2. 新版本 SDK 的迁移基础；
3. 部分接口的数据结构参考。

## 安装

```bash
go get github.com/CriarBrand/qianchuanSDK
```

在代码中导入：

```go
import qianchuan "github.com/CriarBrand/qianchuanSDK"
```

## 快速开始

### 初始化客户端

```go
package main

import (
	"fmt"

	qianchuan "github.com/CriarBrand/qianchuanSDK"
)

func main() {
	const appID int64 = 123456789
	const appSecret = "YOUR_APP_SECRET"

	credentials := qianchuan.NewCredentials(appID, appSecret)
	manager := qianchuan.NewManager(credentials, nil)

	fmt.Printf("manager initialized: %+v\n", manager)
}
```

请勿将真实的 `app_secret`、`access_token` 或其他敏感凭证直接提交到代码仓库。

建议通过环境变量读取：

```go
package main

import (
	"log"
	"os"
	"strconv"

	qianchuan "github.com/CriarBrand/qianchuanSDK"
)

func main() {
	appID, err := strconv.ParseInt(os.Getenv("QIANCHUAN_APP_ID"), 10, 64)
	if err != nil {
		log.Fatalf("invalid QIANCHUAN_APP_ID: %v", err)
	}

	appSecret := os.Getenv("QIANCHUAN_APP_SECRET")
	if appSecret == "" {
		log.Fatal("QIANCHUAN_APP_SECRET is required")
	}

	credentials := qianchuan.NewCredentials(appID, appSecret)
	manager := qianchuan.NewManager(credentials, nil)

	_ = manager
}
```

### 生成授权链接

```go
authURL := manager.OauthConnect(qianchuan.OauthParam{
	AppId:        appID,
	State:        "your_custom_state",
	Scope:        []int64{20120000},
	MaterialAuth: "1",
	RedirectUri:  "https://example.com/oauth/callback",
})

fmt.Println(authURL)
```

实际使用的权限范围应以当前千川开放平台后台及官方文档为准。

### 使用授权码获取 Access Token

```go
tokenResponse, err := manager.OauthAccessToken(
	qianchuan.OauthAccessTokenReq{
		AuthCode: "AUTH_CODE",
	},
)
if err != nil {
	log.Fatalf("get access token failed: %v", err)
}

fmt.Println(tokenResponse.Data.AccessToken)
```

### 刷新 Access Token

```go
refreshResponse, err := manager.OauthRefreshToken(
	qianchuan.OauthRefreshTokenReq{
		RefreshToken: "REFRESH_TOKEN",
	},
)
if err != nil {
	log.Fatalf("refresh access token failed: %v", err)
}

fmt.Println(refreshResponse.Data.AccessToken)
```

## 代码结构

仓库中的主要模块包括：

| 文件或目录            | 说明                 |
| ---------------- | ------------------ |
| `oauth.go`       | OAuth 授权及 Token 管理 |
| `advertiser.go`  | 广告账户相关接口           |
| `ad.go`          | 广告相关基础接口           |
| `ad_campaign.go` | 广告计划相关接口           |
| `ad_creative.go` | 广告创意相关接口           |
| `ad_report.go`   | 广告报表相关接口           |
| `file.go`        | 图片、视频等素材接口         |
| `tools.go`       | 行业、兴趣、行为等工具接口      |
| `manager.go`     | SDK Manager 初始化与管理 |
| `client/`        | HTTP 客户端及请求处理      |
| `auth/`          | 应用凭证相关逻辑           |
| `conf/`          | API 地址及配置          |

## 2026 维护路线图

### 第一阶段：恢复可用性

* 梳理当前代码中的全部接口
* 标记已经废弃或发生变更的接口
* 更新基础请求地址、参数和响应结构
* 修复 README 中已经失效的调用示例
* 建立基础测试和 CI

### 第二阶段：生产化

* 完善超时、重试及限流处理
* 增加结构化错误类型
* 增加分页请求辅助能力
* 增加日志和可观测性支持
* 补充多广告账户使用示例
* 增加报表同步及数据落库示例

### 第三阶段：生态建设

* 发布稳定的语义化版本
* 建立接口兼容性矩阵
* 完善贡献指南
* 整理高频使用场景
* 持续吸收真实业务中的公共需求

## 贡献

欢迎提交 Issue 和 Pull Request。

目前尤其需要以下类型的贡献：

* 提供已经发生变化的接口信息
* 补充新版千川 API 的请求和响应结构
* 修复字段类型或序列化问题
* 补充单元测试
* 改进错误处理
* 完善文档和示例
* 提供不同业务场景下的兼容性反馈

提交接口适配类 Pull Request 时，建议说明：

1. 对应的接口名称；
2. 修改原因；
3. 是否存在不兼容变更；
4. 使用的文档版本或更新时间；
5. 是否已经在测试或生产环境验证。

请勿在 Issue、日志、测试数据或代码中提交：

* App Secret
* Access Token
* Refresh Token
* 广告主敏感信息
* 用户个人信息
* 未脱敏的真实业务数据

## Issue 建议格式

提交 Bug 时，请尽量提供：

```text
Go 版本：
SDK 版本或 Commit：
接口名称：
请求参数（请脱敏）：
返回错误（请脱敏）：
预期行为：
实际行为：
```

提交接口更新建议时，请尽量提供：

```text
接口名称：
接口用途：
现有实现的问题：
需要新增或修改的字段：
是否包含不兼容变更：
```

## 版本策略

项目恢复维护后将采用语义化版本：

```text
MAJOR.MINOR.PATCH
```

* `MAJOR`：包含不兼容的接口调整
* `MINOR`：新增向后兼容的功能或接口
* `PATCH`：向后兼容的问题修复

在首个恢复维护版本发布之前，部分 API 仍可能发生调整。

## 安全说明

* 不要在代码中硬编码应用密钥和 Token
* 不要将真实广告账户数据提交到公开仓库
* 建议为 Token 设置安全的持久化和更新机制
* 调用生产接口前应设置合理的超时
* 请根据平台要求处理限流、授权失效和权限变更
* SDK 使用者应自行确保业务符合平台规则及适用法律

如果发现可能影响用户凭证或数据安全的问题，请不要直接在公开 Issue 中披露完整细节。

## 免责声明

本项目为社区维护的非官方 SDK，与巨量引擎、Ocean Engine、抖音、字节跳动及其关联公司不存在官方隶属或授权关系。

平台接口、字段、权限及规则可能随时发生变化。使用者应以官方平台及官方文档为准，并自行承担使用本项目产生的风险。

## License

本项目基于 [GNU General Public License v3.0](LICENSE) 开源。

---

这个项目曾经因为维护者工作方向的变化暂停更新，但它没有被放弃。

随着真实千川业务重新产生需求，项目也将重新开始迭代。希望它不仅能够解决维护者自己的问题，也能继续为使用 Go 开发千川相关工具的开发者节省一些时间。
