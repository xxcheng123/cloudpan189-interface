package client

import (
	"context"
)

type GetFileInfoResponse struct {
	ResCode         int    `json:"res_code"`
	ResMessage      string `json:"res_message"`
	CreateDate      string `json:"createDate"`
	FileDownloadUrl string `json:"fileDownloadUrl"`
	FilePath        string `json:"filePath"`
	Id              String `json:"id"`
	LastOpTime      int64  `json:"lastOpTime"`
	LastOpTimeStr   string `json:"lastOpTimeStr"`
	Md5             string `json:"md5"`
	Name            string `json:"name"`
	ParentId        String `json:"parentId"`
	Rev             int64  `json:"rev"`
	Size            int    `json:"size"`
}

type GetFileInfoRequest struct {
	FileId  String `url:"fileId"`
	Short   bool   `url:"short,omitempty"` // 是否获取短地址， true：获取短地址格式下载URL 为空或false：获取原始下载URL
	ShareId int64  `url:"shareId,omitempty"`
	Dt      int8   `url:"dt"`
}

type GetFileFileOption func(*GetFileInfoRequest)

func (c *client) GetFileInfo(ctx context.Context, fileId String, opts ...GetFileFileOption) (*GetFileInfoResponse, error) {
	req := &GetFileInfoRequest{
		FileId: fileId,
		Dt:     1,
	}

	for _, opt := range opts {
		opt(req)
	}

	var (
		result = new(GetFileInfoResponse)
	)

	if err := c.get(ctx, ApiGetFileInfo, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
