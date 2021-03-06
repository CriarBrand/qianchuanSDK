package qianchuanSDK

import (
	"fmt"
	"net/http"

	"github.com/CriarBrand/qianchuanSDK/auth"
	"github.com/CriarBrand/qianchuanSDK/client"
	"github.com/CriarBrand/qianchuanSDK/conf"
)

// Manager Manager结构体
type Manager struct {
	client      *client.Client
	Credentials *auth.Credentials
}

// NewCredentials 获取认证
func NewCredentials(appId int64, appSecret string) *auth.Credentials {
	return auth.New(appId, appSecret)
}

// NewManager 创建新的Manager
func NewManager(credentials *auth.Credentials, tr http.RoundTripper) *Manager {
	client := client.DefaultClient
	client.Transport = newTransport(credentials, nil)
	return &Manager{
		client:      &client,
		Credentials: credentials,
	}
}

func (m *Manager) url(format string, args ...interface{}) string {
	return conf.API_HTTP_SCHEME + conf.API_HOST + fmt.Sprintf(format, args...)
}
