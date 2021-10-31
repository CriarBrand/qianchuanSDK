package qianchuanSDK

import (
	"bytes"
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
	AccessToken string `json:"access_token"`
	Body        struct {
		AdvertiserId   int64     `json:"advertiser_id"`
		UploadType     string    `json:"upload_type"`
		ImageSignature string    `json:"image_signature"`
		ImageFile      io.Reader `json:"image_file"`
		ImageUrl       string    `json:"image_url"`
		Filename       string    `json:"filename"`
	} `json:"body"`
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

//----------------------------------------------------------------------------------------------------------------------

// FileVideoAdReq 上传视频素材 的 请求结构体
type FileVideoAdReq struct {
	AccessToken string `json:"access_token"`
	Body        struct {
		AdvertiserId   int64     `json:"advertiser_id"`
		VideoSignature string    `json:"video_signature"`
		VideoFile      io.Reader `json:"video_file"`
		VideoName      string    `json:"-"`
	} `json:"body"`
}

// FileVideoAdRes 上传视频素材 的 响应结构体
type FileVideoAdRes struct {
	QCError
	Data struct {
		VideoId    string `json:"video_id"`
		Size       int64  `json:"size"`
		Width      int64  `json:"width"`
		Height     int64  `json:"height"`
		VideoUrl   string `json:"video_url"`
		Duration   int64  `json:"duration"`
		MaterialId int64  `json:"material_id"`
	} `json:"data"`
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
