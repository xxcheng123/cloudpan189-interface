package client

import (
	"context"
)

type GetUpResourceShareRequest struct {
	IconOption string `url:"iconOption"`
	PageNum    int64  `url:"pageNum"`
	PageSize   int64  `url:"pageSize"`
	UpUserId   string `url:"upUserId"`
	FileName   string `url:"fileName"`
	OrderBy    string `url:"orderBy"`
	Descending bool   `url:"descending"`
}

type GetUpResourceShareRequestOption func(*GetUpResourceShareRequest)

type GetUpResourceShareResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    *ShareFileList `json:"data"`
}

type AccessCount struct {
	PreviewCount  int `json:"previewCount"`
	CopyCount     int `json:"copyCount"`
	DownloadCount int `json:"downloadCount"`
}

type ShareFileInfo struct {
	LastOpTime   string      `json:"lastOpTime"`
	Heat         int         `json:"heat"`
	ShareDate    string      `json:"shareDate"`
	AccessCount  AccessCount `json:"accessCount"`
	Rev          string      `json:"rev"`
	AccessURL    string      `json:"accessURL"`
	ExpireType   int         `json:"expireType"`
	ShareId      int64       `json:"shareId"`
	RedTip       bool        `json:"redTip,omitempty"`
	TopTime      string      `json:"topTime"`
	ShareType    int         `json:"shareType"`
	ExpireTime   int         `json:"expireTime"`
	Folder       int         `json:"folder"`
	Size         int64       `json:"size"`
	IsTop        int         `json:"isTop"`
	Name         string      `json:"name"`
	ReviewStatus int         `json:"reviewStatus"`
	Id           String      `json:"id"`
	CreateDate   string      `json:"createDate"`
	Status       int         `json:"status"`
	Md5          string      `json:"md5"`
}

type ShareFileList struct {
	Count    int64            `json:"count"`
	FileList []*ShareFileInfo `json:"fileList"`
}

func (c *client) GetUpResourceShare(ctx context.Context, upUserId string, pageNum int64, pageSize int64, opts ...GetUpResourceShareRequestOption) (*GetUpResourceShareResponse, error) {
	req := &GetUpResourceShareRequest{
		UpUserId:   upUserId,
		PageNum:    pageNum,
		PageSize:   pageSize,
		OrderBy:    "lastOpTime",
		Descending: true,
	}

	for _, opt := range opts {
		opt(req)
	}

	var (
		result = new(GetUpResourceShareResponse)
	)
	if err := c.get(ctx, ApiGetUpResourceShare, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
