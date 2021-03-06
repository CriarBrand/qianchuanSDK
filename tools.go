// 千川投放相关的部分工具辅助类API
// 巨量开放平台地址：https://open.oceanengine.com/doc/index.html?key=qianchuan&type=api&id=1697459568144388

package qianchuanSDK

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/CriarBrand/qianchuanSDK/conf"
	"net/http"
	"net/url"
)

//----------------------------------------------------------------------------------------------------------------------

// ToolsIndustryGetReq 获取行业列表 的 请求结构体
type ToolsIndustryGetReq struct {
	AccessToken string `json:"access_token"`    // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Level       int64  `json:"level,omitempty"` //只获取某级别数据，1:第一级,2:第二级,3:第三级，默认都返回
	Type        string `json:"type,omitempty"`  //可选值："ADVERTISER"，"AGENT"，"ADVERTISER"为原有广告3.0行业, "AGENT"为代理商行业获取，代理商行业level都为1
}

// ToolsIndustryGetRes 获取行业列表 的 响应结构体
type ToolsIndustryGetRes struct {
	QCError
	Data struct {
		List []ToolsIndustryGetResDetail `json:"list"`
	} `json:"data"`
}
type ToolsIndustryGetResDetail struct {
	IndustryId         int64  `json:"industry_id"`
	IndustryName       string `json:"industry_name"`
	Level              int64  `json:"level"`
	FirstIndustryId    int64  `json:"first_industry_id"`
	FirstIndustryName  string `json:"first_industry_name"`
	SecondIndustryId   int64  `json:"second_industry_id"`
	SecondIndustryName string `json:"second_industry_name"`
	ThirdIndustryId    int64  `json:"third_industry_id"`
	ThirdIndustryName  string `json:"third_industry_name"`
}

// ToolsIndustryGet 获取行业列表
func (m *Manager) ToolsIndustryGet(req ToolsIndustryGetReq) (res *ToolsIndustryGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)

	reqUrl := conf.API_HTTP_SCHEME + conf.API_HOST + conf.API_TOOLS_INDUSTRY_GET
	reqUrl, err = BuildQuery(reqUrl, req, []string{"access_token"})
	if err != nil {
		return nil, err
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", reqUrl, header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// ToolsAwemeCategoryTopAuthorGetReq 查询抖音类目下的推荐达人 的 请求结构体
type ToolsAwemeCategoryTopAuthorGetReq struct {
	AccessToken  string   `json:"access_token"`          // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId int64    `json:"advertiser_id"`         // 广告主ID
	CategoryId   int64    `json:"category_id,omitempty"` // 类目id，一级，二级，三级类目id均可
	Behaviors    []string `json:"behaviors,omitempty"`   // 抖音用户行为类型，详见【附录-抖音达人互动用户行为类型】 默认为空,仅影响覆盖人群数
}

// ToolsAwemeCategoryTopAuthorGetRes 查询抖音类目下的推荐达人 的 响应结构体
type ToolsAwemeCategoryTopAuthorGetRes struct {
	QCError
	Data struct {
		Authors []ToolsAwemeCategoryTopAuthorGetResAuthor `json:"authors"`
	} `json:"data"`
}
type ToolsAwemeCategoryTopAuthorGetResAuthor struct { // 抖音作者名
	AuthorName      string `json:"author_name"`        //抖音作者名
	TotalFansNumStr string `json:"total_fans_num_str"` //粉丝数
	CoverNumStr     string `json:"cover_num_str"`      //覆盖人群数
	LabelId         string `json:"label_id"`           //抖音号id
	AwemeId         string `json:"aweme_id"`           //抖音id
	Avatar          string `json:"avatar"`             //抖音头像
	CategoryName    string `json:"category_name"`      //抖音分类
}

// ToolsAwemeCategoryTopAuthorGet 查询抖音类目下的推荐达人
func (m *Manager) ToolsAwemeCategoryTopAuthorGet(req ToolsAwemeCategoryTopAuthorGetReq) (res *ToolsAwemeCategoryTopAuthorGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)

	reqUrl := conf.API_HTTP_SCHEME + conf.API_HOST + conf.API_TOOLS_AWEME_CATEGORY_TOP_AUTHOR_GET
	reqUrl, err = BuildQuery(reqUrl, req, []string{"access_token"})
	if err != nil {
		return nil, err
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", reqUrl, header, nil)
	//behavior, err := json.Marshal(req.Behaviors)
	//if err != nil {
	//	return nil, err
	//}
	//err = m.client.CallWithJson(context.Background(), &res, "GET",
	//	m.url("%s?advertiser_id=%d&category_id=%d&behaviors=%v",
	//		conf.API_TOOLS_AWEME_CATEGORY_TOP_AUTHOR_GET, req.AdvertiserId, req.CategoryId, string(behavior)), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// ToolsAwemeMultiLevelCategoryGetReq 查询抖音类目列表 的 请求结构体
type ToolsAwemeMultiLevelCategoryGetReq struct {
	AccessToken  string   `json:"access_token"`        // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId int64    `json:"advertiser_id"`       // 广告主ID
	Behaviors    []string `json:"behaviors,omitempty"` // 抖音用户行为类型，详见【附录-抖音达人互动用户行为类型】 默认为空,仅影响覆盖人群数
}

// ToolsAwemeMultiLevelCategoryGetRes 查询抖音类目列表 的 响应结构体
type ToolsAwemeMultiLevelCategoryGetRes struct {
	QCError
	Data struct {
		Categories []ToolsAwemeMultiLevelCategoryGetResCategory `json:"categories"`
	} `json:"data"`
}
type ToolsAwemeMultiLevelCategoryGetResCategory struct { // 抖音作者名
	Id          int64  `json:"id"`
	CoverNumStr string `json:"cover_num_str"`
	FansNumStr  string `json:"fans_num_str"`
	Value       string `json:"value"`
	Children    []struct {
		Id          int64  `json:"id"`
		CoverNumStr string `json:"cover_num_str"`
		FansNumStr  string `json:"fans_num_str"`
		Value       string `json:"value"`
		Children    []struct {
			Id          int64  `json:"id"`
			CoverNumStr string `json:"cover_num_str"`
			FansNumStr  string `json:"fans_num_str"`
			Value       string `json:"value"`
		} `json:"children"`
	} `json:"children"`
}

// ToolsAwemeMultiLevelCategoryGet 查询抖音类目列表
func (m *Manager) ToolsAwemeMultiLevelCategoryGet(req ToolsAwemeMultiLevelCategoryGetReq) (res *ToolsAwemeMultiLevelCategoryGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)

	reqUrl := conf.API_HTTP_SCHEME + conf.API_HOST + conf.API_TOOLS_AWEME_MULTI_LEVEL_CATEGORY_GET
	reqUrl, err = BuildQuery(reqUrl, req, []string{"access_token"})
	if err != nil {
		return nil, err
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", reqUrl, header, nil)

	//behavior, err := json.Marshal(req.Behaviors)
	//if err != nil {
	//	return nil, err
	//}
	//err = m.client.CallWithJson(context.Background(), &res, "GET",
	//	m.url("%s?advertiser_id=%d&behaviors=%v",
	//		conf.API_TOOLS_AWEME_MULTI_LEVEL_CATEGORY_GET, req.AdvertiserId, string(behavior)), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// ToolsInterestActionActionCategoryReq 行为类目查询 的 请求结构体
type ToolsInterestActionActionCategoryReq struct {
	AccessToken  string   `json:"access_token"`  // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId int64    `json:"advertiser_id"` // 广告主ID
	ActionScene  []string `json:"action_scene"`  // 行为场景，详见【附录-行为场景】 允许值: "E-COMMERCE","NEWS","APP"
	ActionDays   int64    `json:"action_days"`   // 行为天数 默认值: 7、15、30、60、90、180、365
}

// ToolsInterestActionActionCategoryRes 行为类目查询 的 响应结构体
type ToolsInterestActionActionCategoryRes struct {
	QCError
	Data []ToolsInterestActionActionCategoryResLevelOne `json:"data"`
}
type ToolsInterestActionActionCategoryResLevelOne struct { // 行为子类目
	Id       string                                         `json:"id"`   //行为类目id
	Num      string                                         `json:"num"`  //数量
	Name     string                                         `json:"name"` //行为类目
	Children []ToolsInterestActionActionCategoryResLevelTwo `json:"children"`
}
type ToolsInterestActionActionCategoryResLevelTwo struct { // 行为子类目
	Id       string                                           `json:"id"`   //行为类目id
	Num      string                                           `json:"num"`  //数量
	Name     string                                           `json:"name"` //行为类目
	Children []ToolsInterestActionActionCategoryResLevelThree `json:"children"`
}
type ToolsInterestActionActionCategoryResLevelThree struct { // 行为子类目
	Id       string                                          `json:"id"`   //行为类目id
	Num      string                                          `json:"num"`  //数量
	Name     string                                          `json:"name"` //行为类目
	Children []ToolsInterestActionActionCategoryResLevelFour `json:"children"`
}
type ToolsInterestActionActionCategoryResLevelFour struct { // 行为子类目
	Id   string `json:"id"`   //行为类目id
	Num  string `json:"num"`  //数量
	Name string `json:"name"` //行为类目
}

// ToolsInterestActionActionCategory 行为类目查询
func (m *Manager) ToolsInterestActionActionCategory(req ToolsInterestActionActionCategoryReq) (res *ToolsInterestActionActionCategoryRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)

	reqUrl := conf.API_HTTP_SCHEME + conf.API_HOST + conf.API_INTEREST_ACTION_ACTION_CATEGORY
	reqUrl, err = BuildQuery(reqUrl, req, []string{})
	if err != nil {
		return nil, err
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", reqUrl, header, nil)

	//actionScene, err := json.Marshal(req.ActionScene)
	//if err != nil {
	//	return nil, err
	//}
	//err = m.client.CallWithJson(context.Background(), &res, "GET",
	//	m.url("%s?advertiser_id=%d&action_scene=%s&action_days=%d",
	//		conf.API_INTEREST_ACTION_ACTION_CATEGORY, req.AdvertiserId, string(actionScene), req.ActionDays), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// ToolsInterestActionActionKeywordReq 行为关键词查询 的 请求结构体
type ToolsInterestActionActionKeywordReq struct {
	AccessToken  string   `json:"access_token"`  // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId int64    `json:"advertiser_id"` // 广告主ID
	QueryWords   string   `json:"query_words"`   // 关键词
	ActionScene  []string `json:"action_scene"`  // 行为场景，详见【附录-行为场景】 允许值: "E-COMMERCE","NEWS","APP"
	ActionDays   int64    `json:"action_days"`   // 行为天数 默认值: 7、15、30、60、90、180、365
}

// ToolsInterestActionActionKeywordRes 行为关键词查询 的 响应结构体
type ToolsInterestActionActionKeywordRes struct {
	QCError
	Data struct {
		List []ToolsInterestActionActionKeywordResDetail `json:"list"`
	} `json:"data"`
}
type ToolsInterestActionActionKeywordResDetail struct { // 词包列表
	Id   string `json:"id"`   //关键词id
	Name string `json:"name"` //关键词名称
	Num  string `json:"num"`  //关键词数目
}

// ToolsInterestActionActionKeyword 行为关键词查询
func (m *Manager) ToolsInterestActionActionKeyword(req ToolsInterestActionActionKeywordReq) (res *ToolsInterestActionActionKeywordRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)

	reqUrl := conf.API_HTTP_SCHEME + conf.API_HOST + conf.API_TOOLS_INTEREST_ACTION_ACTION_KEYWORD
	reqUrl, err = BuildQuery(reqUrl, req, []string{})
	if err != nil {
		return nil, err
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", reqUrl, header, nil)

	//actionScene, err := json.Marshal(req.ActionScene)
	//if err != nil {
	//	return nil, err
	//}
	//err = m.client.CallWithJson(context.Background(), &res, "GET",
	//	m.url("%s?advertiser_id=%d&query_words=%s&action_scene=%s&action_days=%d",
	//		conf.API_TOOLS_INTEREST_ACTION_ACTION_KEYWORD, req.AdvertiserId, url.QueryEscape(req.QueryWords), url.QueryEscape(string(actionScene)), req.ActionDays), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// ToolsInterestActionInterestCategoryReq 兴趣类目查询 的 请求结构体
type ToolsInterestActionInterestCategoryReq struct {
	AccessToken  string `json:"access_token"`  // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId int64  `json:"advertiser_id"` // 广告主ID
}

// ToolsInterestActionInterestCategoryRes 兴趣类目查询 的 响应结构体
type ToolsInterestActionInterestCategoryRes struct {
	QCError
	Data []ToolsInterestActionInterestCategoryResLevelOne `json:"data"`
}
type ToolsInterestActionInterestCategoryResLevelOne struct {
	Id       string                                           `json:"id"`
	Num      string                                           `json:"num"`
	Name     string                                           `json:"name"`
	Children []ToolsInterestActionInterestCategoryResLevelTwo `json:"children"`
}
type ToolsInterestActionInterestCategoryResLevelTwo struct {
	Id       string                                             `json:"id"`
	Num      string                                             `json:"num"`
	Name     string                                             `json:"name"`
	Children []ToolsInterestActionInterestCategoryResLevelThree `json:"children"`
}
type ToolsInterestActionInterestCategoryResLevelThree struct {
	Id       string                                            `json:"id"`
	Num      string                                            `json:"num"`
	Name     string                                            `json:"name"`
	Children []ToolsInterestActionInterestCategoryResLevelFour `json:"children"`
}
type ToolsInterestActionInterestCategoryResLevelFour struct {
	Id   string `json:"id"`
	Num  string `json:"num"`
	Name string `json:"name"`
}

// ToolsInterestActionInterestCategory 兴趣类目查询
func (m *Manager) ToolsInterestActionInterestCategory(req ToolsInterestActionInterestCategoryReq) (res *ToolsInterestActionInterestCategoryRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d",
			conf.API_TOOLS_INTEREST_ACTION_INTEREST_CATEGORY, req.AdvertiserId), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// ToolsInterestActionInterestKeywordReq 兴趣关键词查询 的 请求结构体
type ToolsInterestActionInterestKeywordReq struct {
	AccessToken  string `json:"access_token"`  // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId int64  `json:"advertiser_id"` // 广告主ID
	QueryWords   string `json:"query_words"`   // 关键词
}

// ToolsInterestActionInterestKeywordRes 兴趣关键词查询 的 响应结构体
type ToolsInterestActionInterestKeywordRes struct {
	QCError
	Data struct {
		List []ToolsInterestActionInterestKeywordResDetail `json:"list"`
	} `json:"data"`
}
type ToolsInterestActionInterestKeywordResDetail struct { // 词包列表
	Id   string `json:"id"`   //关键词id
	Name string `json:"name"` //关键词名称
	Num  string `json:"num"`  //关键词数目
}

// ToolsInterestActionInterestKeyword 兴趣关键词查询
func (m *Manager) ToolsInterestActionInterestKeyword(req ToolsInterestActionInterestKeywordReq) (res *ToolsInterestActionInterestKeywordRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d&query_words=%s",
			conf.API_TOOLS_INTEREST_ACTION_INTEREST_KEYWORD, req.AdvertiserId, url.QueryEscape(req.QueryWords)), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// ToolsCreativeWordSelectReq 查询动态创意词包 的 请求结构体
type ToolsCreativeWordSelectReq struct {
	AccessToken     string   `json:"access_token"`      // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId    int64    `json:"advertiser_id"`     // 广告主ID
	CreativeWordIds []string `json:"creative_word_ids"` // 创意词包id列表，如不填默认返回所有创意词包
}

// ToolsCreativeWordSelectRes 查询动态创意词包 的 响应结构体
type ToolsCreativeWordSelectRes struct {
	QCError
	Data struct {
		CreativeWord []ToolsCreativeWordSelectResDetail `json:"creative_word"`
	} `json:"data"`
}
type ToolsCreativeWordSelectResDetail struct { // 词包列表
	CreativeWordId int64    `json:"creative_word_id"`
	Name           string   `json:"name"`
	DefaultWord    string   `json:"default_word"`
	Words          []string `json:"words"`
	ContentType    string   `json:"content_type"`
	MaxWordLen     int64    `json:"max_word_len"`
	Status         string   `json:"status"`
	UserRate       float64  `json:"user_rate"`
}

// ToolsCreativeWordSelect 查询动态创意词包
func (m *Manager) ToolsCreativeWordSelect(req ToolsCreativeWordSelectReq) (res *ToolsCreativeWordSelectRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	params := ""
	if len(req.CreativeWordIds) > 0 {
		creativeWordIds, err := json.Marshal(req.CreativeWordIds)
		if err != nil {
			return nil, err
		}
		params = fmt.Sprintf("&creative_word_ids=%s", creativeWordIds)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d%s",
			conf.API_TOOLS_CREATIVE_WORD_SELECT, req.AdvertiserId, params), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// DmpAudiencesGetReq 查询人群包列表 的 请求结构体
type DmpAudiencesGetReq struct {
	AccessToken         string `json:"access_token"`          // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId        int64  `json:"advertiser_id"`         // 千川广告账户ID
	RetargetingTagsType int64  `json:"retargeting_tags_type"` // 人群包类型，枚举值：0：不限营销目标的平台精选人群包，1：自定义人群包
	Offset              int64  `json:"offset,omitempty"`      // 偏移,类似于SQL中offset(起始为0,翻页时new_offset=old_offset+limit），默认值：0，取值范围:≥ 0
	Limit               int64  `json:"limit,omitempty"`       // 返回数据量，默认值：100，取值范围：1-100
}

// DmpAudiencesGetRes 查询人群包列表 的 响应结构体
type DmpAudiencesGetRes struct {
	QCError
	Data struct {
		RetargetingTags []DmpAudiencesGetResDetail `json:"retargeting_tags"`
		Offset          int64                      `json:"offset"`    //下一次查询的偏移,类似于SQL中offset(起始为0,翻页时new_offset=old_offset+limit），返回0时，代表已查询到最后一页
		TotalNum        int64                      `json:"total_num"` //总的人群包数量
	} `json:"data"`
}
type DmpAudiencesGetResDetail struct { // 人群包列表
	RetargetingTagsId  int64  `json:"retargeting_tags_id"`  //人群包id
	Name               string `json:"name"`                 //人群包名称
	Source             string `json:"source"`               //人群包来源，自定义类详见【附录-DMP相关-人群包来源】，平台精选类返回空值
	Status             int64  `json:"status"`               //人群包状态，详见【附录-DMP相关-人群包状态】
	RetargetingTagsOp  string `json:"retargeting_tags_op"`  //人群包可选的定向规则，枚举值：INCLUDE只支持定向，EXCLUDE只支持排除，ALL支持两种规则。 当source为RETARGETING_TAGS_TYPE_PLATFORM时，只支持INCLUDE或EXCLUDE；当source为RETARGETING_TAGS_TYPE_CUSTOM时，支持ALL
	CoverNum           int64  `json:"cover_num"`            //预估人群包覆盖人群数目
	RetargetingTagsTip string `json:"retargeting_tags_tip"` //人群包说明
	IsCommon           int64  `json:"is_common"`            //0 该人群包不支持通投，1 该人群包支持通投，注意：不支持通投的人群包不能在千川平台创建计划，否则会报错。
}

// DmpAudiencesGet 查询人群包列表
func (m *Manager) DmpAudiencesGet(req DmpAudiencesGetReq) (res *DmpAudiencesGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)

	reqUrl := conf.API_HTTP_SCHEME + conf.API_HOST + conf.API_DMP_AUDIENCES_GET
	reqUrl, err = BuildQuery(reqUrl, req, []string{"access_token"})
	if err != nil {
		return nil, err
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", reqUrl, header, nil)

	//err = m.client.CallWithJson(context.Background(), &res, "GET",
	//	m.url("%s?advertiser_id=%d&retargeting_tags_type=%d&offset=%d&limit=%d",
	//		conf.API_DMP_AUDIENCES_GET, req.AdvertiserId, req.RetargetingTagsType, req.Offset, req.Limit), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

type ToolsAwemeAuthorInfoGetReq struct {
	AccessToken  string   `json:"access_token"`
	AdvertiserId int64    `json:"advertiser_id"`       // 广告主ID
	LabelIds     []int64  `json:"label_ids"`           // 抖音号id列表，取值大小：1～50；label_id即计划中设置的抖音达人账号的id
	Behaviors    []string `json:"behaviors,omitempty"` //抖音用户行为类型 允许值: "FOLLOWED_USER","COMMENTED_USER","LIKED_USER","SHARED_USER"
}

type ToolsAwemeAuthorInfoGetRes struct {
	QCError
	Data ToolsAwemeAuthorInfoGetResData `json:"data"`
}

type ToolsAwemeAuthorInfoGetResData struct {
	Authors []ToolsAwemeAuthorInfoGetResDataAuthor `json:"authors"`
}

type ToolsAwemeAuthorInfoGetResDataAuthor struct {
	AuthorName      string `json:"author_name"`
	LabelId         int64  `json:"label_id"`
	CoverNumStr     string `json:"cover_num_str"`
	TotalFansNumStr string `json:"total_fans_num_str"`
	Avatar          string `json:"avatar"`
	AwemeId         string `json:"aweme_id"`
}

// ToolsAwemeAuthorInfoGet 查询抖音号id对应的达人信息(巨量引擎)
func (m *Manager) ToolsAwemeAuthorInfoGet(req ToolsAwemeAuthorInfoGetReq) (res *ToolsAwemeAuthorInfoGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)

	reqUrl := conf.API_HTTP_SCHEME + conf.API_HOST + conf.API_TOOLS_AWEME_AUTHOR_INFO_GET
	reqUrl, err = BuildQuery(reqUrl, req, []string{"access_token"})
	if err != nil {
		return nil, err
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET", reqUrl, header, nil)
	return res, err
}
