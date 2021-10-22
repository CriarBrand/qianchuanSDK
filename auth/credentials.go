package auth

// Credentials 认证结构体
type Credentials struct {
	AppId     int64
	AppSecret string
}

// New 新的认证
func New(appId int64, appSecret string) *Credentials {
	return &Credentials{
		appId,
		appSecret,
	}
}
