// 账户管理相关API
// 巨量开放平台地址：https://open.oceanengine.com/doc/index.html?key=qianchuan&type=api&id=1697459480882190

package qianchuanSDK

import (
	"context"
	"encoding/json"
	"github.com/CriarBrand/qianchuanSDK/conf"
	"net/http"
)

// AdvertiserListReq 获取已授权的账户（店铺/代理商）-请求
type AdvertiserListReq struct {
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AppId       int64  // 开发者申请的应用APP_ID，可通过“应用管理”界面查看
	Secret      string // 开发者应用的私钥Secret，可通过“应用管理”界面查看（确保填入secret与app_id对应以免报错！）
}

type AdvertiserListResCom struct {
	List []AdvertiserListResData `json:"list"`
}

// AdvertiserListResData 获取已授权的账户（店铺/代理商）-返回
type AdvertiserListResData struct {
	AdvertiserId   int64  `json:"advertiser_id"`   // 账户id
	AdvertiserName string `json:"advertiser_name"` // 账户名称
	IsValid        bool   `json:"is_valid"`        // 授权有效性，返回值：true/false,用于判断当前授权关系是否仍然有效
	AccountRole    string `json:"account_role"`    // 授权账号角色，返回值：PLATFORM_ROLE_QIANCHUAN_AGENT代理商账户、PLATFORM_ROLE_SHOP_ACCOUNT 店铺账户
}

// AdvertiserListRes 获取已授权的账户（店铺/代理商）-返回结构体
type AdvertiserListRes struct {
	QCError
	Data AdvertiserListResCom `json:"data"`
}

// AdvertiserList 获取已授权的账户（店铺/代理商）
func (m *Manager) AdvertiserList(req AdvertiserListReq) (res *AdvertiserListRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&app_id=%d&secret=%s",
		conf.API_ADVERTISER_LIST, req.AccessToken, m.Credentials.AppId, m.Credentials.AppSecret), nil, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// ShopAdvertiserListReq 获取店铺账户关联的广告账户列表-请求
type ShopAdvertiserListReq struct {
	ShopId      int64  // 店铺id
	Page        uint64 // 页码.默认值: 1
	PageSize    uint64 // 页面数据量.默认值: 10， 最大值：100
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type ShopAdvertiserListResCom struct {
	List     []int64  `json:"list"`
	PageInfo PageInfo `json:"page_info"`
}

// ShopAdvertiserListRes 获取店铺账户关联的广告账户列表-返回结构体
type ShopAdvertiserListRes struct {
	QCError
	Data ShopAdvertiserListResCom `json:"data"`
}

// ShopAdvertiserList 获取店铺账户关联的广告账户列表
func (m *Manager) ShopAdvertiserList(req ShopAdvertiserListReq) (res *ShopAdvertiserListRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?shop_id=%d&page=%d&page_size=%d",
		conf.API_SHOP_ADVERTISER_LIST, req.ShopId, req.Page, req.PageSize), header, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// AgentAdvertiserListReq 获取代理商账户关联的广告账户列表-请求
type AgentAdvertiserListReq struct {
	AdvertiserId int64  // 代理商ID
	Page         uint64 // 页码.默认值: 1
	PageSize     uint64 // 页面数据量.默认值: 10， 最大值：100
	AccessToken  string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type AgentAdvertiserListResCom struct {
	List     []int64  `json:"list"`
	PageInfo PageInfo `json:"page_info"`
}

// AgentAdvertiserListRes 获取代理商账户关联的广告账户列表-返回结构体
type AgentAdvertiserListRes struct {
	QCError
	Data AgentAdvertiserListResCom `json:"data"`
}

// AgentAdvertiserList 获取代理商账户关联的广告账户列表
func (m *Manager) AgentAdvertiserList(req AgentAdvertiserListReq) (res *AgentAdvertiserListRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?Agent_id=%d&page=%d&page_size=%d",
		conf.API_AGENT_ADVERTISER_LIST, req.AdvertiserId, req.Page, req.PageSize), header, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// UserInfoReq 获取授权时登录用户信息-请求
type UserInfoReq struct {
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type UserInfoResCom struct {
	ID          int64  `json:"id"`           //用户id
	Email       string `json:"email"`        //邮箱（已经脱敏处理）
	DisplayName string `json:"display_name"` // 用户名
}

// UserInfoRes 获取授权时登录用户信息-返回结构体
type UserInfoRes struct {
	QCError
	Data UserInfoResCom `json:"data"`
}

// UserInfo 获取授权时登录用户信息
func (m *Manager) UserInfo(req UserInfoReq) (res *UserInfoRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?", conf.API_USER_INFO), header, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// ShopAccountInfoReq 获取店铺账户信息
type ShopAccountInfoReq struct {
	ShopIds     []int64 //店铺id列表，一次最多查询10个shop_id信息
	AccessToken string  // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type ShopAccountInfoResCom struct {
	List []ShopAccountInfoResComDetail `json:"list"`
}
type ShopAccountInfoResComDetail struct {
	ShopId   int64  `json:"shop_id"`
	ShopName string `json:"shop_name"`
}

// ShopAccountInfoRes 获取店铺账户信息-返回结构体
type ShopAccountInfoRes struct {
	QCError
	Data ShopAccountInfoResCom `json:"data"`
}

// ShopAccountInfo  获取店铺账户信息
func (m *Manager) ShopAccountInfo(req ShopAccountInfoReq) (res *ShopAccountInfoRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)

	// 接收一个数组并转为string格式
	shopIds, err := json.Marshal(req.ShopIds)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?shop_ids=%s",
		conf.API_SHOP_ACCOUNT_INFO, string(shopIds)), header, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// AgentInfoReq 获取代理商账户信息-请求
type AgentInfoReq struct {
	AdvertiserIds []int64
	Fields        []string
	AccessToken   string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type AgentInfoResData struct {
	AgentId       int64  `json:"agent_id"`
	AgentName     string `json:"agent_name"`
	CustomerId    string `json:"customer_id"`
	CompanyId     string `json:"company_id"`
	CompanyName   string `json:"company_name"`
	AccountStatus string `json:"account_status"`
	CreateTime    string `json:"create_time"`
	Role          string `json:"role"`
}

// AgentInfoRes 获取代理商账户信息-返回结构体
type AgentInfoRes struct {
	QCError
	Data []AgentInfoResData `json:"data"`
}

// AgentInfo 获取代理商账户信息
func (m *Manager) AgentInfo(req AgentInfoReq) (res *AgentInfoRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	// 接收一个数组并转为string格式
	advertiserIds, err := json.Marshal(req.AdvertiserIds)
	if err != nil {
		panic(err)
	}
	fields, err := json.Marshal(req.Fields)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?advertiser_ids=%s&fields=%s",
		conf.API_AGENT_INFO, string(advertiserIds), string(fields)), header, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// AdvertiserPublicInfoReq 获取千川广告账户基础信息-请求
type AdvertiserPublicInfoReq struct {
	AdvertiserIds []int64
	AccessToken   string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type AdvertiserPublicInfoResData struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	Company            string `json:"company"`
	FirstIndustryName  string `json:"first_industry_name"`
	SecondIndustryName string `json:"second_industry_name"`
}

// AdvertiserPublicInfoRes 获取千川广告账户基础信息-返回结构体
type AdvertiserPublicInfoRes struct {
	QCError
	Data []AdvertiserPublicInfoResData `json:"data"`
}

// AdvertiserPublicInfo 获取千川广告账户基础信息
func (m *Manager) AdvertiserPublicInfo(req AdvertiserPublicInfoReq) (res *AdvertiserPublicInfoRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	// 接收一个数组并转为string格式
	advertiserIds, err := json.Marshal(req.AdvertiserIds)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?advertiser_ids=%s",
		conf.API_ADVERTISER_PUBLIC_INFO, string(advertiserIds)), header, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// AdvertiserInfoReq 获取千川广告账户全量信息-请求
type AdvertiserInfoReq struct {
	AdvertiserIds []int64
	Fields        []string
	AccessToken   string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type AdvertiserInfoResData struct {
	ID                      int64  `json:"id"`                        // 广告主ID
	Name                    string `json:"name"`                      // 账户名
	Role                    string `json:"role"`                      // 角色, 详见【附录-广告主角色】
	Status                  string `json:"status"`                    // 状态,详见【附录-广告主状态】
	Address                 string `json:"address"`                   // 地址
	LicenseUrl              string `json:"license_url"`               // 执照预览地址(链接默认1小时内有效)
	LicenseNo               string `json:"license_no"`                // 执照编号
	LicenseProvince         string `json:"license_province"`          // 执照省份
	LicenseCity             string `json:"license_city"`              // 执照城市
	Company                 string `json:"company"`                   // 公司名
	Brand                   string `json:"brand"`                     // 经营类别
	PromotionArea           string `json:"promotion_area"`            // 运营区域
	PromotionCenterProvince string `json:"promotion_center_province"` // 运营省份
	PromotionCenterCity     string `json:"promotion_center_city"`     // 运营城市
	FirstIndustryName       string `json:"first_industry_name"`       // 一级行业名称（新版）
	SecondIndustryName      string `json:"second_industry_name"`      // 二级行业名称（新版）
	Reason                  string `json:"reason"`                    // 审核拒绝原因
	CreateTime              string `json:"create_time"`               // 创建时间
}

// AdvertiserInfoRes 获取千川广告账户全量信息-返回结构体
type AdvertiserInfoRes struct {
	QCError
	Data []AdvertiserInfoResData `json:"data"`
}

// AdvertiserInfo 获取千川广告账户全量信息
func (m *Manager) AdvertiserInfo(req AdvertiserInfoReq) (res *AdvertiserInfoRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	// 接收一个数组并转为string格式
	advertiserIds, err := json.Marshal(req.AdvertiserIds)
	if err != nil {
		panic(err)
	}
	fields, err := json.Marshal(req.Fields)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?advertiser_ids=%s&fields=%s",
		conf.API_ADVERTISER_INFO, string(advertiserIds), string(fields)), header, nil)
	return res, err
}
