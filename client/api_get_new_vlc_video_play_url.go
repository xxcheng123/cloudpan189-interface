package client

import (
	"context"
)

type GetNewVLCVideoPlayURLRequest struct {
	ShareId string `url:"shareId"`
	FileId  String `url:"fileId"`
	Type    int8   `url:"type"`
	Dt      int8   `url:"dt"`
}

type GetNewVLCVideoPlayURLResponse struct {
	ResCode    int    `json:"res_code"`
	ResMessage string `json:"res_message"`
	Normal     struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		URL       string `json:"url"`
		VideoKbps int    `json:"videoKbps"`
	} `json:"normal"`
}

type WithGetNewVLCVideoPlayURLRequestOption func(*GetNewVLCVideoPlayURLRequest)

func (c *client) GetNewVLCVideoPlayURL(ctx context.Context, fileId String, opts ...WithGetNewVLCVideoPlayURLRequestOption) (*GetNewVLCVideoPlayURLResponse, error) {
	req := &GetNewVLCVideoPlayURLRequest{
		FileId: fileId,
		Type:   2,
	}

	for _, opt := range opts {
		opt(req)
	}

	var (
		result = new(GetNewVLCVideoPlayURLResponse)
	)

	if err := c.get(ctx, ApiGetNewVLCVideoPlayURL, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
