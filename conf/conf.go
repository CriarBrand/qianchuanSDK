package conf

const (

	// API_HOST OpenAPI HOST
	API_HOST = "ad.oceanengine.com"

	// API_HTTP_SCHEME 协议
	API_HTTP_SCHEME = "https://"

	// API_OAUTH_CONNECT 生成授权链接
	API_OAUTH_CONNECT = "/openapi/qc/audit/oauth.html"

	// API_OAUTH_ACCESS_TOKEN 获取access_token
	API_OAUTH_ACCESS_TOKEN = "/open_api/oauth2/access_token/"

	// API_OAUTH_REFRESH_TOKEN 刷新access_token
	API_OAUTH_REFRESH_TOKEN = "/open_api/oauth2/refresh_token/"

	// API_ADVERTISER_LIST 获取已授权的账户（店铺/代理商）
	API_ADVERTISER_LIST = "/open_api/oauth2/advertiser/get/"

	// API_SHOP_ADVERTISER_LIST 获取店铺账户关联的广告账户列表
	API_SHOP_ADVERTISER_LIST = "/open_api/v1.0/qianchuan/shop/advertiser/list/"

	// API_AGENT_ADVERTISER_LIST 获取代理商账户关联的广告账户列表
	API_AGENT_ADVERTISER_LIST = "/open_api/2/agent/advertiser/select/"

	// API_USER_INFO 获取授权时登录用户信息
	API_USER_INFO = "/open_api/2/user/info/"

	// API_SHOP_ACCOUNT_INFO 获取店铺账户信息
	API_SHOP_ACCOUNT_INFO = "/open_api/v1.0/qianchuan/shop/get/"

	// API_AGENT_INFO 获取代理商账户信息
	API_AGENT_INFO = "/open_api/2/agent/info/"

	// API_ADVERTISER_PUBLIC_INFO 获取代理商账户信息
	API_ADVERTISER_PUBLIC_INFO = "/open_api/2/advertiser/public_info/"

	// API_ADVERTISER_INFO 获取千川广告账户全量信息
	API_ADVERTISER_INFO = "/open_api/2/advertiser/info/"

	// API_ADVERTISER_REPORT 获取广告账户数据
	API_ADVERTISER_REPORT = "/open_api/v1.0/qianchuan/report/advertiser/get/"

	// API_AD_REPORT 获取广告计划数据
	API_AD_REPORT = "/open_api/v1.0/qianchuan/report/ad/get/"

	// API_creative_REPORT 获取广告创意数据
	API_creative_REPORT = "/open_api/v1.0/qianchuan/report/creative/get/"

	// API_CAMPAIGN_CREATE 广告组创建
	API_CAMPAIGN_CREATE = "/open_api/v1.0/qianchuan/campaign/create/"

	// API_CAMPAIGN_UPDATE 广告组更新
	API_CAMPAIGN_UPDATE = "/open_api/v1.0/qianchuan/campaign/update/"

	// API_BATCH_CAMPAIGN_STATUS_UPDATE 广告组状态更新
	API_BATCH_CAMPAIGN_STATUS_UPDATE = "/open_api/v1.0/qianchuan/batch_campaign_status/update/"

	// API_CAMPAIGN_LIST_GET 广告组列表获取
	API_CAMPAIGN_LIST_GET = "/open_api/v1.0/qianchuan/campaign_list/get/"

	// API_AD_CREATE 创建计划（含创意生成规则）
	API_AD_CREATE = "/open_api/v1.0/qianchuan/ad/create/"

	// API_AD_UPDATE 更新计划（含创意生成规则）
	API_AD_UPDATE = "/open_api/v1.0/qianchuan/ad/update/"

	// API_AD_STATUS_UPDATE 更新计划状态
	API_AD_STATUS_UPDATE = "/open_api/v1.0/qianchuan/ad/status/update/"

	// API_AD_BUDGET_UPDATE 更新计划预算
	API_AD_BUDGET_UPDATE = "/open_api/v1.0/qianchuan/ad/budget/update/"

	// API_AD_BID_UPDATE 更新计划出价
	API_AD_BID_UPDATE = "/open_api/v1.0/qianchuan/ad/bid/update/"

	// API_AD_DETAIL_GET 获取计划详情（含创意信息）
	API_AD_DETAIL_GET = "/open_api/v1.0/qianchuan/ad/detail/get/"

	// API_AD_LIST_GET 获取账户下计划列表（不含创意）
	API_AD_LIST_GET = "/open_api/v1.0/qianchuan/ad/get/"

	// API_AD_REJECT_REASON 获取计划审核建议
	API_AD_REJECT_REASON = "/open_api/v1.0/qianchuan/ad/reject_reason/"
)
