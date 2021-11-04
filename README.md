# qianchuang SDK for Go

巨量引擎开放平台-千川SDK

## 安装
```go
import "github.com/CriarBrand/qianchuanSDK"
```
## 使用
**初始化**
```go
credentials := qianchuanSDK.NewCredentials("CLIENT_KEY", "CLIENT_SECRET")
manager := qianchuanSDK.NewManager(credentials, nil)
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

**获取账户下计划列表（不含创意）** `/ad/get/`
```go
manager.AdListGet(AdListGetReq{
    AdvertiserId: "ADVERTISER_ID",
    AccessToken:  "ACCESS_TOKEN",
    Page:         1,
    PageSize:     20,
    Filtering: AdListGetFiltering{
        MarketingGoal: "LIVE_PROM_GOODS",
    },
})
```

**获取计划详情（含创意信息）** `/ad/detail/get/`
```go
manager.AdDetailGet(AdDetailGetReq{
    AdvertiserId: "ADVERTISER_ID",
    AccessToken:  "ACCESS_TOKEN",
    AdId:         AD_ID,
})
```

**获取计划审核建议** `/ad/reject_reason/`
```go
manager.AdRejectReason(AdRejectReasonReq{
    AdvertiserId: "ADVERTISER_ID",
    AccessToken:  "ACCESS_TOKEN",
    AdIds:        []int64{AD_ID1, AD_ID2},
})
```

**获取账户下创意列表** `/creative/get/`
```go
manager.CreativeGet(CreativeGetReq{
    AccessToken:  string,
    AdvertiserId: string,
    Filtering: CreativeGetReqFiltering{
        MarketingGoal:         string,
        CreativeCreateEndDate: string,    
        Status:                string, 
        CreativeModifyTime:    string,
        AdIds:                 []int64,
        CreativeId:            int64,
        CreativeMaterialMode:  string,
        CampaignId:            int64,
        CreativeCreateEndDate: string,
    },
    Page:     1,
    PageSize: 20,
})
```

**获取创意审核建议** `/creative/reject_reason/`
```go
manager.CreativeRejectReason(CreativeRejectReasonReq{
    AccessToken:  string,
    AdvertiserId: int64,
    CreativeIds:  []int64,
})
```

**获取千川账户下可投商品列表接口** `/product/available/get/`
```go
manager.ProductAvailableGet(ProductAvailableGetReq{
    AccessToken:  string,
    AdvertiserId: int64,
    Page:         1,
    PageSize:     20,
})
```

**获取千川账户下已授权抖音号** `/aweme/authorized/get/`
```go
manager.AwemeAuthorizedGet(AwemeAuthorizedGetReq{
    AccessToken:  string,
    AdvertiserId: int64,
    Page:         1,
    PageSize:     10,
})
```

**获取素材库的图片** `/file/image/get/`
```go
manager.FileImageGet(FileImageGetReq{
    AccessToken:  string,
    AdvertiserId: int64,
    Filtering: FileImageGetReqFiltering{
        ImageIds:    []string,
        MaterialIds: []int64,
        Signatures:  []string,
        Width:       int64,
        Height:      int64,
        Ratio:       []float64,
        StartTime:   string,
        EndTime:     string,
    },
    Page:     1,
    PageSize: 10,
})
```

**获取素材库的视频** `/file/video/get/`
```go
manager.FileVideoGet(FileVideoGetReq{
    AccessToken:  string,
    AdvertiserId: int64,
    Filtering: FileVideoGetReqFiltering{
        Width:       int64,
        Height:      int64,
        Ratio:       []float64,
        VideoIds:    []string,
        MaterialIds: []int64,
        Signatures:  []string,
        StartTime:   string,
        EndTime:     string,
    },
    Page:     1,
    PageSize: 10,
})
```

**获取抖音号下的视频** `/file/video/aweme/get/`
```go
manager.FileVideoAwemeGet(FileVideoAwemeGetReq{
    AccessToken:  string,
    AdvertiserId: int64,
    AwemeId:      int64,
    Filtering: FileVideoAwemeGetReqFiltering{
        ProductId: int64,
    },
    Count: 30,
})
```

**获取行业列表** `/tools/industry/get/`
```go
manager.ToolsIndustryGet(ToolsIndustryGetReq{
    AccessToken: string,
    Level:       int64,
    Type:        string,
})
```

**查询抖音类目下的推荐达人** `/tools/aweme_category_top_author/get/`
```go
manager.ToolsAwemeCategoryTopAuthorGet(ToolsAwemeCategoryTopAuthorGetReq{
    AccessToken:  string,
    AdvertiserId: int64,
    CategoryId:   int64,
    Behaviors:    []string,
})
```

**查询抖音类目列表** `/tools/aweme_multi_level_category/get/`
```go
manager.ToolsAwemeMultiLevelCategoryGet(ToolsAwemeMultiLevelCategoryGetReq{
    AccessToken:  string,
    AdvertiserId: int64,
    Behaviors:    []string{},
})
```

**行为类目查询** `/tools/interest_action/action/category/`
```go
manager.ToolsInterestActionActionCategory(ToolsInterestActionActionCategoryReq{
    AccessToken:  string,
    AdvertiserId: int64,
    ActionScene:  []string,
    ActionDays:   int64,
})
```

**行为关键词查询** `/tools/interest_action/action/keyword/`
```go
manager.ToolsInterestActionActionKeyword(ToolsInterestActionActionKeywordReq{
    AccessToken:  string,
    AdvertiserId: int64,
    QueryWords:   string,
    ActionScene:  []string,
    ActionDays:   int64,
})
```

**兴趣类目查询** `/tools/interest_action/interest/category/`
```go
manager.ToolsInterestActionInterestCategory(ToolsInterestActionInterestCategoryReq{
    AccessToken:  string,
    AdvertiserId: int64,
})
```

**兴趣关键词查询** `/tools/interest_action/interest/keyword/`
```go
manager.ToolsInterestActionInterestKeyword(ToolsInterestActionInterestKeywordReq{
    AccessToken:  string,
    AdvertiserId: int64,
    QueryWords:   string,
})
```

**查询动态创意词包** `/tools/creative_word/select/`
```go
manager.ToolsCreativeWordSelect(ToolsCreativeWordSelectReq{
    AccessToken:     string,
    AdvertiserId:    int64,
    CreativeWordIds: []string,
})
```

**查询人群包列表** `/dmp/audiences/get/`
```go
manager.DmpAudiencesGet(DmpAudiencesGetReq{
    AccessToken:         string,
    AdvertiserId:        int64,
    RetargetingTagsType: int64,
    Offset:              int64,
    Limit:               int64,
})
```