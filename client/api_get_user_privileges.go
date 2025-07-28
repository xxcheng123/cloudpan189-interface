package client

import (
	"context"
)

type GetUserPrivilegesResponse struct {
	ResCode            int     `json:"res_code"`
	ResMessage         string  `json:"res_message"`
	BeginTime          string  `json:"beginTime"`
	ChannelId          string  `json:"channelId"`
	EndTime            string  `json:"endTime"`
	IsNetSDKOpen       string  `json:"isNetSDKOpen"`
	PackOverLimitSize  int64   `json:"packOverLimitSize"`
	PicEditRemainTimes int     `json:"picEditRemainTimes"`
	PicEditTotalTimes  int     `json:"picEditTotalTimes"`
	PreDecomp          int64   `json:"preDecomp"`
	ShareFileNum       int     `json:"shareFileNum"`
	TransChannel       int     `json:"transChannel"`
	TransConcurrent    int     `json:"transConcurrent"`
	TransDayFlow       int64   `json:"transDayFlow"`
	TransDownSpeed     int     `json:"transDownSpeed"` // 可以排查有没有被限速
	TransFileSize      float64 `json:"transFileSize"`
	TransQos           int     `json:"transQos"`
	TransSpeed         int     `json:"transSpeed"`
	TransVideo         int     `json:"transVideo"`
	TryNum             int     `json:"tryNum"`
	UsedDayFlow        int     `json:"usedDayFlow"`
	UserLevel          int     `json:"userLevel"`
	VipExpiredTime     string  `json:"vipExpiredTime"`
}

func (c *client) GetUserPrivileges(ctx context.Context) (*GetUserPrivilegesResponse, error) {
	var (
		result = new(GetUserPrivilegesResponse)
	)

	if err := c.get(ctx, ApiGetUserPrivileges, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}
