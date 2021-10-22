package qianchuanSDK

import "fmt"

// QCError 错误结构体
type QCError struct {
	Code      int64  `json:"code"`                 // 错误码
	Message   string `json:"message"`              // 错误码描述
	RequestId string `json:"request_id,omitempty"` // 错误码描述r
}

func (e *QCError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// NewError 新建错误结构体
func NewError(errorCode int64, description string) *QCError {
	return &QCError{
		Code:    errorCode,
		Message: description,
	}
}
