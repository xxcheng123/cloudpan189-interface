package client

import (
	"context"
)

type GetFileDownloadResponse struct {
	ResCode         int    `json:"res_code"`
	ResMessage      string `json:"res_message"`
	FileDownloadUrl string `json:"fileDownloadUrl"`
}

type GetFileDownloadRequest struct {
	FileId    String `url:"fileId"`
	Short     bool   `url:"short,omitempty"`     // 是否获取短地址， true：获取短地址格式下载URL 为空或false：获取原始下载URL
	forcedGet int8   `url:"forcedGet,omitempty"` // 	1 - 强制获取下载地址（需配置应用权限） 0或空 -普通获取
	ShareId   int64  `url:"shareId,omitempty"`
	Dt        int8   `url:"dt"`
}

type GetFileDownloadOption func(*GetFileDownloadRequest)

func (c *client) GetFileDownload(ctx context.Context, fileId String, opts ...GetFileDownloadOption) (*GetFileDownloadResponse, error) {
	req := &GetFileDownloadRequest{
		FileId: fileId,
		Dt:     1,
	}

	for _, opt := range opts {
		opt(req)
	}

	var (
		result = new(GetFileDownloadResponse)
	)

	if err := c.get(ctx, ApiGetFileDownloadURL, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
