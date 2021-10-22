// Token管理相关API
// 巨量开放平台地址：https://open.oceanengine.com/doc/index.html?key=qianchuan&type=api&id=1697468184666115

package qianchuanSDK

import (
	"context"
	"encoding/json"

	"github.com/CriarBrand/qianchuangSDK/conf"
)

// OauthParam 授权参数
type OauthParam struct {
	AppId        int64   // 应用的app_id
	State        string  // 应用的状态，默认”your_custom_params“，暂时不知道其他选项
	Scope        []int64 // 应用权限范围，形如”[20120000,22000000]“
	MaterialAuth string  // 暂时不知道，默认”1“
	RedirectUri  string  // 重定向链接
	Rid          string  // 暂时不知道
}

// OauthConnect 生成授权链接,获取授权码
func (m *Manager) OauthConnect(param OauthParam) string {
	scope, err := json.Marshal(param.Scope)
	if err != nil {
		panic(err)
	}
	return m.url("%s?app_id=%d&state=%s&scope=%s&material_auth=%s&redirect_uri=%s&rid=%s", conf.API_OAUTH_CONNECT,
		param.AppId, param.State, string(scope), param.MaterialAuth, param.RedirectUri, param.Rid)
}

// ---------------------------------------------------------------------------------------------------------------------

// OauthAccessTokenReq access_token请求
type OauthAccessTokenReq struct {
	Code string // 授权码
}

// OauthAccessTokenResData access_token返回
type OauthAccessTokenResData struct {
	AccessToken           string `json:"access_token"`             // 用于验证权限的token
	ExpiresIn             uint64 `json:"expires_in"`               // access_token剩余有效时间,单位(秒)
	RefreshToken          string `json:"refresh_token"`            // 刷新access_token,用于获取新的access_token和refresh_token，并且刷新过期时间
	RefreshTokenExpiresIn uint64 `json:"refresh_token_expires_in"` // refresh_token剩余有效时间,单位(秒)
}

// OauthAccessTokenRes access_token返回结构体
type OauthAccessTokenRes struct {
	QCError
	Data OauthAccessTokenResData `json:"data"`
}

// OauthAccessToken 获取access_token
func (m *Manager) OauthAccessToken(req OauthAccessTokenReq) (res OauthAccessTokenRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?app_id=%d&secret=%s&grant_type=auth_code&auth_code=%s",
		conf.API_OAUTH_ACCESS_TOKEN, m.Credentials.AppId, m.Credentials.AppSecret, req.Code), nil, nil)
	return res, err
}

// ---------------------------------------------------------------------------------------------------------------------

// OauthRefreshTokenReq 刷新access_token请求
type OauthRefreshTokenReq struct {
	RefreshToken string // 填写通过access_token获取到的refresh_token参数
}

// OauthRefreshTokenResData 刷新access_token返回
type OauthRefreshTokenResData OauthAccessTokenResData

// OauthRefreshTokenRes 刷新access_token返回结构体
type OauthRefreshTokenRes struct {
	QCError
	Data OauthRefreshTokenResData `json:"data"`
}

// OauthRefreshToken 刷新access_token
func (m *Manager) OauthRefreshToken(req OauthRefreshTokenReq) (res OauthRefreshTokenRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?app_id=%d&secret=%s&grant_type=auth_code&refresh_token=%s",
		conf.API_OAUTH_REFRESH_TOKEN, m.Credentials.AppId, m.Credentials.AppSecret, req.RefreshToken), nil, nil)
	return res, err
}
