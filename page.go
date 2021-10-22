package qianchuanSDK

// PageInfo 页码结构体
type PageInfo struct {
	Page        uint64 `json:"page"`         // 页数
	PageSize    uint64 `json:"page_size"`    // 页面大小
	TotalNumber uint64 `json:"total_number"` // 总数
	TotalPage   uint64 `json:"total_page"`   // 总页数
}
