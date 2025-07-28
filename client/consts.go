package client

import "time"

const (
	TVAppKey             = "600100885"
	TVAppSignatureSecret = "fe5734c74c2f96a38157f420b32dc995"
	TvVersion            = "6.5.5"
	AndroidTV            = "FAMILY_TV"
	TvChannelId          = "home02"

	ApiUrl = "https://api.cloud.189.cn"

	UserAgent = "EcloudTV/6.5.5 (PJX110; unknown; home02) Android/35"
	Accept    = "application/json;charset=UTF-8"

	DefaultTimeout = time.Second * 5
)
