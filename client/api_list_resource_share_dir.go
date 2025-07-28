package client

import (
	"context"
	"fmt"
)

type ListResourceShareFileRequest struct {
	ShareId    string `url:"shareId"`
	FileId     String `url:"fileId"`
	UpUserId   string `url:"upUserId"`
	IsFolder   bool   `url:"isFolder"`
	IconOption int    `url:"iconOption"`
	OrderBy    string `url:"orderBy"`
	Descending bool   `url:"descending"`
	PageNum    int    `url:"pageNum"`
	PageSize   int    `url:"pageSize"`
}

type WithListResourceShareDirRequestOption func(*ListResourceShareFileRequest)

func (c *client) ListResourceShareDir(ctx context.Context, upUserId string, shareId int64, fileId String, opts ...WithListResourceShareDirRequestOption) (*ListFilesResponse, error) {
	req := &ListResourceShareFileRequest{
		UpUserId:   upUserId,
		ShareId:    fmt.Sprintf("%d", shareId),
		FileId:     fileId,
		IsFolder:   true,
		IconOption: 5,
		Descending: true,
		OrderBy:    "lastOpTime",
		PageNum:    1,
		PageSize:   100,
	}

	for _, opt := range opts {
		opt(req)
	}

	var (
		result = new(ListFilesResponse)
	)

	if err := c.get(ctx, ApiListResourceShareDir, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
