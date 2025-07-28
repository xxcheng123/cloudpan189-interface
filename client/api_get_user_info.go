package client

import (
	"context"
)

type GetUserInfoResponse struct {
	ResCode         int     `json:"res_code"`
	ResMessage      string  `json:"res_message"`
	Available       int64   `json:"available"`
	Capacity        int64   `json:"capacity"`
	ExtPicAvailable int     `json:"extPicAvailable"`
	ExtPicCapacity  int     `json:"extPicCapacity"`
	ExtPicUsed      int     `json:"extPicUsed"`
	HasFamily       int     `json:"hasFamily"`
	LoginName       string  `json:"loginName"`
	Mail189UsedSize int     `json:"mail189UsedSize"`
	MaxFilesize     float64 `json:"maxFilesize"`
	OrderAmount     int     `json:"orderAmount"`
	ProvinceCode    string  `json:"provinceCode"`
	UEncrypt        string  `json:"uEncrypt"`
	UMd5            string  `json:"uMd5"`
}

func (c *client) GetUserInfo(ctx context.Context) (*GetUserInfoResponse, error) {
	var (
		result = new(GetUserInfoResponse)
	)

	if err := c.get(ctx, ApiGetUserInfo, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}
