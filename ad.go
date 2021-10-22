// 广告计划相关API
// 巨量开放平台地址：https://open.oceanengine.com/doc/index.html?key=qianchuan&type=api&id=1697466251535372

package qianchuanSDK

import (
	"context"
	"github.com/CriarBrand/qianchuangSDK/conf"
	"net/http"
)

// AdCreateReq 获取广告账户数据-请求
type AdCreateReq struct {
	AccessToken string       // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        AdCreateBody // POST请求的data
}

type AdCreateBody struct {
	AdvertiserId    int64                   `json:"advertiser_id"`           // 千川广告主账户id
	MarketingGoal   string                  `json:"marketing_goal"`          // 营销目标，允许值：VIDEO_PROM_GOODS 短视频带货、LIVE_PROM_GOODS 直播带货
	PromotionWay    string                  `json:"promotion_way,omitempty"` // 推广方式 ，目前仅支持专业版，不支持极速版，允许值：STANDARD（默认）
	Name            string                  `json:"name"`                    // 计划名称，长度为1-100个字符，其中1个汉字算2位字符。名称不可重复，否则会报错
	CampaignId      int64                   `json:"campaign_id"`             // 千川广告组id
	AwemeId         int64                   `json:"aweme_id"`
	ProductIds      []int64                 `json:"product_ids"`
	DeliverySetting AdCreateDeliverySetting `json:"delivery_setting"`
	Audience        AdCreateAudience        `json:"audience"`
	AdCreateCreative
}

type AdCreateDeliverySetting struct {
	SmartBidType       string  `json:"smart_bid_type"`                 // 投放场景（出价方式），详见【附录-自动出价类型】，允许值：SMART_BID_CUSTOM控成本投放、SMART_BID_CONSERVATIVE 放量投放控成本投放：控制成本，尽量消耗完预算放量投放：接受成本上浮，尽量消耗更多预算
	FlowControlMode    string  `json:"flow_control_mode,omitempty"`    // 投放速度，详见【附录-计划投放速度类型】仅当 smart_bid_type 为SMART_BID_CUSTOM 时需传值，允许值：FLOW_CONTROL_MODE_FAST 尽快投放（默认值）、FLOW_CONTROL_MODE_BALANCE 均匀投放、FLOW_CONTROL_MODE_SMOOTH 优先低成本，对应千川后台「严格控制成本上限」勾选项
	ExternalAction     string  `json:"external_action"`                // 转化目标短视频带货目的允许值：AD_CONVERT_TYPE_SHOPPING 商品购买、AD_CONVERT_TYPE_QC_FOLLOW_ACTION 粉丝提升、AD_CONVERT_TYPE_QC_MUST_BUY 点赞评论直播带货目的允许值：AD_CONVERT_TYPE_LIVE_ENTER_ACTION 进入直播间、AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION 直播间商品点击、AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION 直播间下单、AD_CONVERT_TYPE_NEW_FOLLOW_ACTION 直播间粉丝提升、AD_CONVERT_TYPE_LIVE_COMMENT_ACTION 直播间评论、AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY直播间成交
	DeepExternalAction string  `json:"deep_external_action,omitempty"` // 深度转化目标，对应千川后台「期待同时优化」注意：1. 仅直播带货场景支持2. 当 smart_bid_type 为SMART_BID_CUSTOM 且 flow_control_mode 为 FLOW_CONTROL_MODE_SMOOTH 亦不支持深度转化目标允许值：AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION 直播间下单若不传，则不生效；若传入，则仅当转化目标为AD_CONVERT_TYPE_LIVE_ENTER_ACTION、AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION 时生效
	Budget             float64 `json:"budget"`                         // 预算，最多支持两位小数当预算模式为日预算时，预算范围是300 - 9999999.99；当预算模式为总预算时，预算范围是max(300,投放天数x100) - 9999999.99
	BudgetMode         string  `json:"budget_mode"`                    // 预算类型（创建后不可修改），详见【附录-预算类型】，允许值：BUDGET_MODE_DAY 日预算，BUDGET_MODE_TOTAL 总预算
	CpaBid             float64 `json:"cpa_bid,omitempty"`              // 转化出价，出价不能大于预算仅当 smart_bid_type 为SMART_BID_CUSTOM 时需传值
	VideoScheduleType  string  `json:"video_schedule_type,omitempty"`  // 短视频投放日期选择方式，仅短视频带货场景需入参，允许值：SCHEDULE_FROM_NOW 从今天起长期投放（总预算模式下不支持）、SCHEDULE_START_END 设置开始和结束日期
	LiveScheduleType   string  `json:"live_schedule_type,omitempty"`   // 直播间投放时段选择方式，仅直播带货场景需入参，允许值：SCHEDULE_TIME_ALLDAY 全天、SCHEDULE_TIME_WEEKLY_SETTING 指定时间段、SCHEDULE_TIME_FIXEDRANGE 固定时长
	StartTime          string  `json:"start_time,omitempty"`           // 投放起始时间，形式如：2017-01-01广告投放起始时间不允许修改。当video_schedule_type为SCHEDULE_START_END 设置开始和结束日期时需传入。当live_schedule_type 为SCHEDULE_TIME_ALLDAY 全天、SCHEDULE_TIME_WEEKLY_SETTING 指定时间段时必填；当 live_schedule_type 为SCHEDULE_TIME_FIXEDRANGE固定时长时不能传入
	EndTime            string  `json:"end_time,omitempty"`             // 投放结束时间，形式如：2017-01-01结束时间不能比起始时间早。当video_schedule_type为SCHEDULE_START_END 设置开始和结束日期时需传入。当live_schedule_type 为SCHEDULE_TIME_ALLDAY 全天、SCHEDULE_TIME_WEEKLY_SETTING 指定时间段时必填；当 live_schedule_type 为SCHEDULE_TIME_FIXEDRANGE固定时长时不能传入
	ScheduleTime       string  `json:"schedule_time,omitempty"`        // 投放时段，当 live_schedule_type 为SCHEDULE_TIME_WEEKLY_SETTING 时生效默认全时段投放，格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。例如：填写"000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000"，则投放时段为周一到周日的11:30~13:30
	ScheduleFixedRange int64   `json:"schedule_fixed_range,omitempty"` // 固定投放时长当 live_schedule_type 为 SCHEDULE_TIME_FIXEDRANGE 时必填；当live_schedule_type 为SCHEDULE_TIME_ALLDAY 全天、SCHEDULE_TIME_WEEKLY_SETTING 指定时间段时不能传入。单位为秒，最小值为1800（0.5小时），最大值为48*1800（24小时），值必须为1800倍数，不然会报错
}

type AdCreateAudience struct {
	District               string   `json:"district,omitempty"`                 // 地域定向类型，配合 city 字段使用，允许值：CITY 省市， COUNTY 区县， NONE 不限默认值为NONE
	City                   []int64  `json:"city,omitempty"`                     // 具体定向的城市列表，当 district 为COUNTY，CITY为必填，枚举值详见【附件-city.json】省市的传法："city" : [12], "district" : "CITY"区县的传法："city" : [130102], "district" : "COUNTY"
	LocationType           string   `json:"location_type,omitempty"`            // 地域定向的用户状态类型，当 district 为COUNTY，CITY为必填，允许值：CURRENT 正在该地区的用户、HOME 居住在该地区的用户、TRAVEL 到该地区旅行的用户、ALL 该地区内的所有用户
	Gender                 string   `json:"gender,omitempty"`                   // 性别，允许值：GENDER_FEMALE 女性， GENDER_MALE 男性，NONE 不限
	Age                    []string `json:"age,omitempty"`                      // 年龄，详见【附录-受众年龄区间】，允许值：AGE_BETWEEN_18_23, AGE_BETWEEN_24_30、AGE_BETWEEN_31_40、AGE_BETWEEN_41_49、AGE_ABOVE_50
	AwemeFanBehaviors      []string `json:"aweme_fan_behaviors,omitempty"`      // 抖音用户行为类型，详见【附录-抖音达人互动用户行为类型】
	AwemeFanBehaviorsDays  string   `json:"aweme_fan_behaviors_days,omitempty"` // 抖音达人互动用户行为天数
	AwemeFanCategories     []int64  `json:"aweme_fan_categories,omitempty"`     // 抖音达人分类ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向），可通过【工具-抖音达人-查询抖音类目列表】接口获取
	AwemeFanAccounts       []int64  `json:"aweme_fan_accounts,omitempty"`       // 抖音达人ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向），可通过【工具-抖音达人-查询抖音类目下的推荐达人】接口获取
	AutoExtendEnabled      int64    `json:"auto_extend_enabled,omitempty"`      // 是否启用智能放量，允许值：0 关闭、1 开启
	AutoExtendTargets      []string `json:"auto_extend_targets,omitempty"`      // 可放开定向列表。当auto_extend_enabled=1 时必填。允许值：AGE 年龄、REGION 地域、GENDER 性别、INTEREST_ACTION 行为兴趣 、CUSTOM_AUDIENCE 更多人群-自定义人群
	Platform               []string `json:"platform,omitempty"`                 // 投放平台列表，允许值：ANDROID、 IOS、不传值为全选
	SmartInterestAction    string   `json:"smart_interest_action,omitempty"`    // 行为兴趣意向定向模式，允许值：RECOMMEND系统推荐，CUSTOM 自定义；不传值则为不限制需要注意：如果设置RECOMMEND，则传入action_scene、action_days、action_categories、action_words、 interest_categories、interest_words字段都无效
	ActionScene            []string `json:"action_scene,omitempty"`             // 行为场景，详见【附录-行为场景】，smart_interest_actionCUSTOM时有效，允许值：E-COMMERCE 电商互动行为、NEWS 资讯互动行为、APP APP推广互动行为
	ActionDays             int64    `json:"action_days,omitempty"`              // 用户发生行为天数，当 smart_interest_action 传 CUSTOM 时有效允许值：7, 15, 30, 60, 90, 180, 365
	ActionCategories       []int64  `json:"action_categories,omitempty"`        // 行为类目词，当 smart_interest_action 传 CUSTOM 时有效行为类目可以通过【工具-行为兴趣词管理-行为类目查询】获取
	ActionWords            []int64  `json:"action_words,omitempty"`             // 行为关键词，当 smart_interest_action 传 CUSTOM 时有效行为关键词可以通过【工具-行为兴趣词管理-行为关键词查询】获取
	InterestCategories     []int64  `json:"interest_categories,omitempty"`      // 兴趣类目词，当 smart_interest_action 传 CUSTOM 时有效兴趣类目可以通过【工具-行为兴趣词管理-兴趣类目查询】获取
	InterestWords          []int64  `json:"interest_words,omitempty"`           // 兴趣关键词，当 smart_interest_action 传 CUSTOM 时有效行为关键词可以通过【工具-行为兴趣词管理-行为关键词查询】获取
	Ac                     []string `json:"ac,omitempty"`                       // 网络类型, 详见【附录-受众网络类型】，允许值:WIFI、2G、3G、4G。 不传值或全传为全选
	RetargetingTagsInclude []int64  `json:"retargeting_tags_include,omitempty"` // 定向人群包id列表，长度限制 0-200。定向人群包可以通过【工具-DMP人群管理-获取人群包列表】获取
	RetargetingTagsExclude []int64  `json:"retargeting_tags_exclude,omitempty"` // 排除人群包id列表，长度限制 0-200。排除人群包可以通过【工具-DMP人群管理-获取人群包列表】获取
	LivePlatformTags       []string `json:"live_platform_tags,omitempty"`       // 直播带货平台精选人群包，当marketing_goal=LIVE_PROM_GOODS时有效，默认为全不选。允许值：LARGE_FANSCOUNT 高关注人群、ABNORMAL_ACTIVE高活跃人群、AWEME_FANS抖音号粉丝
}

type AdCreateCreative struct {
	CreativeMaterialMode          string                                `json:"creative_material_mode"`                     // 创意呈现方式，允许值：CUSTOM_CREATIVE 自定义创意、PROGRAMMATIC_CREATIVE 程序化创意
	FirstIndustryMode             int64                                 `json:"first_industry_mode"`                        // 创意一级行业ID。可从【获取行业列表】接口获取
	SecondIndustryId              int64                                 `json:"second_industry_id"`                         // 创意二级行业ID。可从【获取行业列表】接口获取
	ThirdIndustryId               int64                                 `json:"third_industry_id"`                          // 创意三级行业ID。可从【获取行业列表】接口获取
	AdKeywords                    []string                              `json:"ad_keywords,omitempty"`                      // 创意标签。最多20个标签，且每个标签长度要求为1~20个字符，汉字算2个字符
	CreativeList                  []AdCreateCreativeList                `json:"creative_list,omitempty"`                    // 自定义素材信息
	CreativeAutoGenerate          int64                                 `json:"creative_auto_generate,omitempty"`           // 是否开启「生成更多创意」
	ProgrammaticCreativeMediaList AdCreateProgrammaticCreativeMediaList `json:"programmatic_creative_media_list,omitempty"` // 程序化创意素材信息
	ProgrammaticCreativeTitleList AdCreateProgrammaticCreativeTitleList `json:"programmatic_creative_title_list,omitempty"` // 程序化创意标题信息
	ProgrammaticCreativeCard      AdCreateProgrammaticCreativeCard      `json:"programmatic_creative_card,omitempty"`       // 程序化创意推广卡片信息
	IsHomepageHide                int64                                 `json:"is_homepage_hide,omitempty"`                 // 抖音主页是否隐藏视频
}

// AdCreateCreativeList 广告创意 - creative_list
type AdCreateCreativeList struct {
	ImageMode             string                        `json:"image_mode,omitempty"`              // 创意素材类型
	VideoMaterial         AdCreateCustomVideoMaterial   `json:"video_material,omitempty"`          // 视频类型素材
	ImageMaterial         AdCreateImageMaterial         `json:"image_material,omitempty"`          // 图片类型素材
	TitleMaterial         AdCreateTitleMaterial         `json:"title_material,omitempty"`          // 标题类型素材，若选择了抖音号上的视频，不支持修改标题
	PromotionCardMaterial AdCreatePromotionCardMaterial `json:"promotion_card_material,omitempty"` // 推广卡片素材
}

// AdCreateCustomVideoMaterial 广告创意 - 视频类型素材
type AdCreateCustomVideoMaterial struct {
	VideoId      string `json:"video_id,omitempty"`       // 视频ID
	VideoCoverId string `json:"video_cover_id,omitempty"` // 视频封面ID
	AwemeItemId  int64  `json:"aweme_item_id,omitempty"`  // 抖音视频ID
}

// AdCreateImageMaterial 广告创意 - 图片类型素材
type AdCreateImageMaterial struct {
	ImageIds []string `json:"image_ids,omitempty"` // 图片ID列表
}

// AdCreateTitleMaterial 广告创意 - 标题类型素材，若选择了抖音号上的视频，不支持修改标题
type AdCreateTitleMaterial struct {
	Title        string                 `json:"title,omitempty"`         // 创意标题
	DynamicWords []AdCreateDynamicWords `json:"dynamic_words,omitempty"` // 动态词包对象列表
}

type AdCreateDynamicWords struct {
	WordId      int64  `json:"word_id,omitempty"`      // 动态词包ID
	DictName    string `json:"dict_name,omitempty"`    // 创意词包名称
	DefaultWord string `json:"default_word,omitempty"` // 创意词包默认词
}

// AdCreatePromotionCardMaterial 广告创意 - 推广卡片素材
type AdCreatePromotionCardMaterial struct {
	Title                   string   `json:"title,omitempty"`                     // 推广卡片标题
	SellingPoints           []string `json:"selling_points,omitempty"`            // 推广卡片卖点列表
	ImageId                 string   `json:"image_id,omitempty"`                  // 推广卡片配图
	ActionButton            string   `json:"action_button,omitempty"`             // 推广卡片行动号召按钮文案
	ButtonSmartOptimization int64    `json:"button_smart_optimization,omitempty"` // 是否对行动号召按钮文案启用智能优选
}

// AdCreateProgrammaticCreativeMediaList 广告创意 - 程序化创意素材信息
type AdCreateProgrammaticCreativeMediaList struct {
	ImageMode    string   `json:"image_mode,omitempty"`          // 创意素材类型，支持视频和图片
	VideoId      string   `json:"video_id,omitempty"`            // 视频ID
	VideoCoverId string   `json:"video_cover_id,omitempty"`      // 视频封面ID
	ImageIds     []string `json:"image_ids,omitempty,omitempty"` // 图片ID列表
}

// AdCreateProgrammaticCreativeTitleList 广告创意 - 程序化创意标题信息
type AdCreateProgrammaticCreativeTitleList struct {
	Title        string                 `json:"title,omitempty"`         // 创意标题
	DynamicWords []AdCreateDynamicWords `json:"dynamic_words,omitempty"` // 动态词包对象列表
}

// AdCreateProgrammaticCreativeCard 广告创意 - 程序化创意推广卡片信息
type AdCreateProgrammaticCreativeCard struct {
	PromotionCardTitle                   string   `json:"promotion_card_title,omitempty"`                     // 推广卡片标题，最多7个字
	PromotionCardSellingPoints           []string `json:"promotion_card_selling_points,omitempty"`            // 推广卡片卖点列表，卖点文字长度要求为12~18个字符，汉字算2个字符
	PromotionCardImageId                 string   `json:"promotion_card_image_id,omitempty"`                  // 推广卡片配图，可通过【获取图片素材】接口获得图片素材id
	PromotionCardActionButton            string   `json:"promotion_card_action_button,omitempty"`             // 推广卡片行动号召按钮文案
	PromotionCardButtonSmartOptimization int64    `json:"promotion_card_button_smart_optimization,omitempty"` // 是否对行动号召按钮文案启用智能优选
}

type AdCreateResData struct {
	AdId int64 `json:"ad_id"` // 创建的计划id
}

// AdCreateRes 获取广告账户数据-返回结构体
type AdCreateRes struct {
	QCError
	Data AdCreateResData `json:"data"`
}

// AdCreate 获取广告账户数据
func (m *Manager) AdCreate(req AdCreateReq) (res *AdCreateRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	err = m.client.CallWithJson(context.Background(), &res, "POST",
		m.url("%s?", conf.API_AD_CREATE), header, req.Body)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// AdUpdateReq 获取广告账户数据-请求
type AdUpdateReq struct {
	AccessToken string       // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        AdUpdateBody // POST请求的data
}

type AdUpdateBody struct {
	AdvertiserId    int64                   `json:"advertiser_id"` // 千川广告主账户id
	AdId            int64                   `json:"ad_id"`
	Name            string                  `json:"name"` // 计划名称，长度为1-100个字符，其中1个汉字算2位字符。名称不可重复，否则会报错
	DeliverySetting AdUpdateDeliverySetting `json:"delivery_setting"`
	Audience        AdUpdateAudience        `json:"audience"`
	AdUpdateCreative
}

type AdUpdateAudience AdCreateAudience

type AdUpdateDeliverySetting struct {
	FlowControlMode    string  `json:"flow_control_mode,omitempty"`    // 投放速度，详见【附录-计划投放速度类型】仅当 smart_bid_type 为SMART_BID_CUSTOM 时需传值，允许值：FLOW_CONTROL_MODE_FAST 尽快投放（默认值）、FLOW_CONTROL_MODE_BALANCE 均匀投放、FLOW_CONTROL_MODE_SMOOTH 优先低成本，对应千川后台「严格控制成本上限」勾选项
	Budget             float64 `json:"budget"`                         // 预算，最多支持两位小数当预算模式为日预算时，预算范围是300 - 9999999.99；当预算模式为总预算时，预算范围是max(300,投放天数x100) - 9999999.99
	CpaBid             float64 `json:"cpa_bid,omitempty"`              // 转化出价，出价不能大于预算仅当 smart_bid_type 为SMART_BID_CUSTOM 时需传值
	VideoScheduleType  string  `json:"video_schedule_type,omitempty"`  // 短视频投放日期选择方式，仅短视频带货场景需入参，允许值：SCHEDULE_FROM_NOW 从今天起长期投放（总预算模式下不支持）、SCHEDULE_START_END 设置开始和结束日期
	EndTime            string  `json:"end_time,omitempty"`             // 投放结束时间，形式如：2017-01-01结束时间不能比起始时间早。当video_schedule_type为SCHEDULE_START_END 设置开始和结束日期时需传入。当live_schedule_type 为SCHEDULE_TIME_ALLDAY 全天、SCHEDULE_TIME_WEEKLY_SETTING 指定时间段时必填；当 live_schedule_type 为SCHEDULE_TIME_FIXEDRANGE固定时长时不能传入
	ScheduleTime       string  `json:"schedule_time,omitempty"`        // 投放时段，当 live_schedule_type 为SCHEDULE_TIME_WEEKLY_SETTING 时生效默认全时段投放，格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。例如：填写"000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000"，则投放时段为周一到周日的11:30~13:30
	ScheduleFixedRange int64   `json:"schedule_fixed_range,omitempty"` // 固定投放时长当 live_schedule_type 为 SCHEDULE_TIME_FIXEDRANGE 时必填；当live_schedule_type 为SCHEDULE_TIME_ALLDAY 全天、SCHEDULE_TIME_WEEKLY_SETTING 指定时间段时不能传入。单位为秒，最小值为1800（0.5小时），最大值为48*1800（24小时），值必须为1800倍数，不然会报错
}

type AdUpdateCreative struct {
	CreativeMaterialMode          string                                `json:"creative_material_mode"`                     // 创意呈现方式，允许值：CUSTOM_CREATIVE 自定义创意、PROGRAMMATIC_CREATIVE 程序化创意
	FirstIndustryMode             int64                                 `json:"first_industry_mode"`                        // 创意一级行业ID。可从【获取行业列表】接口获取
	SecondIndustryId              int64                                 `json:"second_industry_id"`                         // 创意二级行业ID。可从【获取行业列表】接口获取
	ThirdIndustryId               int64                                 `json:"third_industry_id"`                          // 创意三级行业ID。可从【获取行业列表】接口获取
	AdKeywords                    []string                              `json:"ad_keywords,omitempty"`                      // 创意标签。最多20个标签，且每个标签长度要求为1~20个字符，汉字算2个字符
	CreativeList                  []AdUpdateCreativeList                `json:"creative_list,omitempty"`                    // 自定义素材信息
	CreativeAutoGenerate          int64                                 `json:"creative_auto_generate,omitempty"`           // 是否开启「生成更多创意」
	ProgrammaticCreativeMediaList AdUpdateProgrammaticCreativeMediaList `json:"programmatic_creative_media_list,omitempty"` // 程序化创意素材信息
	ProgrammaticCreativeTitleList AdUpdateProgrammaticCreativeTitleList `json:"programmatic_creative_title_list,omitempty"` // 程序化创意标题信息
	ProgrammaticCreativeCard      AdUpdateProgrammaticCreativeCard      `json:"programmatic_creative_card,omitempty"`       // 程序化创意推广卡片信息
	IsHomepageHide                int64                                 `json:"is_homepage_hide,omitempty"`                 // 抖音主页是否隐藏视频
}

// AdUpdateCreativeList 广告创意 - creative_list
type AdUpdateCreativeList struct {
	CreativeId            int64                         `json:"creative_id"`                       // 创意ID
	ImageMode             string                        `json:"image_mode,omitempty"`              // 创意素材类型
	VideoMaterial         AdUpdateCustomVideoMaterial   `json:"video_material,omitempty"`          // 视频类型素材
	ImageMaterial         AdUpdateImageMaterial         `json:"image_material,omitempty"`          // 图片类型素材
	TitleMaterial         AdUpdateTitleMaterial         `json:"title_material,omitempty"`          // 标题类型素材，若选择了抖音号上的视频，不支持修改标题
	PromotionCardMaterial AdUpdatePromotionCardMaterial `json:"promotion_card_material,omitempty"` // 推广卡片素材
}

// AdUpdateCustomVideoMaterial 广告创意 - 视频类型素材
type AdUpdateCustomVideoMaterial struct {
	ID int64 `json:"id,omitempty"`
	AdCreateCustomVideoMaterial
}

// AdUpdateImageMaterial 广告创意 - 图片类型素材
type AdUpdateImageMaterial struct {
	ID int64 `json:"id,omitempty"`
	AdCreateImageMaterial
}

// AdUpdateTitleMaterial 广告创意 - 标题类型素材，若选择了抖音号上的视频，不支持修改标题
type AdUpdateTitleMaterial struct {
	ID int64 `json:"id,omitempty"`
	AdCreateTitleMaterial
}

type AdUpdateDynamicWords AdCreateDynamicWords

// AdUpdatePromotionCardMaterial 广告创意 - 推广卡片素材
type AdUpdatePromotionCardMaterial struct {
	ID          int64 `json:"id,omitempty"`           // 素材唯一标识，通过获取计划详情接口可以获取
	ComponentId int64 `json:"component_id,omitempty"` // 组件唯一标识，通过获取计划详情接口可以获取
	AdCreatePromotionCardMaterial
}

// AdUpdateProgrammaticCreativeMediaList 广告创意 - 程序化创意素材信息
type AdUpdateProgrammaticCreativeMediaList AdCreateProgrammaticCreativeMediaList

// AdUpdateProgrammaticCreativeTitleList 广告创意 - 程序化创意标题信息
type AdUpdateProgrammaticCreativeTitleList AdCreateProgrammaticCreativeTitleList

// AdUpdateProgrammaticCreativeCard 广告创意 - 程序化创意推广卡片信息
type AdUpdateProgrammaticCreativeCard AdCreateProgrammaticCreativeCard

type AdUpdateResData struct {
	AdId      int64 `json:"ad_id"` // 修改的计划id
	ErrorList []struct {
		ObjectId     int64  `json:"object_id"`     // 错误对象id
		ObjectType   string `json:"object_type"`   // 错误对象类型，返回值: AD 计划，CREATIVE 创意
		OptType      string `json:"opt_type"`      // 操作类型，返回值：UPDATE 更新，ADD 新建
		ErrorCode    int64  `json:"error_code"`    // 错误码
		ErrorMessage string `json:"error_message"` // 错误信息
	} `json:"error_list"` // 错误list，计划为分块更新，存在部分内容更新失败，部分内容更新成功的的情况。若计划更新成功，则返回为空数组；若更新失败，则返回错误的部分及原因
}

// AdUpdateRes 获取广告账户数据-返回结构体
type AdUpdateRes struct {
	QCError
	Data AdUpdateResData `json:"data"`
}

// AdUpdate 获取广告账户数据
func (m *Manager) AdUpdate(req AdUpdateReq) (res *AdUpdateRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	err = m.client.CallWithJson(context.Background(), &res, "POST",
		m.url("%s?", conf.API_AD_CREATE), header, req.Body)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------
