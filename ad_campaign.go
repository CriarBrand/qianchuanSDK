// 广告组管理相关API
// 巨量开放平台地址：https://open.oceanengine.com/doc/index.html?key=qianchuan&type=api&id=1701977925996558

package qianchuanSDK

import (
	"context"
	"encoding/json"
	"github.com/CriarBrand/qianchuanSDK/conf"
	"net/http"
)

// CampaignCreateReq 获取广告账户数据-请求
type CampaignCreateReq struct {
	AccessToken string             // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        CampaignCreateBody // POST请求的data
}

type CampaignCreateBody struct {
	AdvertiserId  int64   `json:"advertiser_id"`    // 千川广告账户ID
	CampaignName  string  `json:"campaign_name"`    // 广告组名称，长度为1-100个字符，其中1个中文字符算2位
	MarketingGoal string  `json:"marketing_goal"`   // 营销目标，允许值：VIDEO_PROM_GOODS 短视频带货、LIVE_PROM_GOODS 直播带货
	BudgetMode    string  `json:"budget_mode"`      // 预算类型（创建后不可修改），详见【附录-预算类型】，允许值：BUDGET_MODE_DAY 日预算，BUDGET_MODE_INFINITE 预算不限
	Budget        float64 `json:"budget,omitempty"` // 条件必填,广告组预算，最多支持两位小数，当budget_mode为BUDGET_MODE_DAY时必填，且日预算不少于300元
}

type CampaignCreateResData struct {
	CampaignId int64 `json:"campaign_id"` // 创建的广告组id
}

// CampaignCreateRes 获取广告账户数据-返回结构体
type CampaignCreateRes struct {
	QCError
	Data CampaignCreateResData `json:"data"`
}

// CampaignCreate 获取广告账户数据
func (m *Manager) CampaignCreate(req CampaignCreateReq) (res *CampaignCreateRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	err = m.client.CallWithJson(context.Background(), &res, "POST",
		m.url("%s?", conf.API_CAMPAIGN_CREATE), header, req.Body)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// CampaignUpdateReq 获取广告账户数据-请求
type CampaignUpdateReq struct {
	AccessToken string             // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        CampaignUpdateBody // POST请求的data
}

type CampaignUpdateBody struct {
	AdvertiserId int64   `json:"advertiser_id"`           // 千川广告账户ID
	CampaignId   int64   `json:"campaign_id"`             // 广告组名称，长度为1-100个字符，其中1个中文字符算2位
	CampaignName string  `json:"campaign_name,omitempty"` // 广告组名称，长度为1-100个字符，其中1个中文字符算2位,需要注意：广告组名称不修改的话，可不填。填入的话，需与原广告组名称不同，否则报错
	BudgetMode   string  `json:"budget_mode"`             // 预算类型（创建后不可修改），详见【附录-预算类型】，允许值：BUDGET_MODE_DAY 日预算，BUDGET_MODE_INFINITE 预算不限
	Budget       float64 `json:"budget,omitempty"`        // 条件必填,广告组预算，最多支持两位小数，当budget_mode为BUDGET_MODE_DAY时必填，且日预算不少于300元
}

type CampaignUpdateResData struct {
	CampaignId int64 `json:"campaign_id"` // 修改的广告组id
}

// CampaignUpdateRes 获取广告账户数据-返回结构体
type CampaignUpdateRes struct {
	QCError
	Data CampaignUpdateResData `json:"data"`
}

// CampaignUpdate 获取广告账户数据
func (m *Manager) CampaignUpdate(req CampaignUpdateReq) (res *CampaignUpdateRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	err = m.client.CallWithJson(context.Background(), &res, "POST",
		m.url("%s?", conf.API_CAMPAIGN_UPDATE), header, req.Body)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// BatchCampaignStatusUpdateReq 获取广告账户数据-请求
type BatchCampaignStatusUpdateReq struct {
	AccessToken string                        // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        BatchCampaignStatusUpdateBody // POST请求的data
}

type BatchCampaignStatusUpdateBody struct {
	AdvertiserId int64   `json:"advertiser_id"` // 千川广告账户ID
	CampaignIds  []int64 `json:"campaign_ids"`  // 广告组ID，不超过10个，操作更新的广告组ID需要属于千川账户ID否则会报错；
	OptStatus    string  `json:"opt_status"`    // 操作类型，允许值: "ENABLE"：启用, "DELETE"：删除, "DISABLE"：暂停；对于删除的广告组不可进行任何操作。
}

type BatchCampaignStatusUpdateResData struct {
	Success []int64 `json:"success"` // 更新成功的广告组ID列表
	Errors  []struct {
		CampaignId   int64  `json:"campaign_id"`   // 更新失败广告组ID
		ErrorMessage string `json:"error_message"` // 更新失败的原因
	} `json:"errors"` // 更新失败的广告组列表
}

// BatchCampaignStatusUpdateRes 获取广告账户数据-返回结构体
type BatchCampaignStatusUpdateRes struct {
	QCError
	Data BatchCampaignStatusUpdateResData `json:"data"`
}

// BatchCampaignStatusUpdate 获取广告账户数据
func (m *Manager) BatchCampaignStatusUpdate(req BatchCampaignStatusUpdateReq) (res *BatchCampaignStatusUpdateRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	err = m.client.CallWithJson(context.Background(), &res, "POST",
		m.url("%s?", conf.API_BATCH_CAMPAIGN_STATUS_UPDATE), header, req.Body)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// CampaignListGetReq 获取广告账户数据-请求
type CampaignListGetReq struct {
	AdvertiserId int64                 `json:"advertiser_id"`       // 千川广告账户ID
	Page         int64                 `json:"page,omitempty"`      // 页码，默认为1
	PageSize     int64                 `json:"page_size,omitempty"` // 页面大小，默认值: 10， 允许值：10、20、50、100、500、1000
	Filter       CampaignListGetFilter `json:"filter"`              // 过滤器，无过滤条件情况下返回“所有不包含已删除”的广告组列表
	AccessToken  string                `json:"access_token"`        // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type CampaignListGetFilter struct {
	Ids           []int64 `json:"ids,omitempty"`    // 广告组ID列表，目前只支持一个。
	Name          string  `json:"name,omitempty"`   // 广告组名称关键字，长度为1-30个字符，其中1个中文字符算2位
	MarketingGoal string  `json:"marketing_goal"`   // 广告组营销目标，允许值：VIDEO_PROM_GOODS：短视频带货、LIVE_PROM_GOODS：直播带货
	Status        string  `json:"status,omitempty"` // 广告组状态，允许值：ALL：所有包含已删除、ENABLE：启用、DISABLE：暂停、DELETE：已删除。不传入即默认返回“所有不包含已删除”
}

type CampaignListGetResData struct {
	List []struct {
		ID            int64   `json:"id"`             // 广告组ID
		Name          string  `json:"name"`           // 广告组名称
		Budget        float64 `json:"budget"`         // 广告组预算，单位：元，精确到两位小数。
		BudgetMode    string  `json:"budget_mode"`    // 广告组预算类型
		MarketingGoal string  `json:"marketing_goal"` // 广告组营销目标，VIDEO_PROM_GOODS：短视频带货、LIVE_PROM_GOODS：直播带货。
		Status        string  `json:"status"`         // 广告组状态，ALL：所有包含已删除、ENABLE：启用、DISABLE：暂停、DELETE：已删除。
		CreateDate    string  `json:"create_date"`    // 广告组创建日期, 格式：yyyy-mm-dd
	} `json:"list"`
	PageInfo PageInfo `json:"page_info"`
}

// CampaignListGetRes 获取广告账户数据-返回结构体
type CampaignListGetRes struct {
	QCError
	Data CampaignListGetResData `json:"data"`
}

// CampaignListGet 获取广告账户数据
func (m *Manager) CampaignListGet(req CampaignListGetReq) (res *CampaignListGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	// 接收一个结构体并转为string格式
	filter, err := json.Marshal(req.Filter)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d&filter=%s&page=%d&page_size=%d",
			conf.API_CAMPAIGN_LIST_GET, req.AdvertiserId, string(filter), req.Page, req.PageSize), header, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------
