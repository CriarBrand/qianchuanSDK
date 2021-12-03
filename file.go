// 通过接口上传/获取广告投放流程中必须的 图片/视频 等创意素材API
// 巨量开放平台地址：https://open.oceanengine.com/doc/index.html?key=qianchuan&type=api&id=1697466636262411

package qianchuanSDK

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/CriarBrand/qianchuanSDK/conf"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------------------------------------------------

// FileImageAdReq 上传图片素材 的 请求结构体
type FileImageAdReq struct {
	AccessToken string             `json:"access_token"`
	Body        FileImageAdReqBody `json:"body"`
}
type FileImageAdReqBody struct {
	AdvertiserId   int64     `json:"advertiser_id,omitempty"`
	UploadType     string    `json:"upload_type,omitempty"`
	ImageSignature string    `json:"image_signature,omitempty"`
	ImageFile      io.Reader `json:"image_file,omitempty"`
	ImageUrl       string    `json:"image_url,omitempty"`
	Filename       string    `json:"filename,omitempty"`
}

// FileImageAdRes 上传图片素材 的 响应结构体
type FileImageAdRes struct {
	QCError
	Data struct {
		Id         string `json:"id"`
		Size       int64  `json:"size"`
		Width      int64  `json:"width"`
		Height     int64  `json:"height"`
		Url        string `json:"url"`
		Format     string `json:"format"`
		Signature  string `json:"signature"`
		MaterialId int64  `json:"material_id"`
	} `json:"data"`
}

// FileImageAd 上传图片素材
func (m *Manager) FileImageAd(req FileImageAdReq) (res *FileImageAdRes, err error) {
	if req.Body.UploadType == "UPLOAD_BY_URL" {
		header := http.Header{}
		header.Add("Access-Token", req.AccessToken)
		err = m.client.CallWithJson(context.Background(), &res, "POST",
			m.url("%s?", conf.API_FILE_IMAGE_AD), header, req.Body)
		return res, err
	} else {
		var (
			HttpClient = &http.Client{}
		)
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		formFile, err := writer.CreateFormFile("image_file", req.Body.Filename)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(formFile, req.Body.ImageFile)
		if err != nil {
			return nil, err
		}
		_ = writer.WriteField("advertiser_id", strconv.Itoa(int(req.Body.AdvertiserId)))
		_ = writer.WriteField("upload_type", req.Body.UploadType)
		_ = writer.WriteField("image_signature", req.Body.ImageSignature)
		_ = writer.WriteField("image_url", req.Body.ImageUrl)
		_ = writer.WriteField("filename", req.Body.Filename)

		request, err := http.NewRequest("POST", conf.API_FILE_IMAGE_AD, body)
		if err != nil {
			return nil, err
		}
		request.Header.Add("Access-Token", req.AccessToken)
		request.Header.Add("Content-Type", writer.FormDataContentType())

		resp, err := HttpClient.Do(request)
		if err != nil {
			fmt.Println("resp err: ", err)
			return nil, err
		}
		defer resp.Body.Close()
		respData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(respData, res)
		if err != nil {
			return nil, err
		}
		return res, err
	}
}

//----------------------------------------------------------------------------------------------------------------------

// FileVideoAdReq 上传视频素材 的 请求结构体
type FileVideoAdReq struct {
	AccessToken string             `json:"access_token"`
	Body        FileVideoAdReqBody `json:"body"`
}
type FileVideoAdReqBody struct {
	AdvertiserId   int64     `json:"advertiser_id"`
	VideoSignature string    `json:"video_signature"`
	VideoFile      io.Reader `json:"video_file"`
	VideoName      string    `json:"filename"`
}

// FileVideoAdRes 上传视频素材 的 响应结构体
type FileVideoAdRes struct {
	QCError
	Data FileVideoAdResDetail `json:"data"`
}
type FileVideoAdResDetail struct {
	VideoId    string `json:"video_id"`
	Size       int64  `json:"size"`
	Width      int64  `json:"width"`
	Height     int64  `json:"height"`
	VideoUrl   string `json:"video_url"`
	Duration   int64  `json:"duration"`
	MaterialId int64  `json:"material_id"`
}

// FileVideoAd 上传视频素材
func (m *Manager) FileVideoAd(req FileVideoAdReq) (res *FileVideoAdRes, err error) {
	var (
		HttpClient = &http.Client{}
	)
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile("video_file", req.Body.VideoName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(formFile, req.Body.VideoFile)
	if err != nil {
		return nil, err
	}
	_ = writer.WriteField("advertiser_id", strconv.Itoa(int(req.Body.AdvertiserId)))
	_ = writer.WriteField("video_signature", req.Body.VideoSignature)

	request, err := http.NewRequest("POST", conf.API_FILE_VIDEO_AD, body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Access-Token", req.AccessToken)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := HttpClient.Do(request)
	if err != nil {
		fmt.Println("resp err: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respData, res)
	if err != nil {
		return nil, err
	}
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// FileImageGetReq 获取素材库的图片 的 请求结构体
type FileImageGetReq struct {
	AccessToken  string                   // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId int64                    `json:"advertiser_id"` // 千川广告账户ID
	Filtering    FileImageGetReqFiltering `json:"filtering,omitempty"`
	Page         int64                    `json:"page,omitempty"`
	PageSize     int64                    `json:"page_size,omitempty"`
}
type FileImageGetReqFiltering struct {
	ImageIds    []string  `json:"image_ids,omitempty"`    //图片ids，可以根据图片ids（创意中使用的图片key，存在一张图片对应多个image_ids的情况）进行过滤 数量限制：<=100 注意：image_ids、material_ids、signatures只能选择一个进行过滤
	MaterialIds []int64   `json:"material_ids,omitempty"` //素材id列表，可以根据material_ids（素材报表使用的id，一个素材唯一对应一个素材id）进行过滤 数量限制：<=100 注意：image_ids、material_ids、signatures只能选择一个进行过滤
	Signatures  []string  `json:"signatures,omitempty"`   //md5值列表，可以根据素材的md5进行过滤 数量限制：<=100 注意：image_ids、material_ids、signatures只能选择一个进行过滤
	Width       int64     `json:"width,omitempty"`        //图片宽度
	Height      int64     `json:"height,omitempty"`       //图片高度
	Ratio       []float64 `json:"ratio,omitempty"`        //图片宽高比，eg: [1.7, 2.5]，输入1.7则搜索满足宽高比介于1.65-1.75之间的图片，即精度上下浮动0.05
	StartTime   string    `json:"start_time,omitempty"`   //根据视频上传时间进行过滤的起始时间，与end_time搭配使用，格式：yyyy-mm-dd
	EndTime     string    `json:"end_time,omitempty"`     //根据视频上传时间进行过滤的截止时间，与start_time搭配使用，格式：yyyy-mm-dd
}

// FileImageGetRes 获取素材库的图片 的 响应结构体
type FileImageGetRes struct {
	QCError
	Data struct {
		List     []FileImageGetResDetail
		PageInfo PageInfo `json:"page_info"`
	}
}
type FileImageGetResDetail struct { //图片列表
	Id         string `json:"id"`          //图片ID
	MaterialId int64  `json:"material_id"` //素材id，即多合一报表中的素材id，一个素材唯一对应一个素材id
	Size       int64  `json:"size"`        //图片大小
	Width      int64  `json:"width"`       //图片宽度
	Height     int64  `json:"height"`      //图片高度
	Url        string `json:"url"`         //图片预览地址(1小时内有效)，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”
	Format     string `json:"format"`      //图片格式
	Signature  string `json:"signature"`   //图片md5
	CreateTime string `json:"create_time"` //素材的上传时间，格式："yyyy-mm-dd HH:MM:SS"
	Filename   string `json:"filename"`    //素材的文件名
}

// FileImageGet 获取素材库的图片
func (m *Manager) FileImageGet(req FileImageGetReq) (res *FileImageGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	// 接收一个结构体并转为string格式
	filtering, err := json.Marshal(req.Filtering)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d&filtering=%s&page=%d&page_size=%d",
			conf.API_FILE_IMAGE_GET, req.AdvertiserId, string(filtering), req.Page, req.PageSize), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// FileVideoGetReq 获取素材库的视频 的 请求结构体
type FileVideoGetReq struct {
	AccessToken  string                   // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId int64                    `json:"advertiser_id"` // 千川广告账户ID
	Filtering    FileVideoGetReqFiltering `json:"filtering,omitempty"`
	Page         int64                    `json:"page,omitempty"`
	PageSize     int64                    `json:"page_size,omitempty"`
}
type FileVideoGetReqFiltering struct {
	Width       int64     `json:"width,omitempty"`        //视频宽度
	Height      int64     `json:"height,omitempty"`       //视频高度
	Ratio       []float64 `json:"ratio,omitempty"`        //视频宽高比，示例: [1.7, 2.5] 输入1.7则搜索满足宽高比介于1.65-1.75之间的视频，即精度上下浮动0.5
	VideoIds    []string  `json:"video_ids,omitempty"`    //视频ids，示例: ["86adb23eaa21229fc04ef932b5089bb8"] 数量限制：<=100 注意：video_ids、material_ids、signatures只能选择一个进行过滤
	MaterialIds []int64   `json:"material_ids,omitempty"` //素材id列表，可以根据material_ids（素材报表使用的id，一个素材唯一对应一个素材id）进行过滤 数量限制：<=100 注意：video_ids、material_ids、signatures只能选择一个进行过滤
	Signatures  []string  `json:"signatures,omitempty"`   //md5值列表，可以根据素材的md5进行过滤 数量限制：<=100 注意：video_ids、material_ids、signatures只能选择一个进行过滤
	StartTime   string    `json:"start_time,omitempty"`   //根据视频上传时间进行过滤的起始时间，与end_time搭配使用，格式：yyyy-mm-dd
	EndTime     string    `json:"end_time,omitempty"`     //根据视频上传时间进行过滤的截止时间，与start_time搭配使用，格式：yyyy-mm-dd
}

// FileVideoGetRes 获取素材库的视频 的 响应结构体
type FileVideoGetRes struct {
	QCError
	Data struct {
		List     []FileVideoGetResDetail
		PageInfo PageInfo `json:"page_info"`
	}
}
type FileVideoGetResDetail struct { //素材列表
	Id         string   `json:"id"`          //视频ID
	Size       int64    `json:"size"`        //视频大小
	Width      int64    `json:"width"`       //视频宽度
	Height     int64    `json:"height"`      //视频高度
	Url        string   `json:"url"`         //视频地址，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”，链接1小时过期
	Format     string   `json:"format"`      //视频格式
	Signature  string   `json:"signature"`   //视频md5值
	PosterUrl  string   `json:"poster_url"`  //视频首帧截图，仅限同主体进行素材预览查看，若非同主体会返回“素材所属主体与开发者主体不一致无法获取URL”，链接1小时过期
	BitRate    int64    `json:"bit_rate"`    //码率，单位bps
	Duration   float64  `json:"duration"`    //视频时长
	MaterialId int64    `json:"material_id"` //素材id，即多合一报表中的素材id，一个素材唯一对应一个素材id
	Source     string   `json:"source"`      //素材来源，E_COMMERCE:巨量千川，BP:巨量纵横， STAR:星图&即合， ARTHUR:亚瑟， VIDEO_CAPTURE:易拍
	CreateTime string   `json:"create_time"` //素材的上传时间，格式："yyyy-mm-dd HH:MM:SS"
	Filename   string   `json:"filename"`    //素材的文件名
	Labels     []string `json:"labels"`      //视频标签
}

// FileVideoGet 获取素材库的视频
func (m *Manager) FileVideoGet(req FileVideoGetReq) (res *FileVideoGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	// 接收一个结构体并转为string格式
	filtering, err := json.Marshal(req.Filtering)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d&filtering=%s&page=%d&page_size=%d",
			conf.API_FILE_VIDEO_GET, req.AdvertiserId, string(filtering), req.Page, req.PageSize), header, nil)
	return res, err
}

//----------------------------------------------------------------------------------------------------------------------

// FileVideoAwemeGetReq 获取抖音号下的视频 的 请求结构体
type FileVideoAwemeGetReq struct {
	AccessToken  string                        `json:"access_token"`  // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId int64                         `json:"advertiser_id"` // 千川广告账户ID
	AwemeId      int64                         `json:"aweme_id"`      // 需拉取视频的抖音号
	Filtering    FileVideoAwemeGetReqFiltering `json:"filtering,omitempty"`
	Cursor       int64                         `json:"cursor,omitempty"`
	Count        int64                         `json:"count,omitempty"`
}
type FileVideoAwemeGetReqFiltering struct {
	ProductId int64 `json:"product_id,omitempty"` //商品ID，查询关联商品的相应视频，仅短视频带货场景需入参
}

// FileVideoAwemeGetRes 获取抖音号下的视频 的 响应结构体
type FileVideoAwemeGetRes struct {
	QCError
	Data struct {
		List     []FileVideoAwemeGetResDetail
		PageInfo struct { //分页信息
			HasMore int64 `json:"has_more"` //是否有下一页
			Count   int64 `json:"count"`    //过滤后返回的视频数量，注意，此处的数量不一定与入参的count一致，因为存在过滤逻辑
			Cursor  int64 `json:"cursor"`   //下一次分页拉取的游标值
		} `json:"page_info"`
	}
}
type FileVideoAwemeGetResDetail struct { //视频素材列表
	AwemeItemId   string `json:"aweme_item_id"`   //抖音短视频 ID
	VideoCoverUrl string `json:"video_cover_url"` //视频封面图片url
	Width         int64  `json:"width"`           //视频宽度
	Height        int64  `json:"height"`          //视频高度
	Url           string `json:"url"`             //视频地址，链接1小时过期
	Duration      int64  `json:"duration"`        //视频时长
	Title         string `json:"title"`           //抖音中的视频标题
	IsRecommend   int64  `json:"is_recommend"`    //是否推荐 0 不推荐 1 推荐
}

// FileVideoAwemeGet 获取抖音号下的视频
func (m *Manager) FileVideoAwemeGet(req FileVideoAwemeGetReq) (res *FileVideoAwemeGetRes, err error) {
	header := http.Header{}
	header.Add("Access-Token", req.AccessToken)
	// 接收一个结构体并转为string格式
	filtering, err := json.Marshal(req.Filtering)
	if err != nil {
		panic(err)
	}
	err = m.client.CallWithJson(context.Background(), &res, "GET",
		m.url("%s?advertiser_id=%d&aweme_id=%d&filtering=%s&cursor=%d&count=%d",
			conf.API_FILE_VIDEO_AWEME_GET, req.AdvertiserId, req.AwemeId, string(filtering), req.Cursor, req.Count), header, nil)
	return res, err
}
