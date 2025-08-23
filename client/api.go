package client

type ApiURL string

const (
	ApiGetShareInfoByCode = "/open/share/getShareInfoByCode.action"

	ApiGetFileInfo           = "/open/file/getFileInfo.action"
	ApiGetFileDownloadURL    = "/open/file/getFileDownloadUrl.action"
	ApiGetFolderInfo         = "/open/file/getFolderInfo.action"
	ApiGetNewVLCVideoPlayURL = "/open/file/getNewVlcVideoPlayUrl.action"
	ApiGetUpResourceShare    = "/open/share/getUpResourceShare.action"
	ApiGetUserInfo           = "/open/user/getUserInfo.action"
	ApiGetUserPrivileges     = "/open/user/getUserPrivileges.action"
	ApiListResourceShareDir  = "/open/share/listResourceShareDir.action"
	ApiListShareDir          = "/open/share/listShareDir.action"

	ApiSubscribeGetUser = "/open/subscribe/getUser.action"
)

// RequiresToken 是否需要token
func (a ApiURL) RequiresToken() bool {
	switch a {
	case
		ApiGetFileInfo,
		ApiGetFolderInfo,
		ApiGetFileDownloadURL,
		ApiGetNewVLCVideoPlayURL,
		ApiGetUserInfo,
		ApiGetUserPrivileges,
		ApiListResourceShareDir:
		return true
	}

	return false
}
