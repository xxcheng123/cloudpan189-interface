package client

import (
	"context"
)

type ListShareFileRequest struct {
	PageNum        int    `url:"pageNum"`
	PageSize       int    `url:"pageSize"`
	FileId         String `url:"fileId"`
	ShareDirFileId String `url:"shareDirFileId"`
	IsFolder       bool   `url:"isFolder"`
	ShareId        int64  `url:"shareId"`
	ShareMode      int    `url:"shareMode"`
	IconOption     int    `url:"iconOption"`
	OrderBy        string `url:"orderBy"`
	Descending     bool   `url:"descending"`
	AccessCode     string `url:"accessCode"`
}

type ListShareFileResponse struct {
	ResCode    int    `json:"res_code"`
	ResMessage string `json:"res_message"`
	ExpireTime int    `json:"expireTime"`
	ExpireType int    `json:"expireType"`
	FileListAO struct {
		Count        int64        `json:"count"`
		FileList     []FileInfo   `json:"fileList"`
		FileListSize int64        `json:"fileListSize"`
		FolderList   []FolderInfo `json:"folderList"`
	} `json:"fileListAO"`
	LastRev int64 `json:"lastRev"`
}

type WithListShareFileRequestOption func(*ListShareFileRequest)

func (c *client) ListShareDir(ctx context.Context, shareId int64, fileId String, opts ...WithListShareFileRequestOption) (*ListFilesResponse, error) {
	req := &ListShareFileRequest{
		PageNum:        1,
		PageSize:       100,
		FileId:         fileId,
		ShareDirFileId: fileId,
		ShareId:        shareId,
		ShareMode:      5,
		IconOption:     5,
		OrderBy:        "lastOpTime",
		IsFolder:       true,
		Descending:     true,
		AccessCode:     "",
	}

	for _, opt := range opts {
		opt(req)
	}

	var (
		result = new(ListFilesResponse)
	)

	if err := c.get(ctx, ApiListShareDir, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
