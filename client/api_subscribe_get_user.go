package client

import (
	"context"
	"time"
)

type SubscribeGetUserRequest struct {
	UserId string `url:"userId"`
}

type SubscribeGetUserInfo struct {
	Id               int64       `json:"id"`
	UserId           string      `json:"userId"`
	Name             string      `json:"name"`
	Picture          string      `json:"picture"`
	FansNum          int         `json:"fansNum"`
	FollowNum        int         `json:"followNum"`
	Status           int         `json:"status"`
	Ip               interface{} `json:"ip"`
	IsVip            int         `json:"isVip"`
	IsFirstSubscribe int         `json:"isFirstSubscribe"`
	Phone            string      `json:"phone"`
	Brief            interface{} `json:"brief"`
	OrderNum         interface{} `json:"orderNum"`
	ExtendOne        string      `json:"extendOne"`
	ExtendTwo        string      `json:"extendTwo"`
	CreateTime       time.Time   `json:"createTime"`
	UpdateTime       time.Time   `json:"updateTime"`
	PushSwitch       int         `json:"pushSwitch"`
	ExtendThree      interface{} `json:"extendThree"`
	ExtendFour       interface{} `json:"extendFour"`
	IsFollow         interface{} `json:"isFollow"`
}

type SubscribeGetUserResponse struct {
	Code string               `json:"code"`
	Msg  string               `json:"msg"`
	Data SubscribeGetUserInfo `json:"data"`
}

func (c *client) SubscribeGetUser(ctx context.Context, userId string) (*SubscribeGetUserResponse, error) {
	req := &SubscribeGetUserRequest{
		UserId: userId,
	}

	var (
		result = new(SubscribeGetUserResponse)
	)

	if err := c.get(ctx, ApiSubscribeGetUser, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
