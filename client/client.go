package client

import (
	"context"
	"resty.dev/v3"
)

type Client interface {
	WithDebug(flags ...bool) Client
	WithToken(token AuthToken) Client
	WithForceWithToken(flags ...bool) Client
	WithClient(client *resty.Client) Client

	GetShareInfo(ctx context.Context, shareCode string, opts ...GetShareInfoOption) (*GetShareInfoResponse, error)
	GetFileDownload(ctx context.Context, fileId String, opts ...GetFileDownloadOption) (*GetFileDownloadResponse, error)
	GetFileInfo(ctx context.Context, fileId String, opts ...GetFileFileOption) (*GetFileInfoResponse, error)
	GetFolderInfo(ctx context.Context, folder String, opts ...GetFolderFileOption) (*GetFolderInfoResponse, error)
	GetNewVLCVideoPlayURL(ctx context.Context, fileId String, opts ...WithGetNewVLCVideoPlayURLRequestOption) (*GetNewVLCVideoPlayURLResponse, error)
	GetUpResourceShare(ctx context.Context, upUserId string, pageNum int64, pageSize int64, opts ...GetUpResourceShareRequestOption) (*GetUpResourceShareResponse, error)
	GetUserInfo(ctx context.Context) (*GetUserInfoResponse, error)
	GetUserPrivileges(ctx context.Context) (*GetUserPrivilegesResponse, error)
	ListResourceShareDir(ctx context.Context, upUserId string, shareId int64, fileId String, opts ...WithListResourceShareDirRequestOption) (*ListFilesResponse, error)
	ListShareDir(ctx context.Context, shareId int64, fileId String, opts ...WithListShareFileRequestOption) (*ListFilesResponse, error)
	SubscribeGetUser(ctx context.Context, userId string) (*SubscribeGetUserResponse, error)
}

type client struct {
	isDebug        bool
	forceWithToken bool // 有的接口请求时不是必须携带 token 的，如果设置 true，那么会强制携带

	authToken  AuthToken
	httpClient *resty.Client
}

func New() Client {
	return &client{
		httpClient: httpClient,
	}
}
