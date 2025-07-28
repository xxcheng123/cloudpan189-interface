package client

import (
	"context"
)

type GetFolderInfoResponse struct {
	ResCode       int    `json:"res_code"`
	ResMessage    string `json:"res_message"`
	CreateDate    string `json:"createDate"`    // 文件夹创建日期(YYYY-MM-DD hh:mm:ss)
	CreateTime    int64  `json:"createTime"`    // 文件夹创建时间(13位时间戳)
	FileId        String `json:"fileId"`        // 文件夹ID
	FileName      string `json:"fileName"`      // 文件夹名称
	FilePath      string `json:"filePath"`      // 文件夹路径
	LastOpTime    int64  `json:"lastOpTime"`    // 上次修改时间(13位时间戳)
	LastOpTimeStr string `json:"lastOpTimeStr"` // 上次修改时间(YYYY-MM-DD hh:mm:ss)
	ParentId      String `json:"parentId"`      // 父文件夹ID
	Rev           int64  `json:"rev"`           // 文件夹版本号
}

type GetFolderInfoRequest struct {
	Folder  String `url:"folderId"`
	ShareId int64  `url:"shareId,omitempty"`
	Dt      int8   `url:"dt"`
}

type GetFolderFileOption func(*GetFolderInfoRequest)

func (c *client) GetFolderInfo(ctx context.Context, folder String, opts ...GetFolderFileOption) (*GetFolderInfoResponse, error) {
	req := &GetFolderInfoRequest{
		Folder: folder,
		Dt:     1,
	}

	for _, opt := range opts {
		opt(req)
	}

	var (
		result = new(GetFolderInfoResponse)
	)

	if err := c.get(ctx, ApiGetFolderInfo, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
