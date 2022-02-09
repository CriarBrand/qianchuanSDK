// 广告数据报表相关API
// 巨量开放平台地址：https://open.oceanengine.com/doc/index.html?key=qianchuan&type=api&id=1697466345527308

package qianchuanSDK

import (
	"context"
	"encoding/json"
	"github.com/CriarBrand/qianchuanSDK/conf"
	"net/http"
)

// AdvertiserReportReq 获取广告账户数据-请求
type AdvertiserReportReq struct {
	AdvertiserId int64                     `json:"advertiser_id"` // 千川广告主账户id
	StartDate    string                    `json:"start_date"`    // 开始时间，格式 2021-04-05
	EndDate      string                    `json:"end_date"`      // 结束时间，格式 2021-04-05，时间跨度不能超过180天
	Fields       []string                  `json:"fields"`        // 需要查询的消耗指标
	Filtering    AdvertiserReportFiltering `json:"filtering"`
	//MarketingGoal string   // 过滤条件 营销目标，允许值：VIDEO_PROM_GOODS：短视频带货  LIVE_PROM_GOODS：直播间带货  ALL：不限
	AccessToken string `json:"access_token"` // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type AdvertiserReportFiltering struct {
	MarketingGoal string `json:"marketing_goal"`
}
type AdvertiserReportResDetail struct {
	AdvertiserId         int64   `json:"advertiser_id"`            // 广告主id
	StatCost             float64 `json:"stat_cost"`                // 消耗
	ShowCnt              int64   `json:"show_cnt"`                 // 展示次数
	Ctr                  float64 `json:"ctr"`                      // 点击率
	CpmPlatform          float64 `json:"cpm_platform"`             // 平均千次展示费用
	ClickCnt             int64   `json:"click_cnt"`                // 点击次数
	PayOrderCount        int64   `json:"pay_order_count"`          // 成交订单数
	CreateOrderAmount    float64 `json:"create_order_amount"`      // 下单成交金额
	CreateOrderCount     int64   `json:"create_order_count"`       // 下单订单数
	PayOrderAmount       float64 `json:"pay_order_amount"`         // 成交订单金额
	CreateOrderRoi       float64 `json:"create_order_roi"`         // 下单roi
	DyFollow             int64   `json:"dy_follow"`                // 新增粉丝数
	PrepayAndPayOrderRoi float64 `json:"prepay_and_pay_order_roi"` // 支付roi
	PrepayOrderCount     int64   `json:"prepay_order_count"`       // 广告预售订单数
	PrepayOrderAmount    float64 `json:"prepay_order_amount"`      // 广告预售订单金额
}

type AdvertiserReportResData struct {
	List []AdvertiserReportResDetail `json:"list"`
}

// AdvertiserReportRes 获取广告账户数据-返回结构体
type AdvertiserReportRes struct {
	QCError
	Data AdvertiserReportResData `json:"data"`
}

// AdvertiserReport 获取广告账户数据
func (m *Manager) AdvertiserReport(req AdvertiserReportReq) (res *AdvertiserReportRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	// 接收一个数组并转为string格式
	fields, err := json.Marshal(req.Fields)
	if err != nil {
		panic(err)
	}
	// 接收一个结构体并转为string格式
	filtering, err := json.Marshal(req.Filtering)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d&start_date=%s&end_date=%s&fields=%s&filtering=%s",
			conf.API_REPORT_ADVERTISER_GET, req.AdvertiserId, req.StartDate, req.EndDate, string(fields), string(filtering)), header, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// ReportAdGetReq 获取广告计划数据-请求
type ReportAdGetReq struct {
	AdvertiserId int64                `json:"advertiser_id"`         // 千川广告主账户id
	StartDate    string               `json:"start_date"`            // 开始时间，格式 2021-04-05
	EndDate      string               `json:"end_date"`              // 结束时间，格式 2021-04-05，时间跨度不能超过180天
	Fields       []string             `json:"fields"`                // 需要查询的消耗指标
	Filtering    ReportAdGetFiltering `json:"filtering"`             //过滤条件
	OrderField   string               `json:"order_field,omitempty"` // 排序字段
	OrderType    string               `json:"order_type,omitempty"`  // 排序方式，允许值： ASC 升序（默认）、DESC 降序
	Page         int64                `json:"page,omitempty"`        // 页码，默认为1
	PageSize     int64                `json:"page_size,omitempty"`   // 页面大小，默认为10，取值范围：1-500
	AccessToken  string               `json:"access_token"`          // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type ReportAdGetFiltering struct {
	AdIds         []int64 `json:"ad_ids"`         // 广告计划id列表，最多支持100个
	MarketingGoal string  `json:"marketing_goal"` // 过滤条件 营销目标，允许值：VIDEO_PROM_GOODS：短视频带货  LIVE_PROM_GOODS：直播间带货  ALL：不限
}

type ReportAdGetResDetail struct {
	AdvertiserId               int64   `json:"advertiser_id"`                  // 广告主id
	AdId                       int64   `json:"ad_id"`                          // 广告计划id
	StatCost                   float64 `json:"stat_cost"`                      // 消耗
	ShowCnt                    int64   `json:"show_cnt"`                       // 展示次数
	Ctr                        float64 `json:"ctr"`                            // 点击率
	CpmPlatform                float64 `json:"cpm_platform"`                   // 平均千次展示费用
	ClickCnt                   int64   `json:"click_cnt"`                      // 点击次数
	PayOrderCount              int64   `json:"pay_order_count"`                // 成交订单数
	CreateOrderAmount          float64 `json:"create_order_amount"`            // 下单成交金额
	CreateOrderCount           int64   `json:"create_order_count"`             // 下单订单数
	PayOrderAmount             float64 `json:"pay_order_amount"`               // 成交订单金额
	CreateOrderRoi             float64 `json:"create_order_roi"`               // 下单roi
	PrepayAndPayOrderRoi       float64 `json:"prepay_and_pay_order_roi"`       // 支付roi
	PrepayOrderCount           int64   `json:"prepay_order_count"`             // 广告预售订单数
	PrepayOrderAmount          float64 `json:"prepay_order_amount"`            // 广告预售订单金额
	DyFollow                   int64   `json:"dy_follow"`                      // 新增粉丝数
	ConvertCnt                 int64   `json:"convert_cnt"`                    // 转化数
	ConvertCost                int64   `json:"convert_cost"`                   // 转化成本
	ConvertRate                float64 `json:"convert_rate"`                   // 转化率
	DyShare                    int64   `json:"dy_share"`                       // 分享次数。直播间带货：LIVE_PROM_GOODS 不支持该指标
	DyComment                  int64   `json:"dy_comment"`                     // 评论次数。直播间带货：LIVE_PROM_GOODS 不支持该指标
	DyLike                     int64   `json:"dy_like"`                        // 点赞次数。直播间带货：LIVE_PROM_GOODS 不支持该指标
	LivePayOrderCostPerOrder   float64 `json:"live_pay_order_cost_per_order"`  // 成交客单价。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveEnterCnt          int64   `json:"luban_live_enter_cnt"`           // 直播间观看人次。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LiveWatchOneMinuteCount    int64   `json:"live_watch_one_minute_count"`    // 直播间超过1分钟观看人次。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LiveFansClubJoinCnt        int64   `json:"live_fans_club_join_cnt"`        // 直播间新加团人次。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveSlidecartClickCnt int64   `json:"luban_live_slidecart_click_cnt"` // 直播间查看购物车次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveClickProductCnt   int64   `json:"luban_live_click_product_cnt"`   // 直播间商品点击次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveCommentCnt        int64   `json:"luban_live_comment_cnt"`         // 直播间评论次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveShareCnt          int64   `json:"luban_live_share_cnt"`           // 直播间分享次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveGiftCnt           int64   `json:"luban_live_gift_cnt"`            // 直播间打赏次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveGiftAmount        float64 `json:"luban_live_gift_amount"`         // 直播间音浪收入。短视频带货：VIDEO_PROM_GOODS 不支持该指标
}

type ReportAdGetResData struct {
	List     []ReportAdGetResDetail `json:"list"`
	PageInfo PageInfo               `json:"page_info"`
}

// ReportAdGetRes 获取广告计划数据-返回结构体
type ReportAdGetRes struct {
	QCError
	Data ReportAdGetResData `json:"data"`
}

// ReportAdGet 获取广告计划数据
func (m *Manager) ReportAdGet(req ReportAdGetReq) (res *ReportAdGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)

	reqUrl := conf.API_HTTP_SCHEME + conf.API_HOST + conf.API_REPORT_AD_GET
	reqUrl, err = BuildQuery(reqUrl, req, []string{"access_token"})
	if err != nil {
		return nil, err
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", reqUrl, header, nil)

	//// 接收一个数组并转为string格式
	//fields, err := json.Marshal(req.Fields)
	//if err != nil {
	//	panic(err)
	//}
	//// 接收一个结构体并转为string格式
	//filtering, err := json.Marshal(req.Filtering)
	//if err != nil {
	//	panic(err)
	//}
	//// 判断OrderType或OrderField是否为空，如果为空则get参数不加上
	//if req.OrderType == "" || req.OrderField == "" {
	//	err = m.client.CallWithJson(context.Background(), &res, "GET",
	//		m.url("%s?advertiser_id=%d&start_date=%s&end_date=%s&fields=%s&filtering=%s&page=%d&page_size=%d",
	//			conf.API_REPORT_AD_GET, req.AdvertiserId, req.StartDate, req.EndDate, string(fields),
	//			string(filtering), req.Page, req.PageSize), header, nil)
	//	return res, err
	//}
	//err = m.client.CallWithJson(context.Background(), &res, "GET",
	//	m.url("%s?advertiser_id=%d&start_date=%s&end_date=%s&fields=%s&filtering=%s&order_field=%s&order_type=%s&page=%d&page_size=%d",
	//		conf.API_REPORT_AD_GET, req.AdvertiserId, req.StartDate, req.EndDate, string(fields),
	//		string(filtering), req.OrderField, req.OrderType, req.Page, req.PageSize), header, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// ReportCreativeGetReq 获取广告创意数据-请求
type ReportCreativeGetReq struct {
	AdvertiserId int64                      `json:"advertiser_id"`         // 千川广告主账户id
	StartDate    string                     `json:"start_date"`            // 开始时间，格式 2021-04-05
	EndDate      string                     `json:"end_date"`              // 结束时间，格式 2021-04-05，时间跨度不能超过180天
	Fields       []string                   `json:"fields"`                // 需要查询的消耗指标
	Filtering    ReportCreativeGetFiltering `json:"filtering"`             //过滤条件
	OrderField   string                     `json:"order_field,omitempty"` // 排序字段
	OrderType    string                     `json:"order_type,omitempty"`  // 排序方式，允许值： ASC 升序（默认）、DESC 降序
	Page         int64                      `json:"page,omitempty"`        // 页码，默认为1
	PageSize     int64                      `json:"page_size,omitempty"`   // 页面大小，默认为10，取值范围：1-500
	AccessToken  string                     `json:"access_token"`          // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type ReportCreativeGetFiltering struct {
	CreativeIds   []int64 `json:"creative_ids"`   // 广告创意id列表，数量不超过100
	MarketingGoal string  `json:"marketing_goal"` // 过滤条件 营销目标，允许值：VIDEO_PROM_GOODS：短视频带货  LIVE_PROM_GOODS：直播间带货  ALL：不限
}

type ReportCreativeGetResData struct {
	List     []ReportCreativeGetResDataDetail `json:"list"`
	PageInfo PageInfo                         `json:"page_info"`
}
type ReportCreativeGetResDataDetail struct {
	AdvertiserId               int64   `json:"advertiser_id"`                  // 广告主id
	CreativeId                 int64   `json:"creative_id"`                    // 广告创意id
	StatCost                   float64 `json:"stat_cost"`                      // 消耗
	ShowCnt                    int64   `json:"show_cnt"`                       // 展示次数
	Ctr                        float64 `json:"ctr"`                            // 点击率
	CpmPlatform                float64 `json:"cpm_platform"`                   // 平均千次展示费用
	ClickCnt                   int64   `json:"click_cnt"`                      // 点击次数
	PayOrderCount              int64   `json:"pay_order_count"`                // 成交订单数
	CreateOrderAmount          float64 `json:"create_order_amount"`            // 下单成交金额
	CreateOrderCount           int64   `json:"create_order_count"`             // 下单订单数
	PayOrderAmount             float64 `json:"pay_order_amount"`               // 成交订单金额
	CreateOrderRoi             float64 `json:"create_order_roi"`               // 下单roi
	PrepayAndPayOrderRoi       float64 `json:"prepay_and_pay_order_roi"`       // 支付roi
	PrepayOrderCount           int64   `json:"prepay_order_count"`             // 广告预售订单数
	PrepayOrderAmount          float64 `json:"prepay_order_amount"`            // 广告预售订单金额
	DyFollow                   int64   `json:"dy_follow"`                      // 新增粉丝数
	ConvertCnt                 int64   `json:"convert_cnt"`                    // 转化数
	ConvertCost                int64   `json:"convert_cost"`                   // 转化成本
	ConvertRate                float64 `json:"convert_rate"`                   // 转化率
	DyShare                    int64   `json:"dy_share"`                       // 分享次数。直播间带货：LIVE_PROM_GOODS 不支持该指标
	DyComment                  int64   `json:"dy_comment"`                     // 评论次数。直播间带货：LIVE_PROM_GOODS 不支持该指标
	DyLike                     int64   `json:"dy_like"`                        // 点赞次数。直播间带货：LIVE_PROM_GOODS 不支持该指标
	LivePayOrderCostPerOrder   float64 `json:"live_pay_order_cost_per_order"`  // 成交客单价。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveEnterCnt          int64   `json:"luban_live_enter_cnt"`           // 直播间观看人次。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LiveWatchOneMinuteCount    int64   `json:"live_watch_one_minute_count"`    // 直播间超过1分钟观看人次。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LiveFansClubJoinCnt        int64   `json:"live_fans_club_join_cnt"`        // 直播间新加团人次。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveSlidecartClickCnt int64   `json:"luban_live_slidecart_click_cnt"` // 直播间查看购物车次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveClickProductCnt   int64   `json:"luban_live_click_product_cnt"`   // 直播间商品点击次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveCommentCnt        int64   `json:"luban_live_comment_cnt"`         // 直播间评论次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveShareCnt          int64   `json:"luban_live_share_cnt"`           // 直播间分享次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveGiftCnt           int64   `json:"luban_live_gift_cnt"`            // 直播间打赏次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveGiftAmount        float64 `json:"luban_live_gift_amount"`         // 直播间音浪收入。短视频带货：VIDEO_PROM_GOODS 不支持该指标
}

// ReportCreativeGetRes 获取千川广告账户全量信息-返回结构体
type ReportCreativeGetRes struct {
	QCError
	Data ReportCreativeGetResData `json:"data"`
}

// ReportCreativeGet 获取千川广告账户全量信息
func (m *Manager) ReportCreativeGet(req ReportCreativeGetReq) (res *ReportCreativeGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)

	reqUrl := conf.API_HTTP_SCHEME + conf.API_HOST + conf.API_REPORT_CREATIVE_GET
	reqUrl, err = BuildQuery(reqUrl, req, []string{"access_token"})
	if err != nil {
		return nil, err
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", reqUrl, header, nil)

	//// 接收一个数组并转为string格式
	//fields, err := json.Marshal(req.Fields)
	//if err != nil {
	//	panic(err)
	//}
	//// 接收一个结构体并转为string格式
	//filtering, err := json.Marshal(req.Filtering)
	//if err != nil {
	//	panic(err)
	//}
	//// 判断OrderType或OrderField是否为空，如果为空则get参数不加上
	//if req.OrderType == "" || req.OrderField == "" {
	//	err = m.client.CallWithJson(context.Background(), &res, "GET",
	//		m.url("%s?advertiser_id=%d&start_date=%s&end_date=%s&fields=%s&filtering=%s&page=%d&page_size=%d",
	//			conf.API_REPORT_CREATIVE_GET, req.AdvertiserId, req.StartDate, req.EndDate, string(fields),
	//			string(filtering), req.Page, req.PageSize), header, nil)
	//	return res, err
	//}
	//err = m.client.CallWithJson(context.Background(), &res, "GET",
	//	m.url("%s?advertiser_id=%d&start_date=%s&end_date=%s&fields=%s&filtering=%s&order_field=%s&order_type=%s&page=%d&page_size=%d",
	//		conf.API_REPORT_CREATIVE_GET, req.AdvertiserId, req.StartDate, req.EndDate, string(fields),
	//		string(filtering), req.OrderField, req.OrderType, req.Page, req.PageSize), header, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------
