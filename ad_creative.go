// 广告创意管理相关API
// 巨量开放平台地址：https://open.oceanengine.com/doc/index.html?key=qianchuan&type=api&id=1697466268841999

package qianchuanSDK

import (
	"context"
	"encoding/json"
	"github.com/CriarBrand/qianchuanSDK/conf"
	"net/http"
)

//----------------------------------------------------------------------------------------------------------------------

// CreativeStatusUpdateReq 更新创意状态 的 请求结构体
type CreativeStatusUpdateReq struct {
	AccessToken string                      `json:"access_token"`
	Body        CreativeStatusUpdateReqBody `json:"body"`
}
type CreativeStatusUpdateReqBody struct {
	AdvertiserId int64   `json:"advertiser_id"`
	OptStatus    string  `json:"opt_status"`
	CreativeIds  []int64 `json:"creative_ids"`
}

// CreativeStatusUpdateRes 更新创意状态 的 响应结构体
type CreativeStatusUpdateRes struct {
	QCError
	Data struct {
		CreativeIds []int64    `json:"creative_ids"` //更新成功的创意id
		Errors      []struct { //更新失败的创意id和失败原因
			CreativeId   int64  `json:"creative_id"`   //更新失败的创意id
			ErrorMessage string `json:"error_message"` //更新失败的原因
		} `json:"errors"`
	} `json:"data"`
}

// CreativeStatusUpdate 更新创意状态
func (m *Manager) CreativeStatusUpdate(req CreativeStatusUpdateReq) (res *CreativeStatusUpdateRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	err = m.client.CallWithJson(context.Background(), &res, "POST",
		m.url("%s?", conf.API_AD_CREATIVE_UPDATE), header, req.Body)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// CreativeGetReq 获取账户下创意列表 的 请求结构体
type CreativeGetReq struct {
	AccessToken  string                  `json:"access_token"`
	AdvertiserId int64                   `json:"advertiser_id"` //千川广告账户ID
	Filtering    CreativeGetReqFiltering `json:"filtering"`
	Page         int64                   `json:"page"`      //页码，默认值：1
	PageSize     int64                   `json:"page_size"` //页面大小，允许值：10, 20, 50, 100, 500, 1000，默认值：10
}
type CreativeGetReqFiltering struct {
	AdIds                   []int64 `json:"ad_ids"`                     //按计划ID过滤，list长度限制 1-100
	CreativeId              int64   `json:"creative_id"`                //按创意ID过滤
	CreativeMaterialMode    string  `json:"creative_material_mode"`     //按创意呈现方式过滤，允许值： CUSTOM_CREATIVE 自定义创意、PROGRAMMATIC_CREATIVE 程序化创意
	Status                  string  `json:"status"`                     //按创意状态过滤，不传入即默认返回“所有不包含已删除”，其他规则详见【附录-创意查询状态】
	MarketingGoal           string  `json:"marketing_goal"`             //按营销目标过滤，允许值：VIDEO_PROM_GOODS 短视频带货、LIVE_PROM_GOODS 直播带货
	CampaignId              int64   `json:"campaign_id"`                //按广告组ID过滤
	CreativeCreateStartDate string  `json:"creative_create_start_date"` //创意创建开始时间，格式："yyyy-mm-dd"
	CreativeCreateEndDate   string  `json:"creative_create_end_date"`   //创意创建结束时间，与creative_create_start_date搭配使用，格式："yyyy-mm-dd"，时间跨度不能超过180天
	CreativeModifyTime      string  `json:"creative_modify_time"`       //创意修改时间，格式："yyyy-mm-dd HH"
}

// CreativeGetRes 获取账户下创意列表 的 响应结构体
type CreativeGetRes struct {
	QCError
	Data struct {
		List []struct { //创意列表
			AdId               int64    `json:"ad_id"`                //计划ID
			CreativeId         int64    `json:"creative_id"`          //创意ID
			Status             string   `json:"status"`               //创意状态
			OptStatus          string   `json:"opt_status"`           //创意操作状态
			ImageMode          string   `json:"image_mode"`           //创意素材类型
			CreativeCreateTime string   `json:"creative_create_time"` //创意创建时间
			CreativeModifyTime string   `json:"creative_modify_time"` //创意修改时间
			VideoMaterial      struct { //视频素材信息
				VideoId        string `json:"video_id"`         //视频ID
				VideoCoverId   string `json:"video_cover_id"`   //视频封面ID
				AwemeItemId    int64  `json:"aweme_item_id"`    //抖音视频ID
				IsAutoGenerate int64  `json:"is_auto_generate"` //是否为派生创意标识，1：是，0：不是
			} `json:"video_material"`
			ImageMaterial struct { //图片素材信息
				ImageIds       []string `json:"image_ids"`        //图片ID列表
				IsAutoGenerate int64    `json:"is_auto_generate"` //是否为派生创意标识，1：是，0：不是
			} `json:"image_material"`
			TitleMaterial struct { //标题素材信息
				Title        string     `json:"title"` //创意标题
				DynamicWords []struct { //动态词包对象列表
					WordId      int64  //动态词包ID
					DictName    string //创意词包名称
					DefaultWord string //创意词包默认词
				} `json:"dynamic_words"`
			} `json:"title_material"`
		} `json:"list"`
		PageInfo struct { //页面信息
			Page        int64 `json:"page"`         //页码
			PageSize    int64 `json:"page_size"`    //页面大小
			TotalNumber int64 `json:"total_number"` //总数
			TotalPage   int64 `json:"total_page"`   //总页数
		} `json:"page_info"`
	} `json:"data"`
}

// CreativeGet 获取账户下创意列表
func (m *Manager) CreativeGet(req CreativeGetReq) (res *CreativeGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	if err != nil {
		panic(err)
	}
	filtering, err := json.Marshal(req.Filtering)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d&filtering=%s&page=%d&page_size=%d",
			conf.API_AD_CREATIVE_GET, req.AdvertiserId, string(filtering), req.Page, req.PageSize), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// CreativeRejectReasonReq 获取创意审核建议 的 请求结构体
type CreativeRejectReasonReq struct {
	AccessToken  string  `json:"access_token"`
	AdvertiserId int64   `json:"advertiser_id"`
	CreativeIds  []int64 `json:"creative_ids"`
}

// CreativeRejectReasonRes 获取创意审核建议 的 响应结构体
type CreativeRejectReasonRes struct {
	QCError
	Data struct {
		List []struct { //审核详细信息
			CreativeId   int64      `json:"creative_id"` //广告创意id
			AuditRecords []struct { //审核详细内容
				Desc          string   `json:"desc"`           //审核内容，即审核的内容类型，如 视频，图片，标题 等
				Content       string   `json:"content"`        //拒绝内容（文字类型）
				ImageId       int64    `json:"image_id"`       //拒绝内容（图片类型）
				VideoId       int64    `json:"video_id"`       //拒绝内容（视频类型）
				AuditPlatform string   `json:"audit_platform"` //审核来源类型，返回值： AD 广告审核、CONTENT 内容审核
				RejectReason  []string `json:"reject_reason"`  //拒绝原因，可能会有多条
				Suggestion    []string `json:"suggestion"`     //审核建议，可能会有多条
			} `json:"audit_records"` //
		} `json:"list"`
	} `json:"data"`
}

// CreativeRejectReason 获取创意审核建议
func (m *Manager) CreativeRejectReason(req CreativeRejectReasonReq) (res *CreativeRejectReasonRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	if err != nil {
		panic(err)
	}
	creativeIds, err := json.Marshal(req.CreativeIds)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d&creative_ids=%s",
			conf.API_AD_CREATIVE_REJECT, req.AdvertiserId, string(creativeIds)), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------
