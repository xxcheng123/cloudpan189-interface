package client

import (
	"context"
	"github.com/google/uuid"
)

type GetShareInfoRequest struct {
	ShareCode  string `url:"shareCode"`
	AccessCode string `url:"accessCode"`
	UUID       string `url:"uuid"`
}

type GetShareInfoOption func(*GetShareInfoRequest)

type GetShareInfoResponse struct {
	ResCode    int    `json:"res_code"`
	ResMessage string `json:"res_message"`
	AccessCode string `json:"accessCode"`
	Creator    struct {
		IconURL      string `json:"iconURL"`
		NickName     string `json:"nickName"`
		Oper         bool   `json:"oper"`
		OwnerAccount string `json:"ownerAccount"`
		SuperVip     int    `json:"superVip"`
		Vip          int    `json:"vip"`
	} `json:"creator"`
	ExpireTime     int    `json:"expireTime"`
	ExpireType     int    `json:"expireType"`
	FileCreateDate string `json:"fileCreateDate"`
	FileId         String `json:"fileId"`
	FileLastOpTime string `json:"fileLastOpTime"`
	FileName       string `json:"fileName"`
	FileSize       int    `json:"fileSize"`
	FileType       string `json:"fileType"`
	IsFolder       bool   `json:"isFolder"`
	MediaType      int    `json:"mediaType"`
	NeedAccessCode int    `json:"needAccessCode"`
	ReviewStatus   int    `json:"reviewStatus"`
	ShareDate      int64  `json:"shareDate"`
	ShareId        int64  `json:"shareId"`
	ShareMode      int    `json:"shareMode"`
	ShareType      int    `json:"shareType"`
}

func (c *client) GetShareInfo(ctx context.Context, shareCode string, opts ...GetShareInfoOption) (*GetShareInfoResponse, error) {
	req := &GetShareInfoRequest{
		ShareCode: shareCode,
		UUID:      uuid.NewString(),
	}

	for _, opt := range opts {
		opt(req)
	}

	var (
		result = new(GetShareInfoResponse)
	)

	if err := c.get(ctx, ApiGetShareInfoByCode, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
