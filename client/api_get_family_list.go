package client

import "context"

type FamilyInfo struct {
	Count      int    `json:"count"`
	CreateTime string `json:"createTime"`
	FamilyId   String `json:"familyId"`
	RemarkName string `json:"remarkName"`
	Type       int    `json:"type"`
	UseFlag    int    `json:"useFlag"`
	UserRole   int    `json:"userRole"`
	ExpireTime string `json:"expireTime,omitempty"`
}

type GetFamilyListResponse struct {
	FamilyInfoResp []FamilyInfo `json:"familyInfoResp"`
}

func (c *client) GetFamilyList(ctx context.Context) (*GetFamilyListResponse, error) {
	var (
		result = new(GetFamilyListResponse)
	)

	if err := c.get(ctx, ApiGetFamilyList, nil, result); err != nil {
		return nil, err
	}

	return result, nil
}
