# qianchuang SDK for Go

巨量引擎开放平台-千川SDK

## 安装
```go
import "github.com/CriarBrand/qianchuangSDK"
```
## 使用
**初始化**
```go
credentials := qianchuangSDK.NewCredentials("CLIENT_KEY", "CLIENT_SECRET")
manager := qianchuangSDK.NewManager(credentials, nil)
```

**生成授权链接,获取授权码** `/platform/oauth/connect/`
```go
oauthUrl := manager.OauthConnect(douyinGo.OauthParam{
    Scope: "user_info,mobile_alert,video.list,video.data,video.create,video.delete,data.external.user,data.external.item,aweme.share,fans.list,following.list,item.comment,star_top_score_display,fans.data,data.external.fans_source,data.external.fans_favourite,discovery.ent,video.search,video.search.comment,fans.check",
    RedirectUri: "REDIRECT_URI",
})
```

**获取AccessToken** `/oauth/access_token/`
```go
accessToken, err := manager.OauthAccessToken(douyinGo.OauthAccessTokenReq{
    Code: "CODE",
})
```

**刷新access_token** `/oauth/refresh_token/`
```go
manager.OauthRenewRefreshToken(douyinGo.OauthRenewRefreshTokenReq{
    RefreshToken: "REFRESH_TOKEN",
})
```

**刷新refresh_token** `/oauth/renew_refresh_token/`
```go
manager.OauthRenewRefreshToken(douyinGo.OauthRenewRefreshTokenReq{
    RefreshToken: "REFRESH_TOKEN",
})
```

**生成client_token** `/oauth/client_token/`
```go
clientToken, err := manager.OauthClientAccessToken()
```

**获取用户信息** `/oauth/userinfo/`
```go
userInfo, err := manager.OauthUserinfo(douyinGo.OauthUserinfoReq{
    OpenId:      "OPEN_ID",
    AccessToken: "ACCESS_TOKEN",
})

// 解析手机号
mobile, err := manager.DecryptMobile("ENCRYPT_MOBILE")
```

