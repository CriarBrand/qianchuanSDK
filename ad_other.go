// 商品/直播间管理相关API
// 巨量开放平台地址：https://open.oceanengine.com/doc/index.html?key=qianchuan&type=api&id=1697466297171983

package qianchuanSDK

import (
	"context"
	"github.com/CriarBrand/qianchuanSDK/conf"
	"net/http"
)

//----------------------------------------------------------------------------------------------------------------------

// ProductAvailableGetReq 获取可投商品列表接口 的 请求结构体
type ProductAvailableGetReq struct {
	AccessToken  string `json:"access_token"`
	AdvertiserId int64  `json:"advertiser_id,omitempty"`
	Filter       ProductAvailableGetFilter
	Page         int64 `json:"page,omitempty"`
	PageSize     int64 `json:"page_size,omitempty"`
}
type ProductAvailableGetFilter struct {
	ProductIds  []string `json:"access_token,omitempty"`
	ProductName string   `json:"product_name,omitempty"`
}

// ProductAvailableGetRes 获取可投商品列表接口 的 响应结构体
type ProductAvailableGetRes struct {
	QCError
	Data struct {
		PageInfo    PageInfo   `json:"page_info"`
		ProductList []struct { //商品列表
			CategoryName  string  `json:"category_name"`  //分类
			DiscountPrice float64 `json:"discount_price"` //售价，单位：元
			Id            int64   `json:"id"`             //商品id
			Img           string  `json:"img"`            //主图
			Inventory     int64   `json:"inventory"`      //库存
			Name          string  `json:"name"`           //商品名称
			ProductRate   float64 `json:"product_rate"`   //好评率
			SaleTime      string  `json:"sale_time"`      //上架时间
		} `json:"product_list"`
	} `json:"data"`
}

// ProductAvailableGet 获取可投商品列表接口
func (m *Manager) ProductAvailableGet(req ProductAvailableGetReq) (res *ProductAvailableGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d&page=%d&page_size=%d",
			conf.API_PRODUCT_AVAILABLE_GET, req.AdvertiserId, req.Page, req.PageSize), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// AwemeAuthorizedGetReq 获取千川账户下已授权抖音号 的 请求结构体
type AwemeAuthorizedGetReq struct {
	AccessToken  string `json:"access_token"`
	AdvertiserId int64  `json:"advertiser_id,omitempty"`
	Page         int64  `json:"page,omitempty"`
	PageSize     int64  `json:"page_size,omitempty"`
}

// AwemeAuthorizedGetRes 获取千川账户下已授权抖音号 的 响应结构体
type AwemeAuthorizedGetRes struct {
	QCError
	Data struct {
		PageInfo    PageInfo   `json:"page_info"`
		AwemeIdList []struct { //抖音号列表
			AwemeAvatar string   `json:"aweme_avatar"`  //抖音头像
			AwemeId     int64    `json:"aweme_id"`      //抖音id，用于创建计划，拉取抖音号视频素材时入参
			AwemeShowId string   `json:"aweme_show_id"` //抖音号，即客户在手机端上看到的抖音号，若向客户披露抖音号请使用该字段
			AwemeName   string   `json:"aweme_name"`    //抖音号名称
			AwemeStatus string   `json:"aweme_status"`  //抖音号带货状态，返回值： NORMAL可以正常投放 ANCHOR_FORBID带货口碑分过低，暂时无法创建计划 ANCHOR_REACH_UPPER_LIMIT_TODAY带货分过低或暂无带货分，可以创建计划，但无法产生消耗，带货分恢复正常后可正常消耗
			BindType    []string `json:"bind_type"`     //抖音号关系类型
		} `json:"aweme_id_list"`
	} `json:"data"`
}

// AwemeAuthorizedGet 获取千川账户下已授权抖音号
func (m *Manager) AwemeAuthorizedGet(req AwemeAuthorizedGetReq) (res *AwemeAuthorizedGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d&page=%d&page_size=%d",
			conf.API_AWEME_AUTHORIZED_GET, req.AdvertiserId, req.Page, req.PageSize), header, nil)
	return res, err
}
